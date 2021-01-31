package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg3 sync.WaitGroup

func work3(ctx context.Context) error {
	defer wg3.Done()

	for i := 0; i < 1000; i++ {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("Doing some work ", i)

		// we received the signal of cancellation in this channel
		case <-ctx.Done():
			fmt.Println("Cancel the context ", i)
			return ctx.Err()
		}
	}
	return nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	fmt.Println("Hey, I'm going to do some work")

	wg3.Add(1)
	var err error
	go func() {
		err = work3(ctx)
	}()
	if err != nil {
		fmt.Println("work3 exception exit", err)
	}
	wg3.Wait()

	fmt.Println("Finished. I'm going home")
}
