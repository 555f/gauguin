// Code generated by GG version . DO NOT EDIT.

package metrics

import (
	client "github.com/555f/gg/examples/metrics-middleware/internal/client"
	middleware "github.com/555f/gg/examples/metrics-middleware/internal/middleware"
	prometheus "github.com/prometheus/client_golang/prometheus"
	"time"
)

const fooClientBarMethod = "(github.com/555f/gg/examples/metrics-middleware/internal/client.FooClient).BarMethod"
const fooClientBarMethodShort = "(FooClient).BarMethod"

type FooClientMetricMiddleware struct {
	next        client.FooClient
	inRequests  *prometheus.CounterVec
	requests    *prometheus.CounterVec
	errRequests *prometheus.CounterVec
	duration    *prometheus.HistogramVec
}

func (m *FooClientMetricMiddleware) BarMethod(test string) (n int, err error) {
	m.inRequests.With(prometheus.Labels{"method": fooClientBarMethod, "shortMethod": fooClientBarMethodShort}).Inc()
	defer func(now time.Time) {
		m.requests.With(prometheus.Labels{"method": fooClientBarMethod, "shortMethod": fooClientBarMethodShort}).Inc()
		if err != nil {
			m.errRequests.With(prometheus.Labels{"method": fooClientBarMethod, "shortMethod": fooClientBarMethodShort}).Inc()
		}
		m.duration.With(prometheus.Labels{"method": fooClientBarMethod, "shortMethod": fooClientBarMethodShort}).Observe(time.Since(now).Seconds())
	}(time.Now())
	n, err = m.next.BarMethod(test)
	return
}
func LoggingFooClientMiddleware(namespace string, subsystem string) middleware.FooClientMiddleware {
	return func(next client.FooClient) client.FooClient {
		return &FooClientMetricMiddleware{next: next, inRequests: prometheus.NewCounterVec(prometheus.CounterOpts{Namespace: namespace, Subsystem: subsystem, Name: "in_requests_total", Help: "A counter for incoming requests."}, []string{"method", "shortMethod"}), requests: prometheus.NewCounterVec(prometheus.CounterOpts{Namespace: namespace, Subsystem: subsystem, Name: "requests_total", Help: "A counter for complete requests."}, []string{"method", "shortMethod"}), errRequests: prometheus.NewCounterVec(prometheus.CounterOpts{Namespace: namespace, Subsystem: subsystem, Name: "err_requests_total", Help: "A counter for error requests."}, []string{"method", "shortMethod"}), duration: prometheus.NewHistogramVec(prometheus.HistogramOpts{Namespace: namespace, Subsystem: subsystem, Name: "request_duration_histogram_seconds", Help: "A histogram of outgoing request latencies."}, []string{"method", "shortMethod"})}
	}
}
