/*
   private.go
       private methods available with an API ID and an API Key or ReadOnly API Key.
*/

package nicehash

import (
	"fmt"
)

// CurrentBalance : struct to handle callback from GetCurrentBalance
type CurrentBalance struct {
	Result struct {
		BalancePending   string `json:"balance_pending"`
		BalanceConfirmed string `json:"balance_confirmed"`
	} `json:"result"`
	Method string `json:"method"`
}

/*
  Get current confirmed Bitcoin balance.

  Parameters:
    - id - API ID
	- key - API Key or ReadOnly API Key
*/
func (n *Nicehash) GetCurrentBalance() (info CurrentBalance, err error) {

	reqUrl := fmt.Sprintf("?method=balance")
	_, err = n.client.do("GET", reqUrl, "private", &info)
	if err != nil {
		return
	}

	return
}

// AllMyOrders : struct to handle callback from GetAllMyOrders
type AllMyOrders struct {
	Result struct {
		Orders []struct {
			Type          int     `json:"type"`
			BtcAvail      string  `json:"btc_avail"`
			LimitSpeed    string  `json:"limit_speed"`
			PoolUser      string  `json:"pool_user"`
			PoolPort      int     `json:"pool_port"`
			Alive         bool    `json:"alive"`
			Workers       int     `json:"workers"`
			PoolPass      string  `json:"pool_pass"`
			AcceptedSpeed string  `json:"accepted_speed"`
			ID            int     `json:"id"`
			Algo          int     `json:"algo"`
			Price         string  `json:"price"`
			BtcPaid       string  `json:"btc_paid"`
			PoolHost      string  `json:"pool_host"`
			End           int64   `json:"end"`
		} `json:"orders"`
	} `json:"result"`
	Method string `json:"method"`
}

/*
  Get all orders for certain algorithm owned by the customer. Refreshed every 30 seconds.

  Parameters:
    - id - API ID;
		- key - API Key or ReadOnly API Key;
		- location - 0 for Europe (NiceHash), 1 for USA (WestHash);
		- algo - Algorithm marked with ID.
*/
func (n *Nicehash) GetAllMyOrders(location, algo int64) (orders AllMyOrders, err error) {

	reqUrl := fmt.Sprintf("?method=orders.get&my&location=%d&algo=%d", location, algo)
	_, err = n.client.do("GET", reqUrl, "private", &orders)
	if err != nil {
		return
	}

	return
}

// OrderCallBack : struct to handle callback from Orders
type OrderCallBack struct {
	Result struct {
		Success string `json:"success"`
		Error 	string `json:"error,omitempty"`
	} `json:"result"`
	Method string `json:"method"`
}

/*
  Create new order. Only standard orders can be created with use of API.

  Parameters:
		- id - API ID;
		- key - API Key;
		- location - 0 for Europe (NiceHash), 1 for USA (WestHash);
		- algo - Algorithm marked with ID;
		- amount - Pay amount in BTC;
		- price - Price in BTC/GH/day or BTC/TH/day;
		- limit - Speed limit in GH/s or TH/s (0 for no limit);
		- pool_host - Pool hostname or IP;
		- pool_port - Pool port;
		- pool_user - Pool username;
		- pool_pass - Pool password;
		- code - This parameter is optional. You have to provide it if you have 2FA enabled.
*/
func (n *Nicehash) CreateNewOrder(location, algo int64, amount, price, limit float64, pool_host, pool_port, pool_user, pool_pass string, code int64) (order OrderCallBack, err error) {

	reqUrl := fmt.Sprintf("?method=orders.create&location=%d&algo=%d&amount=%f&price=%f&limit=%f&pool_host=%s&pool_port=%s&pool_user=%s&pool_pass=%s", location, algo, amount, price, limit, pool_host, pool_port, pool_user, pool_pass)
	if code != 0 {
		reqUrl += fmt.Sprintf("code=%d", code)
	}
	_, err = n.client.do("POST", reqUrl, "private", &order)
	if err != nil {
		return
	}

	return
}

/*
  Refill order with extra Bitcoins.

  Parameters:
		- id - API ID;
		- key - API Key;
		- location - 0 for Europe (NiceHash), 1 for USA (WestHash);
		- algo - Algorithm marked with ID;
		- order - Order ID;
		- amount - Pay amount in BTC;
*/
func (n *Nicehash) RefillOrder(location, algo, orderID int64, amount float64) (order OrderCallBack, err error) {

	reqUrl := fmt.Sprintf("?method=orders.refill&location=%d&algo=%d&order=%d&amount=%f", location, algo, orderID, amount)
	_, err = n.client.do("POST", reqUrl, "private", &order)
	if err != nil {
		return
	}

	return
}

/*
  Remove existing order.

  Parameters:
		- id - API ID;
		- key - API Key;
		- location - 0 for Europe (NiceHash), 1 for USA (WestHash);
		- algo - Algorithm marked with ID;
		- order - Order ID;
*/
func (n *Nicehash) RemoveOrder(location, algo, orderID int64) (order OrderCallBack, err error) {

	reqUrl := fmt.Sprintf("?method=orders.remove&location=%d&algo=%d&order=%d", location, algo, orderID)
	_, err = n.client.do("POST", reqUrl, "private", &order)
	if err != nil {
		return
	}

	return
}

/*
  Set new price for the existing order. Only increase is possible.

  Parameters:
		- id - API ID;
		- key - API Key;
		- location - 0 for Europe (NiceHash), 1 for USA (WestHash);
		- algo - Algorithm marked with ID;
		- order - Order ID;
		- price - Price in BTC/GH/Day or BTC/TH/Day.
*/
func (n *Nicehash) SetHigherPriceOrder(location, algo, orderID int64, price float64) (order OrderCallBack, err error) {

	reqUrl := fmt.Sprintf("?method=orders.set.price&location=%d&algo=%d&order=%d&price=%f", location, algo, orderID, price)
	_, err = n.client.do("POST", reqUrl, "private", &order)
	if err != nil {
		return
	}

	return
}

/*
  Decrease price for the existing order. Price decrease possible every 10 minutes. Read help for more information. (https://www.nicehash.com/help)

  Parameters:
		- id - API ID;
		- key - API Key;
		- location - 0 for Europe (NiceHash), 1 for USA (WestHash);
		- algo - Algorithm marked with ID;
		- order - Order ID;
*/
func (n *Nicehash) SetLowerPriceOrder(location, algo, orderID int64) (order OrderCallBack, err error) {

	reqUrl := fmt.Sprintf("?method=orders.set.price.decrease&location=%d&algo=%d&order=%d", location, algo, orderID)
	_, err = n.client.do("POST", reqUrl, "private", &order)
	if err != nil {
		return
	}

	return
}

/*
	Set new limit for the existing order.

  Parameters:
		- id - API ID;
		- key - API Key;
		- location - 0 for Europe (NiceHash), 1 for USA (WestHash);
		- algo - Algorithm marked with ID;
		- order - Order ID;
		- limit - Speed limit in GH/s or TH/s (0 for no limit).
*/
func (n *Nicehash) SetOrderLimit(location, algo, orderID int64, limit float64) (order OrderCallBack, err error) {

	reqUrl := fmt.Sprintf("?method=orders.set.limit&location=%d&algo=%d&order=%d&limit=%f", location, algo, orderID, limit)
	_, err = n.client.do("POST", reqUrl, "private", &order)
	if err != nil {
		return
	}

	return
}
