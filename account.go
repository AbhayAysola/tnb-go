package account

import (
	"bytes"
	"crypto"
	"encoding/hex"
	"errors"

	"github.com/kevinburke/nacl/sign"
)

type Account struct {
	SigningKey       sign.PrivateKey
	accountNumber    sign.PublicKey
	SigningKeyHex    string
	AccountNumberHex string
}

func CreateAccount(signingKeyHex string) (Account, error) {
	if signingKeyHex != "" {

		signingKeyByte, err := hex.DecodeString(signingKeyHex)
		if err != nil {
			return Account{}, errors.New("invalid signingKeyHex")
		}
		accountNumber, signingKey, err := sign.Keypair(bytes.NewReader(signingKeyByte))
		if err != nil {
			return Account{}, errors.New("invalid signingKeyHex")
		}
		return Account{SigningKey: signingKey, accountNumber: accountNumber, SigningKeyHex: signingKeyHex, AccountNumberHex: hex.EncodeToString(accountNumber)}, nil
	} else {

		accountNumber, signingKey, err := sign.Keypair(nil)
		if err != nil {
			return Account{}, nil
		}
		return Account{SigningKey: signingKey, accountNumber: accountNumber, SigningKeyHex: (hex.EncodeToString(signingKey))[0:64], AccountNumberHex: hex.EncodeToString(accountNumber)}, nil
	}
}

func (account Account) CreateSignature(message string) (string, error) {
	signatureByte, err := account.SigningKey.Sign(nil, []byte(message), crypto.Hash(0))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(signatureByte), err
}

func VerifySignature(signatureHex string, accountNumberHex string) bool {
	accountNumber, err := hex.DecodeString(accountNumberHex)
	if err != nil {
		return false
	}
	signature, err := hex.DecodeString(signatureHex)
	if err != nil {
		return false
	}
	return sign.PublicKey(accountNumber).Verify(signature)
}

func VerifyKeyPair(signingKeyHex string, accountNumberHex string) bool {
	signingKey, err := hex.DecodeString(signingKeyHex)
	if err != nil {
		return false
	}
	accountNumber, err := hex.DecodeString(accountNumberHex)
	if err != nil {
		return false
	}
	accountNumberTest, _, err := sign.Keypair(bytes.NewReader(signingKey))
	if err != nil {
		return false
	}
	return bytes.Equal(accountNumber, accountNumberTest)
}
