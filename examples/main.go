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

	// PUBLIC METHODS

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

	fmt.Println(res1)

	time.Sleep(3 * time.Second)

	res2, err := client.GetCurrentProfitabilityLocation(0)

	if err != nil {
		panic(err)
	}

	fmt.Println(res2)
	time.Sleep(3 * time.Second)

	res3, err := client.GetAverageProfitability()

	if err != nil {
		panic(err)
	}

	fmt.Println(res3)
	time.Sleep(3 * time.Second)

	res31, err := client.GetCurrentStats()

	if err != nil {
		panic(err)
	}

	fmt.Println(res31)
	time.Sleep(3 * time.Second)

	res4, err := client.GetDetailedStats(0)

	if err != nil {
		panic(err)
	}

	fmt.Println(res4)
	time.Sleep(3 * time.Second)

	res5, err := client.GetWorkerStats(24)

	if err != nil {
		panic(err)
	}

	fmt.Println(res5)
	time.Sleep(3 * time.Second)

	res6, err := client.GetAllOrders(0, 5)

	if err != nil {
		panic(err)
	}

	fmt.Println(res6)
	time.Sleep(3 * time.Second)

	res7, err := client.GetMutiAlgo()

	if err != nil {
		panic(err)
	}

	fmt.Println(res7)
	time.Sleep(3 * time.Second)

	res8, err := client.GetSimpleMutiAlgo()

	if err != nil {
		panic(err)
	}

	fmt.Println(res8)
	time.Sleep(3 * time.Second)

	res9, err := client.GetInfoHashingPower()

	if err != nil {
		panic(err)
	}

	fmt.Println(res9)
	time.Sleep(3 * time.Second)

	// PRIVATE methods

	res10, err := client.GetCurrentBalance()
	if err != nil {
		panic(err)
	}

	fmt.Println(res10)
	time.Sleep(3 * time.Second)

	// 0 for EU, 0 for the 'scrypt' algo
	res11, err := client.GetAllMyOrders(0, 0)
	if err != nil {
		panic(err)
	}

	fmt.Println(res11)
	time.Sleep(3 * time.Second)

	res12, err := client.CreateNewOrder(0, 0, 0.01, 2.9, 0.0, "testpool.com", "3333", "worker", "x", 0)
	if err != nil {
		panic(err)
	}

	fmt.Println(res12)
	time.Sleep(3 * time.Second)

	res13, err := client.RefillOrder(0, 0, 123, 0.01)
	if err != nil {
		panic(err)
	}

	fmt.Println(res13)
	time.Sleep(3 * time.Second)

	res14, err := client.RemoveOrder(0, 0, 1880)
	if err != nil {
		panic(err)
	}

	fmt.Println(res14)
	time.Sleep(3 * time.Second)

	res15, err := client.SetHigherPriceOrder(0, 0, 1881, 2.1)
	if err != nil {
		panic(err)
	}

	fmt.Println(res15)
	time.Sleep(3 * time.Second)

	res16, err := client.SetLowerPriceOrder(0, 0, 1881)
	if err != nil {
		panic(err)
	}

	fmt.Println(res16)
	time.Sleep(3 * time.Second)

	res17, err := client.SetOrderLimit(0, 0, 1881, 1.0)
	if err != nil {
		panic(err)
	}

	fmt.Println(res17)
}
