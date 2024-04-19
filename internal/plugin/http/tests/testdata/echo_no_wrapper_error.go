// Code generated by GG version dev. DO NOT EDIT.

//go:build !gg
// +build !gg

package server

import (
	"context"
	"fmt"
	prometheus "github.com/prometheus/client_golang/prometheus"
	promhttp "github.com/prometheus/client_golang/prometheus/promhttp"
	"net"
	"net/http"
)

func instrumentRoundTripperErrCounter(counter *prometheus.CounterVec, next http.RoundTripper) promhttp.RoundTripperFunc {
	return func(r *http.Request) (*http.Response, error) {
		resp, err := next.RoundTrip(r)
		if err != nil {
			labels := prometheus.Labels{"method": r.Method}
			errType := ""
			switch e := err.(type) {
			default:
				errType = err.Error()
			case net.Error:
				errType += "net."
				if e.Timeout() {
					errType += "timeout."
				}
				switch ee := e.(type) {
				case *net.ParseError:
					errType += "parse"
				case *net.InvalidAddrError:
					errType += "invalidAddr"
				case *net.UnknownNetworkError:
					errType += "unknownNetwork"
				case *net.DNSError:
					errType += "dns"
				case *net.OpError:
					errType += ee.Net + "." + ee.Op
				}
			}
			labels["err"] = errType
			counter.With(labels).Add(1)
		}
		return resp, err
	}
}

type outgoingInstrumentation struct {
	inflight    prometheus.Gauge
	errRequests *prometheus.CounterVec
	requests    *prometheus.CounterVec
	duration    *prometheus.HistogramVec
	dnsDuration *prometheus.HistogramVec
	tlsDuration *prometheus.HistogramVec
}

func (i *outgoingInstrumentation) Describe(in chan<- *prometheus.Desc) {
	i.inflight.Describe(in)
	i.requests.Describe(in)
	i.errRequests.Describe(in)
	i.duration.Describe(in)
	i.dnsDuration.Describe(in)
	i.tlsDuration.Describe(in)
}
func (i *outgoingInstrumentation) Collect(in chan<- prometheus.Metric) {
	i.inflight.Collect(in)
	i.requests.Collect(in)
	i.errRequests.Collect(in)
	i.duration.Collect(in)
	i.dnsDuration.Collect(in)
	i.tlsDuration.Collect(in)
}

type ClientBeforeFunc func(context.Context, *http.Request) (context.Context, error)
type ClientAfterFunc func(context.Context, *http.Response) context.Context
type clientOptions struct {
	ctx    context.Context
	before []ClientBeforeFunc
	after  []ClientAfterFunc
	client *http.Client
}
type ClientOption func(*clientOptions)

