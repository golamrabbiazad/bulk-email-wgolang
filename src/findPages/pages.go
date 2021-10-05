package findpages

import (
	errorhandles "bulk-email/src/errorHandles"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var pages []int

var totalPage int

func FindPages(doc *goquery.Document) int {
	doc.Find("#citylistings > p > a").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		page, err := strconv.Atoi(text)
		errorhandles.CheckError("Cannot convert string to int", err)
		pages = append(pages, page)
	})
	totalPage = len(pages) - 1

	orgPage := pages[totalPage]

	return orgPage
}
