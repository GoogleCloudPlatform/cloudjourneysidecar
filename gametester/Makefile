APPNAME=gametester
BASEDIR = $(shell pwd)

clean:
	-docker stop $(APPNAME)
	-docker rm $(APPNAME)

build:
	gcloud builds submit --timeout=1000 --config cloudbuild.yaml .

dockerrun: builddocker serve


builddocker: clean
	docker build -t $(APPNAME) "$(BASEDIR)/."

serve: 
	docker run --name=$(APPNAME) -d -P -p 8080:8080 $(APPNAME)		