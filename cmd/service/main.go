package main

import (
	"log/slog"
	"os"

	"github.com/ruhulfbr/go-echo-ddd-boilerplate/cmd/service/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		slog.Error("Service run error", "err", err)
		os.Exit(1)
	}
}
