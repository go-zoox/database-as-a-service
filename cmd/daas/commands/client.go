package commands

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-zoox/cli"
	"github.com/go-zoox/database-as-a-service/client"
	"github.com/go-zoox/database-as-a-service/data"
)

// RegistryClient registers the client command
func RegistryClient(app *cli.MultipleProgram) {
	app.Register("client", &cli.Command{
		Name:  "client",
		Usage: " database as a service client",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "server",
				Usage:    "server url",
				Aliases:  []string{"s"},
				EnvVars:  []string{"SERVER"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "engine",
				Usage:    "database engine, e.g. mysql, postgres, sqlite3",
				EnvVars:  []string{"ENGINE"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "dsn",
				Usage:    "database dsn",
				EnvVars:  []string{"DSN"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "statement",
				Usage:    "database statement",
				EnvVars:  []string{"STATEMENT"},
				Required: true,
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
			c := client.New(func(opt *client.Option) {
				if ctx.String("server") != "" {
					opt.Server = ctx.String("server")
				}

				if ctx.String("username") != "" {
					opt.Username = ctx.String("username")
				}

				if ctx.String("password") != "" {
					opt.Password = ctx.String("password")
				}
			})

			result, err := c.Request(&data.Request{
				Engine:    ctx.String("engine"),
				DSN:       ctx.String("dsn"),
				Statement: ctx.String("statement"),
			})
			if err != nil {
				return err
			}

			text, err := json.MarshalIndent(result, "", "  ")
			if err != nil {
				return err
			}

			fmt.Fprintf(os.Stdout, "%s\n", text)

			return nil
		},
	})
}
