package main

import (
	"context"
	"fmt"
	"io"
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

	// update Blog
	UpdateBlog(c)

	// delete Blog
	DeleteBlog(c)

	// list Blog
	ListBlog(c)
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
	blogID := "61fdca68ee66f947c22d23f0"
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

// update Blog
func UpdateBlog(c blogpb.BlogServiceClient) error {
	// update blog
	fmt.Println("Updating the blog")
	blog := &blogpb.Blog{
		Id:       "61fdca68ee66f947c22d23f0",
		AuthorId: "Bangadam",
		Content:  "Content of the blog updated",
		Title:    "Title of the blog updated",
	}
	updateBlogRes, err := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Printf("Blog was updated successfully: %v", updateBlogRes)
	return nil
}

func DeleteBlog(c blogpb.BlogServiceClient) error {
	// delete blog
	fmt.Println("Deleting the blog")
	blogID := "61fdca68ee66f947c22d23f0"
	deleteBlogRes, err := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{Id: blogID})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Printf("Blog was deleted successfully: %v", deleteBlogRes)
	return nil
}

func ListBlog(c blogpb.BlogServiceClient) error {
	// list blog
	fmt.Println("Listing the blog")
	stream, err := c.ListBlog(context.Background(), &blogpb.ListBlogRequest{})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Unexpected error : %v", err)
		}
		fmt.Println(res.GetBlog())
	}

	return nil
}
