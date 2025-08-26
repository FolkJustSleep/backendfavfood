package middleware

import (
	"go-template/data/model"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func DecodeCookie(ctx *fiber.Ctx) (*Token, error) {
	Token := &Token{
		Token:     new(string),
		ExpiresAt: new(int64),
	}
	cookie := ctx.Cookies("token")
	secret := []byte(os.Getenv("JWT_SECRET"))

	user, err := jwt.ParseWithClaims(cookie, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "unexpected signing method")
		}
		return secret, nil
	})
	if err != nil || !user.Valid {
		return nil, ctx.Status(fiber.StatusUnauthorized).JSON(model.Response{
			Status:  401,
			Message: "Unauthorized",
			Data:    nil,
		})
	}
	claims, status := user.Claims.(jwt.MapClaims)
	if !status {
		return nil, ctx.Status(fiber.StatusUnauthorized).JSON(model.Response{
			Status:  401,
			Message: "Unauthorized",
			Data:    nil,
		})
	}
	Token.UserID = claims["user_id"].(string)
	*Token.Token = user.Raw
	*Token.ExpiresAt = int64(claims["exp"].(float64))
	Token.Role = claims["role"].(string)

	return Token, nil
}


func CheckRole(ctx *fiber.Ctx) error {
	Token := &Token{
		Token:     new(string),
		ExpiresAt: new(int64),
	}
	cookie := ctx.Cookies("token")
	secret := []byte(os.Getenv("JWT_SECRET"))

	user, err := jwt.ParseWithClaims(cookie, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "unexpected signing method")
		}
		return secret, nil
	})
	if err != nil || !user.Valid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(model.Response{
			Status:  401,
			Message: "Unauthorized",
			Data:    nil,
		})
	}
	claims, status := user.Claims.(jwt.MapClaims)
	if !status {
		return ctx.Status(fiber.StatusUnauthorized).JSON(model.Response{
			Status:  401,
			Message: "Unauthorized",
			Data:    nil,
		})
	}
	*Token.Token = user.Raw
	*Token.ExpiresAt = int64(claims["exp"].(float64))
	Token.Role = claims["role"].(string)


	if Token.Role != "Manager" && Token.Role != "Admin" {
		return ctx.Status(fiber.StatusForbidden).JSON(model.Response{
			Status:  403,
			Message: "Forbidden",
			Data:    nil,
		})
	}
	return ctx.Next()
}