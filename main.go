package main

import (
import (
	"github.com/AxrorbekDev93/0222/db"
	"github.com/AxrorbekDev93/0222/handlers"
	"github.com/AxrorbekDev93/0222/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"os"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	db.Connect()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("🚀 Сервер Go работает на Render!")
	})

	app.Post("/register", handlers.RegisterUser)
	app.Post("/login", handlers.Login)

	app.Get("/users", middleware.Protect(), handlers.GetUsers)
	app.Patch("/users/:id", middleware.Protect(), handlers.UpdateUserBySuperAdmin)
	app.Patch("/users/:id/status", middleware.Protect(), handlers.UpdateUserStatus)
	app.Get("/users/me", middleware.Protect(), handlers.GetMyProfile)

	app.Get("/locomotives", middleware.Protect(), handlers.GetLocomotives)
	app.Post("/locomotives", middleware.Protect(), handlers.AddLocomotive)
	app.Delete("/locomotives/:id", middleware.Protect(), handlers.DeleteLocomotive)

	app.Get("/diesel-oil", middleware.Protect(), handlers.GetDieselOil)
	app.Post("/diesel-oil", middleware.Protect(), handlers.AddDieselOil)
	app.Delete("/diesel-oil/:id", middleware.Protect(), handlers.DeleteDieselOil)

	app.Get("/depos", handlers.GetDepos)
	app.Post("/depos", middleware.Protect(), handlers.CreateDepo)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	app.Listen(":" + port)
}
