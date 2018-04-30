package page

import (
  "io"
  "strings"
  "github.com/PuerkitoBio/goquery"
  "github.com/hkthanh89/url-preview/app/models"
)

func NormalizeUrl(url string) string {
  if !strings.Contains(url, "http") {
    url = "http://" + url
  }
  return url
}

func GetPreviewInfo(r io.Reader) (models.UrlPreview, error) {
  document, err := goquery.NewDocumentFromReader(r)
  urlPreview := models.UrlPreview{}

  if err != nil {
    return urlPreview, err
  }

  document.Find("meta").Each(func(i int, s *goquery.Selection) {
    if s.AttrOr("property", "") == "og:title" {
      urlPreview.Title = attrContent(s)
    }

    if (s.AttrOr("name", "") == "description") || (s.AttrOr("property", "") == "og:description") {
      urlPreview.Description = attrContent(s)
    }

    if s.AttrOr("property", "") == "og:image" {
      urlPreview.Image = attrContent(s)
    }
  })

  // Get missing data
  if blank(urlPreview.Title) {
    urlPreview.Title = document.Find("head > title").First().Text()
  }

  return urlPreview, nil
}

func attrContent(s *goquery.Selection) string {
  return s.AttrOr("content", "")
}

func blank(s string) bool {
  return len(strings.TrimSpace(s)) == 0
}
