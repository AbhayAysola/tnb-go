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

func TestCreateSignature(t *testing.T) {
	acc, _ := CreateAccount("74938d383000114b75f998f49d369dc98ffdb3bfef0f90744056e0165104b1ed")
	sig, err := acc.CreateSignature("hi")
	if err != nil {
		t.Error("Creating signature results in error: ", err)
	}
	if sig != "7fdc1c0ece3d9c3cb74cf253ab7b339f1684352f33d50f3699884100e63a9c232d37a2ef8934d5dfa4dc4e28dc29ee05cb66c67db84032ad416bfd7929863b0d" {
		t.Error("Signature does not match")
	}
}

func TestVerifySignature(t *testing.T) {
	verified := VerifySignature("7fdc1c0ece3d9c3cb74cf253ab7b339f1684352f33d50f3699884100e63a9c232d37a2ef8934d5dfa4dc4e28dc29ee05cb66c67db84032ad416bfd7929863b0d", "20904d4790c907d507979f0c37f4c8c6d55533495c9c1963e8ec68351bc709d5")
	if !verified {
		t.Error("Valid signature returning false")
	}

	verifiedInvalid := VerifySignature("7fdc1c0ece3d9c3cb74cf25fab7b339f1684352f33d50f3699884100e63a9c232d37a2ef8934d5dfa4dc4e28dc29ee05cb66c67db84032ad416bfd7929863b0d", "20904d4790c907d507979f0c37f4c8c6d55533495c9c1963e8ec68351bc709d5")
	if verifiedInvalid {
		t.Error("Valid signature returning true")
	}
}

func TestVerifyKeyPair(t *testing.T) {
	acc, _ := CreateAccount("")
	if !VerifyKeyPair(acc.SigningKeyHex, acc.AccountNumberHex) {
		t.Error("VerifyKeyPair returning false")
	}
}
