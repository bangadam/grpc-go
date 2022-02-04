package main

import (
	"context"
	"fmt"
	"log"

	"github.com/bangadam/grpc-go/blog/blogpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Blog Client")

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect to Blog : %v", err)
	}

	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	// create Blog
	CreateBlog(c)

	// read Blog
	ReadBlog(c)
}

// create Blog
func CreateBlog(c blogpb.BlogServiceClient) error {
	// create blog
	fmt.Println("Creating the blog")
	blog := &blogpb.Blog{
		AuthorId: "Bangadam",
		Content:  "Content of the blog",
		Title:    "Title of the blog",
	}
	createBlogRes, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Printf("Blog has been created successfully: %v", createBlogRes)
	return nil
}

// read Blog
func ReadBlog(c blogpb.BlogServiceClient) error {
	// read blog
	fmt.Println("Reading the blog")
	blogID := "61fd3f589718940e8c70389e"
	_, err := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{Id: blogID})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}

	readBlogReq := &blogpb.ReadBlogRequest{
		Id: blogID,
	}

	readBlogRes, err := c.ReadBlog(context.Background(), readBlogReq)
	if err != nil {
		fmt.Printf("Error while reading blog: %v", err)
	}

	fmt.Printf("Blog was read successfully: %v", readBlogRes)
	return nil
}
