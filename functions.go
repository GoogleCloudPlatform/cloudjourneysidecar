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
	cTimed, _ := context.WithDeadline(c, time.Now().Add(60*time.Second))

	client, err := google.DefaultClient(cTimed, cloudfunctions.CloudPlatformScope)
	if err != nil {
		return nil, err
	}
	return cloudfunctions.New(client)

}

func ListFunctions(c context.Context) (*cloudfunctions.ListFunctionsResponse, error) {
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
