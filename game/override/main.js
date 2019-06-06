//=============================================================================
// main.js
//=============================================================================

PluginManager.setup($plugins);

window.onload = function() {
    SceneManager.run(Scene_Boot);
    Graphics._switchStretchMode();
    
};

var sidecarversion = "REPLACEVERSION";
var links = {};
buildLinks();

function buildLinks(){
  var wt = "https://console.cloud.google.com?walkthrough_tutorial_id";
  links.tut_dev1 = wt + "=cloudjourney_intro_cf"; 
  links.tut_sys1 = wt + "=cloudjourney_intro_ce"; 
  links.tut_dsc1 = wt + "=cloudjourney_intro_bq"; 
  links.tut_pro1 = wt + "=cloudjourney_intro_project"; 
  links.tut_sys2 = wt + "=cloudjourney_02_sys"; 
}

function setRemoteStatus(label, id){
    var project = $gameActors.actor(5).name().trim();
    var xhr = new XMLHttpRequest();
    xhr.open("GET", "https://"+project+".appspot.com/status", false);
    xhr.send();

    var results = JSON.parse(xhr.responseText);

    for (var i = 0; i < results.length; i++){
        var r = results[i];
        if (r.quest == label && r.complete){
            setProject(id, r.complete);
        }
    }
    return results;
}

function setProject(id, value){
    $gameSwitches.setValue(id, value);
}



function checkHealth(){
    var project = "NOTAVALIDPROJECTNAME";
    if (typeof $gameActors.actor(5).name() != 'undefined'){
      project = $gameActors.actor(5).name().trim();
    }
    var response = {};
    response.works = false; 

    var xhr = new XMLHttpRequest();
    try{
      xhr.open("GET", "https://"+project+".appspot.com/health", false);
      xhr.send();
    } catch(exception) {
        console.log('Bad project.');
        return response;
    }
   

    

    if (xhr.status == 404){
        return response;
    }

    var results = JSON.parse(xhr.responseText);

    if (results.msg == "ok"){
        response.works = true;
    }


    return response; 
}


function checkVersion(){
  var project = "NOTAVALIDPROJECTNAME";
  if (typeof $gameActors.actor(5).name() != 'undefined'){
    project = $gameActors.actor(5).name().trim();
  }

  
  var xhr = new XMLHttpRequest();
  try{
    var endpoint = "https://"+project+".appspot.com/version";
    console.log('Endpoint:', endpoint);
    xhr.open("GET", endpoint, false);
    xhr.send();
  } catch(exception) {
      console.log(exception);
      console.log('Bad project');
      return false;
  }

  if (xhr.status == 404){
      console.log('Version isn\'t set, update.');
      return false;
  }

  var results = JSON.parse(xhr.responseText);

  if (results.update){
      console.log('Backend says update.');
      return false;
  }

  if (results.version == sidecarversion){
    console.log('Version set and matches, no update.');
    return true;
  }

  console.log('For some reason, update.');
  return false; 
}

function popUp(url){
    var div = document.createElement("div");
    var ok = document.createElement("button");
    var p = document.createElement("p");

    ok.innerHTML = "Ok";
    ok.id="okbtn";
    ok.style.backgroundColor=  "#0000FF"; /* Blue */
    ok.style.border=  "none";
    ok.style.color=  "white";
    ok.style.padding=  "10px 32px";
    ok.style.textAlign=  "center";
    ok.style.textDecoration=  "none";
    ok.style.display=  "inline-block";
    ok.style.fontSize=  "16px";

    div.id = "alert_box";
    div.style.top = "10%";
    div.style.left = "40%";
    div.style.height = "100px";
    div.style.width = "300px";
    div.style.position = "absolute";
    div.style.zIndex = 10;
    div.style.backgroundColor = "#dddddd";
    div.style.padding = "5px 10px";
    div.style.border = "1px solid #999";
    div.style.fontFamily = "'Google Sans', Roboto, Helvetica, Arial, sans-serif";

    
    p.innerHTML = "Click Ok to launch Cloud Console";
   
    div.appendChild(p);
    div.appendChild(ok);
    document.getElementsByTagName("body")[0].appendChild(div);
    ok.addEventListener("click", function(){
                                        window.open(url,'_blank');
                                        document.querySelector("#alert_box").remove();
                                    } );
}

function populateProject(){
    var projectid = getAllUrlParams().projectid;

    if (typeof projectid !== "undefined"){
      $gameActors.actor(5).setName(projectid);
    
    }
    
    console.log(projectid);
    
}

function getAllUrlParams(url) {

    // get query string from url (optional) or window
    var queryString = url ? url.split('?')[1] : window.location.search.slice(1);
  
    // we'll store the parameters here
    var obj = {};
  
    // if query string exists
    if (queryString) {
  
      // stuff after # is not part of query string, so get rid of it
      queryString = queryString.split('#')[0];
  
      // split our query string into its component parts
      var arr = queryString.split('&');
  
      for (var i = 0; i < arr.length; i++) {
        // separate the keys and the values
        var a = arr[i].split('=');
  
        // set parameter name and value (use 'true' if empty)
        var paramName = a[0];
        var paramValue = typeof (a[1]) === 'undefined' ? true : a[1];
  
        // (optional) keep case consistent
        paramName = paramName.toLowerCase();
        if (typeof paramValue === 'string') paramValue = paramValue.toLowerCase();
  
        // if the paramName ends with square brackets, e.g. colors[] or colors[2]
        if (paramName.match(/\[(\d+)?\]$/)) {
  
          // create key if it doesn't exist
          var key = paramName.replace(/\[(\d+)?\]/, '');
          if (!obj[key]) obj[key] = [];
  
          // if it's an indexed array e.g. colors[2]
          if (paramName.match(/\[\d+\]$/)) {
            // get the index value and add the entry at the appropriate position
            var index = /\[(\d+)\]/.exec(paramName)[1];
            obj[key][index] = paramValue;
          } else {
            // otherwise add the value to the end of the array
            obj[key].push(paramValue);
          }
        } else {
          // we're dealing with a string
          if (!obj[paramName]) {
            // if it doesn't exist, create property
            obj[paramName] = paramValue;
          } else if (obj[paramName] && typeof obj[paramName] === 'string'){
            // if property does exist and it's a string, convert it to an array
            obj[paramName] = [obj[paramName]];
            obj[paramName].push(paramValue);
          } else {
            // otherwise add the property
            obj[paramName].push(paramValue);
          }
        }
      }
    }
  
    return obj;
  }