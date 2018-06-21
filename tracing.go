package integration

import (
	"github.com/opentracing/opentracing-go"
	"github.com/yoonitio/tracer-go"
)

func init() {
	t := tracer.NewTracer()
	opentracing.SetGlobalTracer(t)
}
