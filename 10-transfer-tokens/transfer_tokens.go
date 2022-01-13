package main

import (
    "context"
    "crypto/ecdsa"
    "fmt"
    "log"
    "math/big"
	"github.com/spf13/viper"
    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/common/hexutil"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/crypto"
    "golang.org/x/crypto/sha3"
    "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	viper.SetConfigName(".env") 
    viper.SetConfigType("toml")
    viper.AddConfigPath(".")

    err := viper.ReadInConfig() // 查找并读取配置文件
    if err != nil { // 处理读取配置文件的错误
      panic(fmt.Errorf("Fatal error config file: %s \n", err))
    }

    client, err := ethclient.Dial("https://kovan.infura.io/v3/" + viper.GetString("INFURA_ID"))
    if err != nil {
        log.Fatal(err)
    }

    privateKey, err := crypto.HexToECDSA(viper.GetString("PRIVATE_KEY"))
    if err != nil {
        log.Fatal(err)
    }

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }

    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatal(err)
    }

    value := big.NewInt(0) // in wei (0 eth)
    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    toAddress := common.HexToAddress("0xa3F2Cf140F9446AC4a57E9B72986Ce081dB61E75")
    tokenAddress := common.HexToAddress("0xbA9Fda75842547bc8259B80A22C5a09B408cFeC8")

    transferFnSignature := []byte("transfer(address,uint256)")
    hash := sha3.NewLegacyKeccak256()
    hash.Write(transferFnSignature)
    methodID := hash.Sum(nil)[:4]
    fmt.Println("methodID: ",hexutil.Encode(methodID)) // 0xa9059cbb

    paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
    fmt.Println("toAddress: ",hexutil.Encode(paddedAddress)) 

    amount := new(big.Int)
    amount.SetString("100", 10) // 1000 tokens
    paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
    fmt.Println("Amount: ",hexutil.Encode(paddedAmount)) 

    var data []byte
    data = append(data, methodID...)
    data = append(data, paddedAddress...)
    data = append(data, paddedAmount...)

    gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
        To:   &toAddress,
        Data: data,
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("gasLimit: ", gasLimit) 

    tx := types.NewTransaction(nonce, tokenAddress, value, 3000000, gasPrice, data)

    chainID, err := client.NetworkID(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
    if err != nil {
        log.Fatal(err)
    }

    err = client.SendTransaction(context.Background(), signedTx)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("tx sent: %s ", signedTx.Hash().Hex()) // tx sent: 0xa56316b637a94c4cc0331c73ef26389d6c097506d581073f927275e7a6ece0bc
}