package main

import (
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"

	"golang.org/x/crypto/chacha20poly1305"
)

func hashToPoint(curve elliptic.Curve, message []byte) *big.Int {
	hash := sha256.Sum256(message)
	h := new(big.Int).SetBytes(hash[:])
	h.Mod(h, curve.Params().N) // Reduce hash to a point on the curve
	return h
}

func encrypt(curve elliptic.Curve, publicKeyX, publicKeyY *big.Int, message []byte) ([]byte, *big.Int, *big.Int) {
	// Generate ephemeral key
	k, _ := rand.Int(rand.Reader, curve.Params().N)

	// Compute C1 = k * G
	C1x, C1y := curve.ScalarBaseMult(k.Bytes())

	// Compute shared key from k * publicKey
	sx, sy := curve.ScalarMult(publicKeyX, publicKeyY, k.Bytes())
	sharedKey := sha256.Sum256(elliptic.Marshal(curve, sx, sy))

	// Encrypt message with shared key
	nonce := make([]byte, chacha20poly1305.NonceSizeX)
	aead, _ := chacha20poly1305.NewX(sharedKey[:])
	encryptedMessage := aead.Seal(nil, nonce, message, nil)

	return encryptedMessage, C1x, C1y
}

func decrypt(curve elliptic.Curve, privateKey, C1x, C1y *big.Int, encryptedMessage []byte) []byte {
	// Compute shared key from privateKey * C1
	sx, sy := curve.ScalarMult(C1x, C1y, privateKey.Bytes())
	sharedKey := sha256.Sum256(elliptic.Marshal(curve, sx, sy))

	// Decrypt message with shared key
	nonce := make([]byte, chacha20poly1305.NonceSizeX)
	aead, _ := chacha20poly1305.NewX(sharedKey[:])
	message, _ := aead.Open(nil, nonce, encryptedMessage, nil)

	return message
}

func main() {
	curve := elliptic.P256()

	// Generate key pair
	privateKeyBytes, publicKeyX, publicKeyY, _ := elliptic.GenerateKey(curve, rand.Reader)
	privateKey := new(big.Int).SetBytes(privateKeyBytes)

	message := []byte("Hello, Hashed ElGamal!")

	fmt.Printf("Original Message: %x\n", message)

	// Encrypt
	encryptedMessage, C1x, C1y := encrypt(curve, publicKeyX, publicKeyY, message)
	fmt.Printf("Encrypted Message: %x, C1(%x, %x)\n", encryptedMessage, C1x, C1y)

	// Decrypt
	decryptedMessage := decrypt(curve, privateKey, C1x, C1y, encryptedMessage)
	fmt.Printf("Decrypted Message: %x\n", decryptedMessage)
}
