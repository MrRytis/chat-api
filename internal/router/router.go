package router

import (
	"github.com/MrRytis/chat-api/internal/handler"
	"github.com/MrRytis/chat-api/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App) {
	api := app.Group("/api")

	apiV1 := api.Group("/v1")

	auth := apiV1.Group("/auth")
	auth.Post("/register", handler.Register)
	auth.Post("/login", handler.Login)
	auth.Post("/logout", middleware.Auth, handler.Logout)
	auth.Post("/refresh", handler.Refresh)

	group := apiV1.Group("/group", middleware.Auth)
	group.Post("/create", handler.CreateGroup)
	group.Get("/list", handler.GetGroupList)
	group.Get("/:group", handler.GetGroup)
	group.Put("/:group", handler.UpdateGroup)
	group.Delete("/:group", handler.DeleteGroup)
	group.Post("/:group/add", handler.AddUserToGroup)
	group.Post("/:group/remove/:user", handler.RemoveUserFromGroup)
}
