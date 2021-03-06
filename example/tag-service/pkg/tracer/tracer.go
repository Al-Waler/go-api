package tracer

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"net/http"
	"time"
)

// NewJaegerTracer 创建JaegerTracer 对象及基本配置
func NewJaegerTracer(serviceName, agentHostPort string) (opentracing.Tracer, io.Closer, error) {
	cfg := &config.Configuration{
		//服务名
		ServiceName: serviceName,
		//取样器，采样模式 const 固定采样
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  agentHostPort,
		},
	}
	// 初始化tracer 对象 opentracing.Tracer 并不是某个供应商的tracer 对象
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return nil, nil, err
	}

	// 设置全局tracer 对象
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer, nil
}

type SpanOption func(span opentracing.Span, r interface{})

func SpanWithError(err error) SpanOption {
	return func(span opentracing.Span, r interface{}) {
		if err != nil {
			ext.Error.Set(span, true)
			span.LogFields(log.String("event", "error"), log.String("msg", err.Error()))
		}
	}
}

func SpanWithLog(arg ...interface{}) SpanOption {
	return func(span opentracing.Span, r interface{}) {
		span.LogKV(arg...)
	}
}

func InjectHttp() SpanOption {
	return func(span opentracing.Span, r interface{}) {
		if req, ok := r.(*http.Request); ok {
			_ = opentracing.GlobalTracer().Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))
		}
	}
}

// SetTag r map
func SetTag() SpanOption {
	return func(span opentracing.Span, r interface{}) {
		if m, ok := r.(map[string]string); ok {
			for k, v := range m {
				span.SetTag(k, v)
			}
		}
	}
}

// Start 函数式选项模式 设置tracer span
func Start(tracer opentracing.Tracer, spanName string, ctx context.Context, r interface{}) (newCtx context.Context, finish func(...SpanOption)) {
	if ctx == nil {
		ctx = context.TODO()
	}
	// tag type
	var tagV string
	switch req := r.(type) {
	case *http.Request:
		tagV = req.Proto
	default:
		tagV = "func"
	}
	span, newCtx := opentracing.StartSpanFromContextWithTracer(ctx, tracer, spanName, opentracing.Tag{
		Key:   string(ext.Component),
		Value: tagV,
	})
	finish = func(option ...SpanOption) {
		for _, o := range option {
			o(span, r)
		}
		span.Finish()
	}
	return
}
