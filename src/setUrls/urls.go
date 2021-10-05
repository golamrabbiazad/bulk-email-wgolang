package seturls

import (
	errorhandles "bulk-email/src/errorHandles"
	"net/http"
)

var Url string

func SetCityUrl(val string, state string) *http.Response {
	Url = "http://publicemailrecords.com/city/" + val + "/" + state

	res, err := http.Get(Url)
	errorhandles.CheckError("Cannot get url", err)
	defer res.Body.Close()

	return res
}
