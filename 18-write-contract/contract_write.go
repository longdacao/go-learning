package main

import (
	"context"
	"crypto/ecdsa"
    "fmt"
    "log"
	"math/big"
	"github.com/spf13/viper"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"

    store "write-contract/Store" // for demo
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

    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    auth := bind.NewKeyedTransactor(privateKey)
    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0)     // in wei
    auth.GasLimit = uint64(300000) // in units
    auth.GasPrice = gasPrice

    address := common.HexToAddress("0x751db09472798ce19f22f90d5097c3250989100a")
    instance, err := store.NewStore(address, client)
    if err != nil {
        log.Fatal(err)
    }

    key := [32]byte{}
    value := [32]byte{}
    copy(key[:], []byte("foo"))
    copy(value[:], []byte("bar"))
    
    tx, err := instance.SetItem(auth, key, value)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870

    result, err := instance.Items(nil, key)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(result[:])) // "bar"
}