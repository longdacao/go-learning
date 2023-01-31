# go-learning  
- 安装 abigen  
```shell
go get github.com/ethereum/go-ethereum
cd $GOPATH/src/github.com/ethereum/go-ethereum/
make
make devtools

## 之后把 cmd/abigen/main.go 编译为可执行文件，然后放到 /usr/local/bin 目录下

abigen --help
```

- 安装 protobuf  
```
brew install protobuf
```

- 编译合约  
```shell
solc Store.sol   --optimize  --bin --abi --output-dir .
```

- 生成 go 文件  
```shell
abigen --abi=Store.abi --pkg=store --out=Store.go
```

- 编译可执行文件  
```shell
go build
```

## 参考文档 
goethereumbook: https://goethereumbook.org/zh/  
以太坊官网推荐: https://ethereum.org/en/developers/docs/programming-languages/golang/  
go 源码: https://github.com/ethereum/go-ethereum  
Perigord 一款 go dapp 开发工具: https://medium.com/decentralize-today/introducing-perigord-golang-tools-for-ethereum-dapp-development-60556c2d9fd  
connect to ethereum with go: https://www.quicknode.com/guides/web3-sdks/how-to-connect-to-ethereum-network-using-go  
Go Ethereum : https://geth.ethereum.org/  
Go 设计与实现: https://draveness.me/golang/docs/part2-foundation/ch05-keyword/golang-make-and-new/  
GO Ethereum 接口: https://pkg.go.dev/github.com/ava-labs/go-ethereum/accounts/keystore#NewKeyStore  
go 代码样例网站: www.codegrepper.com/  

安装 abigen  : 
https://www.itread01.com/fxlpq.html   
https://juejin.cn/post/6968086477729333262  
https://www.jianshu.com/p/ddddc5a8e08a  