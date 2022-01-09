// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package wallet

type Position struct {
	AveragePrice       float64            `json:"averagePrice" bson:"averagePrice"`
	Change             float64            `json:"change" bson:"change"`
	ClosingPrice       float64            `json:"closingPrice" bson:"closingPrice"`
	Commission         float64            `json:"commission" bson:"commission"`
	CostBasis          float64            `json:"costBasis" bson:"costBasis"`
	Dividends          float64            `json:"dividends" bson:"dividends"`
	DividendOperations DividendOperations `json:"dividendOperations" bson:"dividendOperations"`
	Gain               float64            `json:"gain" bson:"gain"`
	ItemType           string             `json:"itemType" bson:"itemType"`
	LastPrice          float64            `json:"lastPrice" bson:"lastPrice"`
	LastYearHigh       float64            `json:"lastYearHigh" bson:"lastYearHigh"`
	LastYearLow        float64            `json:"lastYearLow" bson:"lastYearLow"`
	Name               string             `json:"name" bson:"name"`
	Operations         OperationsList     `json:"operations" bson:"operations"`
	OverallReturn      float64            `json:"overallReturn" bson:"overallReturn"`
	Sector             string             `json:"sector" bson:"sector"`
	Segment            string             `json:"segment" bson:"segment"`
	Shares             float64            `json:"shares" bson:"shares"`
	SubSector          string             `json:"subSector" bson:"subSector"`
	Symbol             string             `json:"symbol" bson:"symbol" validate:"required"`
}

func (pi *Position) Recalculate() {
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
			// To properly calculate the average price we need to remove from
			// the cost basis based on the average price at the time of the
			// sale.
			totalPrice -= (totalPrice / totalShares) * operationShares
			totalPrice += operationCommission
			totalShares -= operationShares
			commission += operationCommission
		}
		// FIXME: duplicate purchase and sale
		var operationDate = s.GetDate()
		var operationItemType = s.GetItemType()
		newDividendItem := []DividendItem{}
		for _, dividendOperation := range pi.DividendOperations.Dividends {
			// FIXME
			if (operationItemType == "stocks" && dividendOperation.Date == nil) ||
				(operationItemType == "fiis" && dividendOperation.DeclaredDate == nil) {
				continue
			}

			//FIXME same year?!
			if operationItemType == "stocks" &&
				(dividendOperation.Date.After(*operationDate) ||
					dividendOperation.Date.Equal(*operationDate)) ||
				operationItemType == "fiis" &&
					(dividendOperation.DeclaredDate.After(*operationDate) ||
						dividendOperation.DeclaredDate.Equal(*operationDate)) {
				pi.Dividends += (dividendOperation.Value * operationShares)
				newDividendItem = append(newDividendItem, dividendOperation)
			}
		}
		// filter dividends
		pi.DividendOperations.Dividends = newDividendItem
		pi.Dividends = roundFloatTwoDecimalPlaces(pi.Dividends)
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
