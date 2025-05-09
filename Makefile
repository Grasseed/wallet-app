# Always execute these targets
.PHONY: build run test verify
# Makefile for wallet-app

APP_NAME=wallet-app

build:
	go build -o $(APP_NAME) .
	chmod +x $(APP_NAME)

run:
	./$(APP_NAME)

test:
	go test -v -count=1 ./test

verify:
	./$(APP_NAME) verify --mnemonic="abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about" --path="m/44'/0'/0'/0/0" --expect=1LqBGSKuX5yYUonjxT5qGfpUsXKYYWeabA
