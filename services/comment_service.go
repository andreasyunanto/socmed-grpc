package services

import (
	"errors"

	"github.com/andreasyunanto/socmed-grpc/pb"
	"github.com/andreasyunanto/socmed-grpc/repositories"
)

func GetCommentByPost(repo *repositories.CommentRepository, id string) (*pb.GetCommentsResponse, error) {
	operationResult, err := repo.GetCommentByPost(id)

	if err != nil {
		return &pb.GetCommentsResponse{Status: false, Data: nil, Message: err.Error()}, errors.New(err.Error())
	}

	comments := make([]*pb.Comment, 0)
	for _, val := range operationResult {
		comments = append(comments, &pb.Comment{
			CommentId: val.CommentId,
			PostId:    val.PostId,
			Comments:  val.Comments,
		})
	}

	c := pb.CommentData{
		Items: comments,
	}

	return &pb.GetCommentsResponse{
		Message: "OK",
		Data:    &c,
		Status:  true,
	}, nil
}
