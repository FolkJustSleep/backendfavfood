package middleware

import (
	// "fmt"
	// "go-template/data/model"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func DecodeCookie(ctx *fiber.Ctx) (*Token, error) {
	Token := &Token{
		Token:     new(string),
		ExpiresAt: new(int64),
	}
	cookie := ctx.Cookies("token")
	if cookie == "" {
	return nil, fiber.NewError(fiber.StatusUnauthorized, ": No token provided")
	}
	secret := []byte(os.Getenv("JWT_SECRET"))

	user, err := jwt.ParseWithClaims(cookie, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized)
		}
		return secret, nil
	})
	if err != nil || !user.Valid {
		return nil, fiber.NewError(fiber.StatusUnauthorized)
	}
	claims, status := user.Claims.(jwt.MapClaims)
	if !status {
		return nil, fiber.NewError(fiber.StatusUnauthorized)
	}
	Token.UserID = claims["user_id"].(string)
	*Token.Token = user.Raw
	*Token.ExpiresAt = int64(claims["exp"].(float64))
	Token.Role = claims["role"].(string)
	if int64(claims["exp"].(float64)) < time.Now().Unix() {
		return nil, fiber.NewError(fiber.StatusUnauthorized, ": Token has expired")
	}
	// fmt.Println("Decoded Token:", Token)
	return Token, nil
}



func CheckRole(ctx *fiber.Ctx) error {
	Token := &Token{
		Token:     new(string),
		ExpiresAt: new(int64),
	}
	cookie := ctx.Cookies("token")
	if cookie == "" {
	return fiber.NewError(fiber.StatusUnauthorized)
	}
	secret := []byte(os.Getenv("JWT_SECRET"))

	user, err := jwt.ParseWithClaims(cookie, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, ": unexpected signing method")
		}
		return secret, nil
	})
	if err != nil || !user.Valid {
		return fiber.NewError(fiber.StatusUnauthorized)
	}
	claims, status := user.Claims.(jwt.MapClaims)
	if !status {
		return fiber.NewError(fiber.StatusUnauthorized)
	}
	*Token.Token = user.Raw
	*Token.ExpiresAt = int64(claims["exp"].(float64))
	Token.Role = claims["role"].(string)


	if Token.Role != "Manager" && Token.Role != "Admin" {
		return fiber.NewError(fiber.StatusForbidden)
	}
	return ctx.Next()
}