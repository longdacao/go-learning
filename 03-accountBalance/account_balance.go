package main

import (
    "context"
    "fmt"
    "log"
    "math"
    "math/big"
	"github.com/spf13/viper"
    "github.com/ethereum/go-ethereum/common"
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

    account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
    balance, err := client.BalanceAt(context.Background(), account, nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(balance) // 25893180161173005034

    blockNumber := big.NewInt(13982691)
    balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(balanceAt) // 25729324269165216042

    fbalance := new(big.Float)
    fbalance.SetString(balanceAt.String())
    ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
    fmt.Println(ethValue) // 25.729324269165216041

    pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
    fmt.Println(pendingBalance) // 25729324269165216042
}