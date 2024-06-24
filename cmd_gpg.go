package main

import (
	"fmt"
	"strings"

	"github.com/ProtonMail/gopenpgp/v2/crypto"
	"github.com/ProtonMail/gopenpgp/v2/helper"
	"github.com/urfave/cli/v2"
)

func gen_gpg(Ctx *cli.Context) error {
	name := Ctx.String("name")
	email := Ctx.String("email")
	passphrase := Ctx.String("passphrase")
	rsaBits := Ctx.Int("rsa-bits")

	// RSA, string
	rsaKey, err := helper.GenerateKey(name, email, []byte(passphrase), "rsa", rsaBits)
	if err != nil {
		panic(err)
	}

	fmt.Println(rsaKey)

	keyRing, err := crypto.NewKeyFromArmoredReader(strings.NewReader(rsaKey))
	if err != nil {
		panic(err)
	}

	publicKey, err := keyRing.GetArmoredPublicKey()
	if err != nil {
		panic(err)
	}
	fmt.Println(publicKey)
	return nil
}
