package main

import (
	"github.com/Teyik0/api-sample-golang/api-prisma/client"
	"github.com/Teyik0/api-sample-golang/api-prisma/handlers"
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

	if err := client.Connect(); err != nil {
		panic(err)
	}

	v1Api := app.Group("/api/v1", apiMiddleware)

	v1Api.Post("/user", handlers.AddUser)
	v1Api.Get("/user", handlers.GetUsers)
	v1Api.Get("/user/:id", handlers.GetUser)
	v1Api.Delete("/user/:id", handlers.DeleteUser)
	v1Api.Put("/user/:id", handlers.UpdateUser)

	log.Fatal(app.Listen(":3000"))
}

func apiMiddleware(c *fiber.Ctx) error {
	return c.Next()
}
