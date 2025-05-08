package test

import (
	"testing"

	"github.com/Grasseed/wallet-app/wallet"

	"github.com/tyler-smith/go-bip39"
)

func TestGenerateMnemonic(t *testing.T) {
	mnemonic, err := wallet.GenerateMnemonic(12)
	if err != nil {
		t.Fatalf("Failed to generate mnemonic: %v", err)
	}
	if !bip39.IsMnemonicValid(mnemonic) {
		t.Errorf("Generated mnemonic is invalid: %s", mnemonic)
	}
	t.Logf("Generated mnemonic: %s", mnemonic)
}

func TestDeriveBTCAddress(t *testing.T) {
	// 測試助記詞與預期的地址（實際上可以根據助記詞產出固定的結果）
	mnemonic := "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about"
	path := "m/44'/0'/0'/0/0"
	address, wif, err := wallet.DeriveBTCAddress(mnemonic, path)
	if err != nil {
		t.Fatalf("Failed to derive address: %v", err)
	}
	if address == "" || wif == "" {
		t.Errorf("Address or WIF is empty")
	}
	t.Logf("Derived address: %s", address)
	t.Logf("Derived WIF: %s", wif)
}

func TestDeriveRandomMnemonic(t *testing.T) {
	mnemonic, err := wallet.GenerateMnemonic(12)
	if err != nil {
		t.Fatalf("Failed to generate mnemonic: %v", err)
	}
	if !bip39.IsMnemonicValid(mnemonic) {
		t.Fatalf("Generated mnemonic is invalid: %s", mnemonic)
	}
	path := "m/44'/0'/0'/0/0"
	address, wif, err := wallet.DeriveBTCAddress(mnemonic, path)
	if err != nil {
		t.Fatalf("Failed to derive address from random mnemonic: %v", err)
	}
	if address == "" || wif == "" {
		t.Errorf("Derived address or WIF from random mnemonic is empty")
	}
	t.Logf("Random mnemonic: %s", mnemonic)
	t.Logf("Derived address: %s", address)
	t.Logf("Derived WIF: %s", wif)
}
