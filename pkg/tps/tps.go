package tps

import (
	sdk "chainmaker.org/chainmaker/sdk-go/v2"
	"chainpress/pkg/sdkop"
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}



//func invoceChaincode(client1,client2 *sdk.ChainClient, loop int, name, method, args string){
func invoceChaincode(client *sdk.ChainClient, loop int, name, method, args string){
	//addr1 := sdkop.UserContractAssetQuery(true)  //true 为node1 ，else node0
	//fmt.Printf("client1 address: %s\n", addr1)

	//addr2 := sdkop.UserContractAssetQuery(client1, client2,false, name, method, args)
	addr2 := sdkop.UserContractAssetQuery(client, false, name, method, args)
	for i := 0; i < loop; i++ {
		sdkop.UserContractAssetInvoke(client, name, method, args, "1", addr2, false) //最后一个参数为是否同步获取交易结果？
	}
	wg.Done()
}

//loop 循环次数    每个并发循环次数
//concurrency 连接网关数  并发数
//func RunTps(loop, concurrency int, name, method, args, sdkPath string)  {
func RunTps() error {
	if method == "" || name == "" || sdkPath == "" {
		return fmt.Errorf("method 、 name、sdkpath not nil")
	}
	clients1:=make([]*sdk.ChainClient,concurrency)
	//clients2:=make([]*sdk.ChainClient,concurrency)
	for i := 0; i < concurrency;i++ {
		clients1[i]=sdkop.Connect_chain(1, sdkPath)
		//clients2[i]=sdkop.Connect_chain(2)
	}
	fmt.Println("============ application-golang starts ============")

	wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		//go invoceChaincode(clients1[i], clients2[i], loop, name, method, args)
		go invoceChaincode(clients1[i], loop, name, method, args)
	}

	timeStart := time.Now().UnixNano()
	wg.Wait()

	timeCount := loop * concurrency
	timeEnd := time.Now().UnixNano()
	count := float64(timeCount)
	timeResult := float64((timeEnd-timeStart)/1e6) / 1000.0
	fmt.Println("Throughput:", timeCount, "Duration:", strconv.FormatFloat(timeResult, 'g', 30, 32)+" s", "TPS:", count/timeResult)
	return nil
}
