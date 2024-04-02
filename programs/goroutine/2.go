//In Go, the "errgroup" package from the "golang.org/x/sync/errgroup" module provides a convenient way to manage and propagate errors across a group of goroutines.
//It is useful when you need to run multiple goroutines concurrently and handle any errors that occur during their execution.

package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func errgroupFunc() {
	// Create a new context and an error group
	ctx := context.Background()
	group, _ := errgroup.WithContext(ctx)

	tasks := []func() error{
		func() error {
			// Simulate a task that takes some time
			fmt.Println("Task 1 completed")
			return nil
		},
		func() error {
			defer func() {
				if r := recover(); r != nil {
					if err, ok := r.(error); ok {
						fmt.Println("ERROR", err)
					} else {
						fmt.Println("error", r)
					}
				}
			}()
			fmt.Println("Task 2 started")
			panic("Task 2 encountered an error")
		},
		func() error {
			// Simulate a task that takes some time
			time.Sleep(1 * time.Second)
			fmt.Println("Task 3 completed")
			return nil
		},
	}

	// Run the tasks concurrently
	for _, task := range tasks {
		task := task // Create a new variable to capture the current task for the goroutine
		group.Go(func() error {
			// Execute the task and return any error
			return task()
		})
	}

	// Wait for all tasks to complete
	if err := group.Wait(); err != nil {
		fmt.Printf("One or more tasks encountered an error: %v\n", err)
	} else {
		fmt.Println("All tasks completed successfully")
	}
}
