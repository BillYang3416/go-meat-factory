package worker

import (
	"fmt"
	"sync"
	"time"

	"github.com/billyang3416/go-meat-factory/internal/meat"
)

func StartWorker(ID int, meatCh <-chan meat.Meat, wg *sync.WaitGroup) {
	defer wg.Done()

	for m := range meatCh {
		// 使用員工編號和肉的編號來確認沒有人取到相同的肉
		fmt.Printf("%d 號員工在 %v 取得編號為 %d 的 %v\n", ID+1, time.Now().Format("2006-01-02 15:04:05"), m.ID, m.Type)

		// 模擬處理肉的時間
		time.Sleep(meat.ProcessingTime(m.Type))

		fmt.Printf("%d 號員工在 %v 處理完編號為 %d 的 %v\n", ID+1, time.Now().Format("2006-01-02 15:04:05"), m.ID, m.Type)
	}
}
