package utils

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "strconv"
)

type Example struct {
  name     string
  input    string
  expected string
}

func TestNormalizeUrl(t *testing.T) {
  var urlExamples = []Example{
    { "Add http if url is missed", "google.com.vn", "http://google.com.vn" },
    { "Don't add http if url already has", "http://google.com.vn", "http://google.com.vn" },
  }

  for _, example := range urlExamples {
    actual := NormalizeUrl(example.input)

    t.Log(example.name)
    assert.Equal(t, example.expected, actual)
  }
}

func TestBlank(t *testing.T) {
  var blankExamples = []Example{
    { "False if string has space at beginning", " string", "false" },
    { "False if string has space at ending", "string ", "false" },
    { "False if string has space at beginning and ending", " string ", "false" },
    { "True if there is no space", "", "true" },
    { "True if there is one space", " ", "true" },
    { "True if there are spaces", " ", "true" },
  }

  for _, example := range blankExamples {
    expected, _ := strconv.ParseBool(example.expected)
    actual := Blank(example.input)

    t.Log(example.name)
    assert.Equal(t, expected, actual)
  }
}