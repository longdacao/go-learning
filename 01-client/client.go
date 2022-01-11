package main

import (
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

    client, err := ethclient.Dial("https://mainnet.infura.io/" + viper.GetString("INFURA_ID"))
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("we have a connection")
    _ = client // we'll use this in the upcoming sections
}