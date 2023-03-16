package utils

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func FiberContextFromContext(ctx context.Context) (*fiber.Ctx, error) {
	fiberContext := ctx.Value("FiberContextKey")
	if fiberContext == nil {
		err := fmt.Errorf("could not retrieve fiber.Ctx")
		return nil, err
	}

	fc, ok := fiberContext.(*fiber.Ctx)
	if !ok {
		err := fmt.Errorf("fiber.Ctx has wrong type")
		return nil, err
	}
	return fc, nil
}
