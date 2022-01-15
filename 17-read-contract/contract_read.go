package main

import (
    "fmt"
    "log"
	"github.com/spf13/viper"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"

	store "read-contract/Store" // for demo
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

    address := common.HexToAddress("0x751db09472798ce19f22f90d5097c3250989100a")
    instance, err := store.NewStore(address, client)
    if err != nil {
        log.Fatal(err)
    }

    version, err := instance.Version(nil)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(version) // "1.0"
}