package main

import (
	"fmt"
	"github.com/pquerna/otp/totp"
)

func main() {

	secret := "UBHOMQSYEIAKIGOK"
	fmt.Println(totp.Validate("198264", secret))
	key, _ := totp.Generate(totp.GenerateOpts{
		Issuer:      "Example.com",
		AccountName: "alice@example.com",
	})

	secret = key.Secret()
	fmt.Println(secret)
}
