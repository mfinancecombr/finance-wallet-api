// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package wallet

import (
	"time"
)

type Stock struct {
	BrokerSlug    string     `json:"brokerSlug" bson:"brokerSlug" validate:"required"`
	Commission    float64    `json:"commission" bson:"commission"`
	Date          *time.Time `json:"date" bson:"date" validate:"required"`
	ID            string     `json:"id,omitempty" bson:"_id,omitempty"`
	ItemType      string     `json:"itemType" bson:"itemType" validate:"required"`
	PortfolioSlug string     `json:"portfolioSlug" bson:"portfolioSlug" validate:"required"`
	Price         float64    `json:"price" bson:"price" validate:"required"`
	Shares        float64    `json:"shares" bson:"shares" validate:"required"`
	Symbol        string     `json:"symbol" bson:"symbol" validate:"required"`
	Type          string     `json:"type" bson:"type" validate:"required"`
}

type StockList []Stock

const StockItemType = "stocks"

func NewStock() *Stock {
	return &Stock{ItemType: StockItemType}
}

func (s Stock) GetPrice() float64 {
	return s.Price
}

func (s Stock) GetShares() float64 {
	return s.Shares
}

func (s Stock) GetComission() float64 {
	return s.Commission
}

func (s Stock) GetType() string {
	return s.Type
}

func (s Stock) GetBrokerSlug() string {
	return s.BrokerSlug
}

func (s Stock) GetCollectionName() string {
	return "operations"
}

func (s Stock) GetItemType() string {
	return StockItemType
}

func (s Stock) GetDate() *time.Time {
	return s.Date
}
