package main

import (
    "github.com/joho/godotenv"
    pb "myGoApp/api"
    "myGoApp/infrastructure/config"
    "myGoApp/service/auth_service"
    "myGoApp/service/user_service"
    "net"
    "fmt"

    "myGoApp/infrastructure/db"

    "myGoApp/repository"

    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
)

func main() {
    if err := godotenv.Load(".env"); err != nil {
        panic("Error loading .env file")
    }

    env, err := config.GetEnvConfig()

    if err != nil {
        panic(err)
    }

    s := grpc.NewServer()
    reflection.Register(s)

    conn, err := db.GetConn(&env)


    userRepo := repository.NewRepo(conn)

    pb.RegisterUserServiceServer(s, user_service.New(userRepo))
    pb.RegisterAuthServiceServer(s, auth_service.New())

    listener, err := net.Listen("tcp", ":8111")
    if err != nil {
        panic(err)
    }

    if err := s.Serve(listener); err != nil {
        fmt.Printf("failed to serve: %v", err)
    }

    if err != nil {
        panic(err)
    }
}