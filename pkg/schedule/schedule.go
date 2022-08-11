package schedule

import (
	sdk "chainmaker.org/chainmaker-sdk-go"
	"chainpress/pkg/cmd"
	"context"
	"sync"
	"time"
)

type CallBack func(i interface{})

type BaseConfig struct {
	Batch    int64
	Interval time.Duration
	Duration time.Duration `mapstructure:"duration"`
	// Wg Semaphore of localWorker
	Wg *sync.WaitGroup
	Args interface{}

	timeoutCtx context.Context
	cancelFunc context.CancelFunc
}

func InitEngine(client *sdk.ChainClient) *BaseConfig {
	var clientWg sync.WaitGroup
	engine := NewEngine(&BaseConfig{
		Batch: int64(*cmd.LoopNum),
		Interval: time.Duration(*cmd.Interval),
		Duration: time.Duration(*cmd.Duration),
		Wg: &clientWg,
		Args: client,
	})
	return engine
}

func NewEngine(config *BaseConfig) *BaseConfig {
	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), config.Duration)
	config.cancelFunc = cancelFunc
	config.timeoutCtx = timeoutCtx
	return config
}

func (b *BaseConfig) Schedule(back CallBack)  {
	defer b.Close()

	totalBatch, batchCount := int(b.Duration/b.Interval), 0
	tick := time.NewTicker(b.Interval*time.Second)
	defer func() {
		tick.Stop()
	}()
	for ; batchCount < totalBatch; batchCount++ {
		<-tick.C
		b.Wg.Add(1)
		for i:=int64(0); i<b.Batch;i++ {
			go back(b.Args)
		}
	}
	b.Wg.Wait()
}

func (b *BaseConfig) Close()  {
	b.cancelFunc()
}
