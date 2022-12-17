package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func main() {
	// Generate a new pair of keys
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	// Encode the private key in PKCS#8 format
	privateKeyPEM := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	// Encode the public key in X.509 format
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		panic(err)
	}
	publicKeyPEM := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	// Alice wants to send a message to Bob
	message := []byte("Hello, Bob!")

	// Alice encrypts the message using Bob's public key
	encryptedMessage, err := rsa.EncryptPKCS1v15(rand.Reader, &privateKey.PublicKey, message)
	if err != nil {
		panic(err)
	}

	// Bob receives the encrypted message and decrypts it using his private key
	decryptedMessage, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encryptedMessage)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(decryptedMessage)) // "Hello, Bob!"
}
