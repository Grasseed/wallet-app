package test

import (
	"testing"

	"github.com/Grasseed/wallet-app/wallet"

	"github.com/tyler-smith/go-bip39"
)

var mnemonic string

func TestGenerateMnemonic(t *testing.T) {
	wordCounts := []int{12, 18, 24}

	for _, wc := range wordCounts {
		mnemonic, err := wallet.GenerateMnemonic(wc)
		if err != nil {
			t.Errorf("Failed to generate %d-word mnemonic: %v", wc, err)
			continue
		}
		if !bip39.IsMnemonicValid(mnemonic) {
			t.Errorf("Generated %d-word mnemonic is invalid: %s", wc, mnemonic)
		}
		t.Logf("Generated %d-word mnemonic: %s", wc, mnemonic)
	}
}

func TestDeriveBTCAddress(t *testing.T) {
	if mnemonic == "" {
		t.Skip("Mnemonic not initialized")
	}
	path := "m/44'/0'/0'/0/0"
	address, wif, err := wallet.DeriveBTCAddress(mnemonic, path, "p2pkh")
	if err != nil {
		t.Fatalf("Failed to derive address: %v", err)
	}
	if address == "" || wif == "" {
		t.Errorf("Address or WIF is empty")
	}
	t.Logf("Derived address: %s", address)
	t.Logf("Derived WIF: %s", wif)
}

func TestDeriveRandomPathAddress(t *testing.T) {
	if mnemonic == "" {
		t.Skip("Mnemonic not initialized")
	}
	path := "m/44'/0'/0'/0/5"
	address, wif, err := wallet.DeriveBTCAddress(mnemonic, path, "p2pkh")
	if err != nil {
		t.Fatalf("Failed to derive address from random path: %v", err)
	}
	if address == "" || wif == "" {
		t.Errorf("Derived address or WIF from random path is empty")
	}
	t.Logf("Random path address: %s", address)
	t.Logf("Derived WIF: %s", wif)
}

func TestDeriveSegwitAddress(t *testing.T) {
	if mnemonic == "" {
		t.Skip("Mnemonic not initialized")
	}
	path := "m/84'/0'/0'/0/0"
	address, wif, err := wallet.DeriveBTCAddress(mnemonic, path, "segwit")
	if err != nil {
		t.Fatalf("Failed to derive segwit address: %v", err)
	}
	if address == "" || wif == "" {
		t.Errorf("Segwit address or WIF is empty")
	}
	t.Logf("Segwit address: %s", address)
	t.Logf("Segwit WIF: %s", wif)
}

func TestDeriveP2SHAddress(t *testing.T) {
	if mnemonic == "" {
		t.Skip("Mnemonic not initialized")
	}
	path := "m/49'/0'/0'/0/0"
	address, wif, err := wallet.DeriveBTCAddress(mnemonic, path, "p2sh")
	if err != nil {
		t.Fatalf("Failed to derive p2sh address: %v", err)
	}
	if address == "" || wif == "" {
		t.Errorf("P2SH address or WIF is empty")
	}
	t.Logf("P2SH address: %s", address)
	t.Logf("P2SH WIF: %s", wif)
}
