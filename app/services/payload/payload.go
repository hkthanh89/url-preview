package payload

import (
  "github.com/hkthanh89/url-preview/app/models"
)

func Success(code int, urlPreview models.UrlPreview) models.Response {
  return models.Response{
    code,
    models.Result{
      urlPreview,
    },
  }
}

func Error(code int, err string) models.ErrorResponse {
  return models.ErrorResponse{
    code,
    err,
  }
}