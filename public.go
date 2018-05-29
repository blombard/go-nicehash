/*
   public.go
       public methods available without API KEY
*/

package nicehash

import (
	"fmt"
	//"time"
)

type BaseInfo struct {
	Result struct {
		APIVersion string `json:"api_version"`
	} `json:"result"`
	Method interface{} `json:"method"`
}

/*
  Ping the API to see if it's alive

  Parameters:
    - None
*/
func (n *Nicehash) Ping() (info BaseInfo, err error) {

	reqUrl := fmt.Sprintf("")
	_, err = n.client.do("GET", reqUrl, "", false, &info)
	if err != nil {
		return
	}

	return
}

type CurrentGlobalStats struct {
	Result struct {
		Stats []struct {
			ProfitabilityAboveLtc string `json:"profitability_above_ltc,omitempty"`
			Price                 string `json:"price"`
			ProfitabilityLtc      string `json:"profitability_ltc,omitempty"`
			Algo                  int64  `json:"algo"`
			Speed                 string `json:"speed"`
			ProfitabilityBtc      string `json:"profitability_btc,omitempty"`
			ProfitabilityAboveBtc string `json:"profitability_above_btc,omitempty"`
			ProfitabilityEth      string `json:"profitability_eth,omitempty"`
			ProfitabilityAboveEth string `json:"profitability_above_eth,omitempty"`
		} `json:"stats"`
	} `json:"result"`
	Method string `json:"method"`
}

/*
  Get current profitability (price) and hashing speed for all algorithms. Refreshed every 30 seconds.

  Parameters:
    - None
*/
func (n *Nicehash) GetCurrentProfitability() (stats CurrentGlobalStats, err error) {

	reqUrl := fmt.Sprintf("?method=stats.global.current")
	_, err = n.client.do("GET", reqUrl, "", false, &stats)
	if err != nil {
		return
	}

	return
}

/*
  Get current profitability (price) and hashing speed for all algorithms. Refreshed every 30 seconds.

  Parameters:
    - location: 0 for Europe, 1 for USA. This parameter is optional and method will return combined statistics if not provided.
*/
func (n *Nicehash) GetCurrentProfitabilityLocation(location int64) (stats CurrentGlobalStats, err error) {

	reqUrl := fmt.Sprintf("?method=stats.global.current&%d", location)
	_, err = n.client.do("GET", reqUrl, "", false, &stats)
	if err != nil {
		return
	}

	return
}

type AverageProfitability struct {
	Result struct {
		Stats []struct {
			Price string `json:"price"`
			Algo  int64  `json:"algo"`
			Speed string `json:"speed"`
		} `json:"stats"`
	} `json:"result"`
	Method string `json:"method"`
}

/*
  Get average profitability (price) and hashing speed for all algorithms in past 24 hours.

  Parameters:
    - None
*/
func (n *Nicehash) GetAverageProfitability() (stats AverageProfitability, err error) {

	reqUrl := fmt.Sprintf("?method=stats.global.24h")
	_, err = n.client.do("GET", reqUrl, "", false, &stats)
	if err != nil {
		return
	}

	return
}

type AverageStats struct {
	Result struct {
		Stats []struct {
			Balance       string `json:"balance"`
			RejectedSpeed string `json:"rejected_speed"`
			Algo          int64  `json:"algo"`
			AcceptedSpeed string `json:"accepted_speed"`
		} `json:"stats"`
		Payments []struct {
			Amount string `json:"amount"`
			Fee    string `json:"fee"`
			TXID   string `json:"TXID"`
			Time   string `json:"time"`
		} `json:"payments"`
		Addr string `json:"addr"`
	} `json:"result"`
	Method string `json:"method"`
}

/*
  Get current stats for provider for all algorithms. Refreshed every 30 seconds. It also returns past 56 payments.

  Parameters:
    - addr - Provider's BTC address.
*/
func (n *Nicehash) GetAverageStats(btc_addr string) (stats AverageStats, err error) {

	reqUrl := fmt.Sprintf("?method=stats.provider&addr=%s", btc_addr)
	_, err = n.client.do("GET", reqUrl, "", false, &stats)
	if err != nil {
		return
	}

	return
}

type DetailedStats struct {
	Method string `json:"method"`
	Result struct {
		Addr    string `json:"addr"`
		Current []struct {
			Algo          int64         `json:"algo"`
			Data          []interface{} `json:"data"`
			Name          string        `json:"name"`
			Profitability string        `json:"profitability"`
			Suffix        string        `json:"suffix"`
		} `json:"current"`
		NhWallet bool `json:"nh_wallet"`
		Past     []struct {
			Algo int64           `json:"algo"`
			Data [][]interface{} `json:"data"`
		} `json:"past"`
		Payments []interface{} `json:"payments"`
	} `json:"result"`
}

