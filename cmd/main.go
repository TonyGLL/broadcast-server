package main

import (
	"context"
	"os"

	"github.com/TonyGLL/broadcast-server/internal"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "broadcast-server",
		Usage: "WebSocket broadcast server",
		Commands: []*cli.Command{
			{
				Name:  "start",
				Usage: "Start the WebSocket server",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "port",
						Aliases: []string{"p"},
						Value:   0,
						Usage:   "Server port",
					},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					port := cmd.Int("port")
					internal.StartServer(port)
					return nil
				},
			},
			{
				Name:  "connect",
				Usage: "Connect to the WebSocket server as a client",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "port",
						Aliases: []string{"p"},
						Value:   0,
						Usage:   "Server port",
					},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					port := cmd.Int("port")
					internal.Connect(port)
					return nil
				},
			},
		},
	}

	cmd.Run(context.Background(), os.Args)
}
