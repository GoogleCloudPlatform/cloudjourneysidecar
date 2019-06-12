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
	build "google.golang.org/api/cloudbuild/v1"
)

func getBuildService(c context.Context) (*build.Service, error) {
	cTimed, cancel := context.WithDeadline(c, time.Now().Add(60*time.Second))
	defer cancel()

	client, err := google.DefaultClient(cTimed, build.CloudPlatformScope)
	if err != nil {
		return nil, err
	}
	return build.New(client)

}


func listBuildOperations(c context.Context) ([]build.Build, error) {
	result := []build.Build{}

	srv, err := getBuildService(c)
	if err != nil {
		return result, err
	}

	id, err := metadata.ProjectID()
	if err != nil {
		return result, err
	}


	build, err := srv.Projects.Builds.List(id).Do()
	if err != nil {
		return result, err
	}

	for _, v := range build.Builds {
		result = append(result, *v)
	}

	return result, nil
}
