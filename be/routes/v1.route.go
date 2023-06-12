package routes

import (

	// "github.com/gofiber/fiber/middleware"
	"github.com/gofiber/fiber/v2"
	"logbook_ta/controllers"
	"logbook_ta/middleware"
)

func userRoute(app *fiber.App) {
	register := app.Group("/api/v1")
	user := register.Group("/user")

	user.Post("/register", controllers.CreateUser)
	user.Post("/login", controllers.LoginController)
	user.Get("/", middleware.Middleware, controllers.GetAllUser)

}

func v1Route(app *fiber.App) {
	v1 := app.Group("/api/v1")

	//	Karya

	karya := v1.Group("/karya")
	karya.Post("/", controllers.CreateKarya)
	karya.Get("/", controllers.GetAllKarya)
	karya.Get("/:id", controllers.GetKaryaById)
	karya.Patch("/:id", controllers.UpdateKaryaById)
	karya.Delete("/:id", controllers.DeleteKaryaById)

	//	LOGBOOK

	logbook := v1.Group("/logbook")
	logbook.Post("/", controllers.CreateLogbook)
	logbook.Get("/", controllers.GetAllLogbook)
	logbook.Get("/:id", controllers.GetLogbookById)
	logbook.Patch("/:id", controllers.UpdateLogbookById)
	logbook.Delete("/:id", controllers.DeleteTLogbookById)

}
