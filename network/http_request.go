package network

import (
	"errors"

	"github.com/imroc/req"
)

var defaultHeader req.Header = req.Header{"User-Agent": "go-tnb-client"}

func GetRequest(url string, params map[string]interface{}, target interface{}) error {
	client := req.New()
	params["format"] = "json"
	res, err := client.Get(url, defaultHeader, req.Param(params))
	if err != nil {
		return errors.New("error making request: " + err.Error())
	}
	err = res.ToJSON(&target)
	if err != nil {
		return errors.New("error parsing response: " + err.Error())
	}
	return nil
}
