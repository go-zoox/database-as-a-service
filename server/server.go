package server

import (
	"context"
	"fmt"
	"time"

	daas "github.com/go-zoox/database-as-a-service"
	"github.com/go-zoox/database-as-a-service/data"
	"github.com/go-zoox/database-as-a-service/engine/mysql"
	"github.com/go-zoox/database-as-a-service/engine/postgres"
	"github.com/go-zoox/database-as-a-service/engine/sqlite3"
	"github.com/go-zoox/zoox"
	"github.com/go-zoox/zoox/defaults"
)

type Server interface {
	Run() (err error)
}

type server struct {
	opt *Option
}

func New(opts ...func(opt *Option)) Server {
	opt := &Option{
		Port: 8080,
		Path: "/",
	}
	for _, o := range opts {
		o(opt)
	}

	return &server{
		opt: opt,
	}
}

func (s *server) Run() (err error) {
	app := defaults.Defaults()

	app.Config.Port = s.opt.Port

	app.Post(
		s.opt.Path,
		func(ctx *zoox.Context) {
			if s.opt.Username == "" || s.opt.Password == "" {
				ctx.Next()
				return
			}

			user, pass, ok := ctx.Request.BasicAuth()
			if !ok {
				ctx.Set("WWW-Authenticate", `Basic realm="go-zoox"`)
				ctx.Status(401)
				return
			}

			if !(user == s.opt.Username && pass == s.opt.Password) {
				ctx.Status(401)
				return
			}

			ctx.Next()
		},
		func(ctx *zoox.Context) {
			request := &data.Request{}
			if err = ctx.BindBody(request); err != nil {
				ctx.Fail(err, 400, "Invalid request payload")
				return
			}

			c, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
			defer cancel()

			switch request.Engine {
			case mysql.Name:
				result, err := mysql.Query(c, request.DSN, request.Statement)
				if err != nil {
					ctx.Fail(err, 500, fmt.Sprintf("Failed to execute query: %s", err))
					return
				}

				ctx.Success(result)
			case postgres.Name:
				result, err := postgres.Query(c, request.DSN, request.Statement)
				if err != nil {
					ctx.Fail(err, 500, fmt.Sprintf("Failed to execute query: %s", err))
					return
				}

				ctx.Success(result)
			case sqlite3.Name:
				result, err := sqlite3.Query(c, request.DSN, request.Statement)
				if err != nil {
					ctx.Fail(err, 500, fmt.Sprintf("Failed to execute query: %s", err))
					return
				}

				ctx.Success(result)
			default:
				ctx.Fail(nil, 400, fmt.Sprintf("unsupported engine: %s", request.Engine))
			}
		})

	app.Get("/", func(ctx *zoox.Context) {
		ctx.Success(zoox.H{
			"title":   "Welcome to Database as a Service",
			"version": daas.Version,
		})
	})

	return app.Run()
}
