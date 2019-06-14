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
	"fmt"

	"cloud.google.com/go/compute/metadata"
	"golang.org/x/oauth2/google"
	run "google.golang.org/api/run/v1alpha1"
)

func getRunService(c context.Context) (*run.APIService, error) {
	cTimed, cancel := context.WithDeadline(c, time.Now().Add(60*time.Second))
	defer cancel()

	client, err := google.DefaultClient(cTimed, run.CloudPlatformScope)
	if err != nil {
		return nil, err
	}
	return run.New(client)

}

func listAllRunServices(c context.Context) ([]*run.Service, error) {
	services := []*run.Service{}

	locations, err := listLocations(c)
	if err != nil {
		return services, err
	}

	for _, location := range locations {

		srvs, err := listRunLocationServices(c, location.LocationId)
		if err != nil {
			return services, err
		}

		for _, srv := range srvs {
			services = append(services, srv)
		}

	}

	return services, nil
}


func listRunLocationServices(c context.Context, locationID string) ([]*run.Service, error) {
	result := []*run.Service{}

	srv, err := getRunService(c)
	if err != nil {
		return result, err
	}

	id, err := metadata.ProjectID()
	if err != nil {
		return result, err
	}

	parent := fmt.Sprintf("projects/%s/locations/%s", id, locationID)

	services, err := srv.Projects.Locations.Services.List(parent).Do()
	if err != nil {
		return result, err
	}


	return services.Items, nil
}


func listLocations(c context.Context) ([]*run.Location, error) {
	result := []*run.Location{}

	srv, err := getRunService(c)
	if err != nil {
		return result, err
	}

	id, err := metadata.ProjectID()
	if err != nil {
		return result, err
	}

	parent := fmt.Sprintf("projects/%s", id)
	locations, err := srv.Projects.Locations.List(parent).Do()
	if err != nil {
		return result, err
	}

	return locations.Locations, nil
}