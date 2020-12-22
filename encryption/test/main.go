package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type RSAUtil struct {
	PublicKeys  map[string]*rsa.PublicKey
	PrivateKeys map[string]*rsa.PrivateKey
}

func NewRSAUtil(publicKeys map[string]string, privateKeys map[string]string) (*RSAUtil, error) {
	rsaUtil := &RSAUtil{}
	rsaUtil.PublicKeys = make(map[string]*rsa.PublicKey)
	rsaUtil.PrivateKeys = make(map[string]*rsa.PrivateKey)

	for key, _ := range publicKeys {
		publicKey := publicKeys[key]
		privateKey := privateKeys[key]

		publicKeyBlocks, err := getKeyData(publicKey)
		if err != nil {
			return nil, err
		}
		rsaPublicKey, err := x509.ParsePKIXPublicKey(publicKeyBlocks.Bytes)
		if err != nil {
			return nil, err
		}
		rsaUtil.PublicKeys[key] = rsaPublicKey.(*rsa.PublicKey)

		privateKeyBlocks, err := getKeyData(privateKey)
		if err != nil {
			return nil, err
		}
		rsaPrivateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBlocks.Bytes)
		if err != nil {
			return nil, err
		}
		rsaUtil.PrivateKeys[key] = rsaPrivateKey
	}

	return rsaUtil, nil
}

func getKeyData(keyBase64 string) (*pem.Block, error) {
	keyData, err := base64.StdEncoding.DecodeString(keyBase64)
	if err != nil {
		return nil, err
	}
	keyString := string(keyData)
	keyString = strings.Replace(keyString, "	", "", -1)

	block, _ := pem.Decode([]byte(keyString))
	return block, nil
}

func (r *RSAUtil) Encrypt(msg []byte, index string) ([]byte, error) {
	hash := sha512.New()
	//key := "index" + index
	cipher, err := rsa.EncryptOAEP(hash, rand.Reader, r.PublicKeys[index], msg, nil)

	return cipher, err
}

func (r *RSAUtil) Decrypt(securedMessage []byte, index string) (string, error) {

	hash := sha512.New()
	//key := "index" + index
	message, err := rsa.DecryptOAEP(hash, rand.Reader, r.PrivateKeys[index], securedMessage, nil)

	return string(message), err
}

func main() {
	publicKeyMap := make(map[string]string)
	privateKeyMap := make(map[string]string)

	env := "dev"

	f, err := os.Open(fmt.Sprintf("../keys/%s/public.txt", env))
	b, err := ioutil.ReadAll(f)
	err = json.Unmarshal(b, &publicKeyMap)

	f, err = os.Open(fmt.Sprintf("../keys/%s/private.txt", env))
	b, err = ioutil.ReadAll(f)
	err = json.Unmarshal(b, &privateKeyMap)

	fmt.Println(err)

	stringToBeEncrypted := "Hello World!"

	rsaUtil, err := NewRSAUtil(publicKeyMap, privateKeyMap)

	for key, _ := range publicKeyMap {

		encryptedData, err := rsaUtil.Encrypt([]byte(stringToBeEncrypted), key)

		decryptedData, err := rsaUtil.Decrypt(encryptedData, key)

		fmt.Printf("%+v %v\n", decryptedData, err)
	}

}
