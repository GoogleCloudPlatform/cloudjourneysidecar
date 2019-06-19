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
	// "fmt"
	"time"

	"cloud.google.com/go/compute/metadata"
	// "golang.org/x/oauth2/google"
	// "google.golang.org/api/option"
	firestore "cloud.google.com/go/firestore"
)

func getFirestoreClient(c context.Context, projectID string) (*firestore.Client, error) {
	cTimed, cancel := context.WithDeadline(c, time.Now().Add(60*time.Second))
	defer cancel()

	client, err := firestore.NewClient(cTimed, projectID)
	if err != nil {
		return nil, err
	}

	return client, nil

}


func listFirestoreCollection(c context.Context, collectionName string) ([]*firestore.DocumentSnapshot,error) {
	results := []*firestore.DocumentSnapshot{}


	projectID, err := metadata.ProjectID()
	if err != nil {
		return results, err
	}

	client, err := getFirestoreClient(c, projectID)
	if err != nil {
		return results, err
	}

	return client.Collection(collectionName).Documents(c).GetAll()
}