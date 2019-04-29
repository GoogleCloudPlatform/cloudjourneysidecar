// Copyright 2019 Google LLC All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/appengine/log"
)

func sendJSON(w http.ResponseWriter, content string) {
	w.Header().Set("Content-Type", "application/json;  charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if content == "null" || content == "[]" {
		w.WriteHeader(404)
		content = "{ \"error\" : \"Not Found\" }"
	}

	fmt.Fprint(w, content)
}

func sendMessage(w http.ResponseWriter, msg string) {
	content := "{ \"msg\" : \"" + msg + "\" }"
	sendJSON(w, content)
}

func handleError(c context.Context, w http.ResponseWriter, err error) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json;  charset=UTF-8")
	w.WriteHeader(500)
	content := "{ \"error\" : \"" + err.Error() + "\" }"
	sendJSON(w, content)

	log.Errorf(c, err.Error())
}
