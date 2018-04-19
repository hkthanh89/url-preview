package models

type Response struct {
  Code   int    `json:"code"`
  Result Result `json:"result"`
}

type Result struct {
  UrlPreview UrlPreview `json:"object"`
}

type ErrorResponse struct {
  Code  int    `json:"code"`
  Error string `json:"error"`
}