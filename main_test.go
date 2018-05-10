package main

import(
  "net/http"
  "net/http/httptest"
  "testing"
  "github.com/gorilla/mux"
  "github.com/stretchr/testify/assert"
  "encoding/json"
  "github.com/hkthanh89/url-preview/app/models"
  "github.com/hkthanh89/url-preview/app/utils"
)

func Router() *mux.Router {
  router := mux.NewRouter()
  router.HandleFunc("/preview", PreviewHandler).Methods("GET")
  return router
}

func TestPreviewHandler(t *testing.T) {
  req, _ := http.NewRequest("GET", "/preview?url=apple.com", nil)
  res := httptest.NewRecorder()
  Router().ServeHTTP(res, req)

  var response models.Response
  json.Unmarshal(res.Body.Bytes(), &response)
  var preview = response.Result.UrlPreview

  assert.Equal(t, 200, response.Code)
  assert.Equal(t, false, utils.Blank(preview.Url))
  assert.Equal(t, false, utils.Blank(preview.Title))
  assert.Equal(t, false, utils.Blank(preview.Description))
  assert.Equal(t, false, utils.Blank(preview.Image))

  // Get response object
  // response = models.Response{}
  // json.NewDecoder(res.Body).Decode(&response)
}