package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json: "desc"`
	Content string `json: "content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint Hit: Homepage")
	fmt.Println("Endpoint Hit: Homepage")
}
func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Homepage")
	json.NewEncoder(w).Encode(Articles)
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	//fmt.Fprintf(w, "%v", string(reqBody))
	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)

}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {

		if article.Id == key {

			json.NewEncoder(w).Encode(article)
		}
	}
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":10004", myRouter))
}

func main() {
	Articles = []Article{
		Article{Id: "1", Title: "Title1", Desc: "desc1", Content: "content1"},
		Article{Id: "2", Title: "Title2", Desc: "desc2", Content: "content2"},
	}
	handleRequests()
}
