package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

var RemoteActivationToken = "4bd4d76f072e11667b1db0c4b5177c2b7ee48a3bae4766a5683063cc560fc60f"

type License struct {
	Key             string
	ActivationToken string
	Activated       bool
}

func main() {
	// Simulate license activation
	license := License{
		Key:             "YOUR_LICENSE_KEY",
		ActivationToken: "",
		Activated:       false,
	}

	// Activate the license
	err := activateLicense(&license)
	if err != nil {
		fmt.Println("Error activating license:", err)
		return
	}

	fmt.Println("License activated successfully!")
	fmt.Println("Activation Token:", license.ActivationToken)
}

func activateLicense(license *License) error {
	// Check if the license is already activated
	if license.Activated {
		fmt.Println("License is already activated.")
		return nil
	}

	// Generate a random activation token
	// activationToken, err := generateActivationToken()
	activationToken := RemoteActivationToken

	// if err != nil {
	// 	return err
	// }

	// Associate the activation token with the license
	license.ActivationToken = activationToken
	license.Activated = true

	// Store the activation information in your database or wherever you're managing licenses

	return nil
}

func generateActivationToken() (string, error) {
	// Generate a random byte slice for the token
	tokenBytes := make([]byte, 16) // 16 bytes = 128 bits
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}

	// Hash the random bytes using SHA-256 to create a fixed-length token
	hash := sha256.New()
	hash.Write(tokenBytes)
	hashBytes := hash.Sum(nil)

	// Convert the hash to a hexadecimal string
	token := hex.EncodeToString(hashBytes)

	return token, nil
}
