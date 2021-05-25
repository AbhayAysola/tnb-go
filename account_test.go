package account

import (
	"testing"
)
func TestCreateAccountFromValidSigningKey(t *testing.T) {
	signingKeyHex := "74938d383000114b75f998f49d369dc98ffdb3bfef0f90744056e0165104b1ed"
	acc, err := CreateAccount(signingKeyHex)
	if err != nil {
		t.Error("Creating account from valid key results in error: ", err)
	}
	if acc.SigningKeyHex != signingKeyHex {
		t.Error("SigningKeyHex does not match")
	}
	if acc.AccountNumberHex != "20904d4790c907d507979f0c37f4c8c6d55533495c9c1963e8ec68351bc709d5" {
		t.Error("AccountNumberHex does not match")
	}
}

func TestCreateAccountFromInvalidSigningKey(t *testing.T) {
	signingKeyHex := "dsads"
	_, err := CreateAccount(signingKeyHex)
	if err == nil {
		t.Error("Creating account from invalid key does not result in error: ", err)
	}
}

func TestCreateAccountRandom(t *testing.T) {
	_, err := CreateAccount("")
	if err != nil {
		t.Error("Creating account randomly results in error: ", err)
	}
}