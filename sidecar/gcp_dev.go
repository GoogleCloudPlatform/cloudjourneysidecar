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
	"time"

	"cloud.google.com/go/compute/metadata"
	"golang.org/x/oauth2/google"
	cloudfunctions "google.golang.org/api/cloudfunctions/v1"
)


func getCloudFunctionsService(c context.Context) (*cloudfunctions.Service, error) {
	cTimed, cancel := context.WithDeadline(c, time.Now().Add(60*time.Second))
	defer cancel()

	client, err := google.DefaultClient(cTimed, cloudfunctions.CloudPlatformScope)
	if err != nil {
		return nil, err
	}
	return cloudfunctions.New(client)

}

func listFunctions(c context.Context) (*cloudfunctions.ListFunctionsResponse, error) {
	funcs := new(cloudfunctions.ListFunctionsResponse)

	srv, err := getCloudFunctionsService(c)
	if err != nil {
		return funcs, err
	}

	id, err := metadata.ProjectID()
	if err != nil {
		return funcs, err
	}

	call := srv.Projects.Locations.Functions.List(fmt.Sprintf("projects/%s/locations/-", id))

	return call.Do()
}
