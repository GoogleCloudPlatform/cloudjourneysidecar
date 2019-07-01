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
	redis "cloud.google.com/go/redis/apiv1"
	"google.golang.org/api/iterator"
	redispb "google.golang.org/genproto/googleapis/cloud/redis/v1"
	"google.golang.org/api/option"
)

func getCloudMemoryStoreClient(c context.Context) (redis.CloudRedisClient, error) {
	cTimed, cancel := context.WithDeadline(c, time.Now().Add(60*time.Second))
	defer cancel()

	client := new(redis.CloudRedisClient)
	var err error

	client, err = redis.NewCloudRedisClient(cTimed, option.WithScopes(redis.DefaultAuthScopes()[0]))
	if err != nil {
		return *client, err
	}
	return *client, nil

}

func listAllMemoryStoreInstances(c context.Context) ([]*redispb.Instance, error) {
	items := []*redispb.Instance{}

	locations, err := listLocations(c)
	if err != nil {
		return items, err
	}

	for _, location := range locations {

		instances, err := listMemoryStoreInstances(c, location.LocationId)
		if err != nil {
			return items, err
		}

		for _, inst := range instances {
			items = append(items, inst)
		}

	}

	return items, nil
}

func listMemoryStoreInstances(c context.Context, locationID string) ([]*redispb.Instance, error) {
	items := []*redispb.Instance{}

	client, err := getCloudMemoryStoreClient(c)
	if err != nil {
		return items, err
	}

	id, err := metadata.ProjectID()
	if err != nil {
		return items, err
	}


	parent := fmt.Sprintf("projects/%s/locations/%s", id, locationID)

	req := &redispb.ListInstancesRequest{
		Parent: parent,
	}
	it := client.ListInstances(c, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return items, err
		}
		items = append(items, resp)
	}


	return items, nil
}	

