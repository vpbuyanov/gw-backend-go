package JWT

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	"github.com/vpbuyanov/gw-backend-go/configs"
	"github.com/vpbuyanov/gw-backend-go/internal/models"
)

const (
	contextKeyUser = "Authorization"
)

func SignedToken(ctx *fiber.Ctx, JWT configs.JWT) error {
	user, ok := ctx.Context().Value("User").(models.User)
	if !ok {
		return ctx.Status(http.StatusBadRequest).JSON(errors.New("not header login"))
	}
	payload := jwt.MapClaims{
		"Issuer":    user.Name,
		"ExpiresAt": jwt.NewNumericDate(time.Now().UTC().Add(24 * time.Hour)),
		"IssuedAt":  jwt.NewNumericDate(time.Now().UTC()),
		"NotBefore": jwt.NewNumericDate(time.Now().UTC()),
		"UserID":    user.UUID,
		"Email":     user.Email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token.Header["kid"] = user.UUID

	secretKey := []byte(JWT.Secret)

	t, err := token.SignedString(secretKey)
	if err != nil {
		return err
	}

	ctx.Set("Authorization", t)
	return ctx.SendStatus(http.StatusOK)
}

func CompareToken(c *fiber.Ctx, JWT configs.JWT) error {
	headers := c.GetReqHeaders()
	tokens, ok := headers[contextKeyUser]
	if !ok || len(tokens) < 1 {
		return c.SendStatus(http.StatusForbidden)
	}
	token := tokens[0]
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		secretKey := []byte(JWT.Secret)

		return secretKey, nil
	})
	if err != nil {
		return c.Status(http.StatusForbidden).JSON(err)
	}

	payload, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return c.SendStatus(http.StatusForbidden)
	}

	t, ok := payload["ExpiresAt"].(float64)
	if !ok {
		return c.SendStatus(http.StatusForbidden)
	}

	now := time.Now().Unix()

	if now > int64(t) {
		return c.SendStatus(http.StatusForbidden)
	}

	c.Locals("UserID", jwtToken.Header["kid"])

	return c.Next()
}
