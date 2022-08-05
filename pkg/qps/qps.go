package qps

import (
	sdk "chainmaker.org/chainmaker-sdk-go"
	"context"
	//sdk ""
	"chainpress/pkg/sdkop"
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

	timeStart := time.Now().UnixNano()

	wg.Add(threadNum)

	for t:= 0; t < threadNum; t++ {
		go syncQps(concurrency, ctx, clients)
	}

	wg.Wait()

	timeCount := threadNum*concurrency
	timeEnd := time.Now().UnixNano()
	count := float64(timeCount)
	timeResult := float64((timeEnd-timeStart)/1e6) / 1000.0
	fmt.Println("ToTalThroughput:", timeCount, "ToTalDuration:", strconv.FormatFloat(timeResult, 'g', 30, 32)+" s", "QPS:", count/timeResult)
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
	}

	wgs.Wait()
	defer wg.Done()
}

func getTxByTxId(client *sdk.ChainClient, txid string, wgs *sync.WaitGroup)  {
	_, err := client.GetTxByTxId(txid)
	if err != nil {
		fmt.Println(err)
	}
	defer wgs.Done()
}