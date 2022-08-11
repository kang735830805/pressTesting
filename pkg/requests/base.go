package requests

import (
	"chainpress/pkg/schedule"
	"context"
	"sync"
)

type Base struct {
	Engine schedule.BaseConfig
	Wg sync.WaitGroup
	CancelFunc context.CancelFunc
	CtxFunc context.Context
	Rate int
}