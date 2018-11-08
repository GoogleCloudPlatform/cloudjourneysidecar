ZONE=us-central1-c
REGION=us-central
BASEDIR = $(shell pwd)


env:
	gcloud config set project $(PROJECT)
	gcloud config set compute/zone $(ZONE)

deploy: env
	export GOPATH=$(GOPATH):$(BASEDIR)
	gcloud app deploy -q

create: env
	gcloud app create --region=$(REGION) -q

install: env create deploy
