package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/Teyik0/api-sample-golang/prisma/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func setLog() {
	fmt.Println("Setting up log file as : ", time.Now().Format("2006-01-02T15-04-05")+".log")
	logFileName := fmt.Sprintf("./logs/%s.log", time.Now().Format("2006-01-02"))
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file : ", err)
	}
	log.SetOutput(logFile)
}

func main() {
	fmt.Println("Setting up the log file as : ", time.Now().Format("2006-01-02")+".log")
	setLog()
	log.Info("Starting the server...")
	app := fiber.New()

	api := app.Group("/api", apiMiddleware) // /api
	v1 := api.Group("/v1", apiMiddleware)   // /api/v1

	/* post, err := client.User.FindUnique(
		db.User.ID.Equals(createUser.ID),
	).Exec(ctx)
	if err != nil {
		fmt.Println("Unable to find a user", err)
		return
	}
	result, _ = json.MarshalIndent(post, "", "  ")
	fmt.Printf("user: %s\n", result) */

	v1.Get("/", func(c *fiber.Ctx) error {
		fmt.Println(time.Now().UTC().String(), "| GET /api/v1 from : ", c.IP())
		return c.SendString(string(result))
	})

	app.Listen(":3000")
}

func apiMiddleware(c *fiber.Ctx) error {
	log.Info("Connecting to database...")
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		fmt.Println("Unable to connect to database: ", err)
		return err
	}
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()
	c.Locals("Prisma client", client)
	return c.Next()
}

func createUser() {
	ctx := context.Background()
	createUser, err := client.User.CreateOne(
		db.User.Username.Set("Teyik02229"),
		db.User.Password.Set("test1234"),
		db.User.AuthToken.Set("abcd"),
	).Exec(ctx)
	if err != nil {
		fmt.Println("Unable to create a user", err)
		return
	}

	result, _ := json.MarshalIndent(createUser, "", "  ")
	fmt.Printf("Created user: %s\n", result)
}
