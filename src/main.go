package main

import (
	citynames "bulk-email/src/cityNames"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	allCities := citynames.CityNames()

	fmt.Println("printing emails")
	for _, val := range allCities {
		res, err := http.Get("http://publicemailrecords.com/city/" + val + "/Arkansas")

		if err != nil {
			log.Fatal(err)
		}

		defer res.Body.Close()

		doc, err := goquery.NewDocumentFromReader(res.Body)

		if err != nil {
			log.Fatal(err)
		}

		doc.Find(".container .tbl-sect").Each(func(i int, s *goquery.Selection) {
			s.Find("div.email.email-data.mk-link").Each(func(j int, q *goquery.Selection) {
				text := q.Text()
				fmt.Println(strings.TrimSpace(text))
			})
		})
	}
}