/*
  Get detailed stats for provider for all algorithms including history data and payments in past 7 days.

  Parameters:
    - addr - Provider's BTC address.
    - from - Get history data from this time (UNIX timestamp). Put 0 by default and it returns the complete history.
*/
func (n *Nicehash) GetDetailedStats(btc_addr string, timestamp int64) (stats DetailedStats, err error) {

	reqUrl := fmt.Sprintf("?method=stats.provider.ex&addr=%s", btc_addr)
	_, err = n.client.do("GET", reqUrl, "", false, &stats)

	if err != nil {
		return
	}

	return
}

type WorkerStats struct {
	Result struct {
		Addr    string          `json:"addr"`
		Workers [][]interface{} `json:"workers"`
		Algo    int64           `json:"algo"`
	} `json:"result"`
	Method string `json:"method"`
}

/*
  Get detailed stats for provider's workers (rigs).

  Parameters:
    - addr - Provider's BTC address.
    - algo - Algorithm marked with ID.
*/
func (n *Nicehash) GetWorkerStats(btc_addr string, algo int64) (stats WorkerStats, err error) {

	reqUrl := fmt.Sprintf("?method=stats.provider.workers&addr=%s&algo=%d", btc_addr, algo)
	_, err = n.client.do("GET", reqUrl, "", false, &stats)

	if err != nil {
		return
	}

	return
}

type AllOrders struct {
	Result struct {
		Orders []struct {
			LimitSpeed    string `json:"limit_speed"`
			Alive         bool   `json:"alive"`
			Price         string `json:"price"`
			ID            int    `json:"id"`
			Type          int    `json:"type"`
			Workers       int    `json:"workers"`
			Algo          int    `json:"algo"`
			AcceptedSpeed string `json:"accepted_speed"`
		} `json:"orders"`
		Timestamp int `json:"timestamp"`
	} `json:"result"`
	Method string `json:"method"`
}

/*
  Get all orders for certain algorithm. Refreshed every 30 seconds.

  Parameters:
    - location - 0 for Europe (NiceHash), 1 for USA (WestHash);
    - algo - Algorithm marked with ID.
*/
func (n *Nicehash) GetAllOrders(location int64, algo int64) (orders AllOrders, err error) {

	reqUrl := fmt.Sprintf("?method=orders.get&location=%d&algo=%d", location, algo)
	_, err = n.client.do("GET", reqUrl, "", false, &orders)

	if err != nil {
		return
	}

	return
}

type MutiAlgo struct {
	Result struct {
		Multialgo []struct {
			DefaultFactor float64 `json:"default_factor"`
			Port          int64   `json:"port"`
			Name          string  `json:"name"`
			Algo          int64   `json:"algo"`
		} `json:"multialgo"`
	} `json:"result"`
	Method string `json:"method"`
}

/*
  Get information about Mult-Algorithm Mining.

  Parameters:
    - None
*/
func (n *Nicehash) GetMutiAlgo() (info MutiAlgo, err error) {

	reqUrl := fmt.Sprintf("?method=multialgo.info")
	_, err = n.client.do("GET", reqUrl, "", false, &info)

	if err != nil {
		return
	}

	return
}

type SimpleMutiAlgo struct {
	Result struct {
		Simplemultialgo []struct {
			Paying string `json:"paying"`
			Port   int64  `json:"port"`
			Name   string `json:"name"`
			Algo   int64  `json:"algo"`
		} `json:"simplemultialgo"`
	} `json:"result"`
	Method string `json:"method"`
}

/*
  Get information about Simple Multi-Algorithm Mining. Read more : https://www.nicehash.com/software-developers

  Parameters:
    - None
*/
func (n *Nicehash) GetSimpleMutiAlgo() (info SimpleMutiAlgo, err error) {

	reqUrl := fmt.Sprintf("?method=multialgo.info")
	_, err = n.client.do("GET", reqUrl, "", false, &info)

	if err != nil {
		return
	}

	return
}

type InfoHashingPower struct {
	Result struct {
		Algorithms []struct {
			DownStep  string `json:"down_step"`
			MinLimit  string `json:"min_limit"`
			SpeedText string `json:"speed_text"`
			Name      string `json:"name"`
			Algo      int64  `json:"algo"`
			Multi     string `json:"multi"`
		} `json:"algorithms"`
		DownTime int64 `json:"down_time"`
	} `json:"result"`
	Method string `json:"method"`
}

/*
  Get needed information for buying hashing power using NiceHashBot : https://github.com/nicehash/NiceHashBot

  Parameters:
    - None
*/
func (n *Nicehash) GetInfoHashingPower() (info InfoHashingPower, err error) {

	reqUrl := fmt.Sprintf("?method=buy.info")
	_, err = n.client.do("GET", reqUrl, "", false, &info)

	if err != nil {
		return
	}

	return
}
