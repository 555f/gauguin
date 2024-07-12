// Code generated by GG version . DO NOT EDIT.

package middleware

import client "github.com/555f/gg/examples/metrics-middleware/internal/client"

type FooClientMiddleware func(client.FooClient) client.FooClient

func FooClientMiddlewareChain(outer FooClientMiddleware, others ...FooClientMiddleware) FooClientMiddleware {
	return func(next client.FooClient) client.FooClient {
		for i := len(others) - 1; i >= 0; i-- {
			next = others[i](next)
		}
		return outer(next)
	}
}

type fooClientBaseMiddleware struct {
	next     client.FooClient
	mediator any
}

func (m *fooClientBaseMiddleware) BarMethod(test string) (n int, err error) {
	defer func() {
		if s, ok := m.mediator.(fooClientBarMethodBaseMiddleware); ok {
			s.BarMethod(test)
		}
	}()
	return m.next.BarMethod(test)
}

type fooClientBarMethodBaseMiddleware interface {
	BarMethod(test string)
}

func FooClientBaseMiddleware(mediator any) FooClientMiddleware {
	return func(next client.FooClient) client.FooClient {
		return &fooClientBaseMiddleware{next: next, mediator: mediator}
	}
}
