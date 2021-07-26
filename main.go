package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
)

func main() {

	userkey := flag.String("key", "foo", "a string")
	inputstr := flag.String("msg", "foo", "a string")
	encrypt := flag.Bool("encrypt", false, "encrypt")
	decrypt := flag.Bool("decrypt", false, "decrypt")
	genkey := flag.Bool("genkey", false, "gen key")
	flag.Parse()

	var outputStr string

	if *genkey {
		keystr := genKeyString()
		fmt.Printf("Encryption key: %s", keystr)
		os.Exit(0)
	}

	if *encrypt {
		outputStr = encryptString(*inputstr, *userkey)
		fmt.Println(outputStr)
		os.Exit(0)
	}
	if *decrypt {
		outputStr = decryptString(*inputstr, *userkey)
		fmt.Println(outputStr)
		os.Exit(0)
	}

}

func genKeyString() (keyStr string) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}

	return hex.EncodeToString(bytes)
}

func encryptString(plainMsg string, encKeyStr string) (encstr string) {

	key, _ := hex.DecodeString(encKeyStr)
	plaintext := []byte(plainMsg)
	block, _ := aes.NewCipher(key)
	aesGCM, _ := cipher.NewGCM(block)

	nonce := make([]byte, aesGCM.NonceSize())

	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)

	return fmt.Sprintf("%x", ciphertext)

}

func decryptString(encMsg string, encKeyStr string) (outputStr string) {

	key, _ := hex.DecodeString(encKeyStr)
	enc, _ := hex.DecodeString(encMsg)
	block, _ := aes.NewCipher(key)
	aesGCM, _ := cipher.NewGCM(block)

	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	plaintext, _ := aesGCM.Open(nil, nonce, ciphertext, nil)

	return fmt.Sprintf("%s", plaintext)
}
