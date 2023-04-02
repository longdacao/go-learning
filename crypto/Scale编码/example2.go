package main

import (
    "context"
    "fmt"
    "github.com/ChainSafe/gossamer/lib/common"
    "github.com/ChainSafe/gossamer/rpc"
    "github.com/ChainSafe/gossamer/subspace"
    "github.com/ChainSafe/gossamer/telemetry"
    "github.com/centrifuge/go-substrate-rpc-client/types"
)

func getCallIndex(functionName string, args []interface{}) ([]byte, error) {
    // Initialize the RPC client
    c, err := NewWSClient("wss://rpc.polkadot.io", "system", types.SupportsDPoS)
    if err != nil {
        return nil, err
    }
    defer c.Close()

    // Retrieve the metadata from the network
    meta, err := c.MetaData(context.Background())
    if err != nil {
        return nil, err
    }

    // Find the function in the metadata
    call, err := types.NewCall(meta, functionName, args)
    if err != nil {
        return nil, err
    }

    // Calculate and return the call index
    callBytes, err := call.Encode()
    if err != nil {
        return nil, err
    }
    return callBytes, nil
}

func main() {
    // Call the transfer function with two arguments: recipient (string) and value (uint128)
    functionName := "balances.transfer"
    args := []interface{}{
        "Alice",
        uint64(1_000_000_000_000),
    }

    callIndex, err := getCallIndex(functionName, args)
    if err != nil {
        panic(err)
    }

    fmt.Println("Call index:", callIndex)
}

// NewWSClient creates a new websocket client with the given parameters
func NewWSClient(addr string, ss string, dp types.DPoSVersions) (*rpc.Client, error) {
    sub := subspace.NewSubspace(common.MustHexToHash("0x" + ss), telemetry.New(telemetry.Config{}))
    c, err := rpc.NewClient(addr, sub, dp)
    if err != nil {
        return nil, err
   
