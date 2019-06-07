PROJECT=$(GOOGLE_CLOUD_PROJECT)
BASEDIR = $(shell pwd)
VERSION=$(shell cat .version)
NEXT=$(shell expr $(VERSION) + 1)

.DEFAULT_GOAL := install


env:
	gcloud config set project $(PROJECT)

deploy: env permissions
	export GOPATH=$(GOPATH):$(BASEDIR)
	gcloud app deploy -q

create: env
	-gcloud app create --region=us-central -q

install: env create deploy permissions check


version: 
	@printf $(NEXT) > .version
	@echo Version is: $(NEXT)
	cp .version $(BASEDIR)/../game


permissions: 
	gcloud projects add-iam-policy-binding $(PROJECT) \
	--member serviceAccount:$(PROJECT)@appspot.gserviceaccount.com \
	--role roles/bigquery.admin


main:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "$(BASEDIR)/main" \
	"$(BASEDIR)/main.go" "$(BASEDIR)/responders.go" \
	"$(BASEDIR)/gcp_compute.go"  "$(BASEDIR)/gcp_dev.go" \
	"$(BASEDIR)/gcp_data.go" 

build: main
	docker build -t helper "$(BASEDIR)/."

check:
	curl https://$(PROJECT).appspot.com/health
