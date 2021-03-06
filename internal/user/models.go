package user

import (
	"context"

	pb "github.com/SmitSheth/Mini-twitter/internal/user/userpb"
)

// UserRepository is the interface for user repositories
type UserRepository interface {
	CreateUser(context.Context, *pb.User) (uint64, error)
	GetUser(context.Context, uint64) (*pb.User, error)
	GetUsers(context.Context, []uint64) ([]*pb.User, error)
	GetAllUsers(context.Context) ([]*pb.User, error)
	GetUserByUsername(context.Context, string) (*pb.User, error)
	GetFollowing(context.Context, uint64) ([]*pb.User, error)
	GetNotFollowing(context.Context, uint64) ([]*pb.User, error)
	UpdateUserAccountInfo(context.Context, *pb.AccountInformation) error
	FollowUser(context.Context, uint64, uint64) error
	UnFollowUser(context.Context, uint64, uint64) error
	DeleteUser(context.Context, uint64) error
	NextUserId() (uint64, error)
}
