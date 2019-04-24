PROJECT=$(GOOGLE_CLOUD_PROJECT)
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

deploy.container: env
	gcloud app deploy -q --image-url=gcr.io/instruqt-shadow/gcpquesthelper

create: env
	gcloud app create --region=$(REGION) -q

install: env create deploy permissions


permissions: 
	gcloud projects add-iam-policy-binding $(PROJECT) --member serviceAccount:$(PROJECT)@appspot.gserviceaccount.com --role roles/bigquery.admin


main:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "$(BASEDIR)/main" "$(BASEDIR)/main.go" "$(BASEDIR)/responders.go" "$(BASEDIR)/vm.go"  "$(BASEDIR)/functions.go" "$(BASEDIR)/bigquery.go" 

build: main
	docker build -t helper "$(BASEDIR)/."


serve: 
	docker run --name=helper -d -P -p 8080:8080  helper

clean:
	-docker stop helper
	-docker rm helper
	-docker rmi helper	