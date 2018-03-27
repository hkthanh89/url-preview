package main

import (
  "encoding/json"
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

type UrlPreview struct {
  Url   string `json:"url"`
}

type Result struct {
  UrlPreview UrlPreview `json:"object"`
}

type Response struct {
  Code   int `json:"code"`
  Result Result `json:"result"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()
  url := query["url"][0]

  // Server response
  resp := Response{
    200,
    Result{
      UrlPreview{
        Url: url,
      },
    },
  }
  response, err := json.Marshal(resp)

  if err != nil {
    log.Fatal(err)
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(200)
  w.Write(response)
}

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/", HomeHandler).Methods("GET")

  log.Fatal(http.ListenAndServe(":8889", router))
}