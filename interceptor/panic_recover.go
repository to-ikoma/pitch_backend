package interceptor

import (
	"context"
	"fmt"
	"log"

	"connectrpc.com/connect"
)

type panicRecoverInterCeptor struct{}

func NewPanicRecoverInterCeptor() *panicRecoverInterCeptor {
	return &panicRecoverInterCeptor{}
}

func (i panicRecoverInterCeptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return func(ctx context.Context, req connect.AnyRequest) (res connect.AnyResponse, err error) {
		defer func() {
			if r := recover(); r != nil {
				log.Print("Panic recovered: %v", r)

				err = connect.NewError(connect.CodeInternal, fmt.Errorf("internal server error: %v", r))
			}
		}()
		return next(ctx, req)
	}
}

func (i panicRecoverInterCeptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return connect.StreamingHandlerFunc(func(ctx context.Context, conn connect.StreamingHandlerConn) (err error) {
		defer func() {
			if r := recover(); r != nil {
				log.Print("Panic recovered: %v", r)

				err = connect.NewError(connect.CodeInternal, fmt.Errorf("internal server error: %v", r))
			}
		}()
		return next(ctx, conn)
	})
}

func (i panicRecoverInterCeptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return connect.StreamingClientFunc(func(
		ctx context.Context,
		spec connect.Spec,
	) connect.StreamingClientConn {
		conn := next(ctx, spec)
		return conn
	})
}
