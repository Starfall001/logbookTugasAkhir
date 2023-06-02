package controllers

import (
	"log"
	"logbook_ta/database"
	"logbook_ta/models/entity"
	"logbook_ta/models/request"
	"logbook_ta/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/utils"
)

func CreateUser(c *fiber.Ctx) error {
	user := request.UserRegisterRequest{}

	// PARSE REQUEST BODY
	if errParse := c.BodyParser(&user); errParse != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "fail to parsing data",
			"error":   errParse.Error(),
		})
	}

	// VALIDATION REQUEST DATA

	validate := validator.New()
	if errValidate := validate.Struct(&user); errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "some data is not valid",
			"error":   errValidate.Error(),
		})
	}

	newUser := entity.User{
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

	hashedPassword, err := utils.HashingPassword(user.Password)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	newUser.Password = hashedPassword

	if errDb := database.DB.Create(&newUser).Error; errDb != nil {
		log.Println("todo.controller.go => CreateTodo :: ", errDb)
		return c.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "user registered successfully",
		"data":    newUser,
	})
}

func GetAllUser(c *fiber.Ctx) error {
	allUser := []entity.User{}

	if err := database.DB.Find(&allUser).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "data transmitted",
		"data":    allUser,
	})
}
