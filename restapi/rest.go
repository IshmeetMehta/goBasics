package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json: "Title"`
	Desc    string `json: "desc"`
	Content string `json: "content"`
}

//let's declare a global Articles array
//that we can populate in our main function to simulate database

var Articles []Article

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to home page")
	fmt.Println("Endpoint Hit: Homepage")
}

func handleRequests() {

	http.HandleFunc("/", homepage)
	//add our artcles route and map it to our
	//returnAllArticles function like so
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnArcticles")
	json.NewEncoder(w).Encode(Articles)
}

func main() {

	Articles = []Article{
		Article{Title: "Hello", Desc: "Desc 1", Content: "Content1"},
		Article{Title: "Hello", Desc: "Desc 2", Content: "Content2"},
	}

	//a1 := Article{Title: "Hello", Desc: "Desc 1", Content: "Content1"}
	//a2 := Article{Title: "Hello", Desc: "Desc 2", Content: "Content2"}

	//Articles = []Article{a1, a2}
	handleRequests()
}
