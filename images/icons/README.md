# Icons

## Icons via Status Invest

(So, don't share icons file)

```bash
URL="https://statusinvest.com.br/acao/companiesnavigation?page=1&size=400"
IMG_URL="https://statusinvest.com.br/img/company/avatar"
for X in $(curl -s $URL | jq -c '.[] | {"symbol": .url, "id": .companyId}'); do
    SYMBOL=$(echo "$X" | jq .symbol | sed 's#/acoes/##' | sed 's/"//g');
    COMPANY_ID=$(echo "$X" | jq .id);
    curl -s $IMG_URL/$COMPANY_ID.jpg --output $SYMBOL.jpg;
done
```
