package wallet

import (
	"fmt"
	"strings"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/tyler-smith/go-bip39"
)

// DeriveBTCAddress 依據助記詞與 BIP44 路徑導出 BTC P2PKH 地址與 WIF 私鑰
func DeriveBTCAddress(mnemonic, path, addrType string) (string, string, error) {
	// 驗證助記詞合法性
	if !bip39.IsMnemonicValid(mnemonic) {
		return "", "", fmt.Errorf("invalid mnemonic")
	}

	// 產生種子
	seed := bip39.NewSeed(mnemonic, "")

	// 建立 master key
	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return "", "", err
	}

	// 處理 BIP44 派生路徑，例如 "m/44'/0'/0'/0/0"
	segments := strings.Split(path, "/")[1:] // 移除 m
	key := masterKey
	for _, seg := range segments {
		hardened := false
		if strings.HasSuffix(seg, "'") {
			hardened = true
			seg = strings.TrimRight(seg, "'")
		}
		index, err := parseUint32(seg)
		if err != nil {
			return "", "", err
		}
		if hardened {
			index += hdkeychain.HardenedKeyStart
		}
		key, err = key.Derive(index)
		if err != nil {
			return "", "", err
		}
	}

	// 取得 EC 私鑰
	privKey, err := key.ECPrivKey()
	if err != nil {
		return "", "", err
	}
	wif, err := btcutil.NewWIF(privKey, &chaincfg.MainNetParams, true)
	if err != nil {
		return "", "", err
	}
	pubKey := privKey.PubKey()

	if addrType == "segwit" {
		witnessProg := btcutil.Hash160(pubKey.SerializeCompressed())
		addr, err := btcutil.NewAddressWitnessPubKeyHash(witnessProg, &chaincfg.MainNetParams)
		if err != nil {
			return "", "", err
		}
		return addr.EncodeAddress(), wif.String(), nil
	} else {
		pubKeyHash := btcutil.Hash160(pubKey.SerializeCompressed())
		addr, err := btcutil.NewAddressPubKeyHash(pubKeyHash, &chaincfg.MainNetParams)
		if err != nil {
			return "", "", err
		}
		return addr.EncodeAddress(), wif.String(), nil
	}
}

// parseUint32 將字串轉為 uint32
func parseUint32(s string) (uint32, error) {
	var i uint32
	_, err := fmt.Sscanf(s, "%d", &i)
	return i, err
}
