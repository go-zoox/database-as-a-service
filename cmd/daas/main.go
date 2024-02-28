package main

import (
	"github.com/go-zoox/cli"
	daas "github.com/go-zoox/database-as-a-service"
	"github.com/go-zoox/database-as-a-service/cmd/daas/commands"
)

func main() {
	app := cli.NewMultipleProgram(&cli.MultipleProgramConfig{
		Name:    "daas",
		Usage:   "database as a service",
		Version: daas.Version,
	})

	// server
	commands.RegistryServer(app)
	// client
	commands.RegistryClient(app)

	app.Run()
}
