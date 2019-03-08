PROJECT=$(shell gcloud config list project --format=flattened | awk 'FNR == 1 {print $$2}')
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
	gcloud projects add-iam-policy-binding $(PROJECT) --member serviceAccount:$(PROJECT)@appspot.gserviceaccount.com --role roles/bigquery.admin
