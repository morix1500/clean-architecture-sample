package main

import (
	"fmt"
	pb "github.com/morix1500/clean-architecture-sample/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50021", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	ctx := context.Background()
	ireq := &pb.InsertRequest{
		Id: 1,
		Title: "hogeTitle",
		Content: "hogeContent",
	}
	_, err = c.Insert(ctx, ireq)
	if err == nil {
		fmt.Println("success")
	} else {
		fmt.Println("insert failed", err)
	}

	sreq := &pb.SelectRequest{
		Id: 1,
	}
	b, err := c.Select(ctx, sreq)
	if err == nil {
		fmt.Println(b)
	} else {
		fmt.Println("select failed", err)
	}
}
