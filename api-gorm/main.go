package main

import (
	"github.com/Teyik0/api-sample-golang/controllers"
	"github.com/Teyik0/api-sample-golang/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// func setLog() {
// 	fmt.Println("Setting up log file as : ", time.Now().Format("2006-01-02T15-04-05")+".log")
// 	logFileName := fmt.Sprintf("logs/%s.log", time.Now().Format("2006-01-02T15-04-05"))
// 	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
// 	if err != nil {
// 		log.Panic("Error opening file : ", err)
// 	}
// 	log.SetOutput(logFile)
// }

func main() {
	// setLog()
	log.Info("Starting the server...")

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	if err := db.Connect(); err != nil {
		log.Fatal("Error connecting to database")
	}

	v1Api := app.Group("/api/v1", apiMiddleware) // /api

	v1Api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	v1Api.Post("/user", controllers.AddUser)
	v1Api.Get("/user", controllers.GetUsers)
	v1Api.Get("/user/:id", controllers.GetUser)
	v1Api.Put("/user/:id", controllers.UpdateUser)
	v1Api.Delete("/user/:id", controllers.DeleteUser)

	log.Fatal(app.Listen(":3000"))
}

func apiMiddleware(c *fiber.Ctx) error {
	return c.Next()
}
