package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	pb "pitch_backend/presentation/pb/api/pitchbackend/v1"
	svc "pitch_backend/presentation/pb/api/pitchbackend/v1/pitchbackendv1connect"

	"connectrpc.com/connect"
	_ "github.com/lib/pq"
	"golang.org/x/net/http2"
)

type helloService struct{}

func NewHelloService() (svc.HelloServiceHandler, error) {
	return &helloService{}, nil
}

func (s *helloService) GetHello(ctx context.Context, req *connect.Request[pb.GetHelloRequest]) (*connect.Response[pb.GetHelloResponse], error) {
	databaseURL := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	var greeting string
	err = db.QueryRow("SELECT 'Hello!'").Scan(&greeting)
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}

	res := &pb.GetHelloResponse{
		Message: greeting,
	}
	return connect.NewResponse(res), nil

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

func main() {

	mux := http.NewServeMux()
	// path, helloHandler := svc.NewHelloServiceHandler(&helloService{})
	// mux.Handle(path, middlewareLogging(helloHandler))
	mux.Handle(svc.NewHelloServiceHandler(
		&helloService{},
		connect.WithInterceptors(NewPanicRecoverInterCeptor(), NewLoggingInterCeptor()),
	))

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	http2.ConfigureServer(server, &http2.Server{})

	log.Println("Starting server on :8080")
	if err := server.ListenAndServeTLS("certs/server.crt", "certs/server.key"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
