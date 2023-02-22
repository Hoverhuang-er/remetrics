package pusher

import (
	"context"
	"crypto/tls"
	"fmt"
	grt "github.com/hashicorp/go-retryablehttp"
	"github.com/panjf2000/ants/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"log"
	"net/http"
	"net/http/httptrace"
	"regexp"
	"sync"
	"time"
)

type PusherFunc interface {
	PushMetrics(ctx context.Context) error
}

type Pusher struct {
	EnableCoroutinesPool    bool
	CoroutinePoolBufferSize int
	CoroutinePoolSize       int
	PushGatewayAddr         string
	metricsPath             string
	PushJobName             string
}

var (
	// HTTP Client Trace and Conn reuse
	clientTrace = &httptrace.ClientTrace{
		GotConn: func(info httptrace.GotConnInfo) {
			log.Printf("GotConn: %v reused", info)
		},
	}
	// Rebuild HTTP Client with GRT
	grtc = func() *grt.Client {
		c := grt.NewClient()
		c.HTTPClient = &nhc
		c.RetryWaitMax = 5 * time.Second
		c.RetryWaitMin = 1 * time.Second
		c.RetryMax = 1
		return c
	}
	// Create HTTP Client
	nhc = http.Client{
		Timeout: 15 * time.Second,
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 30,
			TLSClientConfig: &tls.Config{
				MaxVersion:         tls.VersionTLS13,
				InsecureSkipVerify: true,
			}}}
)

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
	p.metricsPath = fmt.Sprintf("%s", regPath.ReplaceAllString(svc, "_"))
	return nil
}

func (p *Pusher) buildPromGaugeVec(ctx context.Context, keys []string) *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace:   "amscreen",
		Subsystem:   "dynatrace",
		Name:        p.metricsPath,
		Help:        fmt.Sprintf("The job use for %s by collector %s", p.metricsPath, "amscreen"),
		ConstLabels: nil,
	}, keys)
}

func (p *Pusher) DoPush(ctx context.Context, clo prometheus.Collector) error {
	ctx = httptrace.WithClientTrace(ctx, clientTrace)
	ph := push.New(p.PushGatewayAddr, p.PushJobName)
	ph.PushContext(ctx)
	ph.Collector(clo)
	ph.Client(grtc().HTTPClient)
	return ph.Push()
}
