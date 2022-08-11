package qps

import (
	sdk "chainmaker.org/chainmaker-sdk-go"
	clients2 "chainpress/pkg/clients"
	"chainpress/pkg/cmd"
	"chainpress/pkg/requests"
	"chainpress/pkg/schedule"
	"context"

	"fmt"
	"strconv"
	"sync"
	"time"
)

//var wg sync.WaitGroup

type QPS struct {
	requests.Base
}

func newQps(qps *QPS) *QPS {
	return qps
}

func NewQps(client *sdk.ChainClient, ctx context.Context, cancelFunc context.CancelFunc) []*QPS {
	t := make([]*QPS, *cmd.ThreadNum)
	for i,_ := range t {
		t[i] = initQPSWorker(client, ctx, cancelFunc)
	}
	return t
}

func initQPSWorker(client *sdk.ChainClient, ctx context.Context, cancelFunc context.CancelFunc) *QPS {
	//timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(*cmd.Duration))
	q := newQps(
		&QPS{
			requests.Base{
				Engine:     *schedule.InitEngine(client),
				Wg:         sync.WaitGroup{},
				CancelFunc: cancelFunc,
				CtxFunc:    ctx,
			},
		})
	return q
}

func RunQps() (err error) {
	//defer q.cancelFunc()

	if *cmd.ThreadNum < 1 {
		return fmt.Errorf("threadNum should not less 1")
	}

	fmt.Println("============ application-golang starts ============")

	timeStart := time.Now().UnixNano()
	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(*cmd.Duration))

	qps := newQps(&QPS{
		requests.Base{
			Wg:         sync.WaitGroup{},
			CancelFunc: cancelFunc,
			CtxFunc:    timeoutCtx,
		},
	})
	qps.asyncJobs()

	timeCount := (*cmd.LoopNum)*(*cmd.ThreadNum)
	timeEnd := time.Now().UnixNano()
	count := float64(timeCount)
	timeResult := float64((timeEnd-timeStart)/1e6) / 1000.0
	fmt.Println("ToTalThroughput:", timeCount, "ToTalDuration:", strconv.FormatFloat(timeResult, 'g', 30, 32)+" s", "QPS:", count/timeResult)
	return err
}


func (q *QPS) asyncJobs()  {

	defer q.tearDown()

	clients:=clients2.CreateClient()

	engines := NewQps(clients[0], q.CtxFunc, q.CancelFunc)

	totalBatch, batchCount := int(*cmd.Duration / *cmd.Interval), 0

	ticker := time.NewTicker(time.Duration(*cmd.Interval)*time.Second)

	for ; batchCount < totalBatch; batchCount++ {
		<-ticker.C
		var w sync.WaitGroup
		q.Wg.Add(1)
		go func() {
			timeStart := time.Now().UnixNano()

			for _, e := range engines {
				w.Add(1)
				go e.syncQps(e.Engine.Args, &w)
			}
			w.Wait()
			q.Wg.Done()
			//defer q.tearDown()
			timeEnd := time.Now().UnixNano()
			count := float64((*cmd.LoopNum)*(*cmd.ThreadNum))
			timeResult := float64((timeEnd-timeStart)/1e6) / 1000.0
			fmt.Println(timeResult)
			fmt.Println("Throughput:", count, "Duration:", strconv.FormatFloat(timeResult, 'g', 30, 32)+" s", "QPS:", count/timeResult)
		}()

	}
	q.Wg.Wait()
}


func (q *QPS) syncQps(clients interface{}, w *sync.WaitGroup) {
	chainClients := clients.(*sdk.ChainClient)

	for i:=0; i<*cmd.LoopNum; i++ {
		q.Engine.Wg.Add(1)
		go q.getTxByTxId(chainClients, txId)
	}
	q.Engine.Wg.Wait()
	w.Done()
}

func (q *QPS) getTxByTxId(client *sdk.ChainClient, txid string) {
	_, err := client.GetTxByTxId(txid)
	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Println(resp)
	defer q.Engine.Wg.Done()

}

func (q *QPS) tearDown()  {
	defer q.CancelFunc()
	//defer q.Engine.Close()
}
