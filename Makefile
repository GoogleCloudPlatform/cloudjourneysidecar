PROJECT=$(GOOGLE_CLOUD_PROJECT)
BASEDIR = $(shell pwd)
VERSION=$(shell cat .version)
NEXT=$(shell expr $(VERSION) + 1)
TESTURL=https://$(GOOGLE_CLOUD_PROJECT).appspot.com/status\?quest\=
VMNAME=deleteme
.DEFAULT_GOAL := install
DATESUFFIX=$(shell date +"%m-%u-%H-%M")


env:
	gcloud config set project $(PROJECT)

deploy: env permissions
	export GOPATH=$(GOPATH):$(BASEDIR)
	gcloud app deploy -q

create: env
	-gcloud app create --region=us-central -q

install: env create deploy permissions apis check


version: 
	@printf $(NEXT) > .version
	@@echo Version is: $(NEXT)
	cp .version $(BASEDIR)/../game


permissions: 
	gcloud projects add-iam-policy-binding $(PROJECT) \
	--member serviceAccount:$(PROJECT)@appspot.gserviceaccount.com \
	--role roles/bigquery.admin

apis:
	gcloud services enable sqladmin.googleapis.com --project $(PROJECT)
	gcloud services enable translate.googleapis.com --project $(PROJECT)

main:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "$(BASEDIR)/main" \
	"$(BASEDIR)/main.go" "$(BASEDIR)/responders.go" \
	"$(BASEDIR)/gcp_compute.go"  "$(BASEDIR)/gcp_dev.go" \
	"$(BASEDIR)/gcp_data.go" 

build: main
	docker build -t helper "$(BASEDIR)/."

check:
	curl https://$(PROJECT).appspot.com/health




setup.sys_01: 
	@echo ------------------------------------------------------
	@echo Performing Sys_01 setup 
	@echo ------------------------------------------------------
	gcloud compute --project=$(GOOGLE_CLOUD_PROJECT) instances create $(VMNAME) \
	--zone=us-central1-a --machine-type=f1-micro --subnet=default \
	--boot-disk-size=10GB --boot-disk-type=pd-standard \
	--boot-disk-device-name=$(VMNAME)

test.sys_01:
	@echo ------------------------------------------------------
	@echo Performing Sys_01 test 
	@echo ------------------------------------------------------
	curl -v --silent $(TESTURL)checkSys01 2>&1 | grep true

clean.sys_01:
	@echo ------------------------------------------------------
	@echo Performing Sys_01 cleanup 
	@echo ------------------------------------------------------
	-gcloud compute --project=$(GOOGLE_CLOUD_PROJECT) instances delete $(VMNAME) \
	-q --zone=us-central1-a


setup.sys_02: setup.sys_01
	@echo ------------------------------------------------------
	@echo Performing Sys_02 setup 
	@echo ------------------------------------------------------
	gcloud beta compute disks create $(VMNAME)-attach \
	--project=$(GOOGLE_CLOUD_PROJECT) --type=pd-standard --size=500GB \
	--zone=us-central1-a --physical-block-size=4096
	gcloud compute instances attach-disk $(VMNAME) \
      --disk $(VMNAME)-attach --zone us-central1-a

test.sys_02:
	@echo ------------------------------------------------------
	@echo Performing Sys_02 test 
	@echo ------------------------------------------------------
	curl -v --silent $(TESTURL)checkSys02 2>&1 | grep true	 

clean.sys_02:	 
	@echo ------------------------------------------------------
	@echo Performing Sys_02 cleanup 
	@echo ------------------------------------------------------
	-gcloud compute instances detach-disk $(VMNAME) \
      --disk $(VMNAME)-attach --zone us-central1-a -q
	-gcloud beta compute disks delete $(VMNAME)-attach \
	--project=$(GOOGLE_CLOUD_PROJECT) --zone=us-central1-a  -q
	$(MAKE) clean.sys_01


setup.sys_03: setup.sys_01
	@echo ------------------------------------------------------
	@echo Performing Sys_03 setup 
	@echo ------------------------------------------------------
	gcloud compute ssh $(VMNAME) --command="sudo apt update -y" --zone us-central1-a
	gcloud compute ssh $(VMNAME) --command="sudo apt install nginx -y" --zone us-central1-a
	gcloud compute instances stop $(VMNAME) --zone us-central1-a -q
	gcloud compute instances detach-disk $(VMNAME) \
      --disk $(VMNAME) --zone us-central1-a -q
	gcloud compute images create nginx-$(VMNAME) \
	--source-disk-zone us-central1-a --source-disk=$(VMNAME)

