package auth_service

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	pb "myGoApp/api"
)

type Implementation struct {
	pb.UnimplementedAuthServiceServer
}

func New() *Implementation {
	return &Implementation{}
}

func (imp *Implementation) GetJwtToken(ctx context.Context, in *pb.GetJwtTokenRequest) (*pb.GetJwtTokenResponse, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": in.Email,
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