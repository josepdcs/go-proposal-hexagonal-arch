package testutils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func App() *fiber.App {
	return fiber.New(fiber.Config{UnescapePath: true})
}

func AcquireFiberCtx(app *fiber.App) *fiber.Ctx {
	ctx := &fasthttp.RequestCtx{}
	fiberCtx := app.AcquireCtx(ctx)

	return fiberCtx
}

func Shutdown(app *fiber.App) {
	_ = app.Shutdown()
}
