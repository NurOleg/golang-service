package main

import (
    pb "myGoApp/api"
    "myGoApp/service/user_service"
    "net"
    "fmt"

    "myGoApp/infrastructure/db"

    "myGoApp/repository"

    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
)

func main() {
    listener, err := net.Listen("tcp", ":8111")
    if err != nil {
        panic(err)
    }

    s := grpc.NewServer()
    reflection.Register(s)

    conn, err := db.GetConn()


    userRepo := repository.NewRepo(conn)
    userService := user_service.New(userRepo)

    pb.RegisterUserServiceServer(s, userService)

    if err := s.Serve(listener); err != nil {
        fmt.Printf("failed to serve: %v", err)
    }

    if err != nil {
        panic(err)
    }
}