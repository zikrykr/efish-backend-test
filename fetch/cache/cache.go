package cache

import (
	"context"
	"fetch/api/currency"
	"log"
	"net/http"
	"time"
)

var CurrencyConverterResult float64

type cache struct {
	currencyService currency.CurrencyService
}

func NewCache() Cache {
	return &cache{
		currencyService: currency.NewService(&http.Client{Timeout: 10 * time.Second}),
	}
}

type Cache interface {
	Start(t *time.Ticker)
}

func (c *cache) Start(t *time.Ticker) {
	c.Store()

	for {
		select {
		case <-t.C:
			c.Store()
		}

	}
}

func (c *cache) Store() {
	ctx := context.Background()
	result, err := c.currencyService.GetCurrencyConverter(ctx)
	if err != nil {
		log.Println(err)
		return
	}

	CurrencyConverterResult = result.Value
}
