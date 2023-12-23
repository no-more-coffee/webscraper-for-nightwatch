package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	document := ReadDocument("pizza.html")

	selectors := ReadSelectors("group-selectors.json")

	resultFile, err := os.Create("out/result.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer resultFile.Close()

	writer := csv.NewWriter(resultFile)
	defer writer.Flush()

	WriteCsv(document, writer, selectors)
	if err := writer.Error(); err != nil {
		log.Fatal("Error on CSV writer flush:", err)
	}
}

type GroupSelectors struct {
	Base        string `json:"base"`
	Title       string `json:"title"`
	Url         string `json:"url"`
	Description string `json:"description"`
}

func ReadSelectors(path string) map[string]GroupSelectors {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var groupSelectors map[string]GroupSelectors
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&groupSelectors)
	if err != nil {
		log.Fatal("Error decoding JSON:", err)
	}

	return groupSelectors
}

func ReadDocument(path string) *goquery.Document {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		log.Fatal("Error parsing the HTML document:", err)
	}

	return doc
}

func WriteCsv(document *goquery.Document, resultWriter *csv.Writer, selectors map[string]GroupSelectors) {
	header := []string{
		"rank type",
		"rank position",
		"title",
		"url",
		"description",
	}
	if err := resultWriter.Write(header); err != nil {
		log.Fatal("Error writing header to CSV:", err)
	}

	for group, selector := range selectors {
		document.Find(selector.Base).Each(func(position int, selection *goquery.Selection) {
			title := selection.Find(selector.Title).Text()
			url := selection.Find(selector.Url).Text()
			description := selection.Find(selector.Description).Text()

			record := []string{
				group,
				strconv.Itoa(position),
				title,
				url,
				description,
			}
			if err := resultWriter.Write(record); err != nil {
				log.Fatal("Error writing record to CSV:", err)
			}
		})
	}
}
