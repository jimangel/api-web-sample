// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Sample firestore_quickstart demonstrates how to connect to Firestore, and add and list documents.
package main

// [START firestore_setup_client_create]
import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

// [END firestore_setup_client_create]

func PrettyString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Welcome to the API! I READ data at /data & WRITE (generated) data at /post")
}

func data(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Fatalf("%v", err)
	}

	// [START firestore_setup_dataset_read]
	iter := client.Collection("users").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		jsonStr, err := json.Marshal(doc.Data())
		fmt.Fprintf(w, string(jsonStr))
	}
}

func post(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Fatalf("%v", err)
	}

	_, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
		"timestamp":   time.Now(),
		"static-data": "Jim",
		"random":      1912,
	})
	if err != nil {
		log.Fatalf("Failed adding random data: %v", err)
	}

	fmt.Fprintf(w, "Added a new entry to the collection with today's timestamp %s", time.Now())
}

func main() {

	http.HandleFunc("/", homePage)
	http.HandleFunc("/data", data)
	http.HandleFunc("/post", post)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("API_PORT")), nil))
}
