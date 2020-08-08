setup:
	@go get github.com/codegangsta/gin
	@go get

setup-prod:
	@go get

run:
	@FINANCE_WALLETAPI_DEBUG=True gin -b finance-wallet-api -a 8889 -p 3001 -i

clean:
	@find . -name "*.swp" -delete

docs-update:
	@swag init

docker-build:
	@docker build . -t mfinance/finance-wallet-api

docker-run:
	@docker run -p 8889:8889 mfinance/finance-wallet-api
