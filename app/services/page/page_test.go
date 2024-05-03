package page

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Example struct {
	name     string
	input    string
	expected string
}

func TestGetPreviewInfo(t *testing.T) {
	testPreviewInfo(t, titleExamples, "title")
	testPreviewInfo(t, descriptionExamples, "description")
	testPreviewInfo(t, imageExamples, "image")
	testPreviewInfo(t, readingTimeExamples, "readingTime")
}

func testPreviewInfo(t *testing.T, examples []Example, info string) {
	for _, example := range examples {
		r := strings.NewReader(example.input)
		urlPreview, _ := GetPreviewInfo(r)

		var actual string
		switch info {
		case "title":
			actual = urlPreview.Title
		case "description":
			actual = urlPreview.Description
		case "image":
			actual = urlPreview.Image
		case "readingTime":
			actual = urlPreview.ReadingTime
		}

		t.Log(example.name)
		assert.Equal(t, example.expected, actual)
	}
}

var titleExamples = []Example{
	{"og:title exists",
		`<!DOCTYPE html>
    <html>
    <head>
    <meta content="Meta Title" property="og:title"/>
    <title>Title</title>
    </head>
    </html>`,
		"Meta Title",
	},
	{"og:title does not exist",
		`
    <!DOCTYPE html>
    <html>
    <head>
    <title>My Title</title>
    </head>
    </html>
    `,
		"My Title",
	},
}

var descriptionExamples = []Example{
	{"og:description exists",
		`<!DOCTYPE html>
    <html>
    <head>
    <meta content="Description" name="description"/>
    <meta content="Meta description" property="og:description"/>
    </head>
    </html>`,
		"Meta description",
	},
	{"og:description does not exist",
		`
    <!DOCTYPE html>
    <html>
    <head>
    <meta content="Description" name="description"/>
    </head>
    </html>
    `,
		"Description",
	},
}

var imageExamples = []Example{
	{"og:image exists",
		`<!DOCTYPE html>
    <html>
    <head>
    <meta content="http://somedomain/image.jpg" property="og:image"/>
    </head>
    </html>`,
		"http://somedomain/image.jpg",
	},
	{"og:image does not exist",
		`
    <!DOCTYPE html>
    <html>
    <head>
    </head>
    </html>
    `,
		"",
	},
}

var readingTimeExamples = []Example{
	{"webpage has content",
		`<!DOCTYPE html>
    <html>
    <body>
      <p>This is content of website</p>
    </body>
    </html>`,
		"1 min",
	},
	{"webpage has no content",
		`<!DOCTYPE html>
    <html>
    <body>
    </body>
    </html>`,
		"0 min",
	},
}
