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
	page.Get("/user", controller.GetUser)
	
	page.Get("/parfume/:id", controller.GetParfumeID)
	page.Post("/insert", controller.InsertParfume)
	page.Put("parfume/update/:id", controller.UpdateParfume)
	page.Delete("parfume/delete/:id", controller.DeleteParfumeByID)

	page.Get("/user/:id", controller.GetUserID)
	page.Post("/post", controller.InsertUser)
	page.Put("user/put/:id", controller.UpdateUser)
	page.Delete("user/hapus/:id", controller.DeleteUserByID)
	

	page.Get("/docs/*", swagger.HandlerDefault)
}