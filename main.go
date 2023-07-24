package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	ctx, cancel := context.WithTimeout(context.TODO(), 3*time.Second+500*time.Millisecond)
	defer cancel()

	for {
		select {
		case t := <-ticker.C:
			fmt.Println("tick:", t.Format(time.RFC3339Nano))
		case <-ctx.Done():
			return
		}
	}
}
