// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package wallet

import (
	"time"
)

type FII struct {
	BrokerID    string     `json:"brokerId" bson:"brokerId" validate:"required"`
	Commission  float64    `json:"commission" bson:"commission"`
	Date        *time.Time `json:"date" bson:"date" validate:"required"`
	ID          string     `json:"id,omitempty" bson:"_id,omitempty"`
	ItemType    string     `json:"itemType" bson:"itemType" validate:"required"`
	PortfolioID string     `json:"portfolioId" bson:"portfolioId" validate:"required"`
	Price       float64    `json:"price" bson:"price" validate:"required"`
	Shares      float64    `json:"shares" bson:"shares" validate:"required"`
	Symbol      string     `json:"symbol" bson:"symbol" validate:"required"`
}

type FIIList []FII

func NewFII() *FII {
	return &FII{ItemType: "fiis"}
}
