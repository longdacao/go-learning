package main

import (
	"context"

    "fmt"
    "log"
	"github.com/spf13/viper"
    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
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

    client, err := ethclient.Dial("wss://kovan.infura.io/ws/v3/" + viper.GetString("INFURA_ID"))
    if err != nil {
        log.Fatal(err)
    }

    contractAddress := common.HexToAddress("0x2eb3F8280A658f7b36D4634b25225cd0B7e53904")
    query := ethereum.FilterQuery{
        Addresses: []common.Address{contractAddress},
    }

    logs := make(chan types.Log)
    sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
    if err != nil {
        log.Fatal(err)
    }

    for {
        select {
        case err := <-sub.Err():
            log.Fatal(err)
        case vLog := <-logs:
            fmt.Println(vLog) // pointer to event log
        }
    }
}