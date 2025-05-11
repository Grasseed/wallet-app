package wallet

import (
	"github.com/tyler-smith/go-bip39"
)

// GenerateMnemonic 產生 BIP39 助記詞，預設為 12、18 或 24 個詞
func GenerateMnemonic(wordCount int) (string, error) {
	var entropySize int
	switch wordCount {
	case 12:
		entropySize = 128
	case 18:
		entropySize = 192
	case 24:
		entropySize = 256
	default:
		return "", bip39.ErrInvalidMnemonic
	}

	entropy, err := bip39.NewEntropy(entropySize)
	if err != nil {
		return "", err
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", err
	}
	return mnemonic, nil
}
