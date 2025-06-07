package main

import (
	"flag"

	"github.com/TonyGLL/broadcast-server/internal"
)

func main() {
	// Define CLI flags
	start := flag.Bool("start", false, "Start the WebSocket server")
	connect := flag.Bool("connect", false, "Connect to the WebSocket server as a client")
	message := flag.String("message", "", "Message to send when connected as a client")

	// Parse CLI flags
	flag.Parse()

	if *connect {
		internal.Connect(*message)
		return
	}

	if *start {
		internal.StartServer()
	}
}
