package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func dummyFuncWithoutError(ctx context.Context, val string) error {
	fmt.Println("Executing ", val)
	time.Sleep(2 * time.Second)
	fmt.Println("no error in ", val)
	return nil
}

func dummyFuncWithError(ctx context.Context, val string) error {
	fmt.Println("Executing ", val)
	var errFailure = errors.New("some error")
	return errFailure
}

func main() {
	fmt.Println("mainnn")

	group := errgroup.Group{}
	fmt.Println(group)
	group.SetLimit(2)
	fmt.Println(group)

	group.Go(func() error {
		time.Sleep(5 * time.Second)
		fmt.Println("worker 1")
		return nil
	})

	group.Go(func() error {
		time.Sleep(1 * time.Second)
		fmt.Println("worker 2")
		return nil
	})

	group.Go(func() error {
		time.Sleep(2 * time.Second)
		fmt.Println("worker 3")
		return nil
	})
	if err := group.Wait(); err != nil {
		fmt.Printf("errgroup tasks ended up with an error: %v\n", err)
	} else {
		fmt.Println("all works done successfully")
	}

	group2, ctx := errgroup.WithContext(context.Background())

	group2.Go(func() error {
		return dummyFuncWithoutError(ctx, "worker 1")
	})

	group2.Go(func() error {
		return dummyFuncWithError(ctx, "worker 2")
	})

	group2.Go(func() error {
		return dummyFuncWithoutError(ctx, "worker 3")
	})

	if err := group2.Wait(); err != nil {
		fmt.Printf("Error occurred: %v\n", err)
	}

}
