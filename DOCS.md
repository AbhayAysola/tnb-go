# tnb-go

This module is used to interact with thenewboston blockchain network and (will) support all the network methods of the beta implementation.

# Packages

## account

The account package is an offline implementation of an account in thenewboston.

### CreateAccount

This method takes in a single parameter and returns 2 values.

```Golang
// Returns an Account struct with a randomly generated keypair
acc, err := account.CreateAccount("")
if err != nil {
    panic(err)
}
print(acc.SigningKeyHex, acc.AccountNumberHex)
```

```Golang
// Returns an Account struct with the signing key that was passed in
acc, err := account.CreateAccount("signingKeyHex")
if err != nil {
    panic(err)
}
print(acc.SigningKeyHex, acc.AccountNumberHex)
```

### Account.CreateSignature

This method takes in a single parameter and returns 2 values.

It can be used to sign any string for interaction with the network.

```Golang
acc, _ := account.CreateAccount("signingKeyHex")
sig, err := acc.CreateSignature("hi")
if err != nil {
    panic(err)
}
print(sig)
```

### VerifySignature

This method takes in 2 parameters and returns a single value.

It is used to verify the signature of an account.

```Golang
sig := "signatureHere"
verified := account.VerifySignature("sig", "accountNumberOfAccountUsedToSignMessage")
print(verified)
// true or false
```

### VerifyKeyPair

This method takes in 2 parameters and returns a single value.

It is used to check if an accountNumber and signingKey match.

```Golang
matched := account.VerifyKeyPair("signingKeyHex", "accountNumberHex")
print(matched)
// true or false
```
