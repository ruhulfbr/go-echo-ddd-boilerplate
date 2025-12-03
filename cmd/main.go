package main

import (
	"log/slog"
	"os"

	"github.com/ruhulfbr/go-echo-ddd-boilerplate/cmd/server"
)

func main() {
	if err := server.Run(); err != nil {
		slog.Error("Application run error", "err", err)
		os.Exit(1)
	}
}
