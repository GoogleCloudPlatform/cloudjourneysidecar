// Script reads in the contents of Map001, finds the Boss's dialog and 
// makes sure that the field length for project name is 30 and not 16. 
console.log("Insert Sidecar Version");
var fs = require('fs');
var path = require('path');

var inputpath = path.resolve(__dirname, '../.version');
var version = fs.readFileSync(inputpath, 'utf8');

var outputpath = path.resolve(__dirname, '../../output/GCPQuest/www/js/main.js');
var rawdata = fs.readFileSync(outputpath, 'utf8');

var updated = rawdata.replace(/\"REPLACEVERSION\"/g, version);

fs.writeFile(outputpath, updated, 'utf8', callback);

function callback(e){
    console.log("Done")
}

