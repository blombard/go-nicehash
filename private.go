/*
   private.go
       private methods available with an API ID and an API Key or ReadOnly API Key.
*/

package nicehash

import (
	"fmt"
)

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
    - id - API ID;
	- key - API Key or ReadOnly API Key.
*/
func (n *Nicehash) GetCurrentBalance() (info CurrentBalance, err error) {

	addUrl := fmt.Sprintf("?method=balance")
	_, err = n.client.do("GET", "", addUrl, true, &info)
	if err != nil {
		return
	}

	return
}
