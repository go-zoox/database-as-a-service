package commands

import (
	"github.com/go-zoox/cli"
	"github.com/go-zoox/database-as-a-service/server"
)

// RegistryServer registers the server command
func RegistryServer(app *cli.MultipleProgram) {
	app.Register("server", &cli.Command{
		Name:  "server",
		Usage: " database as a service server",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "port",
				Usage:   "server port",
				Aliases: []string{"p"},
				EnvVars: []string{"PORT"},
				Value:   8080,
			},
			&cli.StringFlag{
				Name:  "path",
				Usage: "api path",
			},
			&cli.StringFlag{
				Name:    "username",
				Usage:   "Username for Basic Auth",
				EnvVars: []string{"USERNAME"},
			},
			&cli.StringFlag{
				Name:    "password",
				Usage:   "Password for Basic Auth",
				EnvVars: []string{"PASSWORD"},
			},
		},
		Action: func(ctx *cli.Context) (err error) {
			s := server.New(func(opt *server.Option) {
				if ctx.Int("port") != 0 {
					opt.Port = ctx.Int("port")
				}

				if ctx.String("path") != "" {
					opt.Path = ctx.String("path")
				}

				if ctx.String("username") != "" {
					opt.Username = ctx.String("username")
				}

				if ctx.String("password") != "" {
					opt.Password = ctx.String("password")
				}
			})

			return s.Run()
		},
	})
}
