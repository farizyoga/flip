package main

import (
	"context"
	"flip/internal/app"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	app, _ := app.New(ctx, 100, 4)
	defer cancel()

	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}
