package middleware

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v4"
	"github.com/golang-jwt/jwt/v4"

	"go-template/data/model"
)

type Token struct {
	Token *string `json:"token"`
	UserID string `json:"user_id"`
	Role string `json:"role"`
	ExpiresAt *int64 `json:"expires_at"`
}

func JWTHeaderMiddleware() fiber.Handler{
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: []byte(os.Getenv("JWT_SECRET")),
		},
		
		
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		log.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(model.Response{
			Status: 401,
			Message: "Unauthorized",
			Data: nil,
		})
	},
})


}


func GenerateToken(userID string, role string)  (*Token,error) {
	Token := &Token{
		Token: new(string),
		ExpiresAt: new(int64),
	}
	Token.UserID = userID
	*Token.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()

	SigningKey := []byte(os.Getenv("JWT_SECRET"))
	
	claims := make(jwt.MapClaims)
	claims["user_id"] = userID
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["iat"] = time.Now().Unix()//created at
	claims["nbf"] = time.Now().Unix()//time that can started to use
	token , err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(SigningKey)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	Token.Token = &token
	return Token, nil
}


// decode the token and return the token if you want to check permission create logic here
func DecodeToken(ctx *fiber.Ctx) (*Token,error) { 
	Token := &Token{
		Token: new(string),
	}
	user ,status:= ctx.Locals("user").(*jwt.Token)
	if !status {
		return nil,ctx.Status(fiber.StatusUnauthorized).JSON(model.Response{
			Status: 401,
			Message: "Unauthorized",
			Data: nil,
		})
	}
	claims ,status := user.Claims.(jwt.MapClaims)
	if !status {
		return nil,ctx.Status(fiber.StatusUnauthorized).JSON(model.Response{
			Status: 401,
			Message: "Unauthorized",
			Data: nil,
		})
	}
	Token.UserID = claims["user_id"].(string)
	*Token.Token = user.Raw
	*Token.ExpiresAt = claims["exp"].(int64)
	Token.Role = claims["role"].(string)

	return Token , nil
}

