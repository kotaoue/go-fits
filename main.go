package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/fitness/v1"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var (
	jsonPath  string
	projectID string
)

func init() {
	jsonPath = os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	projectID = os.Getenv("GOOGLE_PROJECT_ID")
}

func main() {
	fmt.Println("go-fits")
	fmt.Printf("jsonPath:%s, projectID: %s\n", projectID, jsonPath)

	explicit(jsonPath, projectID)

	ctx := context.Background()
	fitnessService, err := fitness.NewService(ctx)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(fitnessService)
}

// explicit reads credentials from the specified path.
func explicit(jsonPath, projectID string) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(jsonPath))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	fmt.Println("Buckets:")
	it := client.Buckets(ctx, projectID)
	for {
		battrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(battrs.Name)
	}
}
