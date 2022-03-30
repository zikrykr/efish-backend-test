package fetch

import (
	"context"
	"encoding/json"
	"errors"
	"fetch/config"
	"fetch/model"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"gopkg.in/guregu/null.v4"
)

type service struct {
	client *http.Client
}

func NewService(client *http.Client) FetchService {
	return &service{
		client: client,
	}
}

var conf = config.GetConfig()

type FetchService interface {
	GetResources(ctx context.Context) (*[]model.ResourceData, error)
	GetCurrencyConverter(ctx context.Context) (*model.CurrencyConverter, error)
}

func (svc *service) sendGetRequest(url string, out interface{}) error {
	resp, err := svc.client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New(string(body))
	}

	err = json.Unmarshal([]byte(body), &out)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (svc *service) GetResources(ctx context.Context) (*[]model.ResourceData, error) {
	var (
		url    = conf.ResourceUrl
		resp   []model.ResourceDataResponse
		result []model.ResourceData
	)

	err := svc.sendGetRequest(url, &resp)
	if err != nil {
		return nil, err
	}

	currencyConverter, err := svc.GetCurrencyConverter(ctx)
	if err != nil {
		return nil, err
	}

	for _, v := range resp {
		tempResult := &model.ResourceData{
			UUID:      v.UUID,
			Commodity: v.Commodity,
			Province:  v.Province,
			City:      v.City,
			Size:      v.Size,
			Price:     v.Price,
			ParsedAt:  v.ParsedAt,
			Timestamp: v.Timestamp,
		}

		if v.Price.Valid {
			price, _ := strconv.ParseFloat(v.Price.String, 64)

			priceUSD := price / currencyConverter.Value

			priceUSDStr := strconv.FormatFloat(priceUSD, 'f', 2, 64)

			tempResult.PriceUSD = null.NewString(priceUSDStr, priceUSDStr != "")
		}

		result = append(result, *tempResult)
	}

	return &result, nil
}

func (svc *service) GetCurrencyConverter(ctx context.Context) (*model.CurrencyConverter, error) {
	var (
		url    = fmt.Sprintf("%s?q=USD_IDR&compact=ultra&apiKey=%s", conf.CurrencyConverterUrl, conf.CurrencyConverterApiKey)
		result model.CurrencyConverter
	)

	err := svc.sendGetRequest(url, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
