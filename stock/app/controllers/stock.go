package controllers

import (
	"fmt"
	"github.com/ryuzaki01/go-ms/stock/app/async"
	"github.com/ryuzaki01/go-ms/stock/app/config"
	util "github.com/ryuzaki01/go-ms/stock/app/http"
	"github.com/ryuzaki01/go-ms/stock/app/logs"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

func init() {
	http.Handle("/symbol/", util.Chain(util.APIResourceHandler(stock{})))
}

type stock struct {
	util.APIResourceBase
}

func GetStockData(symbol string) string {
	cfg := config.NewConfig()
	requestUrl := "https://www.alphavantage.co/query?function=OVERVIEW&symbol=" + symbol+ "&apikey=" + cfg.AlphaVantageKey

	logs.Info.Print("starting request to: " + requestUrl)

	resp, err := http.Get(requestUrl)

	if err != nil {
		logs.Error.Print(err.Error())
		return "not found"
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	data, readErr := ioutil.ReadAll(resp.Body)

	if readErr != nil {
		logs.Error.Print(readErr.Error())
	}

	dataStr := string(data)

	logs.Info.Print(dataStr)

	return dataStr
}

func EncryptData(d string) string {
	resp, err := http.PostForm("http://encrypt:3001/encrypt", url.Values{
		"data": { d },
	})

	if err != nil {
		logs.Error.Print(err.Error())
		return "not found"
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	data, readErr := ioutil.ReadAll(resp.Body)

	if readErr != nil {
		logs.Error.Print(readErr.Error())
	}

	return string(data)
}

func (c stock) Get(url string, queries url.Values, body io.Reader) (util.APIStatus, interface{}) {
	if symbol := url[len("/symbol/"):]; len(symbol) != 0 {
		stockFuture := async.Exec(func() interface{} {
			return GetStockData(symbol)
		})
		result := stockFuture.Await()
		resultStr := fmt.Sprintf("%v", result)

		encryptFuture := async.Exec(func() interface{} {
			return EncryptData(resultStr)
		})
		encryptedResult := encryptFuture.Await()

		return util.Success(http.StatusOK), encryptedResult
	}

	return util.Success(http.StatusOK), "no data"
}