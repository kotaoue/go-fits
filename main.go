package main

import (
	"context"
	"fmt"
	"os"
	"strings"

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
	if err := Main(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Main() error {
	ctx := context.Background()

	fmt.Println(strings.Repeat("-", 64))
	if err := getBuckets(ctx); err != nil {
		return err
	}

	fmt.Println(strings.Repeat("-", 64))
	if err := getFit(ctx); err != nil {
		return err
	}

	return nil
}

func getFit(ctx context.Context) error {
	// fitnessService, err := fitness.NewService(ctx, option.WithCredentialsFile(jsonPath))
	fitnessService, err := fitness.NewService(ctx, option.WithCredentialsFile(jsonPath))
	if err != nil {
		return err
	}

	us, err := fitnessService.Users.Sessions.List("me").Do()
	fmt.Println("# Sessions")
	fmt.Printf("- HTTPStatusCode:%d\n", us.HTTPStatusCode)
	fmt.Printf("- SessionLength:%d\n", len(us.Session))
	if err != nil {
		return err
	}
	for _, v := range us.Session {
		fmt.Println(v)
	}

	ds, err := fitnessService.Users.DataSources.List("me").Do()
	fmt.Println("# DataSources")
	fmt.Printf("- HTTPStatusCode:%d\n", ds.HTTPStatusCode)
	fmt.Printf("- SessionLength:%d\n", len(ds.DataSource))
	if err != nil {
		return err
	}
	for _, v := range ds.DataSource {
		fmt.Println(v)
	}

	return nil
}

// https://github.com/googleapis/google-api-go-client/blob/f3223f87150fd81bdaf5cacb8cb3e502f5664ead/examples/fitness.go
func getBuckets(ctx context.Context) error {
	fmt.Println("Get a bucket list using sample code to see if the service account is correct")

	client, err := storage.NewClient(ctx, option.WithCredentialsFile(jsonPath))
	if err != nil {
		return err
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
			return err
		}
		fmt.Println(battrs.Name)
	}

	return nil
}
