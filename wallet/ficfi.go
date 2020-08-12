// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package wallet

type FICFI struct {
    TradableAsset
}

type FICFIList []FICFI

func NewFICFI() *FICFI {
	return &FICFI{TradableAsset:TradableAsset{ItemType: "ficfi"}}
}