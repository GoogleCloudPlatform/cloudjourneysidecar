PROJECT=$(GOOGLE_CLOUD_PROJECT)
REGION=$(CLOUD_JOURNEY_REGION)
BASEDIR = $(shell pwd)
.DEFAULT_GOAL := install


env:
	gcloud config set project $(PROJECT)

deploy: env permissions
	export GOPATH=$(GOPATH):$(BASEDIR)
	gcloud app deploy -q

create: env
	gcloud app create --region=$(REGION) -q

install: env create deploy permissions check


permissions: 
	gcloud projects add-iam-policy-binding $(PROJECT) --member serviceAccount:$(PROJECT)@appspot.gserviceaccount.com --role roles/bigquery.admin


main:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "$(BASEDIR)/main" "$(BASEDIR)/main.go" "$(BASEDIR)/responders.go" "$(BASEDIR)/vm.go"  "$(BASEDIR)/functions.go" "$(BASEDIR)/bigquery.go" 

build: main
	docker build -t helper "$(BASEDIR)/."

check:
	curl https://$(PROJECT).appspot.com/health
