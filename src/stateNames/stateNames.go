package statenames

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func StateNames() []string {
	res, err := http.Get("http://publicemailrecords.com/state/arkansas")

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	var passState []string

	doc.Find("body > div.container > div > div").Each(func(i int, s *goquery.Selection) {
		var stateList []string
		s.Find("h4").Each(func(j int, q *goquery.Selection) {
			source := q.Text()
			states := strings.TrimSpace(source)
			trimmedState := strings.Split(states, " ")
			stateList = append(stateList, trimmedState...)
		})

		passState = append(passState, stateList...)
	})

	return passState
}
