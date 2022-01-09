// Copyright (c) 2022, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package wallet

import (
	"time"
)

// STOCKS ONLY
type DividendItem struct {
	Date         *time.Time `json:"date" bson:"date" validate:"required"`
	DeclaredDate *time.Time `json:"declaredDate" bson:"declaredDate" validate:"required"`
	ItemType     string     `json:"type" bson:"type" validate:"required"`
	PayDate      *time.Time `json:"payDate" bson:"payDate" validate:"required"`
	Value        float64    `json:"value" bson:"value" validate:"required"`
}

type DividendOperations struct {
	Dividends []DividendItem `json:"dividends" bson:"dividends" validate:"required"`
	Symbol    string         `json:"symbol" bson:"symbol" validate:"required"`
}
