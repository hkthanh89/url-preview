package main

import (
  "encoding/json"
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "strings"
  "fmt"
  "github.com/PuerkitoBio/goquery"
)

type UrlPreview struct {
  Url   string `json:"url"`
  Title string `json:"title"`
}

type Result struct {
  UrlPreview UrlPreview `json:"object"`
}

type Response struct {
  Code   int `json:"code"`
  Result Result `json:"result"`
}

type ErrorResponse struct {
  Code int `json:"code"`
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

    // Find information
    title := html.Find("title").First()

    meta := html.Find("meta")
    for _, node := range meta.Nodes {
      // Loop nodes
      fmt.Printf("node=%s attributes=%s \n", node.Data, node.Attr)
      // for _, attr := range node.Attr {
      //   // Loop attributes of a node.
      //   fmt.Printf("Key=%s Value=%s Data=%s \n", attr.Key, attr.Val, node.Data)
      // }
    }

    response = Response{
      res.StatusCode,
      Result{
        UrlPreview{
          Url: url,
          Title: title.Text(),
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