package ctps

import (
	sdk "chainmaker.org/chainmaker-sdk-go"
	"chainpress/pkg/sdkop"
	"encoding/json"
	"fmt"
	"strconv"
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
	clients1:=make([]*sdk.ChainClient,concurrency)
	for i := 0; i < concurrency;i++ {
		clients1[i]=sdkop.Connect_chain(sdkPath)
	}
	fmt.Println("============ application-golang starts ============")

	for i := 0; i < concurrency; i = i + threadNum {

		if concurrency < i + threadNum && concurrency > i {
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
	clients:=make([]*sdk.ChainClient,num)

	for i := 0; i < num;i++ {
		clients[i]=sdkop.Connect_chain(sdkPath)
	}
	wg.Add(num)
	m := sync.Map{}
	ma := make(map[string]string)

	err = json.Unmarshal([]byte(parameter), &ma)
	for i := 0 ; i < num; i++ {
		go InvoceChaincode(clients[i], loop, name, method, ma, m)
	}
	timeStart := time.Now().UnixNano()
	wg.Wait()

	//m.Range(func(key, value interface{}) bool {
	//
	//})
	timeCount := loop * concurrency
	timeEnd := time.Now().UnixNano()
	count := float64(timeCount)
	fmt.Println(count)
	timeResult := float64((timeEnd-timeStart)/1e6) / 1000.0
	fmt.Println("Throughput:", timeCount, "Duration:", strconv.FormatFloat(timeResult, 'g', 30, 32)+" s", "TPS:", count/timeResult)
	return nil
}


//func InvoceChaincode(client1,client2 *sdk.ChainClient, loop int, name, method, args string){
func InvoceChaincode(client *sdk.ChainClient, loop int, name, method string ,paramers map[string]string, m sync.Map){
	var txid string = ""
	addr2 := sdkop.UserContractAssetQuery(client, false, name, method, paramers)
	for i := 0; i < loop; i++ {
		txid = sdkop.UserContractAssetInvoke(client, name, method, "1", addr2, paramers,false) //最后一个参数为是否同步获取交易结果？
	}
	wg.Done()
	if txid != "" {
		tran, _ := client.GetTxByTxId(txid)
		m.Store(txid, tran)
	}
}