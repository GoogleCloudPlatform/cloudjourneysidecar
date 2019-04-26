// Copyright 2019 Google Inc. All Rights Reserved.
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
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"google.golang.org/appengine/log"

	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/", healthHandler)
	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/health", healthHandler)

	appengine.Main()
}

// Status is a quest name and whether or not it has been completed. Additionally
// notes can be provided for more context.
type Status struct {
	Quest    string `json:"quest"`
	Complete bool   `json:"complete"`
	Notes    string `json:"notes"`
}

// StatusList is a slice of Statii.
type StatusList []Status

func statusHandler(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	statii := StatusList{}

	sqIntro, err := checkIntroSys(c)
	if err != nil {
		handleError(c, w, fmt.Errorf("could not check the status of %s: %v", sqIntro.Quest, err))
		return
	}
	statii = append(statii, sqIntro)

	bqIntro, err := checkIntroBigData(c)
	if err != nil {
		handleError(c, w, fmt.Errorf("could not check the status of %s: %v", bqIntro.Quest, err))
		return
	}
	statii = append(statii, bqIntro)

	devIntro, err := checkIntroDev(c)
	if err != nil {
		handleError(c, w, fmt.Errorf("could not check the status of %s: %v", devIntro.Quest, err))
		return
	}
	statii = append(statii, devIntro)

	b, err := json.Marshal(statii)
	if err != nil {
		handleError(c, w, errors.New("Could not marshal the json: "+err.Error()))
		return
	}
	sendJSON(w, string(b))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	sendMessage(w, "ok")
	log.Infof(r.Context(), "Health Check triggered.")
}

func checkIntroSys(c context.Context) (Status, error) {
	s := Status{}
	s.Quest = "intro_sys"

	vms, err := listAllInstances(c)
	if err != nil {
		if strings.Index(err.Error(), "Compute Engine API has not been used") > -1 {
			s.Notes = "API not enabled yet."
			return s, nil
		}
		return s, fmt.Errorf("tut_sys1: %v", err)
	}

	for _, v := range vms.Items {
		s.Notes = "VM is not of type f1-micro"
		if strings.Index(v.MachineType, "f1-micro") > -1 {
			s.Complete = true
			s.Notes = ""
			break
		}

	}
	return s, nil
}

func checkIntroBigData(c context.Context) (Status, error) {
	s := Status{}
	s.Quest = "intro_bigdata"

	jobs, err := listJobs(c)
	if err != nil {
		return s, fmt.Errorf("tut_dsc1: %v", err)
	}

	for _, v := range jobs.Jobs {
		if strings.Index(v.Configuration.Query.Query, "bigquery-public-data.usa_names.usa_1910_2013") > -1 {
			s.Complete = true
			break
		}
	}
	return s, nil
}

func checkIntroDev(c context.Context) (Status, error) {
	s := Status{}
	s.Quest = "intro_dev"

	funcs, err := listFunctions(c)
	if err != nil {
		if strings.Index(err.Error(), "Cloud Functions API has not been used") > -1 {
			s.Notes = "API not enabled yet."
			return s, nil
		}
		return s, fmt.Errorf("tut_dev1: %v", err)
	}

	for _, v := range funcs.Functions {
		if strings.Index(v.Name, "tokenGenerator") > -1 {
			s.Complete = true
			break
		}
	}
	return s, nil
}
