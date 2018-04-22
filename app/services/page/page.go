package page

import (
  "io"
  "strings"
  "github.com/PuerkitoBio/goquery"
  "github.com/hkthanh89/url-preview/app/models"
)

func NormalizeUrl(link string) (url string) {
  if !strings.Contains(link, "http") {
    url = "http://" + link
  }
  return
}

func GetPreviewInfo(url string, r io.Reader) (models.UrlPreview, error) {
  document, err := goquery.NewDocumentFromReader(r)
  urlPreview := models.UrlPreview{Url: url}

  if err != nil {
    return urlPreview, err
  }

  document.Find("meta").Each(func(i int, s *goquery.Selection) {
    if s.AttrOr("property", "") == "og:title" {
      urlPreview.Title = attrContent(s)
    } else {
      // Fallback
      urlPreview.Title = document.Find("head > title").First().Text()
    }

    if (s.AttrOr("name", "") == "description") || (s.AttrOr("property", "") == "og:description") {
      urlPreview.Description = attrContent(s)
    }

    if s.AttrOr("property", "") == "og:image" {
      urlPreview.Image = attrContent(s)
    }
  })

  return urlPreview, nil
}

func attrContent(s *goquery.Selection) string {
  return s.AttrOr("content", "")
}
