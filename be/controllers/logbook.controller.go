package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"logbook_ta/database"
	"logbook_ta/models/entity"
	"logbook_ta/models/request"
)

func CreateLogbook(c *fiber.Ctx) error {
	logbookReq := request.LogbookRequest{}

	// PARSE REQUEST BODY
	if errParse := c.BodyParser(&logbookReq); errParse != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "fail to parsing data",
			"error":   errParse.Error(),
		})
	}

	// VALIDATION REQUEST DATA

	validate := validator.New()
	if errValidate := validate.Struct(&logbookReq); errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "some data is not valid",
			"error":   errValidate.Error(),
		})
	}

	logbook := entity.Logbook{}
	logbook.UserID = logbookReq.UserID
	logbook.Judul = logbookReq.Judul
	logbook.Topik = logbookReq.Topik
	logbook.Pembimbing = logbookReq.Pembimbing

	if errDb := database.DB.Create(&logbook).Error; errDb != nil {
		return c.Status(500).JSON(fiber.Map{

			"message": errDb.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "logbook created",
		"data":    logbook,
	})
}

func GetAllLogbook(c *fiber.Ctx) error {
	logbooks := []entity.Logbook{}

	// Without Middleware
	if err := database.DB.Preload("User").Find(&logbooks).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	logbookCheck := entity.Logbook{}
	if logbookCheck.UserID != 0 {
		logbookCheck.UserID = logbookCheck.UserID
	}
	// With middleware
	// if err := database.DB.Find(&todos).Error; err != nil {
	// 	return c.Status(500).JSON(fiber.Map{
	// 		"message": "internal server error",
	// 	})
	// }

	return c.JSON(fiber.Map{
		"message": "data transmited",
		"data":    logbooks,
	})
}

func GetLogbookById(c *fiber.Ctx) error {
	logbookId := c.Params("id")
	logbook := entity.Logbook{}

	if err := database.DB.Preload("User").First(&logbook, "id = ?", logbookId).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "category not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "data transmitted",
		"data":    logbook,
	})
}

func UpdateLogbookById(c *fiber.Ctx) error {
	logbookUpdateReq := request.LogbookUpdateRequest{}

	// PARSE REQUEST BODY
	if errParse := c.BodyParser(&logbookUpdateReq); errParse != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "fail to parsing data",
			"error":   errParse.Error(),
		})
	}

	// VALIDATION REQUEST DATA

	validate := validator.New()
	if errValidate := validate.Struct(&logbookUpdateReq); errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "some data is not vali",
			"error":   errValidate.Error(),
		})
	}

	logbookId := c.Params("id")
	logbook := entity.Logbook{}

	if err := database.DB.First(&logbook, "id = ?", logbookId).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "logbook not found",
		})
	}

	logbook.Judul = logbookUpdateReq.Judul
	logbook.Topik = logbookUpdateReq.Topik
	logbook.Pembimbing = logbookUpdateReq.Pembimbing

	if errSave := database.DB.Save(&logbook).Error; errSave != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "logbook updated",
		"data":    logbook,
	})
}

func DeleteTLogbookById(c *fiber.Ctx) error {

	logbookID := c.Params("id")
	logbook := entity.Logbook{}

	if err := database.DB.First(&logbook, "id = ?", logbookID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "todo not found",
		})
	}

	if errDelete := database.DB.Delete(&logbook).Error; errDelete != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "logbook deleted",
	})
}
