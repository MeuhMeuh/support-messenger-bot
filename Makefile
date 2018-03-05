.PHONY: ng start

ng:
	ngrok http --subdomain=meuhmeuh 3000
start:
	go run *.go