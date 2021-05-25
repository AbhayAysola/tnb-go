package account

import (
	"encoding/hex"

	"github.com/kevinburke/nacl/sign"
)

type Account struct {
	signingKey sign.PrivateKey
	accountNumber sign.PublicKey
}

func createAccount(signingKeyHex string) Account {
	signingKeyByte, err := hex.DecodeString(signingKeyHex)
	if err != nil {
		panic(err)
	}
	signingKey := sign.PrivateKey(signingKeyByte)
	return Account{signingKey: signingKey, accountNumber: signingKey.Public().(sign.PublicKey)}
}