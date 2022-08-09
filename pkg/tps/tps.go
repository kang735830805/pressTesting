package tps

import (
	"chainmaker.org/chainmaker/pb-go/v2/common"
	sdk "chainmaker.org/chainmaker/sdk-go/v2"
	"chainpress/pkg/sdkop"
	"encoding/json"
	"fmt"
	"github.com/panjf2000/ants/v2"
	"strconv"
	"strings"
	"sync"
	"time"
	"unsafe"
)

var wg = sync.WaitGroup{}


func InvoceChaincode(client *sdk.ChainClient, name, method string, kvs []*common.KeyValuePair, wgs *sync.WaitGroup){
	sdkop.UserContractAssetInvoke(client, name, method, kvs, "1", "", false) //最后一个参数为是否同步获取交易结果？
}


func RunTps() (err error) {

	if method == "" || name == "" || sdkPath == "" {
		return fmt.Errorf("method 、 name、sdkpath not nil")
	}

	fmt.Println("============ application-golang starts ============")

	sdkList := strings.Split(sdkPath, ",")

	clients:=make([]*sdk.ChainClient,len(sdkList))
	for i := 0; i <= len(sdkList)-1;i++ {
		clients[i]=sdkop.Connect_chain(sdkList[i])
	}
	pool, _ := ants.NewPoolWithFunc(loopNum, syncTps)
	defer pool.Release()

	timeStart := time.Now().UnixNano()
	for i:=0; i<loopNum*threadNum; i++ {
		wg.Add(1)

		pool.Invoke(clients)
	}

	wg.Wait()

	timeEnd := time.Now().UnixNano()
	count := float64(threadNum*loopNum)
	timeResult := float64((timeEnd-timeStart)/1e6) / 1000.0
	fmt.Println(timeResult)
	fmt.Println("Throughput:", count, "Duration:", strconv.FormatFloat(timeResult, 'g', 30, 32)+" s", "TPS:", count/timeResult)

	return err
}


func syncTps(clients interface{}) {
	chainClients := clients.([]*sdk.ChainClient)

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

	sdkop.UserContractAssetInvoke(chainClients[sNum], name, method, kvs, "1", "", false)

	defer wg.Done()
}
