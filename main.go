package main

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/api/fitness/v1"
)

func main() {
	fmt.Println("go-fits")

	ctx := context.Background()
	fitnessService, err := fitness.NewService(ctx)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(fitnessService)
}
