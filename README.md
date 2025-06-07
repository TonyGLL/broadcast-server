# Setting Up and Running WebSocket Broadcast Server with Go

https://roadmap.sh/projects/broadcast-server

This project is a **WebSocket Broadcast Server** built in Go. It allows multiple clients to connect via WebSocket and broadcast messages to each other in real time. The server supports both server mode (to start and accept connections) and client mode (to connect and send messages). The client now accepts manual input from the console.

---

## Overview

- **Server Mode**: Starts a WebSocket server on a specified port (default: 3000) and handles client connections on the `/ws` endpoint. When any client sends a message, it is broadcasted to all connected clients.
- **Client Mode**: Connects to the WebSocket server and allows sending messages manually from the command line.

---

## Prerequisites

Ensure you have the following tools installed on your system:

- **Golang** ([Download](https://go.dev/dl/))
- **Make** (Available by default on most Linux/macOS systems, installable on Windows via [Chocolatey](https://chocolatey.org/) or [Scoop](https://scoop.sh/))

---

## Steps to Set Up the Application

### 1. Clone the Repository

## Installation

```bash
git clone https://github.com/TonyGLL/broadcast-server.git
cd broadcast-server
go build -o broadcast-server cmd/main.go
```

## Usage

### Start the websocket server
```bash
./broadcast-server -start
```

### Connect a client and send messages
```bash
./broadcast-server -connect -message="Hello from client"
```