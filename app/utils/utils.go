package utils

import (
  "strings"
)

func NormalizeUrl(url string) string {
  if !strings.Contains(url, "http") {
    url = "http://" + url
  }
  return url
}

func Blank(s string) bool {
  return len(strings.TrimSpace(s)) == 0
}