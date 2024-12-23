package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/snykk/go-fx/controllers"
	"github.com/snykk/go-fx/repositories"
	"github.com/snykk/go-fx/services"
	"go.uber.org/fx"
)

func NewFiberApp() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	return app
}

func RegisterRoutes(app *fiber.App, userController *controllers.UserController) {
	app.Get("/users/:id", userController.GetUserHandler)
}

func AppLifecycle(lc fx.Lifecycle, app *fiber.App) {
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go func() {
				if err := app.Listen(":8080"); err != nil {
					fmt.Println("Error starting app:", err)
				}
			}()
			return nil
		},
		OnStop: func(_ context.Context) error {
			return app.Shutdown()
		},
	})
}

func InitializeApp(lc fx.Lifecycle, app *fiber.App, userController *controllers.UserController) {
	RegisterRoutes(app, userController)
	AppLifecycle(lc, app)
}

func main() {
	fx.New(
		fx.Provide(
			repositories.NewUserRepository,
			services.NewUserService,
			controllers.NewUserController,
			NewFiberApp,
		),
		fx.Invoke(InitializeApp),
	).Run()
}
