package main

import (
	"context"
	"modular-monolith-libraryApp/config"
	"modular-monolith-libraryApp/modules/author"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		// infra: config, pool, bus
		//infraModule,

		// módulos de domínio
		author.Module,
		// book.Module,   ← adicionar conforme implementar
		// user.Module,

		// fiber app — provida aqui para ser injetada em registerRoutes
		fx.Provide(func() *fiber.App {
			return fiber.New()
		}),

		// registra rotas e listeners
		//fx.Invoke(registerRoutes),
		//fx.Invoke(registerEventListeners),

		// sobe o servidor
		fx.Invoke(func(lc fx.Lifecycle, app *fiber.App, cfg *config.Config) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					go func() {
						if err := app.Listen(":" + cfg.ServerPort); err != nil {
							panic(err)
						}
					}()
					return nil
				},
				OnStop: func(ctx context.Context) error {
					return app.Shutdown()
				},
			})
		}),

		// fecha o pool ao encerrar
		fx.Invoke(func(lc fx.Lifecycle, pool *pgxpool.Pool) {
			lc.Append(fx.Hook{
				OnStop: func(ctx context.Context) error {
					pool.Close()
					return nil
				},
			})
		}),
	).Run()
}
