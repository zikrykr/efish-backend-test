package model

import "gopkg.in/guregu/null.v4"

type ResourceDataResponse struct {
	UUID      null.String `json:"uuid"`
	Commodity null.String `json:"komoditas"`
	Province  null.String `json:"area_provinsi"`
	City      null.String `json:"area_kota"`
	Size      null.String `json:"size"`
	Price     null.String `json:"price"`
	ParsedAt  null.String `json:"tgl_parsed"`
	Timestamp null.String `json:"timestamp"`
}

type ResourceData struct {
	UUID      null.String `json:"uuid"`
	Commodity null.String `json:"komoditas"`
	Province  null.String `json:"area_provinsi"`
	City      null.String `json:"area_kota"`
	Size      null.String `json:"size"`
	Price     null.String `json:"price"`
	ParsedAt  null.String `json:"tgl_parsed"`
	Timestamp null.String `json:"timestamp"`
	PriceUSD  null.String `json:"price_usd"`
}

type ResourceDataAggregate struct {
	Province   string    `json:"area_provinsi"`
	Week       int       `json:"minggu"`
	Year       int       `json:"tahun"`
	Min        float64   `json:"min"`
	Max        float64   `json:"max"`
	Avg        float64   `json:"avg"`
	Median     float64   `json:"median"`
	TotalPrice float64   `json:"totalPrice"`
	TotalData  int       `json:"totalData"`
	Prices     []float64 `json:"prices"`
}

type ResourcePrice struct {
	Province string  `json:"area_provinsi"`
	Week     int     `json:"minggu"`
	Year     int     `json:"tahun"`
	Price    float64 `json:"price"`
}
