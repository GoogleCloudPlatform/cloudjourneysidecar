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
	"time"

	"cloud.google.com/go/compute/metadata"
	"golang.org/x/oauth2/google"
	bigquery "google.golang.org/api/bigquery/v2"
)

func getBigQueryService(c context.Context) (*bigquery.Service, error) {
	cTimed, cancel := context.WithDeadline(c, time.Now().Add(60*time.Second))
	defer cancel()

	client, err := google.DefaultClient(cTimed, bigquery.BigqueryScope)
	if err != nil {
		return nil, err
	}
	return bigquery.New(client)

}

func listJobs(c context.Context) (*bigquery.JobList, error) {
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
