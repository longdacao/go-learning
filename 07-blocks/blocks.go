package main

import (
    "context"
    "fmt"
    "log"
	"github.com/spf13/viper"
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
	
    client, err := ethclient.Dial("https://mainnet.infura.io/v3/" + viper.GetString("INFURA_ID"))
    if err != nil {
        log.Fatal(err)
    }

    header, err := client.HeaderByNumber(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(header.Number.String()) 

    block, err := client.BlockByNumber(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Block Number: " , block.Number().Uint64())     
    fmt.Println("Block Time: ",block.Time())       
    fmt.Println("Block Difficulty: ",block.Difficulty().Uint64()) 
    fmt.Println("Block Hash: ", block.Hash().Hex())         
    fmt.Println("Block transactions: ",len(block.Transactions()))  

    count, err := client.TransactionCount(context.Background(), block.Hash())
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("TransactionCount: ", count) // 144
}