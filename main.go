package main

import (
    "github.com/gofibre/fiber/v2"
)

func main() {
    // Start a new fibre app
    app := fiber.New()

    // Listen on PORT 3000
    app.Listen(":3000")
}
