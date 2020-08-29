// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package wallet

import (
	"math"
)

type Portfolio struct {
	CostBasis     float64               `json:"costBasis" bson:"costBasis,omitempty"`
	Gain          float64               `json:"gain" bson:"gain,omitempty"`
	ID            string                `json:"id,omitempty" bson:"_id,omitempty"`
	Items         map[string][]Position `json:"items" bson:"items,omitempty"`
	Name          string                `json:"name" bson:"name" validate:"required"`
	OverallReturn float64               `json:"overallReturn" bson:"overallReturn,omitempty"`
	Slug          string                `json:"slug" bson:"slug" validate:"required"`
}

func (s Portfolio) GetCollectionName() string {
	return "portfolios"
}

func (s Portfolio) GetItemType() string {
	return ""
}

func roundFloatTwoDecimalPlaces(n float64) float64 {
	return math.Ceil(n*100) / 100
}

func (p *Portfolio) Recalculate() {
	if len(p.Items) == 0 {
		return
	}

	costBasis := 0.0
	gain := 0.0
	for _, items := range p.Items {
		for _, item := range items {
			costBasis += item.CostBasis
			gain += item.Gain
		}
	}

	p.CostBasis = roundFloatTwoDecimalPlaces(costBasis)
	p.Gain = roundFloatTwoDecimalPlaces(gain)
	p.OverallReturn = roundFloatTwoDecimalPlaces(p.Gain * 100 / p.CostBasis)
}
