# Cloud Journey
## Game Files

This code is the base game code for the in browser portion of the game.  


## Requirements
### Edit Game
* Install RPGMaker MV
### Deploy
* Check in to master on github 

## To Edit
* Launch RPGMaker MV
* Open game/GCPQuest/Game.rpgproject

## To Edit js
Javascript in the project can get overwritten by RPGMaker, so it helps to make 
sure that work gets edited in a separate folder, namely `override.` Then add to
the Makefile a workflow that replaces the correct js file. 

## To Deploy
* `make deploy`
* `make permissions`

## Javascript Reference
The custom javascript in the game communicates with the App Engine app to 
validate that players have performed the tasks that they are supposed to. 

### `checkHealth()`
#### Parameters
None
#### Description
Ensures that the App Engine side car application is up and running.

#### Return
Boolean

### `setRemoteStatus(label, id)`
#### Parameters
**lable**  The quest to check for completion  
**id** The RPGMaker variable to set to true if the quest is complete 
#### Description
Checks to see if a quest has been completed in the project. 

#### Return
Raw json results from App Engine app.  
Ex: 

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


