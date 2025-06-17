package api

import (
	"context"
	"net/http"
	"time"

	"github.com/elghazx/perpustakaan/domain"
	"github.com/elghazx/perpustakaan/dto"
	"github.com/elghazx/perpustakaan/internal/util"
	"github.com/gofiber/fiber/v2"
)

type customerApi struct {
	customerService domain.CustomerService
}

func NewCustomer(app *fiber.App, customerService domain.CustomerService) {
	ca := customerApi{
		customerService: customerService,
	}
	app.Get("/customers", ca.Index)
	app.Post("/customers", ca.Create)
	app.Put("/customers/:id", ca.Update)
}

func (ca customerApi) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	res, err := ca.customerService.Index(c)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.JSON(dto.CreateResponseSuccess(res))
}

func (ca customerApi) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateCustomerRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}
	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData("validation failed", fails))
	}

	err := ca.customerService.Create(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccess(""))
}

func (ca customerApi) Update(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()
	var req dto.UpdateCustomerRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}
	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData("validation failed", fails))
	}

	// customers/:id
	req.ID = ctx.Params("id")
	err := ca.customerService.Update(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusOK).JSON(dto.CreateResponseSuccess(""))
}
