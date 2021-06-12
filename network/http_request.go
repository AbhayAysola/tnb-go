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
	if res.Response().StatusCode != 200 {
		return errors.New("error making request: " + res.Response().Status)
	}
	err = res.ToJSON(&target)
	if err != nil {
		return errors.New("error parsing response: " + err.Error())
	}
	return nil
}

func PostRequest(url string, params map[string]interface{}, body interface{}, target interface{}) error {
	client := req.New()
	params["format"] = "json"
	res, err := client.Post(url, defaultHeader, req.Param(params), req.BodyJSON(body))
	if err != nil {
		return errors.New("error making request: " + err.Error())
	}
	if res.Response().StatusCode != 200 {
		return errors.New("error making request: " + res.Response().Status)
	}
	err = res.ToJSON(target)
	if err != nil {
		return errors.New("error parsing response: " + err.Error())
	}
	return nil
}

func PatchRequest(url string, params map[string]interface{}, body interface{}, target interface{}) error {
	client := req.New()
	params["format"] = "json"
	res, err := client.Patch(url, defaultHeader, req.Param(params), req.BodyJSON(body))
	if err != nil {
		return errors.New("error making request: " + err.Error())
	}
	if res.Response().StatusCode != 200 {
		return errors.New("error making request: " + res.Response().Status)
	}
	err = res.ToJSON(target)
	if err != nil {
		return errors.New("error parsing response: " + err.Error())
	}
	return nil
}