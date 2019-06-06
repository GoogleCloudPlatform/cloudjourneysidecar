# Cloud Journey
The overarching project for Cloud Journey code.

## Overview
The code consists of  4 directories


| Name  |Folder   |Description   |   
|---|---|---|
|Design Assets   |`\assets` | Various original design files     |
|Game Code   |`\game` | Actual game code that should be edited in RPGMaker      |
|Game Builder   |`\gamebuilder` | Container code that Cloud Build uses to publish from GCR      |
|Instruct Launching Container   |`\instruqt`  | Container for launching game on Instruqt platform     |
|Validator   |`\sidecar`   | App Engine application for valdiating quest completion     |
|Neos Tutorials   |`\tutorial`    | Neos tutorials that power the in console experience of this application     |

## How does it all work
The `game` is hosted on App Engine and is a HTML, JavaScript and CSS app 
written using [RPGMaker MV](http://www.rpgmakerweb.com/products/programs/rpg-maker-mv). 
The easiest way to edit it is using the RPGMaker MV IDE.

The `game` occasionally requires players perform tasks in an actual GCP project. 
When players are prompted to perform a task, they launch a 
[NEOS](http://go/neos) tutorial.  The code for the `tutorials` is where these 
are kept. However it is important to know that they are here for reference, and 
to actually change the tutorials on the website, you have to submit to Google3.  

In order to check that those tasks have been completed there is a `sidecar` 
application that users are prompted to run in their own projects in order to 
interrogate the project about what is currently deployed. 

## What does it take to add a quest. 
Each quest is comprise of 3 distinct parts that need to be added to make it 
work.

1. A NEOS entry in `/tutorials` outlining the task for somone to do. 
1. A `/sidecar` check of conditions in the project that would satisfy the task
1. Diaglow and decision login in `/game` that will introduce the quest and 
handle rewarding success and walking through failure. 
