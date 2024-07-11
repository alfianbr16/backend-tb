package url

import (
	"github.com/alfianbr16/backend-tb/controller"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/swagger" // swagger handler
)

func Web(page *fiber.App) {
	page.Get("/makanan", controller.GetMakanan)
	page.Get("/makanan/:id", controller.GetMakananID)
	page.Post("/ins", controller.InsertDataMakanan)
	page.Put("/update/:id", controller.UpdateData)
	page.Delete("/delete/:id", controller.DeleteMakananByID)

	// page.Get("/docs/*", swagger.HandlerDefault)
}