package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type IndexController struct{}

// yzyrouter[Get:"/index"]
func (c IndexController) Index(ctx *fiber.Ctx) error {
	fmt.Println(ctx.OriginalURL())
	return nil
}