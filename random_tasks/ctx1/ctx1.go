package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	time.Sleep(2999 * time.Millisecond)
	doRequest(ctx)
}

func doRequest(ctx context.Context) {
	newCtx, _ := context.WithTimeout(ctx, time.Second*10)
	timer := time.NewTimer(time.Second * 1)
	select {
	case <-newCtx.Done():
		fmt.Println("Timeout")
	case <-timer.C:
		fmt.Println("Request done")
	}
}
