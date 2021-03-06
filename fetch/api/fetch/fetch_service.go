package fetch

import (
	"context"
	"encoding/json"
	"errors"
	"fetch/api/currency"
	"fetch/cache"
	"fetch/config"
	"fetch/helper"
	"fetch/model"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"

	"gopkg.in/guregu/null.v4"
)

type service struct {
	client         *http.Client
	currenyService currency.CurrencyService
}

func NewService(client *http.Client, currencyService currency.CurrencyService) FetchService {
	return &service{
		client:         client,
		currenyService: currencyService,
	}
}

var conf = config.GetConfig()

type FetchService interface {
	GetResources(ctx context.Context) (*[]model.ResourceData, error)
	GetCurrencyConverter(ctx context.Context) (*model.CurrencyConverter, error)
	GetResourcesAggregate(ctx context.Context) ([]model.ResourceDataAggregate, error)
}

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

func (svc *service) GetResources(ctx context.Context) (*[]model.ResourceData, error) {
	var (
		url    = conf.ResourceUrl
		resp   []model.ResourceDataResponse
		result []model.ResourceData
	)

	err := svc.sendGetRequest(url, &resp)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	currencyConverter, err := svc.GetCurrencyConverter(ctx)
	if err != nil {
		log.Println(err)
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
	if cache.CurrencyConverterResult != 0 {
		return &model.CurrencyConverter{
			Value: cache.CurrencyConverterResult,
		}, nil
	} else {
		return svc.currenyService.GetCurrencyConverter(ctx)
	}
}

func (svc *service) GetResourcesAggregate(ctx context.Context) ([]model.ResourceDataAggregate, error) {
	var (
		result         []model.ResourceDataAggregate
		resourcePrices []model.ResourcePrice
	)

	resources, err := svc.GetResources(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for _, v := range *resources {
		if !v.Province.Valid || !v.ParsedAt.Valid || !v.Price.Valid {
			continue
		}
		date := helper.ParseStringToTime(v.ParsedAt.String)
		year, week := date.ISOWeek()

		price, _ := strconv.ParseFloat(v.Price.String, 64)

		resourcePrices = append(resourcePrices, model.ResourcePrice{
			Province: v.Province.String,
			Week:     week,
			Year:     year,
			Price:    price,
		})
	}

	for _, v := range resourcePrices {
		resourceAggregate, index := findResourceAggregate(result, v.Province, v.Year, v.Week)
		if resourceAggregate == nil {
			result = append(result, model.ResourceDataAggregate{
				Province:   v.Province,
				Year:       v.Year,
				Week:       v.Week,
				Min:        v.Price,
				Max:        v.Price,
				Avg:        v.Price,
				Median:     v.Price,
				TotalPrice: v.Price,
				TotalData:  1,
				Prices:     []float64{v.Price},
			})
		} else {
			totalData := resourceAggregate.TotalData + 1
			totalPrice := resourceAggregate.TotalPrice + v.Price
			prices := append(resourceAggregate.Prices, v.Price)

			min := calcMin(resourceAggregate.Min, v.Price)
			max := calcMax(resourceAggregate.Max, v.Price)
			avg := calcAvg(totalPrice, totalData)

			median := resourceAggregate.Median
			if len(prices) > 0 {
				median = calcMedian(prices, totalData)
			}

			result = append(result[:index], result[index+1:]...)
			result = append(result, model.ResourceDataAggregate{
				Province:   v.Province,
				Year:       v.Year,
				Week:       v.Week,
				Min:        min,
				Max:        max,
				Avg:        avg,
				Median:     median,
				TotalPrice: totalPrice,
				TotalData:  totalData,
				Prices:     prices,
			})
		}
	}

	return result, nil
}

func findResourceAggregate(data []model.ResourceDataAggregate, province string, year, week int) (*model.ResourceDataAggregate, int) {
	var (
		res   *model.ResourceDataAggregate
		index int
	)
	for i, v := range data {
		if v.Province == province && v.Year == year && v.Week == week {
			res = &v
			index = i
			break
		}
	}

	return res, index
}

func calcMin(current float64, target float64) float64 {
	if current <= target {
		return current
	} else {
		return target
	}
}

func calcMax(current, target float64) float64 {
	if current >= target {
		return current
	} else {
		return target
	}
}

func calcAvg(totalPrice float64, totalData int) float64 {
	return totalPrice / float64(totalData)
}

func calcMedian(prices []float64, totalData int) float64 {
	sort.Float64s(prices)

	medianIndex := totalData - 1

	if totalData%2 > 0 {
		return prices[medianIndex/2]
	} else {
		a := prices[medianIndex/2]
		b := prices[(medianIndex/2)+1]

		return (a + b) / 2
	}
}
