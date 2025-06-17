package main

import (
	"github.com/elghazx/perpustakaan/internal/api"
	"github.com/elghazx/perpustakaan/internal/config"
	"github.com/elghazx/perpustakaan/internal/connection"
	"github.com/elghazx/perpustakaan/internal/repository"
	"github.com/elghazx/perpustakaan/internal/service"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cnf := config.Get()
	dbConnection := connection.GetDatabase(cnf.Database)

	app := fiber.New()

	customerRepository := repository.NewCustomer(dbConnection)

	customerService := service.NewCustomer(customerRepository)

	api.NewCustomer(app, customerService)

	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}
