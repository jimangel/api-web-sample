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
	fmt.Fprintf(w, "Welcome to the most basic demo using golang / firestore!\n\nTry: /get-data to read firestore via the contianer API\n\nTry: /post-random-data to add data to firestore via the contianer API\n(note: this causes the web server to send a GET request to the API server to generate a POST vs. create a raw POST")
}

func data(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Host:", os.Getenv("HOST_URL"))
	url := fmt.Sprintf("%s://%s/data", os.Getenv("API_HTTP_S"), os.Getenv("API_URL"))
	resp, err := http.Get(url)
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
	url := fmt.Sprintf("%s://%s:%s/post", os.Getenv("API_HTTP_S"), os.Getenv("API_URL"), os.Getenv("API_PORT"))
	resp, err := http.Get(url)
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
