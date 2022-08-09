package qps

import (
	sdk "chainmaker.org/chainmaker/sdk-go/v2"
	"chainpress/pkg/sdkop"
	"fmt"
	"github.com/panjf2000/ants/v2"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func RunQps() (err error) {

	if threadNum < 1 {
		return fmt.Errorf("threadNum should not less 1")
	}

	fmt.Println("============ application-golang starts ============")
	sdkList := strings.Split(sdkPath, ",")

	clients:=make([]*sdk.ChainClient, len(sdkList))

	for i := 0; i <= len(sdkList)-1;i++ {
		clients[i]=sdkop.Connect_chain(sdkList[i])
	}

	pool, _ := ants.NewPoolWithFunc(loopNum, syncQps)

	defer pool.Release()
	timeStart := time.Now().UnixNano()

	//wg.Add(threadNum*loopNum)

	for i:=0; i<loopNum*threadNum; i++ {
		wg.Add(1)

		pool.Invoke(clients)
	}

	//if err != nil {
	//	fmt.Errorf(err.Error())
	//}

	wg.Wait()

	timeEnd := time.Now().UnixNano()
	count := float64(threadNum*loopNum)
	timeResult := float64((timeEnd-timeStart)/1e6) / 1000.0
	fmt.Println(timeResult)
	fmt.Println("Throughput:", count, "Duration:", strconv.FormatFloat(timeResult, 'g', 30, 32)+" s", "QPS:", count/timeResult)

	return err
}


func syncQps(clients interface{}) {
	//var wgs sync.WaitGroup
	chainClients := clients.([]*sdk.ChainClient)

	sNum := 0
	//wgs.Add(1)
	//for i := 0 ; i < loopNum; i++ {
	//	if sNum > len(chainClients)-1 {
	//		sNum = 0
	//	}
	getTxByTxId(chainClients[sNum], txId)
	//sNum++
	//fmt.Println(i)
	//}
	//wgs.Wait()

	defer wg.Done()
}
//
//func syncQps(num int, ctx context.Context, clients []*sdk.ChainClient) {
//	var wgs sync.WaitGroup
//
//	sNum := 0
//	wgs.Add(num)
//	for i := 0 ; i < num; i++ {
//		if sNum > len(clients)-1 {
//			sNum = 0
//		}
//		go getTxByTxId(clients[sNum], txId, &wgs)
//		sNum++
//		fmt.Println(i)
//	}
//	wgs.Wait()
//
//	defer wg.Done()
//}


//func getTxByTxId(client *sdk.ChainClient, txid string, wgs *sync.WaitGroup) {
func getTxByTxId(client *sdk.ChainClient, txid string) {
	resp, err := client.GetTxByTxId(txid)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(resp)
}
