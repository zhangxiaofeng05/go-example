package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	// Create a new error group and context
	g, _ := errgroup.WithContext(context.Background())

	// Goroutine 1 - Long running task
	g.Go(func() error {
		fmt.Println("Goroutine 1 started")
		<-time.After(3 * time.Second)
		fmt.Println("Goroutine 1 finished")
		return nil
	})

	// Goroutine 2 - Task that fails
	g.Go(func() error {
		fmt.Println("Goroutine 2 started")
		<-time.After(2 * time.Second)
		fmt.Println("Goroutine 2 finished")
		//return errors.New("Goroutine 2 failed")
		return nil
	})

	// Goroutine 3 - Task that succeeds
	g.Go(func() error {
		fmt.Println("Goroutine 3 started")
		<-time.After(1 * time.Second)
		fmt.Println("Goroutine 3 finished")
		return nil
	})

	// Wait for all the goroutines to finish and return the first error, if any
	if err := g.Wait(); err != nil {
		fmt.Printf("At least one goroutine failed: %v\n", err)
		return
	}

	fmt.Println("All goroutines finished successfully")
}
