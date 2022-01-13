package main

import (
    "context"
    "fmt"
    "log"
	"github.com/spf13/viper"
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

    client, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/" + viper.GetString("INFURA_ID"))
    if err != nil {
        log.Fatal(err)
    }

    headers := make(chan *types.Header)
    sub, err := client.SubscribeNewHead(context.Background(), headers)
    if err != nil {
        log.Fatal(err)
    }

    for {
        select {
        case err := <-sub.Err():
            log.Fatal(err)
        case header := <-headers:
			fmt.Println("header hash: ", header.Hash().Hex())

            block, err := client.BlockByHash(context.Background(), header.Hash())
            if err != nil {
                log.Fatal(err)
            }

            fmt.Println("block hash: ",block.Hash().Hex())        
            fmt.Println("block number: ", block.Number().Uint64())  
            fmt.Println("block time: ", block.Time())     
            fmt.Println("block nonce: ",block.Nonce())    
            fmt.Println("block transactions: ",len(block.Transactions())) 
        }
    }
}