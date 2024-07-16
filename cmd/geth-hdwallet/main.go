package main

import (
	"fmt"
	"log"
	"os"

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
	app.Flags = []cli.Flag{}
	app.Commands = cli.Commands{
		&cli.Command{
			Name:   "generate-hd",
			Action: cmdGenerateHD,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "from-mnemonic",
					Usage:    "from mnemonic, only for wallet-type: hd",
					Value:    "",
					Required: false,
				},
				&cli.UintFlag{
					Name:     "start-index",
					Usage:    "start from account index of: m/44'/60'/0'/0/{start-index}",
					Value:    0,
					Required: false,
				},
				&cli.UintFlag{
					Name:     "account-number",
					Usage:    "account number to generate, one of account-index / account-number should be use",
					Value:    1,
					Required: false,
				},
			},
		},
		&cli.Command{
			Name:   "generate-rand",
			Action: cmdGenerateRand,
			Flags: []cli.Flag{
				&cli.UintFlag{
					Name:     "account-number",
					Usage:    "random account number to generate",
					Value:    1,
					Required: false,
				},
			},
		},
	}
	app.Action = runAction

	if err := app.Run(os.Args); err != nil {
		log.Println(err)
	}
}

func runAction(ctx *cli.Context) error {
	fmt.Println("runAction")
	return nil
}
