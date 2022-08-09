package tps

import (
	sdk "chainmaker.org/chainmaker-sdk-go"
	"chainpress/pkg/cmd"
	"chainpress/pkg/sdkop"
	"encoding/json"
	"fmt"
	"github.com/panjf2000/ants/v2"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}


//func InvoceChaincode(client *sdk.ChainClient, name, method string , args map[string]string, wgs *sync.WaitGroup){
func InvoceChaincode(client *sdk.ChainClient, name, method string , args map[string]string){
	sdkop.UserContractAssetInvoke(client, name, method, "1", "", args,false) //最后一个参数为是否同步获取交易结果？
}


func RunTps() (err error) {
	if method == "" || name == "" || *cmd.SdkPath == "" {
		return fmt.Errorf("method 、 name、sdkpath not nil")
	}

	fmt.Println("============ application-golang starts ============")
	sdkList := strings.Split(*cmd.SdkPath, ",")

	clients:=make([]*sdk.ChainClient,len(sdkList))
	for i := 0; i <= len(sdkList)-1;i++ {
		fmt.Println(i)
		clients[i]=sdkop.Connect_chain(sdkList[i])
	}
	pool, _ := ants.NewPoolWithFunc(*cmd.LoopNum, syncTps)
	defer pool.Release()
	timeStart := time.Now().UnixNano()

	for i:=0; i<(*cmd.LoopNum)*(*cmd.ThreadNum); i++ {
		wg.Add(1)

		pool.Invoke(clients)
	}

	wg.Wait()

	timeEnd := time.Now().UnixNano()
	timeCount := (*cmd.LoopNum)*(*cmd.ThreadNum)
	count := float64(timeCount)
	timeResult := float64((timeEnd-timeStart)/1e6) / 1000.0
	fmt.Println(timeResult)
	fmt.Println("Throughput:", timeCount, "Duration:", strconv.FormatFloat(timeResult, 'g', 30, 32)+" s", "TPS:", count/timeResult)
	return err
}


func syncTps(clients interface{}) {

	chainClients := clients.([]*sdk.ChainClient)

	sNum := 0
	m := make(map[string]string)
	err := json.Unmarshal([]byte(parameter), &m)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	InvoceChaincode(chainClients[sNum], name, method, m)
	defer wg.Done()
}
