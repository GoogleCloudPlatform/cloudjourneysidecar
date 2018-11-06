package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/test", testHandler)
	appengine.Main()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, World!")
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	jobs, err := ListFunctions(c)
	if err != nil {
		handleError(c, w, errors.New("Could not retrieve a list of functions in the project: "+err.Error()))
		return
	}

	b, err := json.Marshal(jobs)
	if err != nil {
		handleError(c, w, errors.New("Could not marshal the json: "+err.Error()))
		return
	}
	sendJSON(w, string(b))
}

type Status struct {
	Quest    string `json:"quest"`
	Complete bool   `json:"complete"`
	Notes    string `json:"notes"`
}

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

func checkIntroSys(c context.Context) (Status, error) {
	s := Status{}
	s.Quest = "intro_sys"

	vms, err := ListInstances(c)
	if err != nil {
		return s, fmt.Errorf("SQ_intro: %v", err)
	}

	for _, v := range vms.Items {
		if strings.Index(v.MachineType, "custom-8-43008") > -1 {
			s.Complete = true
			break
		}
	}
	return s, nil
}

func checkIntroBigData(c context.Context) (Status, error) {
	s := Status{}
	s.Quest = "intro_bigdata"

	jobs, err := ListJobs(c)
	if err != nil {
		return s, fmt.Errorf("SQ_intro: %v", err)
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

	funcs, err := ListFunctions(c)
	if err != nil {
		if strings.Index(err.Error(), "Cloud Functions API has not been used") > -1 {
			s.Notes = "API not enabled yet."
			return s, nil
		}
		return s, fmt.Errorf("SQ_intro: %v", err)
	}

	for _, v := range funcs.Functions {
		if strings.Index(v.Name, "tokenGenerator") > -1 {
			s.Complete = true
			break
		}
	}
	return s, nil
}
