# go-nicehash

A Golang wrapper for the [nicehash](https://www.nicehash.com/) API.

## Documentation
Full API Documentation can be found at https://www.nicehash.com/doc-api

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
API_ID := os.Getenv("API_ID")
API_Key    := os.Getenv("API_Key")

func main() {

    client :=  nicehash.New(API_ID, API_Key)

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

TODO
