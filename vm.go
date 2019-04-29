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
	"sync"
	"time"

	"cloud.google.com/go/compute/metadata"
	"golang.org/x/oauth2/google"
	compute "google.golang.org/api/compute/v1"
)

func getService(c context.Context) (*compute.Service, error) {
	cTimed, _ := context.WithDeadline(c, time.Now().Add(60*time.Second))

	client, err := google.DefaultClient(cTimed, compute.ComputeScope)
	if err != nil {
		return nil, err
	}
	return compute.New(client)

}

func listAllInstances(c context.Context) (*compute.InstanceList, error) {
	vms := new(compute.InstanceList)

	zones, err := listZones(c)
	if err != nil {
		return vms, err
	}

	var wg sync.WaitGroup
	wg.Add(len(zones.Items))

	for _, zone := range zones.Items {

		go func(name string) {
			l, err := listInstances(c, name)
			if err != nil {
				return
			}

			for _, vm := range l.Items {
				vms.Items = append(vms.Items, vm)
			}
			wg.Done()
		}(zone.Name)

	}
	wg.Wait()

	return vms, nil
}

func listInstances(c context.Context, zone string) (*compute.InstanceList, error) {
	vms := new(compute.InstanceList)

	srv, err := getService(c)
	if err != nil {
		return vms, err
	}

	id, err := metadata.ProjectID()
	if err != nil {
		return vms, err
	}

	return srv.Instances.List(id, zone).Do()
}

func listZones(c context.Context) (*compute.ZoneList, error) {
	zones := new(compute.ZoneList)

	srv, err := getService(c)
	if err != nil {
		return zones, err
	}

	id, err := metadata.ProjectID()
	if err != nil {
		return zones, err
	}

	return srv.Zones.List(id).Do()
}
