package helloapi

import (
	"context"
	"fmt"
	hello "hello-goakit/gen/hello"

	"github.com/go-kit/kit/log"
)

// hello service example implementation.
// The example methods log the requests and return zero values.
type hellosrvc struct {
	logger log.Logger
}

// NewHello returns the hello service implementation.
func NewHello(logger log.Logger) hello.Service {
	return &hellosrvc{logger}
}

// Hello implements hello.
func (s *hellosrvc) Hello(ctx context.Context, p *hello.HelloPayload) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("hello.hello"))
	return "Hello, " + p.Name + ".", nil
}
