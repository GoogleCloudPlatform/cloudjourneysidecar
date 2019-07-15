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
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	storage "cloud.google.com/go/storage"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func main() {
	http.HandleFunc("/", healthHandler)
	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/version", versionHandler)
	http.HandleFunc("/random", randomHandler)

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

type check interface {
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	statii := StatusList{}



	checks := map[string]func(context.Context) (Status, error){
		"checkSys01": checkSys01,
		"checkSys02": checkSys02,
		"checkSys03": checkSys03,
		"checkSys04": checkSys04,
		"checkSys05": checkSys05,
		"checkData01": checkData01,
		"checkData02": checkData02,
		"checkData03": checkData03,
		"checkData04": checkData04,
		"checkData05": checkData05,
		"checkDev01": checkDev01,
		"checkDev02": checkDev02,
		"checkDev03": checkDev03,
		"checkDev04": checkDev04,
		"checkDev05": checkDev05,
	}

	quest, ok := r.URL.Query()["quest"]
    
    if !ok || len(quest) < 1 {
		for _, f := range checks {

			status, err := f(c)
			if err != nil {
				handleError(c, w, fmt.Errorf("could not check the status of %s: %v", status.Quest, err))
				return
			}
			statii = append(statii, status)
	
		}

    } else{
		status, err := checks[quest[0]](c)
		if err != nil {
			handleError(c, w, fmt.Errorf("could not check the status of %s: %v", status.Quest, err))
			return
		}
		statii = append(statii, status)
	}

	

	b, err := json.Marshal(statii)
	if err != nil {
		handleError(c, w, errors.New("Could not marshal the json: "+err.Error()))
		return
	}
	sendJSON(w, string(b))
}

func randomHandler(w http.ResponseWriter, r *http.Request) {
	c := r.Context()


	items, err := listFunctions(c)
	if err != nil {
		handleError(c, w, fmt.Errorf("could not check the status of %v: %v", items, err))
		return
	}

	b, err := json.Marshal(items)
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

func versionHandler(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	v, err := checkVersion(c)
	if err != nil {
		content := fmt.Sprintf("{\"version\" : %d, \"update\" : %t, \"notes\" : \"%s\"}", v, true, err)
		sendJSON(w, content)
		log.Warningf(r.Context(), "Error in version check %s.", err)
		return
	}

	content := fmt.Sprintf("{\"version\" : %d, \"update\" : %t, \"notes\" : \"%s\"}", v, false, "Version check working as expected")
	sendJSON(w, content)
	log.Infof(r.Context(), "Version check triggered.")
}

func checkSys01(c context.Context) (Status, error) {
	s := Status{}
	s.Quest = "01_sys"

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

func checkSys02(c context.Context) (Status, error) {
	s := Status{}
	s.Quest = "02_sys"

	disks, err := listAllDisks(c)
	if err != nil {
		return s, fmt.Errorf("tut_sys2: %v", err)
	}

	for _, v := range disks.Items {
		s.Notes = "Could not find an unimaged and attached disk"
		if len(v.Users) > 0 && v.SourceImage == "" {
			s.Complete = true
			s.Notes = ""
			break
		}

	}
	return s, nil
}

func checkSys03(c context.Context) (Status, error) {
	s := Status{}
	s.Quest = "03_sys"
	s.Complete = false
	s.Notes = "Could not find a custom image."

	items, err := listImages(c)
	if err != nil {
		return s, fmt.Errorf("tut_sys3: %v", err)
	}

	if len(items.Items) > 0 {
		s.Complete = true
		s.Notes = ""
	}

	return s, nil
}

func checkSys04(c context.Context) (Status, error) {
	s := Status{}
	s.Quest = "04_sys"
	s.Complete = false
	s.Notes = "Could not find an instance template."

	items, err := listTemplates(c)
	if err != nil {
		return s, fmt.Errorf("tut_sys4: %v", err)
	}

	if len(items.Items) > 0 {
		s.Complete = true
		s.Notes = ""
	}

	return s, nil
}

func checkSys05(c context.Context) (Status, error) {
	s := Status{}
	s.Quest = "05_sys"
	s.Complete = false
	s.Notes = "Could not find an instance group and corresponding backend service."

	groups, err := listAllInstanceGroups(c)
	if err != nil {
		return s, fmt.Errorf("tut_sys5: %v", err)
	}

	backends, err := listBackendServices(c)
	if err != nil {
		return s, fmt.Errorf("tut_sys5: %v", err)
	}

	for _, group := range groups {

		for _, item := range backends.Items {
			for _, backend := range item.Backends {
				if backend.Group == group.SelfLink {
					s.Complete = true
					s.Notes = ""
					return s, nil
				}
			}
		}
	}

	return s, nil
}

func checkData01(c context.Context) (Status, error) {
	s := Status{}
	s.Quest = "01_data"

	jobs, err := listJobs(c)
	if err != nil {
		return s, fmt.Errorf("tut_dsc1: %v", err)
	}

	for _, v := range jobs.Jobs {
		if strings.Index(v.Configuration.Query.Query, "bigquery-public-data.usa_names.usa_1910_2013") > -1 {
			t := time.Unix(v.Statistics.EndTime/1000, 0)
			s.Complete = true
			s.Notes = fmt.Sprintf("Query run at: %s", t.Format("Jan 2 2006 at 15:04:05"))

		}
	}
	return s, nil
}

func checkData02(c context.Context) (Status, error) {
	s := Status{}
	s.Quest = "02_data"

	buckets, err := listBuckets(c)
	if err != nil {
		return s, fmt.Errorf("tut_dat2: %v", err)
	}

	bucks := []storage.BucketAttrs{}

	for _, v := range buckets {
		if strings.Index(v.Name, "appspot.com") >= 0 {
			continue
		}
		if strings.Index(v.Name, "_cloudbuild") >= 0 {
			continue
		}
		bucks = append(bucks, v)
	}
	s.Complete = false
	s.Notes = "No buckets found"

	if len(bucks) > 0 {
		s.Complete = true
		s.Notes = ""
	}

	return s, nil
}

func checkData03(c context.Context) (Status, error) {
	s := Status{}
	s.Quest = "03_data"
	s.Notes = "Did not find a Cloud SQL server."

	items, err := listSQLServers(c)
	if err != nil {
		return s, fmt.Errorf("tut_dsc3: %v", err)
	}

	for _, v := range items.Items {
		if v.DatabaseVersion == "MYSQL_5_7" && v.InstanceType == "CLOUD_SQL_INSTANCE" {
			s.Complete = true
			s.Notes = fmt.Sprintf("found %s", v.Name)
		}
	}
	return s, nil
}

func checkData04(c context.Context) (Status, error) {
	s := Status{}
	s.Quest = "04_data"
	s.Notes = "Did not find any Firestore records."

	items, err := listFirestoreCollection(c, "colors")
	if err != nil {
		return s, fmt.Errorf("tut_dsc4: %v", err)
	}

	if len(items) > 0 {
		s.Complete = true
		s.Notes = fmt.Sprintf("found references")
	}

	return s, nil
}

func checkData05(c context.Context) (Status, error) {
	s := Status{}
	s.Quest = "05_data"
	s.Notes = "Did not find any Memorystore instances."

	items, err := listAllMemoryStoreInstances(c)
	if err != nil {
		return s, fmt.Errorf("tut_dsc5: %v", err)
	}

	if len(items) > 0 {
		s.Complete = true
		s.Notes = fmt.Sprintf("found %s", items[0].Name)
	}

	return s, nil
}

func checkDev01(c context.Context) (Status, error) {
	s := Status{}
	s.Quest = "01_dev"

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

func checkDev02(c context.Context) (Status, error) {
	s := Status{}
	s.Quest = "02_dev"

	builds, err := listBuildOperations(c)
	if err != nil {
		return s, fmt.Errorf("tut_dev2: %v", err)
	}

	for _, v := range builds {
		s.Notes = "Could not find 'randomcolor' container image"
		//&& len(v.Artifacts.Images) > 0 && strings.Index(v.Artifacts.Images[0], "randomcolor" ) >= 0
		if v.Artifacts != nil && len(v.Artifacts.Images) > 0 {

			for _, image := range v.Artifacts.Images {
				if strings.Index(image, "randomcolor") >= 0 {
					s.Complete = true
					s.Notes = ""
					break
				}
			}
			if s.Complete {
				break
			}
		}

	}
	return s, nil
}

func checkDev03(c context.Context) (Status, error) {
	s := Status{}
	s.Quest = "03_dev"

	svc, err := listAllRunServices(c)
	if err != nil {
		return s, fmt.Errorf("tut_dev3: %v", err)
	}

	for _, v := range svc {
		s.Notes = "Could not find 'randomcolor' service"
		if v.Status != nil {
			if strings.Index(v.Status.LatestCreatedRevisionName, "randomcolor") >= 0 {
				s.Complete = true
				s.Notes = ""
				break
			}
		}

	}
	return s, nil
}

func checkDev04(c context.Context) (Status, error) {
	s := Status{}
	s.Quest = "04_dev"

	funcs, err := listFunctions(c)
	if err != nil {
		if strings.Index(err.Error(), "Cloud Functions API has not been used") > -1 {
			s.Notes = "API not enabled yet."
			return s, nil
		}
		return s, fmt.Errorf("tut_dev4: %v", err)
	}

	for _, v := range funcs.Functions {
		if v.EventTrigger != nil && v.EventTrigger.EventType == "google.storage.object.finalize" {
			s.Complete = true
			s.Notes = fmt.Sprintf("Found function %s", v.Name)
			break
		}
	}
	return s, nil
}

func checkDev05(c context.Context) (Status, error) {
	s := Status{}
	s.Quest = "05_dev"
	objectName := "translation/test.txt"

	funcs, err := listFunctions(c)
	if err != nil {
		if strings.Index(err.Error(), "Cloud Functions API has not been used") > -1 {
			s.Notes = "API not enabled yet."
			return s, nil
		}
		return s, fmt.Errorf("tut_dev5: %v", err)
	}

	for _, v := range funcs.Functions {
		if v.EventTrigger != nil {
			if v.EventTrigger != nil && v.EventTrigger.EventType == "google.storage.object.finalize" {
				bucket := listLast(v.EventTrigger.Resource, "/")

				exists, err := objectExists(c, bucket, objectName)
				if err != nil {
					if strings.Index(err.Error(), storage.ErrObjectNotExist.Error()) > -1 {
						continue
					}
					return s, fmt.Errorf("could not verify bucket and file %v/%s err: %v", bucket, objectName, err)
				}

				s.Complete = exists
				if exists {
					s.Notes = fmt.Sprintf("Found function %s found bucket %s, found object %s", v.Name, bucket, objectName)
					break
				}

			}
		}
	}
	return s, nil
}

func listLast(string, delimiter string) string {
	ss := strings.Split(string, "/")
	return ss[len(ss)-1]
}

func checkVersion(c context.Context) (int, error) {
	dat, err := ioutil.ReadFile(".version")
	if err != nil {
		return 0, err
	}
	i, err := strconv.Atoi(string(dat))
	if err != nil {
		return 0, err
	}
	return i, nil

}
