PROJECT=gcpquest-2019
ZONE=us-central1-c
REGION=us-central
BASEDIR = $(shell pwd)


env:
	gcloud config set project $(PROJECT)
	gcloud config set compute/zone $(ZONE)

deploy: env
	export GOPATH=$(GOPATH):$(BASEDIR)
	gcloud app deploy -q