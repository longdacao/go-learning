package main

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	j               = "application/json"
	flashbotURL     = "https://relay.flashbots.net"
	stats           = "flashbots_getUserStats"
	flashbotXHeader = "X-Flashbots-Signature"
	p               = "POST"
)

var (
	privateKey, _ = crypto.HexToECDSA(
		"2077296d5884088c34762a5765283fddb57e1c0a157a78d4092eb076db94f23b",
	)
)

func flashbotHeader(signature []byte, privateKey *ecdsa.PrivateKey) string {
	return crypto.PubkeyToAddress(privateKey.PublicKey).Hex() +
		":" + hexutil.Encode(signature)
}

func main() {
	mevHTTPClient := &http.Client{
		Timeout: time.Second * 3,
	}
	currentBlock := big.NewInt(12_900_000)
	params := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  stats,
		"params": []interface{}{
			fmt.Sprintf("0x%x", currentBlock.Uint64()),
		},
	}
	payload, _ := json.Marshal(params)
	req, _ := http.NewRequest(p, flashbotURL, bytes.NewBuffer(payload))
	headerReady, _ := crypto.Sign(
		accounts.TextHash([]byte(hexutil.Encode(crypto.Keccak256(payload)))),
		privateKey,
	)
	req.Header.Add("content-type", j)
	req.Header.Add("Accept", j)
	req.Header.Add(flashbotXHeader, flashbotHeader(headerReady, privateKey))
	resp, _ := mevHTTPClient.Do(req)
	res, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(res))
}
