# Always execute these targets
.PHONY: build run test verify mnemonic mnemonic12 mnemonic18 mnemonic24
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

mnemonic:
	go run main.go mnemonic --words=12

mnemonic12:
	go run main.go mnemonic --words=12

mnemonic18:
	go run main.go mnemonic --words=18

mnemonic24:
	go run main.go mnemonic --words=24
