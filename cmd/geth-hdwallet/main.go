package main

import (
	"fmt"
	"log"
	"os"

	hdwallet "github.com/agol586/go-ethereum-hdwallet"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/urfave/cli/v2"
)

const (
	version = "v0.1"
	gitHash = "-"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	app := cli.NewApp()
	app.Name = "ethereum hdwallet"
	app.Version = fmt.Sprintf("version: %s-%s", version, gitHash)
	app.Description = "ethereum hdwallet generator"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "from-mnemonic",
			Usage:    "from mnemonic",
			Value:    "",
			Required: false,
		},
		&cli.UintFlag{
			Name:     "account-index",
			Usage:    "account index of: m/44'/60'/0'/0/{account-index}",
			Value:    0,
			Required: false,
		},
	}
	app.Action = runAction

	if err := app.Run(os.Args); err != nil {
		log.Println(err)
	}
}

func runAction(ctx *cli.Context) error {
	var mnemonic = ctx.String("from-mnemonic")
	var idx = uint32(ctx.Uint("account-index"))
	if mnemonic == "" {
		m, err := hdwallet.NewMenonic(128, "")
		if err != nil {
			return err
		}
		fmt.Println("mnemonic:", m)
		mnemonic = m.Data()
	}
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return err
	}
	rootPath := accounts.DefaultRootDerivationPath.String()
	hdpath := fmt.Sprintf("%s/%d", rootPath, idx)
	path := hdwallet.MustParseDerivationPath(hdpath)
	account, err := wallet.Derive(path, false)
	if err != nil {
		return err
	}

	privateKey, err := wallet.PrivateKeyHex(account)
	if err != nil {
		return err
	}

	fmt.Println("public address:", account.Address.Hex())
	fmt.Println("private key:", privateKey)
	return nil
}
