package tps

import (
	sdk "chainmaker.org/chainmaker-sdk-go"
	clients2 "chainpress/pkg/clients"
	"chainpress/pkg/cmd"
	"chainpress/pkg/requests"
	"chainpress/pkg/schedule"
	"chainpress/pkg/sdkop"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"
)

type TPS struct {
	requests.Base
}

//var wg = sync.WaitGroup{}
func newTPS(tps *TPS) *TPS {
	return tps
}

func initTPSWorker(client *sdk.ChainClient) *TPS {
	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(*cmd.Duration))
	t := newTPS(
		&TPS{
			requests.Base{
				Engine:     *schedule.InitEngine(client),
				Wg:         sync.WaitGroup{},
				CancelFunc: cancelFunc,
				CtxFunc:    timeoutCtx,
			},
		})
	return t
}

func NewTps(client *sdk.ChainClient) []*TPS {
	t := make([]*TPS, *cmd.ThreadNum)
	for i,_ := range t {
		t[i] = initTPSWorker(client)
	}
	return t
}


//func InvoceChaincode(client *sdk.ChainClient, name, method string , args map[string]string, wgs *sync.WaitGroup){
func (t *TPS) InvoceChaincode(client *sdk.ChainClient, name, method string , args map[string]string){
	sdkop.UserContractAssetInvoke(client, name, method, "1", "", args,false) //最后一个参数为是否同步获取交易结果？
	defer t.Engine.Wg.Done()

}


func RunTps() (err error) {
	if method == "" || name == "" || *cmd.SdkPath == "" {
		return fmt.Errorf("method 、 name、sdkpath not nil")
	}

	fmt.Println("============ application-golang starts ============")
	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(*cmd.Duration))
	timeStart := time.Now().UnixNano()

	tps := newTPS(&TPS{
		requests.Base{
			Wg:         sync.WaitGroup{},
			CancelFunc: cancelFunc,
			CtxFunc:    timeoutCtx,
		},
	})
	tps.asyncJobs()

	timeEnd := time.Now().UnixNano()
	timeCount := (*cmd.LoopNum)*(*cmd.ThreadNum)
	count := float64(timeCount)
	timeResult := float64((timeEnd-timeStart)/1e6) / 1000.0
	fmt.Println(timeResult)
	fmt.Println("Throughput:", timeCount, "Duration:", strconv.FormatFloat(timeResult, 'g', 30, 32)+" s", "TPS:", count/timeResult)
	return err
}


//func (t *TPS) syncTps(clients interface{}, w *sync.WaitGroup, kvs []*common.KeyValuePair) {
//	chainClients := clients.(*sdk.ChainClient)
//
//	for i:=0; i<*cmd.LoopNum; i++ {
//		t.Engine.Wg.Add(1)
//		go t.InvoceChaincode(chainClients, name, method, kvs)
//	}
//	t.Engine.Wg.Wait()
//	w.Done()
//}

func  (t *TPS) syncTps(clients interface{}, w *sync.WaitGroup, p map[string]string) {

	chainClients := clients.(*sdk.ChainClient)

	for i:=0; i<*cmd.LoopNum; i++ {
		t.Engine.Wg.Add(1)
		go t.InvoceChaincode(chainClients, name, method, p)
	}
	t.Engine.Wg.Wait()
	w.Done()}


func (t *TPS) asyncJobs()  {

	defer t.tearDown()

	clients:=clients2.CreateClient()
	engines := NewTps(clients[0])
	totalBatch, batchCount := int(*cmd.Duration / *cmd.Interval), 0

	m := make(map[string]string)
	err := json.Unmarshal([]byte(parameter), &m)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	//
	p := make(map[string]string)
	err = json.Unmarshal([]byte(parameter), &p)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	ticker := time.NewTicker(time.Duration(*cmd.Interval)*time.Second)

	for ; batchCount < totalBatch; batchCount++ {
		<-ticker.C
		var w sync.WaitGroup
		t.Wg.Add(1)
		go func() {
			timeStart := time.Now().UnixNano()

			for _, e := range engines {
				w.Add(1)
				go e.syncTps(e.Engine.Args, &w, p)
			}
			w.Wait()
			t.Wg.Done()
			//defer t.tearDown()
			timeEnd := time.Now().UnixNano()
			count := float64((*cmd.LoopNum)*(*cmd.ThreadNum))
			timeResult := float64((timeEnd-timeStart)/1e6) / 1000.0
			fmt.Println(timeResult)
			fmt.Println("Throughput:", count, "Duration:", strconv.FormatFloat(timeResult, 'g', 30, 32)+" s", "QPS:", count/timeResult)
		}()
	}
	t.Wg.Wait()
}

func (t *TPS) tearDown()  {
	defer t.CancelFunc()
	//defer t.Engine.Close()
}