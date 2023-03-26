package user_service

import (
    "context"
    pb "myGoApp/api"
    repo "myGoApp/repository"
)

type Implementation struct {
    pb.UnimplementedUserServiceServer
    repo repo.UserRepo
}

func New(userRepo *repo.UserRepo) *Implementation {
    return &Implementation{
        repo: *userRepo,
    }
}

func (imp *Implementation) GetById(ctx context.Context, in *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {
    userModel := imp.repo.GetUserById(in.Id)

    if userModel == nil {
        return &pb.GetUserByIdResponse{}, nil
    }
    return &pb.GetUserByIdResponse{
        Email: userModel.Email,
        }, nil
}