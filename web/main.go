package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Welcome to the most basic demo using golang / firestore!\n\nTry: /get-data to read firestore via the container API\n\nTry: /post-random-data to add data to firestore via the contianer API\n(note: this causes the web server to send a GET request to the API server to generate a POST vs. create a raw POST")
}

func getUrl(path string) string {
	// export NO_PORT="true"
	// This removes the port from the web-app calling the api. Very useful if api is behind another LB.
	key, exists := os.LookupEnv("NO_PORT")
	if exists {
		fmt.Print(key)
		url := fmt.Sprintf("%s://%s/%s", os.Getenv("API_HTTP_S"), os.Getenv("API_URL"), path)
		return url
	} else {
		fmt.Print(key)
		url := fmt.Sprintf("%s://%s:%s/%s", os.Getenv("API_HTTP_S"), os.Getenv("API_URL"), os.Getenv("API_PORT"), path)
		return url
	}
}

// 		url := fmt.Sprintf("%s://%s/data", os.Getenv("API_HTTP_S"), os.Getenv("API_URL"))
func data(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Host:", os.Getenv("HOST_URL"))
	full_url := getUrl("data")
	fmt.Print(full_url)
	resp, err := http.Get(full_url)
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	fmt.Fprintf(w, sb)
}

func new(w http.ResponseWriter, r *http.Request) {

	full_url := getUrl("post")
	fmt.Print(full_url)
	resp, err := http.Get(full_url)
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	fmt.Fprintf(w, sb)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/get-data", data)
	http.HandleFunc("/post-random-data", new)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("WEB_PORT")), nil))
}

func main() {
	handleRequests()
}
