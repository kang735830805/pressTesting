package qps

import (
	sdk "chainmaker.org/chainmaker/sdk-go/v2"
	"chainpress/pkg/sdkop"
	"context"
	"fmt"
	"math"
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
	timeStart := time.Now().UnixNano()

	sdkList := strings.Split(sdkPath, ",")

	clients:=make([]*sdk.ChainClient,len(sdkList))

	for i := 0; i <= len(sdkList)-1;i++ {
		clients[i]=sdkop.Connect_chain(sdkList[i])
	}
	for i := 0; i < loop; i = i+(threadNum*concurrency) {
		// todo 进程处理进程内部交易的逻辑
		tNum := threadNum
		con := concurrency
		timeCount := tNum*concurrency

		if loop < i+(threadNum*concurrency) && loop > i  {
			tNum = int(math.Floor(float64((loop-i)/concurrency)))
			wg.Add(tNum)
			timeCount = tNum*concurrency

		} else if loop > i+(threadNum*concurrency) {
			wg.Add(threadNum)
		} else if loop <= i+(threadNum*concurrency) && loop-i < concurrency {
			wg.Add(1)
			con = loop-i
			timeCount = con
			timeCount = 1*concurrency

		} else {
			tNum = int(math.Floor(float64((loop-i)/concurrency)))
			wg.Add(tNum)
			timeCount = con
		}

		for t := 0; t < tNum; t++ {
			go syncQps(con, ctx, clients)
		}

		timeStartLocal := time.Now().UnixNano()
		wg.Wait()

		timeEndLocal := time.Now().UnixNano()
		count := float64(timeCount)
		timeResult := float64((timeEndLocal-timeStartLocal)/1e6) / 1000.0
		fmt.Println(timeResult)
		fmt.Println("Throughput:", timeCount, "Duration:", strconv.FormatFloat(timeResult, 'g', 30, 32)+" s", "QPS:", count/timeResult)
	}

	timeCount := loop
	timeEnd := time.Now().UnixNano()
	count := float64(timeCount)
	timeResult := float64((timeEnd-timeStart)/1e6) / 1000.0
	fmt.Println("ToTalThroughput:", timeCount, "ToTalDuration:", strconv.FormatFloat(timeResult, 'g', 30, 32)+" s", "QPS:", count/timeResult)
	return err
}


func syncQps(num int, ctx context.Context, clients []*sdk.ChainClient) {
	sNum := 0
	for i := 0 ; i <= num; i++ {
		if sNum > len(clients)-1 {
			sNum = 0
		}
		getTxByTxId(clients[sNum], txId)
		sNum++
	}
	wg.Done()
}


func getTxByTxId(client *sdk.ChainClient, txid string)  {
	_, err := client.GetTxByTxId(txid)
	if err != nil {
		fmt.Println(err)
	}
}