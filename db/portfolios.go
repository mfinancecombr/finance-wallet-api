// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package db

import (
	"github.com/mfinancecombr/finance-wallet-api/financeapi"
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func (m *mongoSession) getPortfolioItem(itemType string, year int) (map[string]wallet.PortfolioItem, error) {
	log.Debugf("[DB] Getting portfolio item %s", itemType)
	operationsSymbols, err := m.getOperationsSymbols(bson.M{"itemType": itemType})
	if err != nil {
		return nil, err
	}

	items := map[string]wallet.PortfolioItem{}
	for _, s := range operationsSymbols {
		symbol := s.(string)
		portfolioItem := &wallet.PortfolioItem{}
		// FIXME: one request
		if err := financeapi.GetJSON("/"+itemType+"/"+symbol, portfolioItem); err != nil {
			log.Errorf("Error on get stock item: %s", err)
		}

		operations, err := m.getAllOperationsBySymbol(symbol, itemType, year)
		if err != nil {
			return nil, err
		}

		// FIXME
		broker := ""
		if len(operations) > 0 {
			operation := operations[0]
			if operation != nil {
				broker = operation.GetBrokerSlug()
			}
		}

		portfolioItem.BrokerSlug = broker
		portfolioItem.ItemType = itemType
		portfolioItem.Operations = operations
		portfolioItem.Recalculate()
		items[symbol] = *portfolioItem
	}

	return items, nil
}

// FIXME
func (m *mongoSession) GetPortfolioItems(portfolio *wallet.Portfolio, year int) error {
	log.Debug("[DB] GetPortfolioItems")

	slugs := []string{
		"certificates-of-deposit",
		"ficfi",
		"fiis",
		"stocks",
		"stocks-funds",
		"treasuries-direct",
	}

	portfolio.Items = map[string]wallet.PortfolioItem{}
	for _, slug := range slugs {
		stocks, err := m.getPortfolioItem(slug, year)
		if err != nil {
			log.Errorf("[DB] Error on get portfolio items: %v", err)
			continue
		}
		for symbol, portfolioItem := range stocks {
			portfolio.Items[symbol] = portfolioItem
		}
	}

	portfolio.Recalculate()

	return nil
}
