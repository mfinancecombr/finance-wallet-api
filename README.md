# Finance Wallet

## Dependencies

* golang
* mongodb

### Installing dependencies

```bash
make setup
```

## Run it

```bash
make run
```

## Examples:

* Adding portfolio:
```curlrc
curl \
  http://localhost:8889/api/v1/portfolios \
  -X POST \
  -H 'Content-Type: application/json' \
  -d '{"id": "default", "name": "Default"}'
```

* Adding brokers:
```curlrc
curl \
  http://localhost:8889/api/v1/brokers \
  -X POST \
  -H 'Content-Type: application/json' \
  -d '{"id": "clear", "name": "CLEAR"}'
```

* Adding stocks purchases:
```curlrc
curl \
  http://localhost:8889/api/v1/stocks/purchases \
  -X POST \
  -H 'Content-Type: application/json' \
  -d '{
    "portfolioId": "default", "symbol": "PETR4", "brokerId": "clear",
    "shares": 500, "price": 10, "date": "2020-04-24T00:00:00Z"}'
```

* Adding stocks sales:
```curlrc
curl \
  http://localhost:8889/api/v1/stocks/sales \
  -X POST \
  -H 'Content-Type: application/json' \
  -d '{
    "portfolioId": "default", "symbol": "PETR4", "brokerId": "clear",
    "shares": 100, "price": 15, "date": "2020-06-30T00:00:00Z"}'
```

## Third Party

Favicon uses a picture from [icon-library.com][icon-library]
licensed under [CC0 Public Domain Licence][cco].

[icon-library]: http://icon-library.com/icon/icon-finance-15.html
[cco]: https://creativecommons.org/share-your-work/public-domain/cc0/
