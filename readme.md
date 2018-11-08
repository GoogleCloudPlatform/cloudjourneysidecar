# Companion app for GCP Quest. 

This is the companion app for GCPQuest. It must be installed in your GCP project
to get GCPQuest to work. 

## Steps
* Create New Project (Must limit length to 16 chars)
* Navigate to Compute Engine in Cloud Console - wait for GCE to be activated
* Open Cloud Shell in Project
* run `export PROJECT=[project id]` where project id is the id of the project 
you created.
* run `git clone https://github.com/tpryan/GCPQuest-Companion.git`
* run `cd GCPQuest-Companion`
* run `make`
* browse to `https://[project id].appspot.com/`

Should output: 

```js
[{
  "quest": "intro_sys",
  "complete": false,
  "notes": ""
}, {
  "quest": "intro_bigdata",
  "complete": false,
  "notes": ""
}, {
  "quest": "intro_dev",
  "complete": false,
  "notes": "API not enabled yet."
}]
```
