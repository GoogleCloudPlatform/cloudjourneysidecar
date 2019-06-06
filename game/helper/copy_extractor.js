// Copy will read in all the files with copy, and extract it for review. 
console.log("Copy Extractor");
var fs = require('fs');
var path = require('path');
var outputpath = path.resolve(__dirname, 'test.html');
var files = [];
files[0] = path.resolve(__dirname, '../../output/GCPQuest/www/data/Map001.json');
files[1] = path.resolve(__dirname, '../../output/GCPQuest/www/data/Map002.json');
files[2] = path.resolve(__dirname, '../../output/GCPQuest/www/data/Map003.json');

var rawdata = fs.readFileSync(files[0], 'utf8');
let map = JSON.parse(rawdata);  

var text = "";
for (var i=0; i < files.length; i++){
    console.log("Map: " + i);
    text += "<h1>Map: " + i + "</h1>" + "\n"; 
    text += parseMapFile(files[i]);
}


fs.writeFile(outputpath, text, 'utf8', callback);

function callback(e){
    console.log("Done")
}


function parseMapFile(file){
    var rawdata = fs.readFileSync(file, 'utf8');
    let map = JSON.parse(rawdata);  
    var text= "";


    for (var i=0; i< map.events.length; i++ ){
        
        var event = map.events[i];
        if (event != null && typeof event.name != "undefined"){
            console.log("event:", event.name);
            text += "<h2>" + event.name+ "</h2>" + "\n"; 
            for (var j=0; j< event.pages.length; j++ ){
                var page = event.pages[j];
                // console.log("page:", j);
                text += "<h3>Page: " +  j + "</h3>" + "\n"; 
                for (var k=0; k< page.list.length; k++ ){
                    var item = page.list[k];
                    if (item.code == 401){

                        
                        text += "<p>" +  item.parameters[0].replace(">", "</strong>: ").replace("<", "<strong>"); + "</p>" + "\n"; 
                        // console.log("copy:", item.parameters[0]); 
                    }

                }
            }
        }
    }
    return text;
}

