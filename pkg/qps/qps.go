package qps

import (
	sdk "chainmaker.org/chainmaker/sdk-go/v2"
	"chainpress/pkg/sdkop"
	"context"
	"fmt"
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
	ctx := context.Background()

	fmt.Println("============ application-golang starts ============")
	sdkList := strings.Split(sdkPath, ",")

	clients:=make([]*sdk.ChainClient,len(sdkList))

	for i := 0; i <= len(sdkList)-1;i++ {
		clients[i]=sdkop.Connect_chain(sdkList[i])
	}

	wg.Add(threadNum)

	timeStart := time.Now().UnixNano()

	for t := 0; t < threadNum; t++ {
		go syncQps(concurrency, ctx, clients)
	}

	wg.Wait()
	timeEnd := time.Now().UnixNano()
	count := float64(threadNum*concurrency)
	timeResult := float64((timeEnd-timeStart)/1e6) / 1000.0
	fmt.Println(timeResult)
	fmt.Println("Throughput:", count, "Duration:", strconv.FormatFloat(timeResult, 'g', 30, 32)+" s", "QPS:", count/timeResult)

	return err
}


func syncQps(num int, ctx context.Context, clients []*sdk.ChainClient) {
	var wgs sync.WaitGroup

	sNum := 0
	wgs.Add(num)
	for i := 0 ; i < num; i++ {
		if sNum > len(clients)-1 {
			sNum = 0
		}
		go getTxByTxId(clients[sNum], txId, &wgs)
		sNum++
		fmt.Println(i)
	}
	wgs.Wait()

	defer wg.Done()
}


func getTxByTxId(client *sdk.ChainClient, txid string, wgs *sync.WaitGroup) {
	resp, err := client.GetTxByTxId(txid)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(resp)
	defer wgs.Done()
}
