package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"logbook_ta/database"
	"logbook_ta/models/entity"
	"logbook_ta/models/request"
	"logbook_ta/utils"
	"time"
)

func LoginController(ctx *fiber.Ctx) error {
	loginRequest := new(request.UserLoginRequest)

	// PARSE REQUEST BODY
	if errParse := ctx.BodyParser(&loginRequest); errParse != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "fail to parsing data",
			"error":   errParse.Error(),
		})
	}
	// VALIDATION REQUEST DATA
	validate := validator.New()
	if errValidate := validate.Struct(loginRequest); errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "some data is not valid",
			"error":   errValidate.Error(),
		})
	}
	// CHECK VALIDATION USER DATA
	var user entity.User
	if err := database.DB.First(&user, "username = ?", loginRequest.Username).Error; err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong username",
		})
	}

	// CHECK VALIDATION PASSWORD DATA
	isValid := utils.CheckPasswordHash(loginRequest.Password, user.Password)
	if !isValid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong password",
		})
	}

	// GENERATE TOKEN
	claims := jwt.MapClaims{}
	claims["username"] = user.Username
	claims["password"] = user.Password
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()

	token, err := utils.GenerateToken(&claims)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credentials",
		})
	}

	return ctx.JSON(fiber.Map{
		"token": token,
	})
}
