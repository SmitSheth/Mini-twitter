package main

import (
	"fmt"
	"log"
	"net"

	"github.com/SmitSheth/Mini-twitter/internal/config"
	"github.com/SmitSheth/Mini-twitter/internal/post"
	pb "github.com/SmitSheth/Mini-twitter/internal/post/postpb"
	"github.com/SmitSheth/Mini-twitter/internal/post/storage/etcd"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	// Read Config
	runtimeViper := viper.New()
	config.NewRuntimeConfig(runtimeViper, ".")
	postStorage, _ := etcd.NewClient([]string{"http://localhost:2379", "http://localhost:22379", "http://localhost:32379"})
	defer postStorage.Close()
	postRepo := etcd.NewPostRepository(postStorage)
	// Start server
	lis, err := net.Listen("tcp", "localhost:"+runtimeViper.GetStringSlice("postservice.ports")[0])
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	fmt.Println("Server running on port", "localhost:"+runtimeViper.GetStringSlice("postservice.ports")[0])
	pb.RegisterPostServiceServer(s, post.GetPostServiceServer(&postRepo))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
