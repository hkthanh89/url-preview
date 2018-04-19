package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
  pageService "github.com/hkthanh89/url-preview/app/services/page"
  "github.com/hkthanh89/url-preview/app/models"
)

type Result struct {
	UrlPreview models.UrlPreview `json:"object"`
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

  url := pageService.NormalizeUrl(query["url"][0])

	// Get html
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// Server response
	var response interface{}
	if res.StatusCode == 200 {
    // Valid url
    fmt.Println("-- start getting info")
    urlPreview, err := pageService.GetPreviewInfo(url, res.Body)
    if err != nil {
      log.Fatal(err)
    }

		response = Response{
			res.StatusCode,
			Result{
				urlPreview,
			},
		}
	} else {
    // Invalid url
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

func blank(s string) bool {
  return len(strings.TrimSpace(s)) == 0
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8889", router))
}
