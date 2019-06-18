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
	"fmt"

	"cloud.google.com/go/compute/metadata"
	"golang.org/x/oauth2/google"
	compute "google.golang.org/api/compute/v1"
)

func getComputeEngineService(c context.Context) (*compute.Service, error) {
	cTimed, cancel := context.WithDeadline(c, time.Now().Add(60*time.Second))
	defer cancel()

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

func listAllDisks(c context.Context) (*compute.DiskList, error) {
	disks := new(compute.DiskList)

	zones, err := listZones(c)
	if err != nil {
		return disks, err
	}

	var wg sync.WaitGroup
	wg.Add(len(zones.Items))

	for _, zone := range zones.Items {

		go func(name string) {
			l, err := listDisks(c, name)
			if err != nil {
				return
			}

			for _, disk := range l.Items {
				disks.Items = append(disks.Items, disk)
			}
			wg.Done()
		}(zone.Name)

	}
	wg.Wait()

	return disks, nil
}

func listInstances(c context.Context, zone string) (*compute.InstanceList, error) {
	vms := new(compute.InstanceList)

	srv, err := getComputeEngineService(c)
	if err != nil {
		return vms, err
	}

	projectID, err := metadata.ProjectID()
	if err != nil {
		return vms, err
	}

	return srv.Instances.List(projectID, zone).Do()
}

func listZones(c context.Context) (*compute.ZoneList, error) {
	zones := new(compute.ZoneList)

	srv, err := getComputeEngineService(c)
	if err != nil {
		return zones, err
	}

	projectID, err := metadata.ProjectID()
	if err != nil {
		return zones, err
	}

	return srv.Zones.List(projectID).Do()
}

func listRegions(c context.Context) (*compute.RegionList, error) {
	regions := new(compute.RegionList)

	srv, err := getComputeEngineService(c)
	if err != nil {
		return regions, err
	}

	projectID, err := metadata.ProjectID()
	if err != nil {
		return regions, err
	}

	return srv.Regions.List(projectID).Do()
}

func listDisks(c context.Context, zone string) (*compute.DiskList, error) {
	disks := new(compute.DiskList)

	srv, err := getComputeEngineService(c)
	if err != nil {
		return disks, err
	}

	projectID, err := metadata.ProjectID()
	if err != nil {
		return disks, err
	}

	return srv.Disks.List(projectID, zone).Do()
}


func listImages(c context.Context) (*compute.ImageList, error) {
	disks := new(compute.ImageList)

	srv, err := getComputeEngineService(c)
	if err != nil {
		return disks, err
	}

	projectID, err := metadata.ProjectID()
	if err != nil {
		return disks, err
	}

	return srv.Images.List(projectID).Do()
}

func listTemplates(c context.Context) (*compute.InstanceTemplateList, error) {
	items := new(compute.InstanceTemplateList)

	srv, err := getComputeEngineService(c)
	if err != nil {
		return items, err
	}

	projectID, err := metadata.ProjectID()
	if err != nil {
		return items, err
	}

	return srv.InstanceTemplates.List(projectID).Do()
}



func listAllInstanceGroups(c context.Context) ([]*compute.InstanceGroup, error) {
	items := []*compute.InstanceGroup{}

	zones, err := listZones(c)
	if err != nil {
		return items, err
	}

	var wg sync.WaitGroup
	wg.Add(len(zones.Items))


	for _, zone := range zones.Items {

		go func(name string) {

			l, err := listInstanceGroups(c, name)
			if err != nil {
				return 
			}

			for _, g := range l.Items {
				items = append(items, g)
			}
			wg.Done()
		}(zone.Name)

	}
	wg.Wait()

	return items, nil
}


func listInstanceGroups(c context.Context, zone string) (*compute.InstanceGroupList, error) {
	items := new(compute.InstanceGroupList)

	srv, err := getComputeEngineService(c)
	if err != nil {
		return items, err
	}

	projectID, err := metadata.ProjectID()
	if err != nil {
		return items, err
	}

	return srv.InstanceGroups.List(projectID, zone).Do()
}





func listAllInstanceGroupManagers(c context.Context) ([]*compute.InstanceGroupManager, error) {
	items := []*compute.InstanceGroupManager{}

	zones, err := listZones(c)
	if err != nil {
		return items, err
	}

	var wg sync.WaitGroup
	wg.Add(len(zones.Items))


	for _, zone := range zones.Items {

		go func(name string) {

			l, err := listInstanceGroupManagers(c, name)
			if err != nil {
				return 
			}

			for _, g := range l.Items {
				items = append(items, g)
			}
			wg.Done()
		}(zone.Name)

	}
	wg.Wait()

	return items, nil
}


func listInstanceGroupManagers(c context.Context, zone string) (*compute.InstanceGroupManagerList, error) {
	items := new(compute.InstanceGroupManagerList)

	srv, err := getComputeEngineService(c)
	if err != nil {
		return items, err
	}

	projectID, err := metadata.ProjectID()
	if err != nil {
		return items, err
	}

	return srv.InstanceGroupManagers.List(projectID, zone).Do()
}



func listBackendServices(c context.Context) (*compute.BackendServiceList, error) {
	items := new(compute.BackendServiceList)

	srv, err := getComputeEngineService(c)
	if err != nil {
		return items, fmt.Errorf("getComputeEngineService: %s", err)
	}

	projectID, err := metadata.ProjectID()
	if err != nil {
		return items, fmt.Errorf("getMetaDataProjecID: %s", err)
	}


	return srv.BackendServices.List(projectID).Do()
}	