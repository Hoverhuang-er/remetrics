package pusher

import (
	"context"
	"fmt"
	"github.com/panjf2000/ants/v2"
	"log"
	"regexp"
	"sync"
)

type PusherFunc interface {
	PushMetrics(ctx context.Context) error
}

type Pusher struct {
	EnableCoroutinesPool    bool
	CoroutinePoolBufferSize int
	CoroutinePoolSize       int
	PushGatewayAddr         string
	MetricsPath             string
}

func (p *Pusher) PushMetrics(ctx context.Context) error {
	// TODO: push metrics to prometheus
	var pl *ants.Pool
	if p.EnableCoroutinesPool {
		if p.CoroutinePoolBufferSize != 0 {
			p.CoroutinePoolSize = p.CoroutinePoolBufferSize + p.CoroutinePoolSize
		}
		pcp, err := ants.NewPool(p.CoroutinePoolSize, ants.WithMaxBlockingTasks(1))
		if err != nil {
			log.Printf("new ants pool error: %v", err)
			return err
		}
		pl = pcp
	}
	errchan := make(chan error, 1)
	wg := sync.WaitGroup{}
	for i := 0; i < p.CoroutinePoolSize; i++ {
		wg.Add(1)
		go func(pi int) {
			if err := pl.Submit(func() {
				// Submit task to pool + ctx + wg + pi
			}); err != nil {
				log.Printf("submit task error: %v", err)
			}
		}(i)
	}
	wg.Wait()
	if cap(errchan) > 0 {
		log.Printf("error channel length: %d, \t %v", cap(errchan), errchan)
		return <-errchan
	}
	return nil
}

func (p *Pusher) buildMetrics(ctx context.Context, svc string) error {
	regPath, err := regexp.Compile(`[/.-]`)
	if err != nil {
		return err
	}
	p.MetricsPath = fmt.Sprintf("dynatrace_%s_", regPath.ReplaceAllString(svc, "_"))
	return nil
}
