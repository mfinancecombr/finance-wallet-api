// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package wallet

import (
	"math"
)

type PortfolioItem struct {
	AveragePrice  float64        `json:"averagePrice" bson:"averagePrice"`
	BrokerID      string         `json:"brokerId" bson:"brokerId"`
	Change        float64        `json:"change" bson:"change"`
	ClosingPrice  float64        `json:"closingPrice" bson:"closingPrice"`
	Commission    float64        `json:"commission" bson:"commission"`
	CostBasics    float64        `json:"costBasics" bson:"costBasics"`
	Gain          float64        `json:"gain" bson:"gain"`
	OverallReturn float64        `json:"overallReturn" bson:"overallReturn"`
	ItemType      string         `json:"itemType" bson:"itemType"`
	LastPrice     float64        `json:"lastPrice" bson:"lastPrice"`
	LastYearHigh  float64        `json:"lastYearHigh" bson:"lastYearHigh"`
	LastYearLow   float64        `json:"lastYearLow" bson:"lastYearLow"`
	Name          string         `json:"name" bson:"name"`
	Operations    OperationsList `json:"operations" bson:"operations"`
	Sector        string         `json:"sector" bson:"sector"`
	Segment       string         `json:"segment" bson:"segment"`
	Shares        float64        `json:"shares" bson:"shares"`
	SubSector     string         `json:"subSector" bson:"subSector"`
}

type Portfolio struct {
	CostBasics    float64                  `json:"costBasics" bson:"costBasics"`
	Gain          float64                  `json:"gain" bson:"gain"`
	ID            string                   `json:"id" bson:"_id" validate:"required"`
	Items         map[string]PortfolioItem `json:"items" bson:"items"`
	Name          string                   `json:"name" bson:"name" validate:"required"`
	OverallReturn float64                  `json:"overallReturn" bson:"overallReturn"`
}

func roundFloatTwoDecimalPlaces(n float64) float64 {
	return math.Ceil(n*100) / 100
}

func (p *Portfolio) Recalculate() {
	if len(p.Items) == 0 {
		return
	}

	costBasics := 0.0
	gain := 0.0
	for _, item := range p.Items {
		costBasics += item.CostBasics
		gain += item.Gain
	}

	p.CostBasics = roundFloatTwoDecimalPlaces(costBasics)
	p.Gain = roundFloatTwoDecimalPlaces(gain)
	p.OverallReturn = roundFloatTwoDecimalPlaces(p.Gain * 100 / p.CostBasics)
}

func (pi *PortfolioItem) Recalculate() {
	commission := 0.0
	totalPrice := 0.0
	totalShares := 0.0

	for _, s := range pi.Operations {
		var operationPrice = s.(TradableAsset).Price
		var operationShares = s.(TradableAsset).Shares
		var operationCommission = s.(TradableAsset).Commission
		var operationType = s.(TradableAsset).Type
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
		pi.CostBasics = roundFloatTwoDecimalPlaces(totalPrice)
		pi.AveragePrice = roundFloatTwoDecimalPlaces(pi.CostBasics / pi.Shares)

		// FIXME
		if pi.ItemType == "stocks" || pi.ItemType == "fiis" {
			gain := (pi.Shares * pi.LastPrice) - pi.CostBasics
			pi.Gain = roundFloatTwoDecimalPlaces(gain)
			pi.OverallReturn = roundFloatTwoDecimalPlaces((gain * 100) / pi.CostBasics)
		} else {
			pi.Gain = 0
			pi.OverallReturn = 0
		}
	}
}