func WithContext(ctx context.Context) ClientOption {
	return func(o *clientOptions) {
		o.ctx = ctx
	}
}
func WithClient(client *http.Client) ClientOption {
	return func(o *clientOptions) {
		o.client = client
	}
}
func WithProm(namespace string, subsystem string, reg prometheus.Registerer, constLabels map[string]string) ClientOption {
	return func(o *clientOptions) {
		i := &outgoingInstrumentation{inflight: prometheus.NewGauge(prometheus.GaugeOpts{Namespace: namespace, Subsystem: subsystem, Name: "in_flight_requests", Help: "A gauge of in-flight outgoing requests for the client.", ConstLabels: constLabels}), requests: prometheus.NewCounterVec(prometheus.CounterOpts{Namespace: namespace, Subsystem: subsystem, Name: "requests_total", Help: "A counter for outgoing requests from the client.", ConstLabels: constLabels}, []string{"method", "code"}), errRequests: prometheus.NewCounterVec(prometheus.CounterOpts{Namespace: namespace, Subsystem: subsystem, Name: "err_requests_total", Help: "A counter for outgoing error requests from the client."}, []string{"method", "err"}), duration: prometheus.NewHistogramVec(prometheus.HistogramOpts{Namespace: namespace, Subsystem: subsystem, Name: "request_duration_histogram_seconds", Help: "A histogram of outgoing request latencies.", Buckets: prometheus.DefBuckets, ConstLabels: constLabels}, []string{"method", "code"}), dnsDuration: prometheus.NewHistogramVec(prometheus.HistogramOpts{Namespace: namespace, Subsystem: subsystem, Name: "dns_duration_histogram_seconds", Help: "Trace dns latency histogram.", Buckets: prometheus.DefBuckets, ConstLabels: constLabels}, []string{"method", "code"}), tlsDuration: prometheus.NewHistogramVec(prometheus.HistogramOpts{Namespace: namespace, Subsystem: subsystem, Name: "tls_duration_histogram_seconds", Help: "Trace tls latency histogram.", Buckets: prometheus.DefBuckets, ConstLabels: constLabels}, []string{"method", "code"})}
		trace := &promhttp.InstrumentTrace{}
		o.client.Transport = instrumentRoundTripperErrCounter(i.errRequests, promhttp.InstrumentRoundTripperInFlight(i.inflight, promhttp.InstrumentRoundTripperCounter(i.requests, promhttp.InstrumentRoundTripperTrace(trace, promhttp.InstrumentRoundTripperDuration(i.duration, o.client.Transport)))))
		err := reg.Register(i)
		if err != nil {
			panic(err)
		}
	}
}
func Before(before ...ClientBeforeFunc) ClientOption {
	return func(o *clientOptions) {
		o.before = append(o.before, before...)
	}
}
func After(after ...ClientAfterFunc) ClientOption {
	return func(o *clientOptions) {
		o.after = append(o.after, after...)
	}
}

type EchoClientNoWrapperErrorControllerClient struct {
	target string
	opts   *clientOptions
}
type EchoClientNoWrapperErrorControllerFooRequest struct {
	c      *EchoClientNoWrapperErrorControllerClient
	client *http.Client
	opts   *clientOptions
	params struct {
		a string
	}
}

func (r *EchoClientNoWrapperErrorControllerClient) Foo(a string) (err error) {
	err = r.FooRequest(a).Execute()
	return
}
func (r *EchoClientNoWrapperErrorControllerClient) FooRequest(a string) *EchoClientNoWrapperErrorControllerFooRequest {
	m := &EchoClientNoWrapperErrorControllerFooRequest{client: r.opts.client, opts: &clientOptions{ctx: context.TODO()}, c: r}
	m.params.a = a
	return m
}
func (r *EchoClientNoWrapperErrorControllerFooRequest) Execute(opts ...ClientOption) (err error) {
	for _, o := range opts {
		o(r.opts)
	}
	ctx, cancel := context.WithCancel(r.opts.ctx)
	path := fmt.Sprintf("/foo/%s", r.params.a)
	req, err := http.NewRequest("GET", r.c.target+path, nil)
	if err != nil {
		cancel()
		return
	}
	before := append(r.c.opts.before, r.opts.before...)
	for _, before := range before {
		ctx, err = before(ctx, req)
		if err != nil {
			cancel()
			return
		}
	}
	resp, err := r.client.Do(req)
	if err != nil {
		cancel()
		return
	}
	after := append(r.c.opts.after, r.opts.after...)
	for _, after := range after {
		ctx = after(ctx, resp)
	}
	defer resp.Body.Close()
	defer cancel()
	if resp.StatusCode > 399 {
		err = fmt.Errorf("http error %d", resp.StatusCode)
		return
	}
	return nil
}
func NewEchoClientNoWrapperErrorControllerClient(target string, opts ...ClientOption) *EchoClientNoWrapperErrorControllerClient {
	c := &EchoClientNoWrapperErrorControllerClient{target: target, opts: &clientOptions{client: http.DefaultClient}}
	for _, o := range opts {
		o(c.opts)
	}
	return c
}
