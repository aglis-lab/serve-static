package main

import (
	"flag"
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

var (
	port = flag.String("p", "3000", "port of serve static")
	path = flag.String("d", ".", "where the project or static file is located")
)

func main() {
	flag.Parse()

	app := fiber.New()

	app.Use("/", static.New(*path, static.Config{NotFoundHandler: func(c fiber.Ctx) error {
		fmt.Println("File Not found")
		c.Response().SetStatusCode(fiber.StatusOK)
		return c.SendFile(fmt.Sprintf("%s/%s", *path, "index.html"), fiber.SendFile{Compress: true, ByteRange: true})
	}}))

	app.Use("*", func(c fiber.Ctx) error {
		_, err := c.WriteString("404")
		return err
	})

	app.Listen(fmt.Sprintf(":%s", *port))
}
