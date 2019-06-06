// Script reads in the contents of Map001, finds the Boss's dialog and 
// makes sure that the field length for project name is 30 and not 16. 
console.log("Input Length Fix");
var fs = require('fs');
var path = require('path');
var outputpath = path.resolve(__dirname, '../../output/GCPQuest/www/data/Map001.json');
var rawdata = fs.readFileSync(outputpath, 'utf8');
let map = JSON.parse(rawdata);  



for (var i=0; i< map.events.length; i++ ){
    var event = map.events[i];
    if (map.events[i] != null && typeof map.events[i].name != "undefined" && map.events[i].name == "Boss"){
        console.log("Found it");
        for (var j=0; j< map.events[i].pages[j].list.length; j++ ){
            for (var k=0; k< map.events[i].pages[j].list.length; k++ ){
                // Code 303 refers to the event item "Name Input Processing"
                if (map.events[i].pages[j].list[k].code == 303){
                    console.log("Changed it");
                    // Changing it to have 30 characters instead of RPGMAKEr imposed limit of 16. 
                    map.events[i].pages[j].list[k].parameters[1] = 30;
                }
            }
        }    
    }
}

var json = JSON.stringify(map);
fs.writeFile(outputpath, json, 'utf8', callback);

function callback(e){
    console.log("Done")
}