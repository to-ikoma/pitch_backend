package interceptor

import (
	"context"
	"log"

	"connectrpc.com/connect"
)

type loggingInterCeptor struct{}

func NewLoggingInterCeptor() *loggingInterCeptor {
	return &loggingInterCeptor{}
}

func (i loggingInterCeptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
		log.Printf("Request: %s, Addr: %s", req.Spec().Procedure, req.Peer().Addr)
		return next(ctx, req)
	}
}

func (i loggingInterCeptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return connect.StreamingHandlerFunc(func(ctx context.Context, conn connect.StreamingHandlerConn) error {
		log.Printf("Request: %s, Addr: %s", conn.Spec().Procedure, conn.Peer().Addr)
		return next(ctx, conn)
	})
}

func (i loggingInterCeptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return connect.StreamingClientFunc(func(
		ctx context.Context,
		spec connect.Spec,
	) connect.StreamingClientConn {
		conn := next(ctx, spec)
		return conn
	})
}

// 基本
// type Middleware func(http.handler) http.handler

// func middlewareLogging(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log.Printf("Request: %s %s", r.Method, r.URL.Path)
// 		next.ServeHTTP(w, r)
// 	})
// }

// func panicRecoverMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		defer func() {
// 			if r := recover(); r != nil {
// 				log.Print("Panic recovered: %v", r)
// 			}
// 		}()
// 		next.ServeHTTP(w, r)
// 	})
// }

// main()
// path, helloHandler := svc.NewHelloServiceHandler(&helloService{})
// mux.Handle(path, middlewareLogging(helloHandler))
