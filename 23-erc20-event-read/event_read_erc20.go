package main

import (
    "context"
    "fmt"
    "log"
    "math/big"
    "strings"
	"github.com/spf13/viper"
    token "erc20-event-read/token" // for demo
    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
)

// LogTransfer ..
type LogTransfer struct {
    From   common.Address
    To     common.Address
    Tokens *big.Int
}

// LogApproval ..
type LogApproval struct {
    TokenOwner common.Address
    Spender    common.Address
    Tokens     *big.Int
}

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


    // 0x Protocol (ZRX) token address
    contractAddress := common.HexToAddress("0x7820eEE0b99097b8a7B150C332917a3D8603126f")
    query := ethereum.FilterQuery{
        FromBlock: big.NewInt(29270309),
        ToBlock:   big.NewInt(29270332),
        Addresses: []common.Address{
            contractAddress,
        },
    }

    logs, err := client.FilterLogs(context.Background(), query)
    if err != nil {
        log.Fatal(err)
    }

    contractAbi, err := abi.JSON(strings.NewReader(string(token.TokenABI)))
    if err != nil {
        log.Fatal(err)
    }

    logTransferSig := []byte("Transfer(address,address,uint256)")
    LogApprovalSig := []byte("Approval(address,address,uint256)")
    logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
    logApprovalSigHash := crypto.Keccak256Hash(LogApprovalSig)

    for _, vLog := range logs {
        fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
        fmt.Printf("Log Index: %d\n", vLog.Index)

        switch vLog.Topics[0].Hex() {
        case logTransferSigHash.Hex():
            fmt.Printf("Log Name: Transfer\n")

            var transferEvent LogTransfer

            transfer, err := contractAbi.Unpack("Transfer", vLog.Data)
            if err != nil {
                log.Fatal(err)
            }

            transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
            transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

            fmt.Printf("From: %s\n", transferEvent.From.Hex())
            fmt.Printf("To: %s\n", transferEvent.To.Hex())
            fmt.Printf("Value: %s\n", transfer[0])

        case logApprovalSigHash.Hex():
            fmt.Printf("Log Name: Approval\n")

            var approvalEvent LogApproval

            approval, err := contractAbi.Unpack("Approval", vLog.Data)
            if err != nil {
                log.Fatal(err)
            }

            approvalEvent.TokenOwner = common.HexToAddress(vLog.Topics[1].Hex())
            approvalEvent.Spender = common.HexToAddress(vLog.Topics[2].Hex())
            fmt.Println(approval)

            fmt.Printf("Token Owner: %s\n", approvalEvent.TokenOwner.Hex())
            fmt.Printf("Spender: %s\n", approvalEvent.Spender.Hex())
            fmt.Printf("Value: %s\n", approval[0])
        }

        fmt.Printf("\n\n")
    }
}