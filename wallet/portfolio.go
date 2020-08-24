// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package wallet

import (
	"math"
)

type PortfolioItem struct {
	AveragePrice  float64        `json:"averagePrice" bson:"averagePrice"`
	BrokerSlug    string         `json:"brokerSlug" bson:"brokerSlug"`
	Change        float64        `json:"change" bson:"change"`
	ClosingPrice  float64        `json:"closingPrice" bson:"closingPrice"`
	Commission    float64        `json:"commission" bson:"commission"`
	CostBasis     float64        `json:"costBasis" bson:"costBasis"`
	Gain          float64        `json:"gain" bson:"gain"`
	ItemType      string         `json:"itemType" bson:"itemType"`
	LastPrice     float64        `json:"lastPrice" bson:"lastPrice"`
	LastYearHigh  float64        `json:"lastYearHigh" bson:"lastYearHigh"`
	LastYearLow   float64        `json:"lastYearLow" bson:"lastYearLow"`
	Name          string         `json:"name" bson:"name"`
	Operations    OperationsList `json:"operations" bson:"operations"`
	OverallReturn float64        `json:"overallReturn" bson:"overallReturn"`
	Sector        string         `json:"sector" bson:"sector"`
	Segment       string         `json:"segment" bson:"segment"`
	Shares        float64        `json:"shares" bson:"shares"`
	SubSector     string         `json:"subSector" bson:"subSector"`
}

type Portfolio struct {
	CostBasis     float64                  `json:"costBasis" bson:"costBasis,omitempty"`
	Gain          float64                  `json:"gain" bson:"gain,omitempty"`
	ID            string                   `json:"id,omitempty" bson:"_id,omitempty"`
	Items         map[string]PortfolioItem `json:"items" bson:"items,omitempty"`
	Name          string                   `json:"name" bson:"name" validate:"required"`
	OverallReturn float64                  `json:"overallReturn" bson:"overallReturn,omitempty"`
	Slug          string                   `json:"slug" bson:"slug" validate:"required"`
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
	for _, item := range p.Items {
		costBasis += item.CostBasis
		gain += item.Gain
	}

	p.CostBasis = roundFloatTwoDecimalPlaces(costBasis)
	p.Gain = roundFloatTwoDecimalPlaces(gain)
	p.OverallReturn = roundFloatTwoDecimalPlaces(p.Gain * 100 / p.CostBasis)
}

func (pi *PortfolioItem) Recalculate() {
	commission := 0.0
	totalPrice := 0.0
	totalShares := 0.0

	for _, s := range pi.Operations {
		var operationPrice = s.GetPrice()
		var operationShares = s.GetShares()
		var operationCommission = s.GetComission()
		var operationType = s.(Tradable).GetType()
		if operationType == "purchase" {
			totalPrice += (operationPrice * operationShares) + operationCommission
			totalShares += operationShares
			commission += operationCommission
		} else {
			// To properly calculate the average price we need to remove from the cost basis
			// based on the average price at the time of the sale.
			totalPrice -= (totalPrice / totalShares) * operationShares
			totalPrice += operationCommission
			totalShares -= operationShares
			commission += operationCommission
		}
	}

	pi.Shares = totalShares
	if pi.Shares > 0 {
		pi.Commission = roundFloatTwoDecimalPlaces(commission)
		pi.CostBasis = roundFloatTwoDecimalPlaces(totalPrice)
		pi.AveragePrice = roundFloatTwoDecimalPlaces(pi.CostBasis / pi.Shares)

		// FIXME
		if pi.ItemType == "stocks" || pi.ItemType == "fiis" {
			gain := (pi.Shares * pi.LastPrice) - pi.CostBasis
			pi.Gain = roundFloatTwoDecimalPlaces(gain)
			pi.OverallReturn = roundFloatTwoDecimalPlaces((gain * 100) / pi.CostBasis)
		} else {
			pi.Gain = 0
			pi.OverallReturn = 0
		}
	}
}
