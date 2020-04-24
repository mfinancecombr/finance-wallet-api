setup:
	@go get github.com/codegangsta/gin
	@go get

setup-prod:
	@go get

run:
	@FINANCE_WALLETAPI_DEBUG=True gin -b finance-wallet-api -a 8889 -p 3001 -i

clean:
	@find . -name "*.swp" -delete
