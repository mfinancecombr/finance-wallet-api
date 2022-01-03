// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package wallet

import (
	"time"
)

type TreasuryDirect struct {
	BrokerSlug        string     `json:"brokerSlug" bson:"brokerSlug" validate:"required"`
	Commission        float64    `json:"commission" bson:"commission"`
	Date              *time.Time `json:"date" bson:"date" validate:"required"`
	FixedInterestRate float64    `json:"fixedInterestRate" bson:"fixedInterestRate" validate:"required"`
	ID                string     `json:"id,omitempty" bson:"_id,omitempty"`
	ItemType          string     `json:"itemType" bson:"itemType" validate:"required"`
	PortfolioSlug     string     `json:"portfolioSlug" bson:"portfolioSlug" validate:"required"`
	Price             float64    `json:"price" bson:"price" validate:"required"`
	Shares            float64    `json:"shares" bson:"shares" validate:"required"`
	Symbol            string     `json:"symbol" bson:"symbol" validate:"required"`
	Type              string     `json:"type" bson:"type" validate:"required"`
}

type TreasuryDirectList []TreasuryDirect

const TreasuryDirectItemType = "treasury-direct"

func NewTreasuryDirect() *TreasuryDirect {
	return &TreasuryDirect{ItemType: TreasuryDirectItemType}
}

func (s TreasuryDirect) GetPrice() float64 {
	return s.Price
}

func (s TreasuryDirect) GetShares() float64 {
	return s.Shares
}

func (s TreasuryDirect) GetComission() float64 {
	return s.Commission
}

func (s TreasuryDirect) GetType() string {
	return s.Type
}

func (s TreasuryDirect) GetBrokerSlug() string {
	return s.BrokerSlug
}

func (s TreasuryDirect) GetCollectionName() string {
	return "operations"
}

func (s TreasuryDirect) GetItemType() string {
	return TreasuryDirectItemType
}
