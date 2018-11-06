package main

import (
	"context"
	"time"

	"cloud.google.com/go/compute/metadata"
	"golang.org/x/oauth2/google"
	bigquery "google.golang.org/api/bigquery/v2"
)

func getBigQueryService(c context.Context) (*bigquery.Service, error) {
	cTimed, _ := context.WithDeadline(c, time.Now().Add(60*time.Second))

	client, err := google.DefaultClient(cTimed, bigquery.BigqueryScope)
	if err != nil {
		return nil, err
	}
	return bigquery.New(client)

}

func ListJobs(c context.Context) (*bigquery.JobList, error) {
	jobs := new(bigquery.JobList)

	srv, err := getBigQueryService(c)
	if err != nil {
		return jobs, err
	}

	id, err := metadata.ProjectID()
	if err != nil {
		return jobs, err
	}

	call := srv.Jobs.List(id)
	call.AllUsers(true)
	call.Projection("full")

	return call.Do()
}
