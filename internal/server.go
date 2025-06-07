package internal

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

// Upgrade incoming HTTP requests to WebSocket connections
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, // Allow all origins
}

// Track connected WebSocket clients
var clients = make(map[*websocket.Conn]bool)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade GET request to a WebSocket connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("upgrade error:", err)
		return
	}
	fmt.Println("User connected")
	defer ws.Close()

	clients[ws] = true

	for {
		// Read incoming message
		_, msg, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("read error:", err)
			delete(clients, ws)
			break
		}

		// Broadcast message to all connected clients
		for client := range clients {
			if err := client.WriteMessage(websocket.TextMessage, msg); err != nil {
				fmt.Println("broadcast error:", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func StartServer() {
	http.HandleFunc("/ws", handleConnections)

	fmt.Println("WebSocket server started on :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

func Connect(initialMessage string) {
	// Handle OS interrupt signals
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: "localhost:3000", Path: "/ws"}
	log.Printf("Connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	// Goroutine to receive messages from the server
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("Received: %s", message)
		}
	}()

	// Optionally send the initial message
	if initialMessage != "" {
		err := c.WriteMessage(websocket.TextMessage, []byte(initialMessage))
		if err != nil {
			log.Println("initial write:", err)
			return
		}
	}

	// Read from stdin and send to server
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			text := scanner.Text()
			if err := c.WriteMessage(websocket.TextMessage, []byte(text)); err != nil {
				log.Println("write error:", err)
				return
			}
		}
		if err := scanner.Err(); err != nil {
			log.Println("input error:", err)
		}
	}()

	// Wait for either done or interrupt signal
	for {
		select {
		case <-done:
			log.Println("connection closed by server")
			return
		case <-interrupt:
			log.Println("Interrupt received. Closing connection...")

			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("close message error:", err)
				return
			}

			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
