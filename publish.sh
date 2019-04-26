
echo cp to $CLOUD_JOURNEY_SIDECAR_PATH
$(cp -r $PWD/* $CLOUD_JOURNEY_SIDECAR_PATH)
echo git adding to $CLOUD_JOURNEY_SIDECAR_PATH
(cd $CLOUD_JOURNEY_SIDECAR_PATH && git add * )
(cd $CLOUD_JOURNEY_SIDECAR_PATH && git commit -m "syncing from source" )