test.sys_03:
	@echo ------------------------------------------------------
	@echo Performing Sys_03 test 
	@echo ------------------------------------------------------
	curl -v --silent $(TESTURL)checkSys03 2>&1 | grep true

clean.sys_03: 
	@echo ------------------------------------------------------
	@echo Performing Sys_03 cleanup 
	@echo ------------------------------------------------------
	-gcloud beta compute images delete nginx-$(VMNAME) \
	--project=$(GOOGLE_CLOUD_PROJECT)   -q
	-gcloud beta compute disks delete $(VMNAME) \
	--project=$(GOOGLE_CLOUD_PROJECT) --zone=us-central1-a  -q
	$(MAKE) clean.sys_01

setup.sys_04: setup.sys_03
	@echo ------------------------------------------------------
	@echo Performing Sys_04 setup 
	@echo ------------------------------------------------------
	gcloud compute --project=$(GOOGLE_CLOUD_PROJECT) instance-templates \
	create $(VMNAME)-template --machine-type=g1-small \
	--tags=http-server --image=nginx-$(VMNAME) --image-project=$(GOOGLE_CLOUD_PROJECT) \
	--boot-disk-size=10GB --boot-disk-type=pd-standard \
	--boot-disk-device-name=$(VMNAME)-template

test.sys_04:
	@echo ------------------------------------------------------
	@echo Performing Sys_04 test 
	@echo ------------------------------------------------------
	curl -v --silent $(TESTURL)checkSys04 2>&1 | grep true	 

clean.sys_04: clean.sys_03	 
	@echo ------------------------------------------------------
	@echo Performing Sys_04 cleanup 
	@echo ------------------------------------------------------
	-gcloud compute --project=$(GOOGLE_CLOUD_PROJECT) instance-templates \
	delete $(VMNAME)-template -q


setup.sys_05: setup.sys_04
	@echo ------------------------------------------------------
	@echo Performing Sys_05 setup 
	@echo ------------------------------------------------------
	gcloud compute --project=$(GOOGLE_CLOUD_PROJECT) health-checks create tcp \
	"deleteme-healthcheck" --timeout "5" --check-interval "10" \
	--unhealthy-threshold "3" --healthy-threshold "2" --port "80"

	gcloud beta compute --project=$(GOOGLE_CLOUD_PROJECT) instance-groups \
	managed create deleteme-mig --base-instance-name=deleteme-mig \
	--template=$(VMNAME)-template --size=1 --zone=us-central1-a \
	--health-check=deleteme-healthcheck --initial-delay=300

	gcloud beta compute --project "$(GOOGLE_CLOUD_PROJECT)" instance-groups \
	managed set-autoscaling "deleteme-mig" --zone "us-central1-a" \
	--cool-down-period "60" --max-num-replicas "10" --min-num-replicas "5" \
	--target-cpu-utilization "0.6"

	gcloud compute addresses create lb-ip-cr --ip-version=IPV4 --global \
	--project "$(GOOGLE_CLOUD_PROJECT)"
	
	gcloud compute addresses create lb-ipv6-cr --ip-version=IPV6 --global \
	--project "$(GOOGLE_CLOUD_PROJECT)"
	
	gcloud compute instance-groups managed set-named-ports \
	deleteme-mig --named-ports http:80 --zone us-central1-a \
	--project "$(GOOGLE_CLOUD_PROJECT)"
	
	gcloud compute http-health-checks create deleteme-hc-lb \
	--project "$(GOOGLE_CLOUD_PROJECT)"
	gcloud compute backend-services create deleteme-be --http-health-checks \
	deleteme-hc-lb --global --project "$(GOOGLE_CLOUD_PROJECT)"

	gcloud compute url-maps create deleteme-map --default-service deleteme-be
	gcloud compute target-http-proxies create deleteme-proxy \
	--url-map deleteme-map --project "$(GOOGLE_CLOUD_PROJECT)"

	gcloud compute forwarding-rules create deleteme-forward --global \
	--ports 80 --target-http-proxy deleteme-proxy \
	--project "$(GOOGLE_CLOUD_PROJECT)"

	gcloud compute backend-services add-backend deleteme-be \
	--instance-group=deleteme-mig --instance-group-zone us-central1-a --global \
	--project "$(GOOGLE_CLOUD_PROJECT)"

