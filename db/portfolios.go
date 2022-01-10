// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package db

import (
	"encoding/json"
	"fmt"

	"github.com/mfinancecombr/finance-wallet-api/financeapi"
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func (m *mongoSession) getPositionsByItemType(itemType string, year int) ([]wallet.Position, error) {
	log.Debugf("[DB] Getting portfolio item %s", itemType)
	operationsSymbols, err := m.getOperationsSymbols(bson.M{"itemType": itemType})
	if err != nil {
		return nil, err
	}

	// Get all data by itemType
	query := ""
	for _, s := range operationsSymbols {
		query += fmt.Sprintf("symbols=%s&", s)
	}
	tempPosition := &map[string][]wallet.Position{}
	url := fmt.Sprintf("/%s/?%s", itemType, query)
	if err := financeapi.GetJSON(url, tempPosition); err != nil {
		log.Warnf("Error on get %s symbols: %v", itemType, err)
	}

	// Convert to map of symbols
	symbolsMap := map[string]wallet.Position{}
	for _, a := range *tempPosition {
		for _, item := range a {
			symbolsMap[item.Symbol] = item
		}
	}

	// FIXME
	urls := []string{}
	for _, s := range operationsSymbols {
		if itemType != "stocks" && itemType != "fiis" {
			continue
		}
		urlDividend := fmt.Sprintf("/%s/dividends/%s", itemType, s)
		urls = append(urls, urlDividend)
	}

	items := []wallet.Position{}
	for _, s := range operationsSymbols {
		symbol := s.(string)
		operations, err := m.getAllOperationsBySymbol(symbol, itemType, year)
		if err != nil {
			return nil, err
		}

		var position wallet.Position
		if val, ok := symbolsMap[symbol]; ok {
			position = val
		} else {
			position = wallet.Position{}
		}

		position.Symbol = symbol
		position.ItemType = itemType
		position.Operations = operations

		if len(urls) > 0 {
			results := financeapi.GetAsyncJSON(urls)
			tempDividends := wallet.DividendOperations{}
			urlDividend := fmt.Sprintf("/%s/dividends/%s", itemType, symbol)
			result := results[urlDividend]
			defer result.Response.Body.Close()
			// FIXME
			json.NewDecoder(result.Response.Body).Decode(&tempDividends)
			position.DividendOperations = tempDividends
		}

		position.Recalculate()
		items = append(items, position)
	}

	return items, nil
}

func (m *mongoSession) GetPortfolioData(portfolio *wallet.Portfolio, year int) error {
	log.Debug("[DB] GetPositions")
	itemTypes, err := m.getItemTypes()
	if err != nil {
		return err
	}
	portfolio.Items = map[string][]wallet.Position{}
	for _, itemType := range itemTypes {
		kind := itemType.(string)
		positions, err := m.getPositionsByItemType(kind, year)
		if err != nil {
			log.Errorf("[DB] Error on get portfolio items: %v", err)
			continue
		}
		portfolio.Items[kind] = positions
	}
	portfolio.Recalculate()
	return nil
}
