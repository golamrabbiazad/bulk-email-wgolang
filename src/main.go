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

var data = [][]string{}

func main() {
	allCities := citynames.CityNames()

	fmt.Println("printing emails")

	file, err := os.Create("result.csv")

	checkError("Cannot create file", err)

	defer file.Close()

	for _, val := range allCities {

		res, err := http.Get("http://publicemailrecords.com/city/" + val + "/Arkansas")
		data = append(data, strings.Fields(val))

		checkError("Cannot get url", err)

		defer res.Body.Close()

		doc, err := goquery.NewDocumentFromReader(res.Body)

		checkError("response body not parsed", err)

		doc.Find(".container .tbl-sect").Each(func(i int, s *goquery.Selection) {
			s.Find("div.email.email-data.mk-link").Each(func(j int, q *goquery.Selection) {
				text := q.Text()
				data = append(data, strings.Fields(strings.TrimSpace(text)))
			})
		})

	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	fmt.Println("Ready to export in CSV")
	for _, val1 := range data {
		err := writer.Write(val1)
		checkError("Cannot write to file", err)
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
