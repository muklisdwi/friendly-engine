package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Net Http")
	// BasicWebServer()

	Articles = []Article{
		Article{
			Id:      "1",
			Title:   "Story 1",
			Desc:    "About Story 1",
			Content: "Summary of Story 1",
		},
		Article{
			Id:      "2",
			Title:   "Story 2",
			Desc:    "About STory 2",
			Content: "Summary of Story 2",
		},
	}

	handlerApi()

	// handleStatic()
}

func handleStatic() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Fatal(http.ListenAndServe(":9009", nil))
}

func handlerApi() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Endpoint hit: / (GET)")
		fmt.Fprintf(w, "homepage")
	}).Methods("GET")

	myRouter.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Endpoint hit: /articles (GET)")
		json.NewEncoder(w).Encode(Articles)
	}).Methods("GET")

	myRouter.HandleFunc("/article/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Endpoint hit: /article/{id} (GET)")
		vars := mux.Vars(r)
		key := vars["id"]
		for i, l := range Articles {
			if l.Id == key {
				json.NewEncoder(w).Encode(l)
				break
			}
			if i == len(Articles)-1 {
				fmt.Fprintf(w, "Article not Found")
			}
		}
	}).Methods("GET")

	myRouter.HandleFunc("/article", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Endpoint hit: /article (POST)")
		reqBody, _ := ioutil.ReadAll(r.Body)
		var article Article
		json.Unmarshal(reqBody, &article)
		Articles = append(Articles, article)
		json.NewEncoder(w).Encode(article)
	}).Methods("POST")

	myRouter.HandleFunc("/article/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Endpoint hit: /article/{id} (DELETE)")
		vars := mux.Vars(r)
		key := vars["id"]
		for i, l := range Articles {
			if l.Id == key {
				Articles = append(Articles[:i], Articles[i+1:]...)
				break
			}
		}
		json.NewEncoder(w).Encode(Articles)
	}).Methods("DELETE")

	myRouter.HandleFunc("/article/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Endpoint hit: /article/{id} (PUT)")
		reqBody, _ := ioutil.ReadAll(r.Body)
		var article Article
		json.Unmarshal(reqBody, &article)
		vars := mux.Vars(r)
		key := vars["id"]
		for i, l := range Articles {
			if l.Id == key {
				Articles[i].Title = article.Title
				Articles[i].Desc = article.Desc
				Articles[i].Content = article.Content
				break
			}
		}
		json.NewEncoder(w).Encode(Articles)
	}).Methods("PUT")

	http.ListenAndServe(":8081", myRouter)
}

type Article struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func BasicWebServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		str := strings.Replace(r.URL.Path, "/", "", 1)
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(str))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
