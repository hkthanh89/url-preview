package main

import (
  "encoding/json"
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "fmt"
  "io/ioutil"
  "strings"
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

  if !strings.Contains(url, "http") {
    url = "http://" + url
  }

  // Get html
  resp, err := http.Get(url)
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()

  html, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%s", html)

  // Server response
  response := Response{
    resp.StatusCode,
    Result{
      UrlPreview{
        Url: url,
      },
    },
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