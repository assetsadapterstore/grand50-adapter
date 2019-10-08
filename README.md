# grand50-adapter

grand50-adapter继承了bitcoin-adapter，主要修改了如下内容：

- 重写了Symbol = "G50"。
- 重写了addressDecoder，实现了G50地址编码。

## 如何测试

openwtester包下的测试用例已经集成了openwallet钱包体系，创建conf文件，新建G50.ini文件，编辑如下内容：

```ini

# RPC Server Type，0: CoreWallet RPC; 1: Explorer API
rpcServerType = 0
# node api url, if RPC Server Type = 0, use bitcoin core full node
;serverAPI = "http://127.0.0.1:8333/"
# RPC Authentication Username
rpcUser = "user"
# RPC Authentication Password
rpcPassword = "password"
# Is network test?
isTestNet = false
# support segWit
supportSegWit = false
# minimum transaction fees
minFees = "0.00001"
# Cache data file directory, default = "", current directory: ./data
dataDir = ""

```

## 资料介绍

### 区块浏览器

http://explorer.g50.io/