package main

import (
	"log"
	"net/http"

	"pitch_backend/infrastructure/psql/psqlrepo"
	"pitch_backend/interceptor"
	"pitch_backend/usecase"

	"pitch_backend/presentation/controller"
	svc "pitch_backend/presentation/pb/api/pitchbackend/v1/pitchbackendv1connect"

	"connectrpc.com/connect"
	_ "github.com/lib/pq"
	"github.com/samber/do"
	"golang.org/x/net/http2"
)

func main() {
	injector := do.New()
	do.Provide(injector, psqlrepo.NewHelloRepository)
	do.Provide(injector, usecase.NewHelloUsecase)
	do.Provide(injector, controller.NewHelloController)

	mux := http.NewServeMux()
	mux.Handle(svc.NewHelloServiceHandler(
		do.MustInvoke[svc.HelloServiceHandler](injector),
		connect.WithInterceptors(interceptor.NewPanicRecoverInterCeptor(), interceptor.NewLoggingInterCeptor()),
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
