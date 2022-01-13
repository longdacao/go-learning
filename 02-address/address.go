package main

import (
    "fmt"

    "github.com/ethereum/go-ethereum/common"
)

func main() {
    address := common.HexToAddress("0x3238f24e7C752398872B768Ace7dd63c54CfEFEc")

    fmt.Println(address.Hex())        // 0x3238f24e7C752398872B768Ace7dd63c54CfEFEc
    fmt.Println(address.Hash().Hex()) // 0x00000000000000000000000071c7656ec7ab88b098defb751b7401b5f6d8976f
    fmt.Println(address.Bytes())      // [113 199 101 110 199 171 136 176 152 222 251 117 27 116 1 181 246 216 151 111]
}