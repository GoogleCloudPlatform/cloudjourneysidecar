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

	"cloud.google.com/go/compute/metadata"
	"google.golang.org/api/iterator"
	storage "cloud.google.com/go/storage"
)


func getCloudStorageClient(c context.Context) (storage.Client, error) {
	client := new(storage.Client)
	var err error

	client, err = storage.NewClient(c)
	if err != nil {
		return *client, err
	}
	return *client, nil

}

func listBuckets(c context.Context) ([]storage.BucketAttrs, error) {

	buckets := []storage.BucketAttrs{}

	client, err := getCloudStorageClient(c)
	if err != nil {
		return buckets, err
	}

	id, err := metadata.ProjectID()
	if err != nil {
		return buckets, err
	}

	it := client.Buckets(c, id)
	for {
        battrs, err := it.Next()
        if err == iterator.Done {
                break
        }
        if err != nil {
                return nil, err
        }
        buckets = append(buckets, *battrs)
}

	return buckets, nil
}
