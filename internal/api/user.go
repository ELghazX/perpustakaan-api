package api

import (
	"context"
	"net/http"
	"time"

	"github.com/elghazx/perpustakaan/domain"
	"github.com/elghazx/perpustakaan/dto"
	"github.com/gofiber/fiber/v2"
)

type authApi struct {
	AuthService domain.AuthService
}

func NewAuth(app *fiber.App, authService domain.AuthService) {
	aa := authApi{
		AuthService: authService,
	}

	app.Post("/auth", aa.Login)
}

func (aa authApi) Login(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.AuthRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	res, err := aa.AuthService.Login(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusOK).
		JSON(dto.CreateResponseSuccess(res))
}
