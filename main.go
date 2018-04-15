package main

import (
	"encoding/json"
	_ "fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

type UrlPreview struct {
	Url         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type Result struct {
	UrlPreview UrlPreview `json:"object"`
}

type Response struct {
	Code   int    `json:"code"`
	Result Result `json:"result"`
}

type ErrorResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	url := query["url"][0]

	if !strings.Contains(url, "http") {
		url = "http://" + url
	}

	// Get html
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// Server response
	var response interface{}
	if res.StatusCode == 200 { // Valid url
		html, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		var title, description, image string

		meta := html.Find("meta")
		for _, node := range meta.Nodes {
			// Loop nodes
			doc := goquery.NewDocumentFromNode(node)
			val, exists := doc.Selection.Attr("property")

			if exists == true {
				if val == "og:title" {
					title, _ = doc.Selection.Attr("content")
				}

				if val == "og:description" {
					description, _ = doc.Selection.Attr("content")
				}

				if val == "og:image" {
					image, _ = doc.Selection.Attr("content")
				}
			}
		}

		response = Response{
			res.StatusCode,
			Result{
				UrlPreview{
					Url:         url,
					Title:       title,
					Description: description,
					Image:       image,
				},
			},
		}
	} else { // Invalid url
		response = ErrorResponse{
			400,
			"Invalid URL",
		}
	}

	data, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8889", router))
}
