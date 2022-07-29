package tps

import (
	sdk "chainmaker.org/chainmaker-sdk-go"
	"chainpress/pkg/sdkop"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}


//func InvoceChaincode(client1,client2 *sdk.ChainClient, loop int, name, method, args string){
func InvoceChaincode(client *sdk.ChainClient, loop int, name, method string , args map[string]string){
	addr2 := sdkop.UserContractAssetQuery(client, false, name, method, args)
	for i := 0; i < loop; i++ {
		sdkop.UserContractAssetInvoke(client, name, method, "1", addr2, args,false) //最后一个参数为是否同步获取交易结果？
	}
	wg.Done()
}

//loop 循环次数    每个并发循环次数
//concurrency 连接网关数  并发数
//func RunTps(loop, concurrency int, name, method, args, sdkPath string)  {
func RunTps() (err error) {
	if method == "" || name == "" || sdkPath == "" {
		return fmt.Errorf("method 、 name、sdkpath not nil")
	}

	fmt.Println("============ application-golang starts ============")

	for i := 0; i < concurrency; i= i+threadNum {

		if concurrency < i + threadNum && concurrency > i  {
			err = syncTps(concurrency-i)
		} else if concurrency > threadNum {
			err = syncTps(threadNum)
		} else {
			err = syncTps(concurrency)
		}
	}

	return err
}


func syncTps(num int) (err error) {
	sdkList := strings.Split(sdkPath, ",")

	clients:=make([]*sdk.ChainClient,len(sdkList))
	fmt.Println(len(sdkList)-1)
	for i := 0; i <= len(sdkList)-1;i++ {
		fmt.Println(i)
		clients[i]=sdkop.Connect_chain(1, sdkList[i])
	}
	wg.Add(num)
	timeStart := time.Now().UnixNano()
	sNum := 0
	m := make(map[string]string)
	err = json.Unmarshal([]byte(args), &m)
	for i := 0 ; i <= num; i++ {
		if sNum >= len(sdkList)-1 {
			sNum = 0
		}
		go InvoceChaincode(clients[sNum], loop, name, method, m)
		sNum++
	}

	wg.Wait()

	timeCount := loop * concurrency
	timeEnd := time.Now().UnixNano()
	count := float64(timeCount)
	timeResult := float64((timeEnd-timeStart)/1e6) / 1000.0
	//timeResult := float64(timeEnd-timeStart)
	fmt.Println("Throughput:", timeCount, "Duration:", strconv.FormatFloat(timeResult, 'g', 30, 32)+" s", "TPS:", count/timeResult)
	return err
}