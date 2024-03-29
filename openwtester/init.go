package openwtester

import (
	"github.com/assetsadapterstore/grand50-adapter/grand50"
	"github.com/blocktree/openwallet/log"
	"github.com/blocktree/openwallet/openw"
)

func init() {
	//注册钱包管理工具
	log.Notice("Wallet Manager Load Successfully.")
	openw.RegAssets(grand50.Symbol, grand50.NewWalletManager())
}
