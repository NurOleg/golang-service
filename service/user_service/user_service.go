package user_service

import (
    "context"
    //"github.com/davecgh/go-spew/spew"
    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"
    pb "myGoApp/api"
    "myGoApp/models"
    repo "myGoApp/repository"
    "errors"
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

func (imp *Implementation) GetJwtToken(ctx context.Context, in *pb.GetJwtTokenRequest) (*pb.GetJwtTokenResponse, error) {
    authRes, err := imp.auth(in.Email, in.Password)

    if err != nil {
        return nil, err
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "email": authRes.Email,
        })

    hmacSampleSecret := []byte("golang-service-test")

    tokenString, err := token.SignedString(hmacSampleSecret)

    if err != nil {
        return nil, err
    }

    return &pb.GetJwtTokenResponse{
        Token: tokenString,
        }, nil
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

func (imp *Implementation) auth(email string, pass string) (*models.User, error) {
    userModel := imp.repo.GetUserByEmail(email)

    if userModel == nil {
        return nil, errors.New("User this creds doesnt exists")
    }

    err := bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(pass))

    if err != nil {
        return nil, errors.New("User this creds doesnt exists")
    }

    return userModel, nil
}