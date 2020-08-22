package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/SmitSheth/Mini-twitter/internal/auth/authentication"
	"github.com/SmitSheth/Mini-twitter/internal/auth/service"
	etcd "github.com/SmitSheth/Mini-twitter/internal/auth/storage/etcd"
	"github.com/SmitSheth/Mini-twitter/internal/config"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {

	// Read Config
	runtimeViper := viper.New()
	config.NewRuntimeConfig(runtimeViper, ".")

	// storage := runtimeViper.GetStringSlice("storage.storage")[0]
	storage := "etcd"

	// Start server
	lis, err := net.Listen("tcp", "localhost:"+runtimeViper.GetStringSlice("authservice.ports")[0])
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	fmt.Println("Server running on port", "localhost:"+runtimeViper.GetStringSlice("authservice.ports")[0])
	if storage == "etcd" {
		client, _ := etcd.NewClient([]string{"http://localhost:2379", "http://localhost:22379", "http://localhost:32379"})
		pb.RegisterAuthenticationServer(s, service.GetEtcdAuthServer(client))
	} else {
		pb.RegisterAuthenticationServer(s, service.GetAuthServer())
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