test.sys_05:
	@echo ------------------------------------------------------
	@echo Performing Sys_05 test 
	@echo ------------------------------------------------------
	curl -v --silent $(TESTURL)checkSys05 2>&1 | grep true	 	

	

clean.sys_05:
	@echo ------------------------------------------------------
	@echo Performing Sys_05 cleanup 
	@echo ------------------------------------------------------
	-gcloud compute backend-services remove-backend deleteme-be \
	--instance-group=deleteme-mig --instance-group-zone us-central1-a --global \
	--project "$(GOOGLE_CLOUD_PROJECT)"

	-gcloud compute forwarding-rules delete deleteme-forward --global -q \
	--project "$(GOOGLE_CLOUD_PROJECT)"
	
	-gcloud compute target-http-proxies delete deleteme-proxy -q \
	--project "$(GOOGLE_CLOUD_PROJECT)"
	
	-gcloud compute url-maps delete deleteme-map -q \
	--project "$(GOOGLE_CLOUD_PROJECT)"
	
	-gcloud compute backend-services delete deleteme-be -q --global \
	--project "$(GOOGLE_CLOUD_PROJECT)"
	
	-gcloud compute http-health-checks delete deleteme-hc-lb -q
	
	-gcloud compute addresses delete lb-ip-cr   \
	--project "$(GOOGLE_CLOUD_PROJECT)" -q --global
	
	-gcloud compute addresses delete lb-ipv6-cr \
	--project "$(GOOGLE_CLOUD_PROJECT)" -q --global
	
	
	-gcloud beta compute --project=$(GOOGLE_CLOUD_PROJECT) instance-groups \
	managed delete deleteme-mig --zone us-central1-a -q
	-gcloud compute --project=$(GOOGLE_CLOUD_PROJECT) health-checks delete \
	"deleteme-healthcheck" -q

	$(MAKE) clean.sys_04	 



sys: sys_01 sys_02 sys_03 sys_04 sys_05
sys_01: 
	$(MAKE) setup.sys_01 test.sys_01 clean.sys_01
sys_02: 
	$(MAKE) setup.sys_02 test.sys_02 clean.sys_02
sys_03: 
	$(MAKE) setup.sys_03 test.sys_03 clean.sys_03
sys_04: 
	$(MAKE) setup.sys_04 test.sys_04 clean.sys_04
sys_05: 
	$(MAKE) setup.sys_05 test.sys_05 clean.sys_05

sys_clean: clean.sys_01 clean.sys_02 clean.sys_03 clean.sys_04 clean.sys_05





setup.dev_01: 
	@echo ------------------------------------------------------
	@echo Performing dev_01 setup 
	@echo ------------------------------------------------------
	cd ../questcode/dev01 && gcloud functions deploy tokenGenerator \
	--project=$(GOOGLE_CLOUD_PROJECT) --runtime nodejs8 --trigger-http

test.dev_01:
	@echo ------------------------------------------------------
	@echo Performing dev_01 test 
	@echo ------------------------------------------------------
	curl -v --silent $(TESTURL)checkDev01 2>&1 | grep true

clean.dev_01:
	@echo ------------------------------------------------------
	@echo Performing dev_01 cleanup 
	@echo ------------------------------------------------------
	-gcloud functions delete tokenGenerator --project=$(GOOGLE_CLOUD_PROJECT) -q



setup.dev_02: 
	@echo ------------------------------------------------------
	@echo Performing dev_02 setup 
	@echo ------------------------------------------------------
	-cd ../questcode/dev02 && gcloud builds submit . \
	--tag gcr.io/$(GOOGLE_CLOUD_PROJECT)/randomcolor  

test.dev_02:
	@echo ------------------------------------------------------
	@echo Performing dev_02 test 
	@echo ------------------------------------------------------
	curl -v --silent $(TESTURL)checkDev02 2>&1 | grep true	


