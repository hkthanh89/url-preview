package page

import (
	"fmt"
	"io"
	"math"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/hkthanh89/url-preview/app/models"
	"github.com/hkthanh89/url-preview/app/utils"
)

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
	if utils.Blank(urlPreview.Title) {
		urlPreview.Title = document.Find("head > title").First().Text()
	}

	wordsCount := 0
	document.Find("p").Each(func(i int, s *goquery.Selection) {
		wordsCount += len(strings.Split(s.Text(), " "))
	})

	wordsPerMinute := 200
	minutesReading := math.Ceil(float64(wordsCount) / float64(wordsPerMinute))
	urlPreview.ReadingTime = fmt.Sprintf("%d min", int(minutesReading))

	return urlPreview, nil
}

func attrContent(s *goquery.Selection) string {
	return s.AttrOr("content", "")
}
