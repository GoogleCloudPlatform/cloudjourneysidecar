BASEDIR = $(shell pwd)
PROJECT=$(CLOUD_JOURNEY_PROJECT)
BUCKET=$(CLOUD_JOURNEY_BUCKET)
SITEPATH=$(BASEDIR)/../output/GCPQuest/www

env:
	gcloud config set project $(PROJECT)

serve: 
	cd $(SITEPATH) && (http-server -s -o  &) 

version: 
	@printf $(NEXT) > .version
	@echo Version is: $(NEXT)
	cp .version $(BASEDIR)/../sidecar


overwrite:
	cp $(BASEDIR)/override/app.yaml $(SITEPATH)/../app.yaml
	cp $(BASEDIR)/override/app_stage.yaml $(SITEPATH)/../app_stage.yaml
	cp $(BASEDIR)/override/main.go $(SITEPATH)/../main.go
	cp $(BASEDIR)/override/index.html $(SITEPATH)/index.html
	cp $(BASEDIR)/override/main.js $(SITEPATH)/js/main.js
	cp $(BASEDIR)/override/package.json $(SITEPATH)/js/package.json
	cd $(BASEDIR)/helper && node input_length_fix.js
	cd $(BASEDIR)/helper && node analytics_insert.js
	cd $(BASEDIR)/helper && node replace_version.js

clean:  
	-rm -r $(SITEPATH)/*

copy:
	-rm -rf $(SITEPATH)/
	mkdir -p $(SITEPATH)/GCPQuest/www
	cp -r $(BASEDIR)/GCPQuest/* $(SITEPATH)/

compress:
	-cd $(SITEPATH)/ && pngquant -f *.png */*.png */*/*.png

deploy: clean copy overwrite compress
	-rm $(BASEDIR)/.DS_Store
	cd $(SITEPATH)/../ && gcloud app deploy -q --project $(PROJECT)

stage: clean copy overwrite compress
	-rm $(BASEDIR)/.DS_Store
	cd $(SITEPATH)/../ && gcloud app deploy app_stage.yaml -q --project $(PROJECT)	

test: serve
	cd test && npm test

prep: clean copy overwrite perms

perms: 
	-chmod -R g+rwxs /workspace/
	-cd test && npm install

testclean: 
	-kill $(shell lsof -t -i:8080)