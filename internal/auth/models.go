package auth

import (
	"context"

	pb "github.com/SmitSheth/Mini-twitter/internal/auth/authentication"
)

type AuthRepository interface {
	CheckAuthentication(context.Context, *pb.UserCredential) (*pb.IsAuthenticated, error)
	AddCredential(context.Context, *pb.UserCredential) (*pb.Void, error)
	GetAuthToken(context.Context, *pb.UserId) (*pb.AuthToken, error)
	RemoveAuthToken(context.Context, *pb.AuthToken) (*pb.Void, error)
	GetUserId(context.Context, *pb.AuthToken) (*pb.UserId, error)
}
