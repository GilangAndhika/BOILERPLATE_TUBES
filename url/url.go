package url

import (
	"github.com/barganakukuhraditya/BOILERPLATE_TUBES/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger" // swagger handler
)

func Web(page *fiber.App) {
	// page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	// page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth

	page.Get("/", controller.Sink)
	page.Post("/", controller.Sink)
	page.Put("/", controller.Sink)
	page.Patch("/", controller.Sink)
	page.Delete("/", controller.Sink)
	page.Options("/", controller.Sink)

	page.Get("/parfume", controller.GetParfume)
	page.Post("/insert", controller.InsertParfume)
	page.Put("/edit", controller.UpdateParfume)
	page.Delete("/delete", controller.DeleteParfumeByID)

	page.Get("/docs/*", swagger.HandlerDefault)
}