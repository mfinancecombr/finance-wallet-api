// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package wallet

type FII struct {
    TradableAsset
}

type FIIList []FII

func NewFII() *FII {
	return &FII{TradableAsset:TradableAsset{ItemType: "fiis"}}
}