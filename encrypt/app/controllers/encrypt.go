package controllers

import (
	"github.com/ryuzaki01/go-ms/encrypt/app/config"
	"github.com/ryuzaki01/go-ms/encrypt/app/misc"
	"github.com/ryuzaki01/go-ms/stock/app/logs"
	"io"
	"io/ioutil"
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
	cfg := config.NewConfig()
	res, err := ioutil.ReadAll(body)

	if err != nil {
		return util.FailSimple(http.StatusBadRequest), err
	}

	rawStr := string(res)
	encryptedStr := misc.Encrypt(rawStr, cfg.AESKey)

	logs.Info.Print(misc.Decrypt(encryptedStr, cfg.AESKey))

	return util.Success(http.StatusOK), encryptedStr
}