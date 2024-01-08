package main

import (
	"context"
	"fmt"

	"github.com/Ethea2/popio-server/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("Failed to start the application: ", err)
	}
}

// hello comment!

// still deciding what database I hsould use.
// I want an SQL application but can't figure out how to host it if ever.
