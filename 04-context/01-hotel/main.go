package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*3)

	defer cancel()

	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("booking canceled. Timeout reached.")
		return
	case <-time.After(1 * time.Second):
		fmt.Println("Hotel booked.")
	}
}