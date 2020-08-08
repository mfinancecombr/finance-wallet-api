// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package wallet

import (
	"math"

	log "github.com/sirupsen/logrus"
)

type PortfolioItem struct {
	AveragePrice  float64       `json:"averagePrice" bson:"averagePrice"`
	BrokerID      string        `json:"brokerId" bson:"brokerId"`
	Change        float64       `json:"change" bson:"change"`
	ClosingPrice  float64       `json:"closingPrice" bson:"closingPrice"`
	Commission    float64       `json:"commission" bson:"commission"`
	CostBasics    float64       `json:"costBasics" bson:"costBasics"`
	Gain          float64       `json:"gain" bson:"gain"`
	OverallReturn float64       `json:"overallReturn" bson:"overallReturn"`
	ItemType      string        `json:"itemType" bson:"itemType"`
	LastPrice     float64       `json:"lastPrice" bson:"lastPrice"`
	LastYearHigh  float64       `json:"lastYearHigh" bson:"lastYearHigh"`
	LastYearLow   float64       `json:"lastYearLow" bson:"lastYearLow"`
	Name          string        `json:"name" bson:"name"`
	Purchases     PurchasesList `json:"purchases" bson:"purchases"`
	Sales         SalesList     `json:"sales" bson:"sales"`
	Sector        string        `json:"sector" bson:"sector"`
	Segment       string        `json:"segment" bson:"segment"`
	Shares        float64       `json:"shares" bson:"shares"`
	SubSector     string        `json:"subSector" bson:"subSector"`
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
	salesPrice := 0.0
	salesShares := 0.0

	// FIXME: duplicated
	for _, s := range pi.Sales {
		var salePrice float64
		var saleShares float64
		var saleCommission float64
		switch itemType := s.(type) {
		case *Stock:
			salePrice = s.(*Stock).Price
			saleShares = s.(*Stock).Shares
			saleCommission = s.(*Stock).Commission
		case *FII:
			salePrice = s.(*FII).Price
			saleShares = s.(*FII).Shares
			saleCommission = s.(*FII).Commission
		case *CertificateOfDeposit:
			salePrice = s.(*CertificateOfDeposit).Price
			saleShares = s.(*CertificateOfDeposit).Shares
			saleCommission = s.(*CertificateOfDeposit).Commission
		case *TreasuryDirect:
			salePrice = s.(*TreasuryDirect).Price
			saleShares = s.(*TreasuryDirect).Shares
			saleCommission = s.(*TreasuryDirect).Commission
		case *StockFund:
			salePrice = s.(*StockFund).Price
			saleShares = s.(*StockFund).Shares
			saleCommission = s.(*StockFund).Commission
		case *FICFI:
			salePrice = s.(*FICFI).Price
			saleShares = s.(*FICFI).Shares
			saleCommission = s.(*FICFI).Commission
		default:
			log.Errorf("Item type '%s' not found", itemType)
		}

		salesPrice += (salePrice * saleShares) + saleCommission
		salesShares += saleShares
		commission += saleCommission
	}

	purchasesPrice := 0.0
	purchasesShares := 0.0

	// FIXME: duplicated
	for _, p := range pi.Purchases {
		var purchasePrice float64
		var purchaseShares float64
		var purchaseCommission float64
		switch itemType := p.(type) {
		case *Stock:
			purchasePrice = p.(*Stock).Price
			purchaseShares = p.(*Stock).Shares
			purchaseCommission = p.(*Stock).Commission
		case *FII:
			purchasePrice = p.(*FII).Price
			purchaseShares = p.(*FII).Shares
			purchaseCommission = p.(*FII).Commission
		case *CertificateOfDeposit:
			purchasePrice = p.(*CertificateOfDeposit).Price
			purchaseShares = p.(*CertificateOfDeposit).Shares
			purchaseCommission = p.(*CertificateOfDeposit).Commission
		case *TreasuryDirect:
			purchasePrice = p.(*TreasuryDirect).Price
			purchaseShares = p.(*TreasuryDirect).Shares
			purchaseCommission = p.(*TreasuryDirect).Commission
		case *StockFund:
			purchasePrice = p.(*StockFund).Price
			purchaseShares = p.(*StockFund).Shares
			purchaseCommission = p.(*StockFund).Commission
		case *FICFI:
			purchasePrice = p.(*FICFI).Price
			purchaseShares = p.(*FICFI).Shares
			purchaseCommission = p.(*FICFI).Commission
		default:
			log.Errorf("Item type '%s' not found", itemType)
		}

		purchasesPrice += (purchasePrice * purchaseShares) + purchaseCommission
		purchasesShares += purchaseShares
		commission += commission
	}

	pi.Shares = purchasesShares - salesShares
	if pi.Shares > 0 {
		pi.Commission = roundFloatTwoDecimalPlaces(commission)
		pi.CostBasics = roundFloatTwoDecimalPlaces(purchasesPrice - salesPrice + commission)
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
