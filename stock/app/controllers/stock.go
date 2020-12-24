package controllers

import (
	"fmt"
	"github.com/ryuzaki01/go-ms/stock/app/async"
	util "github.com/ryuzaki01/go-ms/stock/app/http"
	"github.com/ryuzaki01/go-ms/stock/app/models"
	"io"
	"net/http"
	"net/url"
)

func init() {
	http.Handle("/symbol/", util.Chain(util.APIResourceHandler(stock{})))
}

type stock struct {
	util.APIResourceBase
}

func (c stock) Get(url string, queries url.Values, body io.Reader) (util.APIStatus, interface{}) {
	if symbol := url[len("/symbol/"):]; len(symbol) != 0 {
		stockFuture := async.Exec(func() interface{} {
			return models.GetStock(symbol)
		})
		result := stockFuture.Await()
		resultStr := fmt.Sprintf(`"{"data" : "%v"}`, result)

		encryptFuture := async.Exec(func() interface{} {
			return models.PostEncrypt(resultStr)
		})
		encryptedResult := encryptFuture.Await()

		return util.Success(http.StatusOK), encryptedResult
	}

	return util.Success(http.StatusOK), "no data"
}