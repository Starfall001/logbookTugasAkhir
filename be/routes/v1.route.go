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
	//
	//// Karya
	//
	todo := v1.Group("/karya")

	todo.Post("/", controllers.CreateKarya)

	todo.Get("/", controllers.GetAllKarya)
	//
	todo.Get("/:id", controllers.GetKaryaById)
	//
	//todo.Patch("/:id", controllers.UpdateTodoById)
	//
	//todo.Delete("/:id", controllers.DeleteTodoById)
	//
	//// CATEGORY
	//
	//category := v1.Group("/category")
	//
	//category.Post("/", controllers.CreateCategory)
	//
	//category.Get("/", controllers.GetAllCategory)
	//
	//category.Get("/:id", controllers.GetCategoryById)
	//
	//category.Delete("/:id", controllers.DeleteCategoryById)
}
