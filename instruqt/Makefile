BASEDIR = $(shell pwd)

app: clean build serve

main:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "$(BASEDIR)/main" "$(BASEDIR)/main.go"

build: main
	docker build -t redirect "$(BASEDIR)/."
	
serve: 
	docker run --name=redirect -d -P -p 8080:80 \
	-e INSTRUCT_GCP_PROJECTID=$(CLOUD_JOURNEY_PROJECT) \
	-e PUBLIC_URL=$(CLOUD_JOURNEY_PUBLIC_URL) \
	redirect

clean:
	-docker stop redirect
	-docker rm redirect
	-docker rmi redirect