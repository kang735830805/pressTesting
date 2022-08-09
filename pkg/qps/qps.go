package qps

import (
	sdk "chainmaker.org/chainmaker-sdk-go"
	"chainpress/pkg/cmd"
	"github.com/panjf2000/ants/v2"

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

	if *cmd.ThreadNum < 1 {
		return fmt.Errorf("threadNum should not less 1")
	}

	fmt.Println("============ application-golang starts ============")

	sdkList := strings.Split(*cmd.SdkPath, ",")

	clients:=make([]*sdk.ChainClient,len(sdkList))

	for i := 0; i <= len(sdkList)-1;i++ {
		clients[i]=sdkop.Connect_chain(sdkList[i])
	}

	pool, _ := ants.NewPoolWithFunc(*cmd.LoopNum, syncQps)

	defer pool.Release()
	timeStart := time.Now().UnixNano()

	for i:=0; i<(*cmd.LoopNum)*(*cmd.ThreadNum); i++ {
		wg.Add(1)

		pool.Invoke(clients)
	}

	wg.Wait()

	timeCount := (*cmd.LoopNum)*(*cmd.ThreadNum)
	timeEnd := time.Now().UnixNano()
	count := float64(timeCount)
	timeResult := float64((timeEnd-timeStart)/1e6) / 1000.0
	fmt.Println("ToTalThroughput:", timeCount, "ToTalDuration:", strconv.FormatFloat(timeResult, 'g', 30, 32)+" s", "QPS:", count/timeResult)
	return err
}


func syncQps(clients interface{}) {
	chainClients := clients.([]*sdk.ChainClient)

	sNum := 0

	getTxByTxId(chainClients[sNum], txId)

	defer wg.Done()
}


func getTxByTxId(client *sdk.ChainClient, txid string) {
		resp, err := client.GetTxByTxId(txid)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(resp)
}
