package lcd

import (
	"github.com/irisnet/explorer/backend/conf"
	"github.com/weichang-bianjie/irishub-sdk-go"
	"github.com/weichang-bianjie/irishub-sdk-go/types"
	"time"
	"path/filepath"
)

var (
	client sdk.Client
)

func init() {
	path := filepath.Join("/User/user/go/", "test")
	client = sdk.NewClient(types.ClientConfig{
		NodeURI: conf.Get().Hub.NodeUrl,
		Network: types.Testnet,
		ChainID: conf.Get().Hub.ChainId,
		Mode:    types.Commit,
		Timeout: time.Duration(10 * time.Second),
		Fee: []types.DecCoin{
			{Denom: "stake", Amount: types.NewDecimal(1)},
		},
		Gas:       2000000,
		KeyDAO:    types.NewMemoryDB(), //default keybase
		StoreType: types.PrivKey,
		Level:     "info",
		DBRootDir: path,
	})

	//types.SetNetwork(types.Mainnet)

}
