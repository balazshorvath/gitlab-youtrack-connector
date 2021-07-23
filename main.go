package main

import (
	"time"

	srv "github.com/balazshorvath/go-srv"
)

func main() {
	srv.CreateAndRunServer(New, 5*time.Second)
}
