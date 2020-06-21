// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package db

import (
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *mongoSession) InsertPortfolio(d interface{}) (*mongo.InsertOneResult, error) {
	log.Debug("[DB] InsertPortfolio")
	return m.collection.InsertOne(portfoliosCollection, d)
}

func (m *mongoSession) DeletePortfolioByID(id string) (*mongo.DeleteResult, error) {
	log.Debug("[DB] DeletePortfolioByID")
	q := bson.M{"_id": id}
	return m.collection.DeleteOne(portfoliosCollection, q)
}

func (m *mongoSession) UpdatePortfolio(id string, d interface{}) (*mongo.UpdateResult, error) {
	log.Debug("[DB] UpdatePortfolio")
	f := bson.D{{"_id", id}}
	u := bson.D{{"$set", d}}
	return m.collection.UpdateOne(portfoliosCollection, f, u)
}

func (m *mongoSession) GetAllPortfolios() ([]wallet.Portfolio, error) {
	log.Debug("[DB] GetAllPortfolios")
	results, err := m.collection.FindAll(portfoliosCollection, bson.M{})
	if err != nil {
		return nil, err
	}
	portfolioTemp := []wallet.Portfolio{}
	for _, result := range results {
		bsonBytes, _ := bson.Marshal(result)
		portfolio := wallet.Portfolio{}
		bson.Unmarshal(bsonBytes, &portfolio)
		portfolioTemp = append(portfolioTemp, portfolio)
	}
	return portfolioTemp, nil
}

func (m *mongoSession) GetPortfolioByID(id string) (*wallet.Portfolio, error) {
	log.Debug("[DB] GetPortfolioByID")
	h := &wallet.Portfolio{}
	query := bson.M{"_id": id}
	if err := m.collection.FindOne(portfoliosCollection, query, h); err != nil {
		return nil, err
	}
	if h.ID == "" {
		return nil, nil
	}
	return h, nil
}

// FIXME: maybe unnecessary
func contains(s []interface{}, e interface{}) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// FIXME: maybe unnecessary
func mergeSlices(a, b []interface{}) []interface{} {
	for _, valB := range b {
		if !contains(a, valB) {
			a = append(a, valB)
		}
	}

	return a
}

func (m *mongoSession) getPortfolioItem(itemType string, year, month int) (map[string]wallet.PortfolioItem, error) {
	log.Debugf("[DB] Getting portfolio item %s", itemType)
	purchasesSymbols, err := m.getPurchasesSymbols(bson.M{"itemType": itemType})
	if err != nil {
		return nil, err
	}

	// FIXME: maybe unnecessary
	salesSymbols, err := m.getSalesSymbols(bson.M{"itemType": itemType})
	if err != nil {
		return nil, err
	}

	// FIXME: maybe unnecessary
	uniqueSymbols := mergeSlices(purchasesSymbols, salesSymbols)

	items := map[string]wallet.PortfolioItem{}
	for _, s := range uniqueSymbols {
		symbol := s.(string)

		purchases, err := m.getAllPurchasesBySymbol(symbol, itemType, year, month)
		if err != nil {
			return nil, err
		}

		sales, err := m.getAllSalesBySymbol(symbol, itemType, year, month)
		if err != nil {
			return nil, err
		}

		// FIXME
		broker := ""
		if len(purchases) > 0 {
			purchase := purchases[0]
			if purchase != nil {
				// FIXME: duplicated
				switch itemType {
				case "stocks":
					broker = purchase.(*wallet.Stock).BrokerID
				case "fiis":
					broker = purchase.(*wallet.FII).BrokerID
				case "certificates-of-deposit":
					broker = purchase.(*wallet.CertificateOfDeposit).BrokerID
				case "treasuries-direct":
					broker = purchase.(*wallet.TreasuryDirect).BrokerID
				case "stocks-funds":
					broker = purchase.(*wallet.StockFund).BrokerID
				case "ficfi":
					broker = purchase.(*wallet.FICFI).BrokerID
				default:
					log.Errorf("Item type '%s' not found", itemType)
				}
			}
		}

		portfolioItem := &wallet.PortfolioItem{}
		portfolioItem.BrokerID = broker
		portfolioItem.ItemType = itemType
		portfolioItem.Purchases = purchases
		portfolioItem.Sales = sales
		items[symbol] = *portfolioItem
	}

	return items, nil
}

// FIXME
func (m *mongoSession) GetPortfolioItems(portfolio *wallet.Portfolio, year, month int) error {
	log.Debug("[DB] GetPortfolioItems")

	purchasesItemType, err := m.GetPurchasesItemType()
	if err != nil {
		return err
	}

	salesItemType, err := m.GetSalesItemType()
	if err != nil {
		return err
	}

	slugs := mergeSlices(purchasesItemType, salesItemType)

	portfolio.Items = map[string]wallet.PortfolioItem{}
	for _, slug := range slugs {
		stocks, err := m.getPortfolioItem(slug.(string), year, month)
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
