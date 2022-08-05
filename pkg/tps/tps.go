package tps

import (
	"chainmaker.org/chainmaker/pb-go/v2/common"
	sdk "chainmaker.org/chainmaker/sdk-go/v2"
	"chainpress/pkg/sdkop"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
	"unsafe"
)

var wg = sync.WaitGroup{}


func InvoceChaincode(client *sdk.ChainClient, name, method string, kvs []*common.KeyValuePair, wgs *sync.WaitGroup){
	sdkop.UserContractAssetInvoke(client, name, method, kvs, "1", "", false) //最后一个参数为是否同步获取交易结果？
	wgs.Done()
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
		clients[i]=sdkop.Connect_chain(sdkList[i])
	}

	wg.Add(threadNum)

	timeStart := time.Now().UnixNano()

	for t:= 0; t < threadNum; t++ {
		go syncTps(concurrency, ctx, clients)
	}

	wg.Wait()

	timeEnd := time.Now().UnixNano()
	count := float64(threadNum*concurrency)
	timeResult := float64((timeEnd-timeStart)/1e6) / 1000.0
	fmt.Println(timeResult)
	fmt.Println("Throughput:", count, "Duration:", strconv.FormatFloat(timeResult, 'g', 30, 32)+" s", "TPS:", count/timeResult)

	return err
}


func syncTps(num int, ctx context.Context, clients []*sdk.ChainClient) {
	var wgs sync.WaitGroup
	m := make(map[string]string)
	err := json.Unmarshal([]byte(parameter), &m)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	kvs := []*common.KeyValuePair{}

	for k,v := range m {
		kvs = append(kvs, &common.KeyValuePair{
			Key: k,
			Value: *(*[]byte)(unsafe.Pointer(&v)),
		})
	}

	sNum := 0
	wgs.Add(num)
	for i := 0 ; i < num; i++ {
		if sNum > len(clients)-1 {

			sNum = 0
		}
		go InvoceChaincode(clients[sNum], name, method, kvs, &wgs)
		sNum++
	}
	wgs.Wait()
	defer wg.Done()
}
