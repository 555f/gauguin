// Code generated by GG version . DO NOT EDIT.

package middleware

import (
	controller "github.com/555f/gg/examples/jsonrpc-service/internal/usecase/controller"
	dto "github.com/555f/gg/examples/jsonrpc-service/pkg/dto"
)

type ProfileControllerMiddleware func(controller.ProfileController) controller.ProfileController

func ProfileControllerMiddlewareChain(outer ProfileControllerMiddleware, others ...ProfileControllerMiddleware) ProfileControllerMiddleware {
	return func(next controller.ProfileController) controller.ProfileController {
		for i := len(others) - 1; i >= 0; i-- {
			next = others[i](next)
		}
		return outer(next)
	}
}

type profileControllerBaseMiddleware struct {
	next     controller.ProfileController
	mediator any
}

func (m *profileControllerBaseMiddleware) Create(token string, firstName string, lastName string, address string) (profile *dto.Profile, err error) {
	defer func() {
		if s, ok := m.mediator.(profileControllerCreateBaseMiddleware); ok {
			s.Create(token, firstName, lastName, address)
		}
	}()
	return m.next.Create(token, firstName, lastName, address)
}
func (m *profileControllerBaseMiddleware) Remove(id string) (err error) {
	defer func() {
		if s, ok := m.mediator.(profileControllerRemoveBaseMiddleware); ok {
			s.Remove(id)
		}
	}()
	return m.next.Remove(id)
}

type profileControllerCreateBaseMiddleware interface {
	Create(token string, firstName string, lastName string, address string)
}
type profileControllerRemoveBaseMiddleware interface {
	Remove(id string)
}

func ProfileControllerBaseMiddleware(mediator any) ProfileControllerMiddleware {
	return func(next controller.ProfileController) controller.ProfileController {
		return &profileControllerBaseMiddleware{next: next, mediator: mediator}
	}
}
