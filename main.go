package main

import (
	"net/http"

	"github.com/elghazx/perpustakaan/dto"
	"github.com/elghazx/perpustakaan/internal/api"
	"github.com/elghazx/perpustakaan/internal/config"
	"github.com/elghazx/perpustakaan/internal/connection"
	"github.com/elghazx/perpustakaan/internal/repository"
	"github.com/elghazx/perpustakaan/internal/service"
	jwtMid "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cnf := config.Get()
	dbConnection := connection.GetDatabase(cnf.Database)

	app := fiber.New()
	jwtmidd := jwtMid.New(jwtMid.Config{
		SigningKey: jwtMid.SigningKey{Key: []byte(cnf.Jwt.Key)},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(http.StatusUnauthorized).
				JSON(dto.CreateResponseError("endpoint perlu token, silahkan login dulu"))
		},
	})

	customerRepository := repository.NewCustomer(dbConnection)
	userRepository := repository.NewUser(dbConnection)

	customerService := service.NewCustomer(customerRepository)
	authService := service.NewAuth(cnf, userRepository)

	api.NewCustomer(app, customerService, jwtmidd)
	api.NewAuth(app, authService)

	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}
