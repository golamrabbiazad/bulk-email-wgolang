package main

import (
	citynames "bulk-email/src/cityNames"
	errorhandles "bulk-email/src/errorHandles"
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type ScrapeFormat struct {
	email []string
}

var data [][]string

// var cityData [][]string

// var emailData [][]string

func main() {
	newMapData := make(map[string]ScrapeFormat)

	allCities := citynames.CityNames()
	fmt.Println("collected cities")

	data = append(data, allCities)

	for idx, val := range allCities {
		res, err := http.Get("http://publicemailrecords.com/city/" + val + "/Oklahoma")
		errorhandles.CheckError("can't get url response ", err)

		defer res.Body.Close()

		doc, err := goquery.NewDocumentFromReader(res.Body)
		errorhandles.CheckError("response body not parsed ", err)

		if idx == 2 {
			break
		}

		// set default value = 1
		totalPage := 1
		// var totalPage int

		doc.Find("#citylistings > p > a:last-child").Each(func(i int, s *goquery.Selection) {
			stringVal := strings.TrimSpace(s.Text())
			page, err := strconv.Atoi(stringVal)
			errorhandles.CheckError("can't convert to interger ", err)

			totalPage = page
		})

		fmt.Println("\nfound " + strconv.Itoa(totalPage) + " page...\n")

		var text string
		var textArr []string

		doc.Find("div.email.email-data.mk-link").Each(func(j int, q *goquery.Selection) {
			text = findEmail(q)
			textArr = append(textArr, text)
			newMapData[val] = ScrapeFormat{
				email: textArr,
			}
		})

		fmt.Println(newMapData)

		// cityData = append(cityData, textArr)

		// emailData = append(emailData, textArr)

		// for i := 1; i <= totalPage; i++ {
		// 	newRes, err := http.Get("http://publicemailrecords.com/city/" + val + "/Oklahoma" + "?page=" + strconv.Itoa(i))
		// 	errorhandles.CheckError("Cannot get url ", err)

		// 	defer res.Body.Close()

		// 	newDoc, err := goquery.NewDocumentFromReader(newRes.Body)
		// 	errorhandles.CheckError("response body not parsed ", err)

		// appendEmails(newDoc)

		// var text string
		// var textArr []string

		// newDoc.Find("div.email.email-data.mk-link").Each(func(j int, q *goquery.Selection) {
		// 	text = findEmail(q)
		// 	textArr = append(textArr, text)
		// })

		// cityData = append(cityData, textArr)
		// }

	}

	file, err := os.Create("./states/oklahoma.csv")
	errorhandles.CheckError("Cannot create file", err)

	defer file.Close()

	writer := csv.NewWriter(file)

	defer writer.Flush()

	fmt.Println("\nReady to export in CSV")

	for key, value := range newMapData {
		r := make([]string, 0, 1+len(value.email))

		r = append(r, key)
		r = append(r, value.email...)
		err := writer.Write(r)
		if err != nil {
			errorhandles.CheckError("error writing record to file", err)
		}
	}
	writer.Flush()
}

// func appendEmails(doc *goquery.Document) {
//

// }

func findEmail(card *goquery.Selection) string {
	var text string
	if strings.Contains(strings.TrimSpace(card.Text()), "@") {
		text = strings.TrimSpace(card.Text())
	}

	fmt.Println(text)
	return text
}
