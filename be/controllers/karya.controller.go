package controllers

import (
	"github.com/go-playground/validator/v10"
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

	// VALIDATION REQUEST DATA

	validate := validator.New()
	if errValidate := validate.Struct(&karyaRequest); errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "some data is not valid",
			"error":   errValidate.Error(),
		})
	}
	karya := entity.Karya{}

	tanggalSubmit, err := time.Parse("2006-01-02", karyaRequest.TanggalSubmit)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid tanggalSubmit value"})
	}

	tanggalPublish, err := time.Parse("2006-01-02", karyaRequest.TanggalPublish)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid tanggalPublish value"})
	}

	karya.Link = karyaRequest.Link
	karya.TanggalSubmit = tanggalSubmit
	karya.TanggalPublish = tanggalPublish
	karya.CreatedAt = time.Now()

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

func UpdateKaryaById(c *fiber.Ctx) error {
	karyaUpdateReq := request.KaryaUpdateRequest{}

	// PARSE REQUEST BODY
	if errParse := c.BodyParser(&karyaUpdateReq); errParse != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "fail to parsing data",
			"error":   errParse.Error(),
		})
	}

	// VALIDATION REQUEST DATA

	validate := validator.New()
	if errValidate := validate.Struct(&karyaUpdateReq); errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "some data is not vali",
			"error":   errValidate.Error(),
		})
	}

	karyaId := c.Params("id")
	karya := entity.Karya{}

	if err := database.DB.First(&karya, "id = ?", karyaId).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "todo not found",
		})
	}

	tanggalSubmit, err := time.Parse("2006-01-02", karyaUpdateReq.TanggalSubmit)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid tanggalSubmit value"})
	}

	tanggalPublish, err := time.Parse("2006-01-02", karyaUpdateReq.TanggalPublish)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid tanggalPublish value"})
	}

	karya.Link = karyaUpdateReq.Link
	karya.TanggalSubmit = tanggalSubmit
	karya.TanggalPublish = tanggalPublish

	if errSave := database.DB.Save(&karya).Error; errSave != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "logbook updated",
		"data":    karya,
	})
}

func DeleteKaryaById(c *fiber.Ctx) error {

	karyaId := c.Params("id")
	karya := entity.Karya{}

	if err := database.DB.First(&karya, "id = ?", karyaId).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "todo not found",
		})
	}

	if errDelete := database.DB.Delete(&karya).Error; errDelete != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "karya deleted",
	})
}
