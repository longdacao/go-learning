package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

const alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

func Base58Encode(input []byte) string {
	var result []byte
	x := big.NewInt(0).SetBytes(input)
	base := big.NewInt(int64(len(alphabet)))
	zero := big.NewInt(0)
	for x.Cmp(zero) > 0 {
		mod := new(big.Int)
		x.DivMod(x, base, mod)
		result = append(result, alphabet[mod.Int64()])
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}

// Base58Decode decodes a Base58 string to a byte slice
func Base58Decode(input string) []byte {
	x := big.NewInt(0)
	base := big.NewInt(int64(len(alphabet)))
	for _, r := range input {
		index := bytes.IndexByte([]byte(alphabet), byte(r))
		if index < 0 {
			return nil
		}
		x.Mul(x, base)
		x.Add(x, big.NewInt(int64(index)))
	}
	result := x.Bytes()
	result = append(bytes.Repeat([]byte{0x00}, len(input)-len(result)), result...)
	return result
}

// SS58Encode encodes a byte slice to an SS58 string
func SS58Encode(network byte, input []byte) string {
	// Add the network prefix
	buffer := bytes.Buffer{}
	buffer.WriteByte(network)
	buffer.Write(input)

	// Calculate the checksum
	hash := sha256.Sum256(buffer.Bytes())
	fmt.Println("hash=====", hash)
	hash = sha256.Sum256(hash[:])
	fmt.Println("hash=====", hash)
	checksum := hash[:2]
	fmt.Println("checksum=====", checksum)

	// Add the checksum and convert to Base58
	buffer.Write(checksum)
	return Base58Encode(buffer.Bytes())
}

// SS58Decode decodes an SS58 string to a byte slice and network prefix
func SS58Decode(input string) (byte, []byte, error) {
	// Decode from Base58
	data := Base58Decode(input)

	// Verify the checksum
	if len(data) < 3 {
		return 0, nil, fmt.Errorf("invalid SS58 format")
	}
	buffer := bytes.Buffer{}
	buffer.Write(data[:len(data)-2])
	hash := sha256.Sum256(buffer.Bytes())
	hash = sha256.Sum256(hash[:])
	if !bytes.Equal(hash[:2], data[len(data)-2:]) {
		return 0, nil, fmt.Errorf("invalid SS58 checksum")
	}

	// Extract the network prefix and return the result
	network := data[0]
	result := data[1 : len(data)-2]
	return network, result, nil
}

func main() {
	fmt.Println(Base58Encode([]byte{'4', 'z', 'z', 'z', 'z'}))
}
