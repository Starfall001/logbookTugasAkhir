package controllers

import (
	"github.com/gofiber/fiber/v2"
	"logbook_ta/database"
	"logbook_ta/models/entity"
	"logbook_ta/models/request"
)

func CreateKarya(c *fiber.Ctx) error {
	karyaRequest := request.KaryaRequest{}

	if errParse := c.BodyParser(&karyaRequest); errParse != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "fail to parsing data",
			"error":   errParse.Error(),
		})
	}

	karya := entity.Karya{}
	karya.TanggalSubmit = karyaRequest.TanggalSubmit
	karya.TanggalPublish = karyaRequest.TanggalPublish
	karya.Link = karyaRequest.Link

	if errDb := database.DB.Create(&karyaRequest).Error; errDb != nil {
		//log.Println("category.controller.go => CreateCategory :: ", errDb)
		return c.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "karya submitted",
		"data":    karyaRequest,
	})
}

func GetKaryaById(c *fiber.Ctx) error {
	karyaId := c.Params("id")
	karya := entity.Karya{}

	if err := database.DB.Preload("User").First(&karya, "id = ?", karyaId).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "karya not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "data transmitted",
		"data":    karya,
	})
}

func GetAllKarya(c *fiber.Ctx) error {
	karyas := []entity.Karya{}

	if err := database.DB.Preload("User").Find(&karyas).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}
	return c.JSON(fiber.Map{
		"message": "data transmited",
		"data":    karyas,
	})
}
