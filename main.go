package main

import "fmt"

// Bookstore

func main() {
	// 连接数据库
	db, err := NewDB("bookstore.db")
	if err != nil {
		fmt.Printf("connect db failed,err:%v\n", err)
		return
	}
	db.DB()
	// 启动gRPC服务
	// 启动gRPC-Gateway服务
}
