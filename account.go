package account

import (
	"bytes"
	"encoding/hex"

	"github.com/kevinburke/nacl/sign"
)

type Account struct {
	signingKey sign.PrivateKey
	accountNumber sign.PublicKey
	SigningKeyHex string
	AccountNumberHex string
}

func CreateAccount(signingKeyHex string) (Account, error) {
	if signingKeyHex != "" {

		signingKeyByte, err := hex.DecodeString(signingKeyHex)
		if err != nil {
			return Account{}, err
		}
		accountNumber, signingKey, err := sign.Keypair(bytes.NewReader(signingKeyByte))
		if err != nil {
			return Account{}, err
		}
		return Account{signingKey: signingKey, accountNumber: accountNumber, SigningKeyHex: signingKeyHex, AccountNumberHex: hex.EncodeToString(accountNumber)}, nil
	} else {
		accountNumber, signingKey, err := sign.Keypair(nil)
		if err != nil {
			return Account{}, nil
		}
		return Account{signingKey: signingKey, accountNumber: accountNumber, SigningKeyHex: (hex.EncodeToString(signingKey))[0:64], AccountNumberHex: hex.EncodeToString(accountNumber)}, nil
	}
}