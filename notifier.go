package main

import (
	"flag"
	"fmt"
	"github.com/carnellj/gws-ttk-notifier/service"
)

var (
	port = flag.String("port", "3002", "Port service starts on. Default port number is 3002")
)

func main() {
	flag.Parse()

	fmt.Printf("Starting server on port %s", *port)

	server := service.NewServer()
	server.Run(":" + *port)
}
