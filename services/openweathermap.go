package services

import (
	"encoding/json"
	"fmt"
	"github.com/gojek/heimdall/v7/httpclient"
	"io/ioutil"
	"time"
)

type temperature struct {
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64     `json:"temp_max"`
		Pressure  float64     `json:"pressure"`
		Humidity  float64     `json:"humidity"`
	} `json:"main"`
}

const timeout = 10 * time.Second
const retryCount = 2
const apiKEY = "b77e07f479efe92156376a8b07640ced" // API key provided by Ifood, doesn't seems necessary to have it in ENV.

func getUrl(city string) string {
	return fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKEY)
}

func getHttpClient() *httpclient.Client {
	return httpclient.NewClient(
		httpclient.WithHTTPTimeout(timeout),
		httpclient.WithRetryCount(retryCount),
	)
}

func GetCityByName(city string) (temp *temperature, err error) {
	client := getHttpClient()
	url := getUrl(city)
	res, err := client.Get(url, nil)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &temp)

	if err != nil {
		return nil, err
	}

	return temp, err
}
