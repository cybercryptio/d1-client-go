// Copyright 2022 CYBERCRYPT
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client_test

import (
	"context"
	"fmt"
	"log"

	gclient "github.com/cybercryptio/d1-client-go/d1-generic"
	client "github.com/cybercryptio/d1-client-go/d1-storage"
	pbstorage "github.com/cybercryptio/d1-client-go/d1-storage/protobuf/storage"
)

func Example_withPerRPCCredentials() {
	// Create a new D1 Storage client providing the hostname, and optionally, the client connection level and per RPC credentials.
	client, err := client.NewStorageClient(endpoint,
		gclient.WithTransportCredentials(creds),
		gclient.WithPerRPCCredentials(
			gclient.NewStandalonePerRPCToken(endpoint, uid, password, creds),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// Store sensitive data in encrypted form.
	storeResponse, err := client.Storage.Store(
		ctx,
		&pbstorage.StoreRequest{
			Plaintext:      []byte("secret data"),
			AssociatedData: []byte("metadata"),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve the stored data.
	retrieveResponse, err := client.Storage.Retrieve(
		ctx,
		&pbstorage.RetrieveRequest{
			ObjectId: storeResponse.ObjectId,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("plaintext:%q associated_data:%q",
		retrieveResponse.Plaintext,
		retrieveResponse.AssociatedData,
	)
	// Output: plaintext:"secret data" associated_data:"metadata"
}
