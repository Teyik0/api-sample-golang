package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/Teyik0/api-sample-golang/api-prisma/client"
	"github.com/Teyik0/api-sample-golang/api-prisma/prisma/db"
	"github.com/gofiber/fiber/v2"
	"github.com/labstack/gommon/log"
)

func AddUser(c *fiber.Ctx) error {
	log.Info(time.Now().UTC().String(), "| POST /api/v1/user from : ", c.IP(), " | ", c.Path())

	// Get user from request body
	var userResp db.UserModel

	if err := c.BodyParser(&userResp); err != nil {
		fmt.Println(err)
		return c.Status(400).JSON(err.Error())
	}

	// Check if user already exists
	user, err := client.Prisma.User.FindUnique(db.User.Username.Equals(userResp.Username)).Exec(context.Background())
	if user != nil {
		return c.Status(400).JSON("User already exists")
	}

	// Create user
	ctx := context.Background()

	createdPost, err := client.Prisma.User.CreateOne(
		db.User.Username.Set(userResp.Username),
		db.User.Password.Set(userResp.Password),
	).Exec(ctx)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(createdPost)
}

func GetUsers(c *fiber.Ctx) error {
	log.Info(time.Now().UTC().String(), "| GET /api/v1/user from : ", c.IP(), " | ", c.Path())

	ctx := context.Background()

	users, err := client.Prisma.User.FindMany().Exec(ctx)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	log.Info(time.Now().UTC().String(), "| GET /api/v1/user/", c.Params("id"), " from : ", c.IP(), " | ", c.Path())

	ctx := context.Background()

	user, err := client.Prisma.User.FindUnique(db.User.ID.Equals(c.Params("id"))).Exec(ctx)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	log.Info(time.Now().UTC().String(), "| GET /api/v1/user/", c.Params("id"), " from : ", c.IP(), " | ", c.Path())

	//Get user from request body
	var userResp db.UserModel
	if err := c.BodyParser(&userResp); err != nil {
		fmt.Println(err)
		return c.Status(400).JSON(err.Error())
	}

	ctx := context.Background()

	user, err := client.Prisma.User.FindUnique(db.User.ID.Equals(c.Params("id"))).Update(
		db.User.Username.Set(userResp.Username),
		db.User.Password.Set(userResp.Password),
	).Exec(ctx)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	log.Info(time.Now().UTC().String(), "| GET /api/v1/user/", c.Params("id"), " from : ", c.IP(), " | ", c.Path())

	ctx := context.Background()

	user, err := client.Prisma.User.FindUnique(db.User.ID.Equals(c.Params("id"))).Delete().Exec(ctx)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(user)
}
