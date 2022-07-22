package server

import (
	"context"

	"github.com/andreasyunanto/socmed-grpc/pb"
	"github.com/andreasyunanto/socmed-grpc/repositories"
	"github.com/andreasyunanto/socmed-grpc/services"
)

type PostServer struct{}

func (s *PostServer) GetPostById(ctx context.Context, req *pb.GetPostRequest) (*pb.GetPostResponse, error) {

	postId := req.PostId

	return services.GetPostById(&repositories.PostRepository{}, postId)

}
