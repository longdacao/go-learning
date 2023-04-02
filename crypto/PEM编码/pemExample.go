package main

import (
	"encoding/pem"
	"fmt"
)

func EncodeToPEM(data []byte, blockType string) ([]byte, error) {
	block := &pem.Block{
		Type:  blockType,
		Bytes: data,
	}
	return pem.EncodeToMemory(block), nil
}

func DecodeFromPEM(pemData []byte) ([]byte, error) {
	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block containing private key")
	}
	return block.Bytes, nil
}
