# go-nicehash

A Golang wrapper for the [nicehash](https://www.nicehash.com/) API.

## Documentation
Full API Documentation can be found at https://www.nicehash.com/doc-api

All APIs return JSON format. Nicehash suggests to wait at least 3 seconds between API calls. Current API version: 1.2.7

## Installation
```shell
go get github.com/blombard/go-nicehash
```

## Importing

```go
import (
    "github.com/blombard/go-nicehash"
)
```

## Setup

Creating a client:

```go
package main

import (
	"os"

	"github.com/blombard/go-nicehash"
)

//not required if you only use the public API
API_ID    := os.Getenv("API_ID")
API_Key   := os.Getenv("API_Key")
//not required if you don't use a wallet
BTC_addr  := os.Getenv("BTC_addr")

func main() {

    //for the public API
    client :=  nicehash.New("", "", "")

    //for the public API with your wallet
    client :=  nicehash.New("", "", BTC_addr)

    //for the private API
    client :=  nicehash.New(API_ID, API_Key, BTC_addr)

}
```

## PUBLIC

### Get average profitability

```go
// Get average profitability (price) and hashing speed for all algorithms in past 24 hours.
res, err := client.GetAverageProfitability()

if err != nil {
  panic(err)
}
fmt.Println(res)
```

## PRIVATE

### Get current confirmed balance

```go
// Get current confirmed Bitcoin balance.
res, err := client.GetCurrentBalance()

if err != nil {
  panic(err)
}
fmt.Println(res)
```

## Other Examples

See ```examples/main.go``` for other examples.