setup.dev_03: 
	@echo ------------------------------------------------------
	@echo Performing dev_03 setup 
	@echo ------------------------------------------------------
	gcloud beta run deploy randomcolor --allow-unauthenticated \
	--image gcr.io/$(GOOGLE_CLOUD_PROJECT)/randomcolor  

test.dev_03:
	@echo ------------------------------------------------------
	@echo Performing dev_03 test 
	@echo ------------------------------------------------------
	curl -v --silent $(TESTURL)checkDev03 2>&1 | grep true		

clean.dev_03: 
	@echo ------------------------------------------------------
	@echo Performing dev_03 cleanup 
	@echo ------------------------------------------------------
	-gcloud beta run services delete randomcolor	\
	--project=$(GOOGLE_CLOUD_PROJECT) -q


setup.dev_04: 
	@echo ------------------------------------------------------
	@echo Performing dev_04 setup 
	@echo ------------------------------------------------------
	gsutil mb gs://$(GOOGLE_CLOUD_PROJECT)
	cd ../questcode/dev04 && gcloud functions deploy bucketloader \
	--project=$(GOOGLE_CLOUD_PROJECT) --runtime nodejs8 \
	--trigger-bucket=$(GOOGLE_CLOUD_PROJECT)

test.dev_04:
	@echo ------------------------------------------------------
	@echo Performing dev_04 test 
	@echo ------------------------------------------------------
	curl -v --silent $(TESTURL)checkDev04 2>&1 | grep true	

clean.dev_04: 
	@echo ------------------------------------------------------
	@echo Performing dev_04 cleanup 
	@echo ------------------------------------------------------
	-gcloud functions delete bucketloader --project=$(GOOGLE_CLOUD_PROJECT) -q
	-gsutil rb gs://$(GOOGLE_CLOUD_PROJECT)


setup.dev_05: 
	@echo ------------------------------------------------------
	@echo Performing dev_05 setup 
	@echo ------------------------------------------------------
	gsutil mb gs://$(GOOGLE_CLOUD_PROJECT)
	gsutil iam ch serviceAccount:$(GOOGLE_CLOUD_PROJECT)@appspot.gserviceaccount.com:objectCreator gs://$(GOOGLE_CLOUD_PROJECT)
	cd ../questcode/dev05 && gcloud functions deploy translateFiles \
	--project=$(GOOGLE_CLOUD_PROJECT) --runtime nodejs8 \
	--trigger-bucket=$(GOOGLE_CLOUD_PROJECT)
	gsutil cp ../questcode/dev05/test.txt gs://$(GOOGLE_CLOUD_PROJECT)


test.dev_05:
	@echo ------------------------------------------------------
	@echo Performing dev_05 test 
	@echo ------------------------------------------------------
	curl -v --silent $(TESTURL)checkDev05 2>&1 | grep true		

clean.dev_05: 
	@echo ------------------------------------------------------
	@echo Performing dev_05 cleanup 
	@echo ------------------------------------------------------
	-gcloud functions delete translateFiles --project=$(GOOGLE_CLOUD_PROJECT) -q
	-gsutil rm -rf gs://$(GOOGLE_CLOUD_PROJECT)/*
	-gsutil rb gs://$(GOOGLE_CLOUD_PROJECT)	

sleep: 
	sleep 30

dev: dev_01 dev_02 dev_03 dev_04 dev_05
dev_01: 
	$(MAKE) setup.dev_01 test.dev_01 clean.dev_01
dev_02: 
	$(MAKE) setup.dev_02 test.dev_02 
dev_03: 
	$(MAKE) setup.dev_03 test.dev_03 clean.dev_03
dev_04: 
	$(MAKE) setup.dev_04 test.dev_04 clean.dev_04
dev_05: 
	$(MAKE) setup.dev_05 sleep test.dev_05 clean.dev_05

dev_clean: clean.dev_01 clean.dev_02 clean.dev_03 clean.dev_04 clean.dev_05	



setup.data_01: 
	@echo ------------------------------------------------------
	@echo Performing data_01 setup 
	@echo ------------------------------------------------------
	bq query --nouse_legacy_sql 'SELECT name, SUM(number) as total \
	FROM `bigquery-public-data.usa_names.usa_1910_2013` \
	WHERE year = 1976 GROUP BY name, gender ORDER BY total \
	DESC LIMIT 10'
	--project=$(GOOGLE_CLOUD_PROJECT) 

