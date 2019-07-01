# Cloud Journey
## Sidecar application
This is the companion app for Cloud Journey. 
It must be installed in your GCP project to get Cloud Journey to work. 

### Requirements
* Create ENV variables
* `export CLOUD_JOURNEY_ZONE=[Zone to install the app in ]`
* `export CLOUD_JOURNEY_REGION=[Region to install the app in ]`
* `export CLOUD_JOURNEY_PROJECT=[ID of a project that you want to use for testing]`

## Steps
* Create New Project (Must limit length to 30 chars)
* Open Cloud Shell in Project
* run `git clone https://github.com/tpryan/GCPQuest-Companion.git`
* run `cd GCPQuest-Companion`
* run `make`
* browse to `https://[project id].appspot.com/status`

Should output: 

```js
[{
  "quest": "01_sys",
  "complete": false,
  "notes": "API not enabled yet."
}, {
  "quest": "01_data",
  "complete": false,
  "notes": ""
}, {
  "quest": "01_dev",
  "complete": false,
  "notes": "API not enabled yet."
}]
```

If that's working **congrats**, you're all set to play. 

## What does it do? 
The code in this repo when launched as an App Engine Application will 
interogate your project using Public GCP API's and try to figure out if you 
have properly completed the steps in Cloud Joruney's tutorials. It will try and 
figure out 3 things:

* Is there an f1micro Compute Engine Instance
* Does it have an attached disk
* Is there there a Cloud Function named tokengenerator
* Was a BigQuery query run that searched *bigquery-public-data.usa_names.usa_1910_2013*

"This is not an official Google Project."
