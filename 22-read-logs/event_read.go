package main

import (
    "context"
    "fmt"
    "log"
    "math/big"
    "strings"
	"github.com/spf13/viper"
    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"

    store "read-logs/Store" // for demo
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

    contractAddress := common.HexToAddress("0x751db09472798ce19f22f90d5097c3250989100a")
    query := ethereum.FilterQuery{
        FromBlock: big.NewInt(29259367),
        ToBlock:   big.NewInt(29259367),
        Addresses: []common.Address{
            contractAddress,
        },
    }

    logs, err := client.FilterLogs(context.Background(), query)
    if err != nil {
        log.Fatal(err)
    }

    contractAbi, err := abi.JSON(strings.NewReader(string(store.StoreABI)))
    if err != nil {
        log.Fatal(err)
    }

    for _, vLog := range logs {
        fmt.Println(vLog.BlockHash.Hex()) // 0x3404b8c050aa0aacd0223e91b5c32fee6400f357764771d0684fa7b3f448f1a8
        fmt.Println(vLog.BlockNumber)     // 29259384
        fmt.Println(vLog.TxHash.Hex())    // 0x280201eda63c9ff6f305fcee51d5eb86167fab40ca3108ec784e8652a0e2b1a6

		type Item struct {
			Key [32]byte
			Value [32]byte
		}
		
		fmt.Println(vLog.Data) 
        eventPack, err := contractAbi.Unpack("ItemSet", vLog.Data)
        if err != nil {
            log.Fatal(err)
        }

		event := Item{
			Key: eventPack[0].([32]byte),
			Value: eventPack[1].([32]byte),
		}

		
        fmt.Println("Key: ", string(event.Key[:]))   // foo
        fmt.Println("Value: ", string(event.Value[:])) // bar

        var topics [4]string
        for i := range vLog.Topics {
            topics[i] = vLog.Topics[i].Hex()
        }

        fmt.Println("EventSignature: ", topics[0]) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
    }

    eventSignature := []byte("ItemSet(bytes32,bytes32)")
    hash := crypto.Keccak256Hash(eventSignature)
    fmt.Println("EventSignature: ", hash.Hex()) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
}