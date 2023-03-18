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

func AuthedUserFromContext(ctx context.Context) (string, error) {
	user_id := ctx.Value("user_id")

	if user_id == nil || user_id == "" {
		return "", fmt.Errorf("couldn't retrieve user id")
	}

	user_ids, ok := user_id.(string)

	if !ok {
		return "", fmt.Errorf("user_id is corrupted i guess?")
	}

	return user_ids, nil
}
