package main

import (
	"github.com/lucastomic/infrastructure-/internal/infrastructure"
	"github.com/lucastomic/infrastructure-/internal/logging"
	"github.com/lucastomic/infrastructure-/internal/middleware"
	"github.com/lucastomic/infrastructure-/internal/server"
)

func main() {
	logger := logging.NewLogrusLogger()
	webprocessor := infrastructure.New(logger)
	server := server.New(":3002", webprocessor, logger, []middleware.Middleware{})
	server.Run()
}
