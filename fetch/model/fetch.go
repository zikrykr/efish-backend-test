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
