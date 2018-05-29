/*
   nicehash.go
       Wrapper for the Nicehash API

   Authors:
       Bapt Lombard
*/
package nicehash

const (
	BaseUrl = "https://api.nicehash.com/api"
)

type Nicehash struct {
	client *Client
}

func New(API_id, API_key string) *Nicehash {
	client := NewClient(API_id, API_key)
	return &Nicehash{client}
}
