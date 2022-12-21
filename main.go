package main

import (
	"app/core/application"
	"fmt"
	"os"
)

func main() {
	app, err := application.Create()
	if err != nil {
		fmt.Printf("[error] %v\n", err)
		os.Exit(1)
	}

	/* wait until app shutdown to close db connection */
	defer app.Singletons.Database.Disconnect()

	/* spin up application web server */
	app.Server.Start()
}
