package main

import (
	"crypto/ecdsa"
	"fmt"

	hdwallet "github.com/agol586/go-ethereum-hdwallet"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"
)

func cmdGenerateHD(ctx *cli.Context) error {
	var (
		mnemonic = ctx.String("from-mnemonic")
		idx      = uint32(ctx.Uint("start-index"))    // start index
		num      = uint32(ctx.Uint("account-number")) // account number
	)
	if mnemonic == "" {
		m, err := hdwallet.NewMenonic(128, "")
		if err != nil {
			return err
		}
		fmt.Println("auto gen", m.String())
		mnemonic = m.Data()
	}
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return err
	}

	pathLen := len(accounts.DefaultBaseDerivationPath)
	basePath := make(accounts.DerivationPath, pathLen)
	copy(basePath, accounts.DefaultBaseDerivationPath)
	privKeys := make([]*ecdsa.PrivateKey, 0, num)
	for ; idx < num; idx++ {
		basePath[pathLen-1] = idx
		path := basePath
		privKey, err := wallet.PrivateKeyByPath(path)
		if err != nil {
			return err
		}
		privKeys = append(privKeys, privKey)
	}

	for i, privKey := range privKeys {
		publicKeyECDSA, _ := privKey.Public().(*ecdsa.PublicKey)
		address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
		fmt.Printf("keypair %d private: %s , public address: %v \n", i, hexutil.Encode(crypto.FromECDSA(privKey))[2:], address)
	}
	return nil
}

func cmdGenerateRand(ctx *cli.Context) error {
	num := uint32(ctx.Uint("account-number"))
	for i := uint32(0); i < num; i++ {

	}
	return nil
}
