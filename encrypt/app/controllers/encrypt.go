package controllers

import (
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

	return util.Success(http.StatusOK), response
}