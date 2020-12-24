package controllers

import (
	"github.com/ryuzaki01/go-ms/encrypt/app/misc"
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

type dataRequest struct {
	Data string    `json:"data,string"`
}

func (c encrypt) Post(url string, queries url.Values, body io.Reader) (util.APIStatus, interface{}) {
	req := &dataRequest{}
	if err := misc.ReadMBJSON(body, req, 100); err != nil {
		logs.Error.Printf("Could not decode response body as a json. Error: %v", err)
		return util.Fail(http.StatusInternalServerError, err.Error()), nil
	}

	return util.Success(http.StatusOK), req.Data
}