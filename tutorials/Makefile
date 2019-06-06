BASEDIR = $(shell pwd)
GITFOLDER=$(CLOUD_JOURNEY_TUTORIALS_PATH)

deploy: 
	cp -r $(BASEDIR)/* $(GITFOLDER)
	cd $(CLOUD_JOURNEY_TUTORIALS_PATH) && \
	git add * && \
	git commit -m "syncing from source" 
	

publish:
	cd $(CLOUD_JOURNEY_TUTORIALS_PATH) && git push