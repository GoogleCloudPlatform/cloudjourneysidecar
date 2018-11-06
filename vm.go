package main

import (
	"context"
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

func ListInstances(c context.Context) (*compute.InstanceList, error) {
	vms := new(compute.InstanceList)

	srv, err := getService(c)
	if err != nil {
		return vms, err
	}

	id, err := metadata.ProjectID()
	if err != nil {
		return vms, err
	}

	return srv.Instances.List(id, "us-central1-c").Do()
}
