# wallet-app 🪙

> A lightweight command-line Bitcoin wallet built with Go.  
> 使用 Go 語言開發的輕量級命令列 BTC 錢包工具，支援助記詞生成、地址導出與簽名功能。

---

## ✨ Features

- ✅ Generate BIP39 mnemonic phrases
- ✅ Derive HD wallet addresses using BIP32/BIP44 (P2PKH)
- ✅ Export private keys in WIF format
- ✅ (Upcoming) Offline Bitcoin transaction signing

---

## 📦 Installation

```bash
git clone https://github.com/Grasseed/wallet-app.git
cd wallet-app
go build -o wallet
```

---

## 🧪 Usage Examples

### Generate Mnemonic
```bash
./wallet mnemonic
```

### Derive BTC Address
```bash
./wallet derive --mnemonic="..." --path="m/44'/0'/0'/0/0"
```

---

## 📁 Project Structure

```
wallet-app/
├── main.go                # CLI entry point
├── wallet/
│   ├── mnemonic.go        # BIP39 mnemonic utilities
│   └── keygen.go          # BIP32/BIP44/P2SH derivation logic
├── test/
│   └── wallet_test.go     # Unit tests for address and mnemonic generation
├── Makefile              # Build and test automation
├── go.mod                # Go module definitions
└── README.md             # Project documentation
```

---

## 📄 License

MIT License © 2025 Grasseed
