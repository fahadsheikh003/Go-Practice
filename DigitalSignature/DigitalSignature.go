package main

/*
References:
- https://gist.github.com/miguelmota/3ea9286bd1d3c2a985b67cac4ba2130a
- https://pkg.go.dev/crypto/rsa#SignPKCS1v15
- https://pkg.go.dev/crypto/rsa#VerifyPKCS1v15
*/

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

// GenerateKeyPair generates a new key pair
func GenerateKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey) {
	privkey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		panic(err)
	}
	return privkey, &privkey.PublicKey
}

// PrivateKeyToBytes private key to bytes
func PrivateKeyToBytes(priv *rsa.PrivateKey) []byte {
	privBytes := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(priv),
		},
	)

	return privBytes
}

// PublicKeyToBytes public key to bytes
func PublicKeyToBytes(pub *rsa.PublicKey) []byte {
	pubASN1, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		panic(err)
	}

	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubASN1,
	})

	return pubBytes
}

// BytesToPrivateKey bytes to private key
func BytesToPrivateKey(priv []byte) *rsa.PrivateKey {
	block, _ := pem.Decode(priv)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		fmt.Println("is encrypted pem block")
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			panic(err)
		}
	}
	key, err := x509.ParsePKCS1PrivateKey(b)
	if err != nil {
		panic(err)
	}
	return key
}

// BytesToPublicKey bytes to public key
func BytesToPublicKey(pub []byte) *rsa.PublicKey {
	block, _ := pem.Decode(pub)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		fmt.Println("is encrypted pem block")
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			panic(err)
		}
	}
	ifc, err := x509.ParsePKIXPublicKey(b)
	if err != nil {
		panic(err)
	}
	key, ok := ifc.(*rsa.PublicKey)
	if !ok {
		panic("not ok")
	}
	return key
}

// EncryptWithPublicKey encrypts data with public key
func EncryptWithPublicKey(msg []byte, pub *rsa.PublicKey) []byte {
	hash := sha512.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, pub, msg, nil) // this function will return ciphertext and an error if something goes wrong
	if err != nil {
		panic(err)
	}
	return ciphertext
}

// DecryptWithPrivateKey decrypts data with private key
func DecryptWithPrivateKey(ciphertext []byte, priv *rsa.PrivateKey) []byte {
	hash := sha512.New()
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, priv, ciphertext, nil) // this function will return plaintext and an error if something goes wrong
	if err != nil {
		panic(err)
	}
	return plaintext
}

// SignWithPrivateKey signs data with private key
func SignWithPrivateKey(msg []byte, priv *rsa.PrivateKey) []byte {
	hashed := sha256.Sum256(msg)
	signature, err := rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA256, hashed[:]) // this function will return signature and an error if something goes wrong
	if err != nil {
		panic(err)
	}
	return signature
}

// VerifyWithPublicKey verifies data with public key
func VerifyWithPublicKey(msg []byte, signature []byte, pub *rsa.PublicKey) bool {
	hashed := sha256.Sum256(msg)
	err := rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed[:], signature) // this function will return nil if the signature is valid and an error otherwise
	if err != nil {
		return false
	}
	return true
}

func main() {
	msg := []byte("Hello World")

	priv, pub := GenerateKeyPair(2048)
	signature := SignWithPrivateKey(msg, priv)
	verified := VerifyWithPublicKey(msg, signature, pub)

	fmt.Printf("signature verified: %v\n", verified)
}
