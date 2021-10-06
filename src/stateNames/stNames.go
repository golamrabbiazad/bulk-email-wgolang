package statenames

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func StateNames() []string {
	var passState []string

	res, err := http.Get("statenames")

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	doc.Find("body > div.container > div > div ").Each(func(i int, s *goquery.Selection) {
		var stateList []string
		s.Find("h4 > a").Each(func(j int, q *goquery.Selection) {
			source, _ := q.Attr("href")
			stateList = append(stateList, source)
		})

		passState = append(passState, stateList...)
	})

	fmt.Println(strconv.Itoa(len(passState)) + " states are ready to pass on...")
	return passState
}
