package main

import (
	citynames "bulk-email/src/cityNames"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var data [][]string

var emailData [][]string

func main() {
	allCities := citynames.CityNames()
	fmt.Println("collected cities")

	data = append(data, allCities)

	for idx, val := range allCities {
		res, err := http.Get("http://publicemailrecords.com/city/" + val + "/Arkansas")
		checkError("Cannot get url", err)

		defer res.Body.Close()

		if idx == 2 {
			break
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)

		checkError("response body not parsed", err)

		appendEmails(doc)

	}

	file, err := os.Create("result.csv")
	checkError("Cannot create file", err)

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	data = append(data, emailData...)
	fmt.Println(data)

	fmt.Println("\nReady to export in CSV")

	for _, row := range data {
		if err := writer.Write(row); err != nil {
			checkError("error writing record to file", err)
		}
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func appendEmails(doc *goquery.Document) {
	var textArr []string
	doc.Find(".container .tbl-sect").Each(func(i int, s *goquery.Selection) {
		s.Find("div.email.email-data.mk-link").Each(func(j int, q *goquery.Selection) {
			text := findEmail(q)
			textArr = append(textArr, text)
		})
	})
	emailData = append(emailData, textArr)
}

func findEmail(card *goquery.Selection) string {
	var text string
	if strings.Contains(strings.TrimSpace(card.Text()), "@") {
		text = strings.TrimSpace(card.Text())
	}

	fmt.Println(text)
	return text
}
