package main

import (
	"context"
	"fmt"
	"keploy-test/applications"
	"os"
	"os/signal"
)

func main() {
	app := applications.New()
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Printf("failed to start app : %v", err)
	}
}
