package controllers

import "github.com/gofiber/fiber/v2"

func HelloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World! X")
}

func ShowEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString("Show Employee X " + id)
}

func ListEmployee(c *fiber.Ctx) error {
	return c.SendString("List Employee X")
}
