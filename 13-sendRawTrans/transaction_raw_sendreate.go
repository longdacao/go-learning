package main

import (
    "context"
    "encoding/hex"
    "fmt"
    "log"
    "github.com/spf13/viper"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/rlp"
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

    rawTx := "f8680d84b2d05e0082520894a3f2cf140f9446ac4a57e9b72986ce081db61e7585e8d4a510008077a0c57db522d4c988d012c4348bc0b81609b966d61bc8c7ec9467f58af3796b9cbfa002daced94e01e46f58d8b0405facb06de70004cae25ceda1e3dad8405d13c64e"

    rawTxBytes, err := hex.DecodeString(rawTx)

    tx := new(types.Transaction)
    rlp.DecodeBytes(rawTxBytes, &tx)

    err = client.SendTransaction(context.Background(), tx)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0xc429e5f128387d224ba8bed6885e86525e14bfdc2eb24b5e9c3351a1176fd81f
}