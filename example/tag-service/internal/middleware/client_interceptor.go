package middleware

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"tag-service/global"
	"tag-service/pkg/metatext"
	"time"
)

func defaultContextTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	var cancel context.CancelFunc
	//ctx.Deadline() 检查是否设置时间
	if _, ok := ctx.Deadline(); !ok {
		defaultTimeout := 1 * time.Second
		ctx, cancel = context.WithTimeout(ctx, defaultTimeout)
	}
	return ctx, cancel
}

func UnaryContextTimeout() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		ctx, cancel := defaultContextTimeout(ctx)
		if cancel != nil {
			defer cancel()
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func StreamContextTimeout() grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		ctx, cancel := defaultContextTimeout(ctx)
		if cancel != nil {
			defer cancel()
		}
		return streamer(ctx, desc, cc, method, opts...)
	}
}

func ClientTracing() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		var parentCtx opentracing.SpanContext
		var spanOpts []opentracing.StartSpanOption
		var parentSpan = opentracing.SpanFromContext(ctx)
		if parentSpan != nil {
			parentCtx = parentSpan.Context()
			spanOpts = append(spanOpts, opentracing.ChildOf(parentCtx))
		}
		spanOpts = append(spanOpts, []opentracing.StartSpanOption{
			opentracing.Tag{Key: string(ext.Component), Value: "grpc"},
			ext.SpanKindRPCClient,
		}...)
		span := global.Tracer.StartSpan(method, spanOpts...)
		defer global.Closer.Close() //2 推出结束前发送agent
		defer span.Finish() //1

		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		}
		_ = global.Tracer.Inject(span.Context(), opentracing.TextMap, metatext.MetadataTextMap{md})
		newCtx := opentracing.ContextWithSpan(metadata.NewOutgoingContext(ctx, md), span)
		return invoker(newCtx, method, req, reply, cc, opts...)
	}
}
