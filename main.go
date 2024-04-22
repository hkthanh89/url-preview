package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	pageService "github.com/hkthanh89/url-preview/app/services/page"
	"github.com/hkthanh89/url-preview/app/services/payload"
	"github.com/hkthanh89/url-preview/app/utils"
	"golang.org/x/time/rate"
)

func PreviewHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	url := utils.NormalizeUrl(query["url"][0])

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
		urlPreview, err := pageService.GetPreviewInfo(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		urlPreview.Url = url

		response = payload.Success(200, urlPreview)
	} else {
		// Invalid url
		response = payload.Error(400, "Invalid URL")
	}

	data, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)
}

var limiter = rate.NewLimiter(10, 5)

func rateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "Too Many Requests!", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/preview", PreviewHandler).Methods("GET")
	router.Use(rateLimit)

	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
