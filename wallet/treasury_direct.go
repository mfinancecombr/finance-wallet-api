// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package wallet

import "time"

type TreasuryDirect struct {
	TradableAsset
	DueDate           *time.Time `json:"dueDate" bson:"dueDate" validate:"required"`
	FixedInterestRate float64    `json:"fixedInterestRate" bson:"fixedInterestRate" validate:"required"`
}

type TreasuryDirectList []TreasuryDirect

func NewTreasuryDirect() *TreasuryDirect {
	return &TreasuryDirect{TradableAsset: TradableAsset{ItemType: "treasury-direct"}}
}
