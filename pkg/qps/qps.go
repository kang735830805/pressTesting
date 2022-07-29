package qps

import (
	sdk "chainmaker.org/chainmaker-sdk-go"
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

	fmt.Println("============ application-golang starts ============")
	timeStart := time.Now().UnixNano()

	for i := 0; i < loop; i = i + threadNum {

		if loop < i + threadNum && loop > i {
			err = syncTps(loop-i)
		} else if loop > threadNum {
			err = syncTps(threadNum)
		} else {
			err = syncTps(loop)
		}

	}
	timeCount := loop
	timeEnd := time.Now().UnixNano()
	count := float64(timeCount)
	//timeResult := float64(timeEnd-timeStart)
	timeResult := float64((timeEnd-timeStart)/1e6) / 1000.0
	fmt.Println("ToTalThroughput:", timeCount, "ToTalDuration:", strconv.FormatFloat(timeResult, 'g', 30, 32)+" s", "AvgTPS:", count/timeResult)
	return err
}

func syncTps(num int) error {
	sdkList := strings.Split(sdkPath, ",")

	clients:=make([]*sdk.ChainClient,len(sdkList))

	for i := 0; i <= len(sdkList)-1;i++ {
		clients[i]=sdkop.Connect_chain(1, sdkList[i])
	}
	wg.Add(num)
	//timeStart := time.Now().UnixNano()
	sNum := 0
	for i := 0 ; i <= num; i++ {
		//go InvoceChaincode(clients[i], loop, name, method, args, m)
		if sNum >= len(sdkList)-1 {
			sNum = 0
		}
		go getTxByTxId(clients[sNum], txId)
		sNum++
	}
	timeStart := time.Now().UnixNano()
	wg.Wait()

	timeCount := loop
	timeEnd := time.Now().UnixNano()
	//timeEnd := time.Now().Unix()
	count := float64(timeCount)
	timeResult := float64((timeEnd-timeStart)/1e6) / 1000.0
	fmt.Println(timeResult)
	fmt.Println("Throughput:", timeCount, "Duration:", strconv.FormatFloat(timeResult, 'g', 30, 32)+" s", "QPS:", count/timeResult)
	return nil
}


func getTxByTxId(client *sdk.ChainClient, txid string)  {
	_, err := client.GetTxByTxId(txid)
	if err != nil {
		fmt.Println(err)
	}
	wg.Done()
}