test.data_01:
	@echo ------------------------------------------------------
	@echo Performing data_01 test 
	@echo ------------------------------------------------------
	curl -v --silent $(TESTURL)checkData01 2>&1 | grep true



setup.data_02: 
	@echo ------------------------------------------------------
	@echo Performing data_02 setup 
	@echo ------------------------------------------------------
	gsutil mb gs://$(GOOGLE_CLOUD_PROJECT)

test.data_02:
	@echo ------------------------------------------------------
	@echo Performing data_02 test 
	@echo ------------------------------------------------------
	curl -v --silent $(TESTURL)checkData02 2>&1 | grep true	

clean.data_02: 
	@echo ------------------------------------------------------
	@echo Performing data_02 cleanup 
	@echo ------------------------------------------------------
	-gsutil rm -rf gs://$(GOOGLE_CLOUD_PROJECT)/*
	-gsutil rb gs://$(GOOGLE_CLOUD_PROJECT)


setup.data_03: 
	@echo ------------------------------------------------------
	@echo Performing data_03 setup 
	@echo ------------------------------------------------------
	gcloud sql instances create cloudsqltest-$(DATESUFFIX) \
	--project=$(GOOGLE_CLOUD_PROJECT)

test.data_03:
	@echo ------------------------------------------------------
	@echo Performing data_03 test 
	@echo ------------------------------------------------------
	curl -v --silent $(TESTURL)checkData03 2>&1 | grep true	

clean.data_03: 
	@echo ------------------------------------------------------
	@echo Performing data_03 cleanup 
	@echo ------------------------------------------------------
	-gcloud sql instances delete cloudsqltest-$(DATESUFFIX) \
	--project=$(GOOGLE_CLOUD_PROJECT) -q

data_03: 
	@echo ------------------------------------------------------
	@echo Performing data_03 setup 
	@echo ------------------------------------------------------
	gcloud sql instances create cloudsqltest-$(DATESUFFIX) \
	--project=$(GOOGLE_CLOUD_PROJECT)

	@echo ------------------------------------------------------
	@echo Performing data_03 test 
	@echo ------------------------------------------------------
	curl -v --silent $(TESTURL)checkData03 2>&1 | grep true	

	@echo ------------------------------------------------------
	@echo Performing data_03 cleanup 
	@echo ------------------------------------------------------
	-gcloud sql instances delete cloudsqltest-$(DATESUFFIX) \
	--project=$(GOOGLE_CLOUD_PROJECT) -q

setup.data_04: 
	@echo ------------------------------------------------------
	@echo Performing data_04 setup 
	@echo ------------------------------------------------------
	go run $(BASEDIR)/../questcode/data04/add/add.go

test.data_04:
	@echo ------------------------------------------------------
	@echo Performing data_04 test 
	@echo ------------------------------------------------------
	curl -v --silent $(TESTURL)checkData04 2>&1 | grep true		

clean.data_04: 
	@echo ------------------------------------------------------
	@echo Performing data_04 cleanup 
	@echo ------------------------------------------------------
	-go run $(BASEDIR)/../questcode/data04/remove/remove.go

setup.data_05: 
	@echo ------------------------------------------------------
	@echo Performing data_05 setup 
	@echo ------------------------------------------------------
	gcloud redis instances create deleteme --size=2 --region=us-central1 --project=$(GOOGLE_CLOUD_PROJECT)

test.data_05:
	@echo ------------------------------------------------------
	@echo Performing data_05 test 
	@echo ------------------------------------------------------
	curl -v --silent $(TESTURL)checkData05 2>&1 | grep true		

clean.data_05: 
	@echo ------------------------------------------------------
	@echo Performing data_05 cleanup 
	@echo ------------------------------------------------------
	-gcloud redis instances delete deleteme  --region=us-central1 --project=$(GOOGLE_CLOUD_PROJECT) -q


data: data_01 data_02 data_03 data_04 data_05
data_01: 
	$(MAKE) setup.data_01 test.data_01 
data_02: 
	$(MAKE) setup.data_02 test.data_02 clean.data_02
data_04: 
	$(MAKE) setup.data_04 test.data_04 clean.data_04
data_05: 
	$(MAKE) setup.data_05 sleep test.data_05 clean.data_05

data_clean: clean.data_01 clean.data_02 clean.data_03 clean.data_04 clean.data_05		