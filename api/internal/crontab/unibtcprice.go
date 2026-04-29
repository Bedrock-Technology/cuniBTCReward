package crontab

import (
	"context"
	"cuniBTCReward/api/internal/config"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpc"
)

type BitcoinResp struct {
	Bitcoin struct {
		Usd uint64 `json:"usd"`
	} `json:"universal-btc"`
}

type CoinGeckoUniBTC struct {
	config      config.Config
	uniBTCPrice uint64
	lock        sync.RWMutex
}

func NewCoinGeckoUniBTC(c config.Config) *CoinGeckoUniBTC {
	return &CoinGeckoUniBTC{
		config: c,
	}
}

func (c *CoinGeckoUniBTC) GetUniBTCPrice() uint64 {
	c.lock.RLock()
	price := c.uniBTCPrice
	c.lock.RUnlock()
	return price
}

func (c *CoinGeckoUniBTC) CoinGeckoUniBTCCron() {
	//bitcoin
	params := url.Values{}
	params.Add("vs_currencies", "usd")
	params.Add("ids", "universal-btc")
	params.Add("x_cg_demo_api_key", c.config.CoinGecoKey)
	resp, err := httpc.Do(context.Background(),
		http.MethodGet, fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?%s", params.Encode()), nil)
	if err != nil {
		logx.Errorf("get coin gecko error")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logx.Errorf("not ok")
		return
	}

	// 3. Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logx.Errorf("get coin gecko err")
		return
	}
	bitcoin := BitcoinResp{}
	err = json.Unmarshal(body, &bitcoin)
	if err != nil {
		logx.Errorf("umarshal err")
		return
	}
	c.lock.Lock()
	c.uniBTCPrice = bitcoin.Bitcoin.Usd
	c.lock.Unlock()
	logx.Infof("Set uniBTC price:%d", bitcoin.Bitcoin.Usd)
}
