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

func main() {
	stateName := "Oklahoma"

	newMapData := make(map[string]ScrapeFormat)

	allCities := citynames.CityNames()

	fmt.Println("collected cities")

	for _, val := range allCities {
		res, err := http.Get("http://publicemailrecords.com/city/" + val + "/" + stateName)
		errorhandles.CheckError("can't get url response ", err)

		defer res.Body.Close()

		doc, err := goquery.NewDocumentFromReader(res.Body)
		errorhandles.CheckError("response body not parsed ", err)

		// set default value = 1
		totalPage := 1

		doc.Find("#citylistings > p > a:last-child").Each(func(i int, s *goquery.Selection) {
			stringVal := strings.TrimSpace(s.Text())
			page, err := strconv.Atoi(stringVal)
			errorhandles.CheckError("can't convert to interger ", err)

			totalPage = page
		})

		fmt.Println("\nfound " + strconv.Itoa(totalPage) + " page...\n")

		var text string
		var textArr []string

		for i := 1; i <= totalPage; i++ {
			newRes, err := http.Get("http://publicemailrecords.com/city/" + val + "/" + stateName + "?page=" + strconv.Itoa(i))
			errorhandles.CheckError("Cannot get url ", err)

			defer res.Body.Close()

			newDoc, err := goquery.NewDocumentFromReader(newRes.Body)
			errorhandles.CheckError("response body not parsed ", err)

			newDoc.Find("div.email.email-data.mk-link").Each(func(j int, q *goquery.Selection) {
				text = findEmail(q)
				textArr = append(textArr, strings.Fields(text)...)
			})
		}
		newMapData[val] = ScrapeFormat{
			email: textArr,
		}

		if err := os.Mkdir("./states/"+stateName+"/", 0755); err != nil && !os.IsExist(err) {
			errorhandles.CheckError("can't create directory ", err)
		}
	}

	fmt.Println(newMapData)

	for key, value := range newMapData {
		file, err := os.Create("./states/" + stateName + "/" + key + ".csv")
		errorhandles.CheckError("Cannot create file", err)

		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		fmt.Println("\nReady to export in CSV")

		for _, valdx := range value.email {
			if err := writer.Write(strings.Fields(valdx)); err != nil {
				errorhandles.CheckError("can't writes values ", err)
			}
		}
	}
}

func findEmail(card *goquery.Selection) string {
	var text string
	if strings.Contains(strings.TrimSpace(card.Text()), "@") {
		text = strings.TrimSpace(card.Text())
	}

	fmt.Println(text)
	return text
}

// stored data as this format
// [{ Achille [example@gmail.com example@yahoo.com]} {Ada [example@gmail.com example@yahoo.com]}]
