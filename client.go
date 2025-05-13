package memogram

import (
	v1pb "github.com/usememos/memos/proto/gen/api/v1"
	"google.golang.org/grpc"
)

type MemosClient struct {
	WorkspaceService v1pb.WorkspaceServiceClient
	AuthService      v1pb.AuthServiceClient
	UserService      v1pb.UserServiceClient
	MemoService      v1pb.MemoServiceClient
	ResourceService  v1pb.ResourceServiceClient
}

func NewMemosClient(conn *grpc.ClientConn) *MemosClient {
	return &MemosClient{
		WorkspaceService: v1pb.NewWorkspaceServiceClient(conn),
		AuthService:      v1pb.NewAuthServiceClient(conn),
		UserService:      v1pb.NewUserServiceClient(conn),
		MemoService:      v1pb.NewMemoServiceClient(conn),
		ResourceService:  v1pb.NewResourceServiceClient(conn),
	}
}
