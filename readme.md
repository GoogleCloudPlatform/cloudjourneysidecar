# Companion app for GCP Quest. 

Thanks for agreeing to test GCPQuest. This is the companion app for GCPQuest. 
It must be installed in your GCP project to get GCPQuest to work. 

## Steps
* Create New Project (Must limit length to 16 chars)
* Open Cloud Shell in Project
* run `git clone https://github.com/tpryan/GCPQuest-Companion.git`
* run `cd GCPQuest-Companion`
* run `make`
* browse to `https://[project id].appspot.com/status`

Should output: 

```js
[{
  "quest": "intro_sys",
  "complete": false,
  "notes": "API not enabled yet."
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

If that's working **congrats**, you're all set to play. 