package citynames

import (
	statenames "bulk-email/src/stateNames"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func CityNames() []string {
	var cityList []string
	var passCity []string

	allStates := statenames.StateNames()

	fmt.Println("wait for a while to collect cities...")

	// for _, val := range allStates {
	res, err := http.Get(string(allStates[0]))

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	doc.Find("body > div.container > div > div").Each(func(i int, s *goquery.Selection) {

		s.Find("h4").Each(func(j int, q *goquery.Selection) {
			source := q.Text()
			cities := strings.TrimSpace(source)
			trimmedCities := strings.Split(cities, " ")
			cityList = append(cityList, trimmedCities...)
		})
	})
	// }
	passCity = append(passCity, cityList...)

	fmt.Println("Cities are now ready to pass on...")

	return passCity
}
