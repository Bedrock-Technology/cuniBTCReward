package crontab

import (
	"context"
	"cuniBTCReward/api/internal/config"
	"encoding/json"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpc"
)

type BitcoinResp struct {
	Data []struct {
		ID    int     `json:"id"`
		Price float64 `json:"price"`
	} `json:"data"`
	Status struct {
		Timestamp    time.Time `json:"timestamp"`
		ErrorCode    string    `json:"error_code"`
		ErrorMessage string    `json:"error_message"`
		Elapsed      int       `json:"elapsed"`
		CreditCount  int       `json:"credit_count"`
	} `json:"status"`
}
type Request struct {
	Ids []int `form:"ids"`
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
	//uniBTC,ids=36175
	request := Request{
		Ids: []int{36175},
	}
	resp, err := httpc.Do(context.Background(),
		http.MethodGet, "https://pro-api.coinmarketcap.com/public-api/v1/simple/price", request)
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
	if len(bitcoin.Data) == 0 {
		logx.Errorf("data len=0")
		return
	}
	c.lock.Lock()
	c.uniBTCPrice = uint64(bitcoin.Data[0].Price)
	c.lock.Unlock()
	logx.Infof("Set uniBTC price:%d", uint64(bitcoin.Data[0].Price))
}
