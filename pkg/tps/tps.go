package tps

import (
	sdk "chainmaker.org/chainmaker-sdk-go"
	"chainpress/pkg/sdkop"
	"context"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}


func InvoceChaincode(client *sdk.ChainClient, name, method string , args map[string]string){
	sdkop.UserContractAssetInvoke(client, name, method, "1", "", args,false) //最后一个参数为是否同步获取交易结果？
}


func RunTps() (err error) {
	if method == "" || name == "" || sdkPath == "" {
		return fmt.Errorf("method 、 name、sdkpath not nil")
	}

	fmt.Println("============ application-golang starts ============")
	ctx := context.Background()
	sdkList := strings.Split(sdkPath, ",")

	clients:=make([]*sdk.ChainClient,len(sdkList))
	for i := 0; i <= len(sdkList)-1;i++ {
		fmt.Println(i)
		clients[i]=sdkop.Connect_chain(sdkList[i])
	}

	timeStart := time.Now().UnixNano()

	for i := 0; i < loop; i = i+(threadNum*concurrency) {
		// todo 进程处理进程内部交易的逻辑

		if loop <= i+(threadNum*concurrency) && loop > i  {
			wg.Add(int(math.Floor(float64((loop-i)/concurrency))))

			for t := 0; t < int(math.Floor(float64((loop-i)/concurrency))); t++ {
				go syncTps(concurrency, ctx, clients)
			}
		} else if loop > i+(threadNum*concurrency) {
			wg.Add(threadNum)

			for t:= 0; t < threadNum; t++ {
				go syncTps(concurrency, ctx, clients)
			}
		} else {
			wg.Add(int(math.Floor(float64((loop-i)/concurrency))))

			for t:= 0; t < int(math.Floor(float64((loop-i)/concurrency))); t++ {
				go syncTps(concurrency, ctx, clients)
			}
		}
		wg.Wait()
	}


	timeCount := loop
	timeEnd := time.Now().UnixNano()
	count := float64(timeCount)
	timeResult := float64((timeEnd-timeStart)/1e6) / 1000.0
	fmt.Println("Throughput:", timeCount, "Duration:", strconv.FormatFloat(timeResult, 'g', 30, 32)+" s", "TPS:", count/timeResult)
	return err
}


func syncTps(num int, ctx context.Context, clients []*sdk.ChainClient) {

	timeStart := time.Now().UnixNano()
	sNum := 0
	m := make(map[string]string)
	err := json.Unmarshal([]byte(parameter), &m)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	for i := 0 ; i < num; i++ {
		if sNum > len(clients)-1 {
			sNum = 0
		}
		InvoceChaincode(clients[sNum], name, method, m)
		sNum++
	}


	timeCount := num
	timeEnd := time.Now().UnixNano()
	count := float64(timeCount)
	timeResult := float64((timeEnd-timeStart)/1e6) / 1000.0
	fmt.Println("Throughput:", timeCount, "Duration:", strconv.FormatFloat(timeResult, 'g', 30, 32)+" s", "TPS:", count/timeResult)
	wg.Done()
}