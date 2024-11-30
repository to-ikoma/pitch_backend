package main

import (
	"context"
	"database/sql"
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

func main() {

	mux := http.NewServeMux()
	mux.Handle(svc.NewHelloServiceHandler(&helloService{}))

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
