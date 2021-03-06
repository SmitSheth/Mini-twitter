package post

import (
	"context"

	pb "github.com/SmitSheth/Mini-twitter/internal/post/postpb"
)

// PostRepository is the interface for post repositories
type PostRepository interface {
	CreatePost(context.Context, *pb.Post) (uint64, error)
	GetPost(context.Context, uint64) (*pb.Post, error)
	GetPosts(context.Context, []uint64) ([]*pb.Post, error)
	GetPostsByAuthor(context.Context, []uint64) ([]*pb.Post, error)
	UpdatePost(context.Context, pb.Post) error
	DeletePost(context.Context, uint64) error
}
