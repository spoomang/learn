package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"log"
	"strings"
)

func main() {

	//1.
	pubKey := ""
	privateKey := ""

	data, _ := base64.StdEncoding.DecodeString(pubKey)

	key := string(data)
	key = strings.Replace(key, "	", "", -1)

	fmt.Println(key)

	block, _ := pem.Decode([]byte(key))
	if block == nil || block.Type != "PUBLIC KEY" {
		log.Fatal("failed to decode PEM block containing public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Got a %T \n", pub)

	text := "Asdf13"

	ciphertext, _ := rsa.EncryptOAEP(sha512.New(), rand.Reader, pub.(*rsa.PublicKey), []byte(text), nil)

	fmt.Println(base64.StdEncoding.EncodeToString(ciphertext))

	data1, _ := base64.StdEncoding.DecodeString(privateKey)

	key1 := string(data1)
	key1 = strings.Replace(key1, "	", "", -1)

	//fmt.Println(key1)

	block1, _ := pem.Decode([]byte(key1))
	//	fmt.Println(block1.Type)
	if block1 == nil || block1.Type != "RSA PRIVATE KEY" {
		log.Fatal("failed to decode PEM block containing private key")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block1.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(priv)
	plaintext, _ := rsa.DecryptOAEP(sha512.New(), rand.Reader, priv, ciphertext, nil)

	fmt.Println(string(plaintext))

}

func BytesToPrivateKey(priv []byte) *rsa.PrivateKey {
	block, _ := pem.Decode(priv)
	fmt.Println(block)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		fmt.Println("is encrypted pem block")
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			fmt.Println(err)
		}
	}
	key, err := x509.ParsePKCS1PrivateKey(b)
	if err != nil {
		fmt.Println(err)
	}
	return key
}
