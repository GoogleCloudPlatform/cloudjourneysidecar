ZONE=us-central1-c
REGION=us-central
BASEDIR = $(shell pwd)
.DEFAULT_GOAL := install


env:
	gcloud config set project $(PROJECT)
	gcloud config set compute/zone $(ZONE)

deploy: env
	export GOPATH=$(GOPATH):$(BASEDIR)
	gcloud app deploy -q

create: env
	gcloud app create --region=$(REGION) -q

install: env create deploy permissions


permissions: 
	gcloud iam service-accounts add-iam-policy-binding $(PROJECT)@appspot.gserviceaccount.com --member=serviceAccount:$(PROJECT)@appspot.gserviceaccount.com --role=roles/owner
