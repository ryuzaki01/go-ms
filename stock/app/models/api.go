package models

import (
	"github.com/ryuzaki01/go-ms/stock/app/config"
)

type daoEncrypt struct {
	Header   APIHeader `json:"header"`
	Response string    `json:"response"`
}

// GetStock retrieves a specified stock data from alpha vantage
//  @param symbol string
//  @return response string
func GetStock(symbol string) string {
	cfg := config.NewConfig()
	response := ""
	if fetch("GET", "https://www.alphavantage.co/query?function=OVERVIEW&symbol=" + symbol + "&apikey=" + cfg.AlphaVantageKey, "", response) == nil {
		return response
	}
	return "not found"
}

// PostEncrypt retrieves encrypted string
//  @param data string
//  @return response string
func PostEncrypt(data string) string {
	res := &daoEncrypt{}
	if fetch("POST", "http://encrypt:3001/encrypt/", data, res) == nil {
		return res.Response
	}
	return "not found"
}