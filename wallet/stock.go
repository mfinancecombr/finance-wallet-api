// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package wallet

type Stock struct {
	TradableAsset
}

type StockList []Stock

func NewStock() *Stock {
	return &Stock{TradableAsset: TradableAsset{ItemType: "stocks"}}
}
