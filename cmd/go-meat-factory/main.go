package main

import (
	"fmt"
	"sync"

	"github.com/billyang3416/go-meat-factory/internal/distributer"
	"github.com/billyang3416/go-meat-factory/internal/meat"
	"github.com/billyang3416/go-meat-factory/internal/supplier"
)

const (
	numOfWorker  = 5
	totalMeat    = 22
	numOfBeef    = 10
	numOfPork    = 7
	numOfChicken = 5
)

var Version = "developement"

func main() {

	fmt.Println("程式版本: ", Version)

	// 1. 進貨肉
	meats := supplier.SupplyMeat(totalMeat, numOfBeef, numOfPork, numOfChicken)

	var wg sync.WaitGroup
	// 使用 buffered channel 增加吞吐量，會比 unbuffered channel 快
	meatCh := make(chan meat.Meat, len(meats))
	// 2. 先啟動員工準備從 receive channel 取出肉來處理，並且處理中或處理完成後都需要印出時間
	distributer.StartWorkers(numOfWorker, meatCh, &wg)

	// 3. 將肉放到 send channel
	distributer.DistributeToWorker(meatCh, meats)

	// 4. 等待所有員工都處理完肉
	wg.Wait()
}
