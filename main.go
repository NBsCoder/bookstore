package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc/credentials/insecure"

	"github.com/NBsCoder/bookstore/pb"

	"google.golang.org/grpc"
)

// Bookstore

func main() {
	// 连接数据库
	db, err := NewDB("bookstore.db")
	if err != nil {
		fmt.Printf("connect db failed,err:%v\n", err)
		return
	}
	// 创建server
	srv := server{
		bs: &bookstore{db: db},
	}
	// 启动gRPC服务
	l, err1 := net.Listen("tcp", ":8972")
	if err1 != nil {
		fmt.Printf("Listen to 8972 failed! err:%v\n", err1)
		return
	}
	s := grpc.NewServer()
	// 注册服务
	pb.RegisterBookstoreServer(s, &srv)
	go func() {
		fmt.Println(s.Serve(l))
	}()
	// 启动gRPC-Gateway服务
	conn, err2 := grpc.DialContext(
		context.Background(),
		"127.0.0.1:8972",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err2 != nil {
		fmt.Printf("grpc conn failed! err:%v\n", err2)
		return
	}
	gwmux := runtime.NewServeMux()
	pb.RegisterBookstoreHandler(context.Background(), gwmux, conn)
	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	fmt.Println("grpc-gateway server on :8090")

	err3 := gwServer.ListenAndServe()
	if err3 != nil {
		fmt.Printf("gwServer listenAndServer to :8090 failed! err:%v\n", err3)
		return
	}
}
