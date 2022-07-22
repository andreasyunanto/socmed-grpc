package server

import (
	"context"

	"github.com/andreasyunanto/socmed-grpc/pb"
	"github.com/andreasyunanto/socmed-grpc/repositories"
	"github.com/andreasyunanto/socmed-grpc/services"
)

type CommentServer struct{}

func (s *CommentServer) GetCommentByPost(ctx context.Context, req *pb.GetCommentsRequest) (*pb.GetCommentsResponse, error) {

	postId := req.PostId

	return services.GetCommentByPost(&repositories.CommentRepository{}, postId)

}
