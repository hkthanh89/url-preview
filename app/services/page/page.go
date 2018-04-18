package page

import (
  "strings"
  "github.com/PuerkitoBio/goquery"
)

func NormalizeUrl(link string) (url string) {
  if !strings.Contains(link, "http") {
    url = "http://" + link
  }
  return
}

func GetPreviewInfo(document *goquery.Document) (url, title, description, image string) {
  document.Find("meta").Each(func(i int, s *goquery.Selection) {
    if s.AttrOr("property", "") == "og:url" {
      url = s.AttrOr("content", "")
    }

    if s.AttrOr("property", "") == "og:title" {
      title = s.AttrOr("content", "")
    } else {
      // Fallback
      title = document.Find("head > title").First().Text()
    }

    if (s.AttrOr("name", "") == "description") || (s.AttrOr("property", "") == "og:description") {
      description = s.AttrOr("content", "")
    }

    if s.AttrOr("property", "") == "og:image" {
      image = s.AttrOr("content", "")
    }
  })

  return
}

