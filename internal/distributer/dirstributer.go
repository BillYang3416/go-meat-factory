package distributer

import (
	"sync"

	"github.com/billyang3416/go-meat-factory/internal/meat"
	"github.com/billyang3416/go-meat-factory/internal/worker"
)

// 不用管理 state，所以使用 package-level function 即可

func DistributeToWorker(meatCh chan<- meat.Meat, meats []meat.Meat) {

	for _, m := range meats {
		meatCh <- m
	}

	close(meatCh)
}

func StartWorkers(workerCount int, meatCh <-chan meat.Meat, wg *sync.WaitGroup) {
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker.StartWorker(i, meatCh, wg)
	}
}
