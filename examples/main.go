/*
   main.go
       Example
*/

package main

import (
	"fmt"
	"time"

	"github.com/blombard/go-nicehash"
)

var (
	ID       = "123456"
	KEY      = "a8dd01d7-a4de-44ad-8e8f-ddaa7a2e0bd4"
	BTC_addr = "17a212wdrvEXWuipCV5gcfxdALfMdhMoqh"
)

func main() {

	client := nicehash.New(ID, KEY, BTC_addr)

	res0, err := client.GetCurrentBalance()
	if err != nil {
		panic(err)
	}

	fmt.Println(res0)

	time.Sleep(3 * time.Second)

	res, err := client.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println(res)

	time.Sleep(3 * time.Second)

	res1, err := client.GetCurrentProfitability()

	if err != nil {
		panic(err)
	}

	fmt.Println(res1, "\n \n")

	time.Sleep(3 * time.Second)

	res2, err := client.GetCurrentProfitabilityLocation(0)

	if err != nil {
		panic(err)
	}

	fmt.Println(res2, "\n \n")
	time.Sleep(3 * time.Second)

	res3, err := client.GetAverageProfitability()

	if err != nil {
		panic(err)
	}

	fmt.Println(res3, "\n \n")
	time.Sleep(3 * time.Second)

	res31, err := client.GetCurrentStats()

	if err != nil {
		panic(err)
	}

	fmt.Println(res31, "\n \n")
	time.Sleep(3 * time.Second)

	res4, err := client.GetDetailedStats(0)

	if err != nil {
		panic(err)
	}

	fmt.Println(res4, "\n \n")
	time.Sleep(3 * time.Second)

	res5, err := client.GetWorkerStats(24)

	if err != nil {
		panic(err)
	}

	fmt.Println(res5, "\n \n")
	time.Sleep(3 * time.Second)

	res6, err := client.GetAllOrders(0, 5)

	if err != nil {
		panic(err)
	}

	fmt.Println(res6, "\n \n")
	time.Sleep(3 * time.Second)

	res7, err := client.GetMutiAlgo()

	if err != nil {
		panic(err)
	}

	fmt.Println(res7, "\n \n")
	time.Sleep(3 * time.Second)

	res8, err := client.GetSimpleMutiAlgo()

	if err != nil {
		panic(err)
	}

	fmt.Println(res8, "\n \n")
	time.Sleep(3 * time.Second)

	res9, err := client.GetInfoHashingPower()

	if err != nil {
		panic(err)
	}

	fmt.Println(res9, "\n \n")
	time.Sleep(3 * time.Second)
}
