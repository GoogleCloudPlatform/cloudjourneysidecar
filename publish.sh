
echo cp to $CLOUD_JOURNEY_SIDECAR_PATH
$(cp -r $PWD/* $CLOUD_JOURNEY_SIDECAR_PATH)
echo git adding to $CLOUD_JOURNEY_SIDECAR_PATH
cd $CLOUD_JOURNEY_SIDECAR_PATH && git add * && git commit -m "syncing from source"
# deploy: 
# 	cp -r $(BASEDIR)/* $(GITFOLDER)
# 	cd $(CLOUD_JOURNEY_TUTORIALS_PATH) && \
# 	git add * && \
# 	git commit -m "syncing from source" 
	

# publish:
# 	cd $(CLOUD_JOURNEY_TUTORIALS_PATH) && git push