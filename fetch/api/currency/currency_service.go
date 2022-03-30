package currency

import (
	"context"
	"encoding/json"
	"errors"
	"fetch/config"
	"fetch/model"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type service struct {
	client *http.Client
}

func NewService(client *http.Client) CurrencyService {
	return &service{
		client: client,
	}
}

type CurrencyService interface {
	GetCurrencyConverter(ctx context.Context) (*model.CurrencyConverter, error)
}

var conf = config.GetConfig()

func (svc *service) sendGetRequest(url string, out interface{}) error {
	resp, err := svc.client.Get(url)
	if err != nil {
		log.Println(err)
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return err
	}

	if resp.StatusCode != 200 {
		log.Println(errors.New(string(body)))
		return errors.New(string(body))
	}

	err = json.Unmarshal([]byte(body), &out)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (svc *service) GetCurrencyConverter(ctx context.Context) (*model.CurrencyConverter, error) {
	var (
		url    = fmt.Sprintf("%s?q=USD_IDR&compact=ultra&apiKey=%s", conf.CurrencyConverterUrl, conf.CurrencyConverterApiKey)
		result model.CurrencyConverter
	)

	err := svc.sendGetRequest(url, &result)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &result, nil
}
