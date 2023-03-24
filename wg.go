package gcurl

import (
	"context"
	"sync"
)

type Wg struct {
	sync.WaitGroup
}

func (w *Wg) RunApi(ctx1 context.Context, callback func(ctx2 context.Context)) {
	w.Add(1)
	go func(ctx context.Context) {
		defer func() {
			w.Done()
		}()
		callback(ctx)
	}(ctx1)
}

func NewWg() *Wg {
	return &Wg{}
}
