// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package wallet

import "time"

type Tradable interface {
	GetBrokerSlug() string
	GetComission() float64
	GetDate() *time.Time
	GetItemType() string
	GetPrice() float64
	GetShares() float64
	GetType() string
}
