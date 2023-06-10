package controllers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"logbook_ta/database"
	"logbook_ta/models/entity"
	"logbook_ta/models/request"
	"time"
)

func CreateKarya(c *fiber.Ctx) error {
	karyaRequest := request.KaryaRequest{}

	if errParse := c.BodyParser(&karyaRequest); errParse != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "fail to parsing data",
			"error":   errParse.Error(),
		})
	}

	tanggalSubmit, err := time.Parse("2006-01-02", karyaRequest.TanggalSubmit)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid tanggalSubmit value"})
	}

	tanggalPublish, err := time.Parse("2006-01-02", karyaRequest.TanggalPublish)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid tanggalPublish value"})
	}

	karya := entity.Karya{}
	karya.Link = karyaRequest.Link
	karya.TanggalSubmit = tanggalSubmit
	karya.TanggalPublish = tanggalPublish

	if errDb := database.DB.Preload("User").Create(&karyaRequest).Error; errDb != nil {
		log.Println(errDb)
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
