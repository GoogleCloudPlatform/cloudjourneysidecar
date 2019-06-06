// Script reads in the contents of Map001, finds the Boss's dialog and 
// makes sure that the field length for project name is 30 and not 16. 
console.log("Insert Google Analytics Code");
var fs = require('fs');
var path = require('path');
var outputpath = path.resolve(__dirname, '../../output/GCPQuest/www/index.html');
var rawdata = fs.readFileSync(outputpath, 'utf8');

var updated = rawdata.replace(/GA_MEASUREMENT_ID/g, process.env.CLOUD_JOURNEY_GACODE);

fs.writeFile(outputpath, updated, 'utf8', callback);

function callback(e){
    console.log("Done")
}

