// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package wallet

type StockFund struct {
	TradableAsset
}

type StockFundList []StockFund

func NewStockFund() *StockFund {
	return &StockFund{TradableAsset: TradableAsset{ItemType: "stocks-funds"}}
}
