package clients

import (
	sdk "chainmaker.org/chainmaker-sdk-go"
	"chainpress/pkg/cmd"
	"chainpress/pkg/sdkop"
	"strings"
)

func CreateClient() (clients []*sdk.ChainClient) {

	sdkList := strings.Split(*cmd.SdkPath, ",")

	clients=make([]*sdk.ChainClient,len(sdkList))

	for i := 0; i <= len(sdkList)-1;i++ {
		clients[i]=sdkop.Connect_chain(sdkList[i])
	}
	return clients
}
