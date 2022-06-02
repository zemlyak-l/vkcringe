package api

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/gorilla/schema"
	"github.com/valyala/fasthttp"
)

const (
	baseApiUrl     = "https://api.vk.com/method/"
	baseApiVersion = "5.131"
)

type Api struct {
	Params  string
	Url     string
	Client  *fasthttp.Client
	Encoder *schema.Encoder
}

func NewApi(token string) *Api {
	// Init Api struct
	params := fmt.Sprintf(
		"?access_token=%s&v=%s", token, baseApiVersion,
	)
	return &Api{
		Url:    baseApiUrl,
		Params: params,
		Client: &fasthttp.Client{
			ReadTimeout:              5 * time.Second,
			WriteTimeout:             5 * time.Second,
			MaxIdleConnDuration:      time.Minute,
			NoDefaultUserAgentHeader: true,
		},
		Encoder: schema.NewEncoder(),
	}
}

func (api *Api) Method(methodName string, data map[string]string, target interface{}) error {
	// Call to vk api method
	resUrl := api.Url + methodName + api.Params

	urlParams := url.Values{}
	for key, value := range data {
		urlParams.Add(key, value)
	}

	urlEncoded := urlParams.Encode()
	reqEntityBytes := []byte(urlEncoded)

	return api.Post(
		resUrl,
		reqEntityBytes,
		target,
	)
}

func (api *Api) Request(methodName string, data interface{}, target interface{}) error {
	// Call to vk api method
	resUrl := api.Url + methodName + api.Params
	urlData := url.Values{}
	if err := api.Encoder.Encode(data, urlData); err != nil {
		return err
	}
	urlEncoded := urlData.Encode()
	encodedData := []byte(urlEncoded)

	return api.Post(
		resUrl,
		encodedData,
		target,
	)
}

func (api *Api) Post(url string, data []byte, target interface{}) error {
	// Create POST request
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/x-www-form-urlencoded")
	req.SetRequestURI(url)
	req.SetBody(data)

	api.Client.Do(req, resp)
	body := resp.Body()

	if target == nil || len(body) == 0 {
		return nil
	}

	return json.Unmarshal(body, target)
}
