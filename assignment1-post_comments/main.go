package main

import (
	"context"
	"posts/cmd"
	"log"
	"net/http"
	"os"
)

func main() {
	//handlers.HandleRequests()
	if err := cmd.Run(os.Args, os.Stdout); err != nil {
		switch err {
		case context.Canceled:
			// not considered error
		case http.ErrServerClosed:
			// not considered error
		default:
			log.Fatalf("could not run application: %v", err)
		}
	}
}
