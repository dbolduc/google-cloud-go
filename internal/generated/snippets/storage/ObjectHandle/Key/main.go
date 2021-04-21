// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// [START storage_generated_storage_ObjectHandle_Key]

package main

import (
	"context"

	"cloud.google.com/go/storage"
)

var secretKey []byte

func main() {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: handle error.
	}
	obj := client.Bucket("my-bucket").Object("my-object")
	// Encrypt the object's contents.
	w := obj.Key(secretKey).NewWriter(ctx)
	if _, err := w.Write([]byte("top secret")); err != nil {
		// TODO: handle error.
	}
	if err := w.Close(); err != nil {
		// TODO: handle error.
	}
}

// [END storage_generated_storage_ObjectHandle_Key]