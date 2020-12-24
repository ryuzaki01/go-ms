package controllers

import (
	"github.com/ryuzaki01/go-ms/stock/app/logs"
	"io"
	"net/http"
	"net/url"

	util "github.com/ryuzaki01/go-ms/encrypt/app/http"
)
func init() {
	http.Handle("/encrypt/", util.Chain(util.APIResourceHandler(encrypt{})))
}

type encrypt struct {
	util.APIResourceBase
}

func (c encrypt) Post(url string, queries url.Values, body io.Reader) (util.APIStatus, interface{}) {
	response := "test"
	logs.Info.Print(body)

	return util.Success(http.StatusOK), response
}