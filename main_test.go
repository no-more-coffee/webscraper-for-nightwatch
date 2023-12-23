package main

import (
	"bytes"
	"encoding/csv"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

var DOCUMENT = `
	<body>
		<div class="organic-results">
			<div class="result-o">
				<div class="name">Organic 1</div>
				<div class="url">organic1@1.com</div>
				<div class="description">First organic result</div>
			</div>
			<div class="result-o">
				<div class="name">Organic 2</div>
				<div class="url">organic2@2.com</div>
			</div>
		</div>
		<div class="local-results">
			<div class="result-l">
				<div class="name">Local 1</div>
				<div class="description">First local result</div>
			</div>
			<div class="result-l">
				<div class="url">local1@.com</div>
				<div class="description">First local result</div>
			</div>
			<div class="result-l">
			</div>
		</div>
	</body>
`

var SELECTORS = map[string]GroupSelectors{
	"organic": {
		Base:        "div.organic-results > div.result-o",
		Title:       "div.name",
		Url:         "div.url",
		Description: "div.description",
	},
	"local": {
		Base:        "div.local-results > div.result-l",
		Title:       "div.name",
		Url:         "div.url",
		Description: "div.description",
	},
}

var EXPECTED = "rank type,rank position,title,url,description" +
	"\norganic,0,Organic 1,organic1@1.com,First organic result\n" +
	"organic,1,Organic 2,organic2@2.com,\n" +
	"local,0,Local 1,,First local result\n" +
	"local,1,,local1@.com,First local result\n" +
	"local,2,,,\n"

func TestWriteCsv(t *testing.T) {
	documentBuffer := bytes.NewBufferString(DOCUMENT)
	doc, err := goquery.NewDocumentFromReader(documentBuffer)
	if err != nil {
		t.Fatal("Error parsing the HTML document:", err)
	}

	resultBuffer := new(bytes.Buffer)
	writer := csv.NewWriter(resultBuffer)

	WriteCsv(doc, writer, SELECTORS)
	writer.Flush()
	result := resultBuffer.String()

	if result != EXPECTED {
		t.Errorf("WriteCsv(). Expected: %q, received: %q", EXPECTED, result)
	}
}
