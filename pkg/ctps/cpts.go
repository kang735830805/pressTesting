package ctps

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

func RunCTps() (err error) {
	if method == "" || name == "" || sdkPath == "" {
		return fmt.Errorf("method 、 name、sdkpath not nil")
	}
	if threadNum < 1 {
		return fmt.Errorf("threadNum should not less 1")
	}

	fmt.Println("============ application-golang starts ============")
	ctx := context.Background()
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

		for t:= 0; t < tNum; t++ {
			go syncTps(concurrency, ctx, clients)
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

	fmt.Println("Throughput:", timeCount, "Duration:", strconv.FormatFloat(timeResult, 'g', 30, 32)+" s", "TPS:", count/timeResult)
	return err
}

func syncTps(num int, ctx context.Context, clients []*sdk.ChainClient) error {

	m := sync.Map{}

	sNum := 0
	for i := 0 ; i < num; i++ {
		if sNum > len(clients)-1 {
			sNum = 0
		}
		InvoceChaincode(clients[i], loop, name, method, parameter, m)
		sNum++
	}

	return nil
}


func InvoceChaincode(client *sdk.ChainClient, loop int, name, method, args string, m sync.Map){
	var txid string = ""
	for i := 0; i < loop; i++ {
		txid = sdkop.UserContractAssetInvoke(client, name, method, args, "1", "", false) //最后一个参数为是否同步获取交易结果？
	}
	wg.Done()
	if txid != "" {
		tran, _ := client.GetTxByTxId(txid)
		m.Store(txid, tran)
	}
}
