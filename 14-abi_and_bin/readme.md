# 合约 和 ABI  
### 操作步骤  
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

## 参考文档  
安装 abigen: 
https://www.itread01.com/fxlpq.html   
https://juejin.cn/post/6968086477729333262  
https://www.jianshu.com/p/ddddc5a8e08a  
