package controllers

import (
	"io"
	"net/http"
	"net/url"

	util "github.com/ryuzaki01/go-ms/stock/app/http"
)
func init() {
	http.Handle("/symbol/", util.Chain(util.APIResourceHandler(stock{})))
}

type stock struct {
	util.APIResourceBase
}

func (c stock) Get(url string, queries url.Values, body io.Reader) (util.APIStatus, interface{}) {
	if symbol := url[len("/symbol/"):]; len(symbol) != 0 {
		stock := symbol

		return util.Success(http.StatusOK), stock
	}

	return util.Success(http.StatusOK), "no data"
}