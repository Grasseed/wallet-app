package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Grasseed/wallet-app/wallet"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	cmd := os.Args[1]

	switch cmd {
	case "mnemonic":
		wordCount := 12
		for _, arg := range os.Args[2:] {
			if strings.HasPrefix(arg, "--words=") {
				fmt.Sscanf(strings.TrimPrefix(arg, "--words="), "%d", &wordCount)
				if wordCount != 12 && wordCount != 18 && wordCount != 24 {
					fmt.Println("Invalid word count. Must be 12, 18, or 24.")
					return
				}
			}
		}
		mnemonic, err := wallet.GenerateMnemonic(wordCount)
		if err != nil {
			fmt.Println("Error generating mnemonic:", err)
			return
		}
		fmt.Println("Mnemonic:", mnemonic)

	case "derive":
		var mnemonic, path, addrType string
		addrType = "p2pkh"
		for _, arg := range os.Args[2:] {
			if strings.HasPrefix(arg, "--mnemonic=") {
				mnemonic = strings.TrimPrefix(arg, "--mnemonic=")
			} else if strings.HasPrefix(arg, "--path=") {
				path = strings.TrimPrefix(arg, "--path=")
			} else if strings.HasPrefix(arg, "--type=") {
				addrType = strings.TrimPrefix(arg, "--type=")
			}
		}
		if mnemonic == "" || path == "" {
			fmt.Println("Missing --mnemonic or --path argument")
			printUsage()
			return
		}
		address, wif, err := wallet.DeriveBTCAddress(mnemonic, path, addrType)
		if err != nil {
			fmt.Println("Error deriving BTC address:", err)
			return
		}
		fmt.Println("BTC Address:", address)
		fmt.Println("WIF Private Key:", wif)

	case "verify":
		var mnemonic, path, expect, addrType string
		addrType = "p2pkh"
		for _, arg := range os.Args[2:] {
			if strings.HasPrefix(arg, "--mnemonic=") {
				mnemonic = strings.TrimPrefix(arg, "--mnemonic=")
			} else if strings.HasPrefix(arg, "--path=") {
				path = strings.TrimPrefix(arg, "--path=")
			} else if strings.HasPrefix(arg, "--expect=") {
				expect = strings.TrimPrefix(arg, "--expect=")
			} else if strings.HasPrefix(arg, "--type=") {
				addrType = strings.TrimPrefix(arg, "--type=")
			}
		}
		if mnemonic == "" || path == "" || expect == "" {
			fmt.Println("Missing --mnemonic, --path or --expect argument")
			printUsage()
			return
		}
		address, _, err := wallet.DeriveBTCAddress(mnemonic, path, addrType)
		if err != nil {
			fmt.Println("Error deriving BTC address:", err)
			return
		}
		fmt.Println("Derived Address:", address)
		if address == expect {
			fmt.Println("✅ Matched")
		} else {
			fmt.Println("❌ Not matched")
		}

	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  mnemonic               Generate a 12-word mnemonic phrase")
	fmt.Println("                         Optionally add --words=12|18|24 to specify mnemonic length")
	fmt.Println("  derive --mnemonic=MNEMONIC --path=PATH [--type=p2pkh|segwit|p2sh]")
	fmt.Println("                         Derive BTC address and WIF private key from mnemonic and derivation path")
	fmt.Println("  verify --mnemonic=MNEMONIC --path=PATH --expect=ADDRESS [--type=p2pkh|segwit|p2sh]")
	fmt.Println("                         Verify that derived address matches expected address")
}
