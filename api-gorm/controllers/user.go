// @/controllers/user.go
package controllers

import (
	"fmt"
	"time"

	"github.com/Teyik0/api-sample-golang/db"
	"github.com/Teyik0/api-sample-golang/entities"
	"github.com/gofiber/fiber/v2"
	"github.com/labstack/gommon/log"
)

func AddUser(c *fiber.Ctx) error {
	fmt.Println(time.Now().UTC().String(), "| POST /api/v1/user from : ", c.IP())
	log.Info("| POST /api/v1/user from : ", c.IP(), " | ", c.Path())

	user := new(entities.User)

	if err := c.BodyParser(user); err != nil {
		fmt.Println(err)
		return c.Status(400).JSON(err.Error())
	}

	tx := db.Database.Create(&user)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return c.Status(400).JSON(tx.Error)
	}

	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "User has been created", "data": user})
}

func GetUsers(c *fiber.Ctx) error {
	fmt.Println(time.Now().UTC().String(), "| GET /api/v1/user from : ", c.IP())
	log.Info("| GET /api/v1/user from : ", c.IP(), " | ", c.Path())
	var users []entities.User

	db.Database.Find(&users)
	return c.Status(200).JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	fmt.Println(time.Now().UTC().String(), "| GET /api/v1/user/:id from : ", c.IP())
	log.Info("| GET /api/v1/user/:id from : ", c.IP(), " | ", c.Path())

	id := c.Params("id")
	var user entities.User

	result := db.Database.Find(&user, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	fmt.Println(time.Now().UTC().String(), "| PUT /api/v1/user/:id from : ", c.IP())
	log.Info("| PUT /api/v1/user/:id from : ", c.IP(), " | ", c.Path())

	user := new(entities.User)
	id := c.Params("id")

	if err := c.BodyParser(user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Database.Where("id = ?", id).Updates(&user)

	return c.Status(200).JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	fmt.Println(time.Now().UTC().String(), "| DELETE /api/v1/user/:id from : ", c.IP())
	log.Info("| DELETE /api/v1/user/:id from : ", c.IP(), " | ", c.Path())

	user := new(entities.User)
	id := c.Params("id")

	result := db.Database.Delete(&user, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(204)
}
