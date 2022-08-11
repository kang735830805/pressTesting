package qps

import (
	clients2 "chainpress/pkg/clients"
	"chainpress/pkg/cmd"
	"chainpress/pkg/requests"
	"chainpress/pkg/schedule"
	"context"
	"sync"
	"testing"
	"time"
)

func Test(t *testing.T)  {
	ctxFunc, cancelFunc := context.WithTimeout(context.Background(), time.Duration(*cmd.Duration))
	*cmd.SdkPath = "/Users/kangxiye/IdeaProjects/chaos/masterpress/1/pressTesting/sdk_config.yml"
	clients:=clients2.CreateClient()

	qps :=
		&QPS{
		requests.Base{
			Wg: sync.WaitGroup{},
			CancelFunc: cancelFunc,
			CtxFunc: ctxFunc,
			Engine: *schedule.InitEngine(clients[0]),
		},

		}
	qps.asyncJobs()

}
