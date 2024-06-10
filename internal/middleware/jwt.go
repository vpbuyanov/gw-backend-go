package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	"github.com/vpbuyanov/gw-backend-go/internal/configs"
	"github.com/vpbuyanov/gw-backend-go/internal/models"
)

const (
	contextKeyUser = "Authorization"
	contextReset   = "Reset"
)

// SignedToken - подписание JWT для авторизированного пользователя токена
func SignedToken(ctx *fiber.Ctx) error {
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
		"IsAdmin":   user.IsAdmin,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token.Header["kid"] = user.UUID

	path := "./config.yaml"
	cfg := configs.MustConfig(&path)

	secretKey := []byte(cfg.JWT.SecretAuth)

	t, err := token.SignedString(secretKey)
	if err != nil {
		return err
	}

	ctx.Set(contextKeyUser, t)
	return ctx.SendStatus(http.StatusOK)
}

// SignedShortToken - подписание короткого JWT для восстановления пароля
func SignedShortToken(ctx *fiber.Ctx) error {
	user, ok := ctx.Context().Value("User").(models.User)
	if !ok {
		return ctx.Status(http.StatusBadRequest).JSON(errors.New("not header login"))
	}
	payload := jwt.MapClaims{
		"Issuer":    "Reset password",
		"ExpiresAt": jwt.NewNumericDate(time.Now().UTC().Add(1 * time.Hour)),
		"UserID":    user.UUID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token.Header["kid"] = user.UUID

	path := "./config.yaml"
	cfg := configs.MustConfig(&path)

	secretKey := []byte(cfg.JWT.SecretReset)

	t, err := token.SignedString(secretKey)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ctx.Set(contextReset, t)
	return ctx.SendStatus(http.StatusOK)
}

// CompareToken - проверка корректности JWT токена
func CompareToken(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()
	tokens, ok := headers[contextKeyUser]
	if !ok || len(tokens) < 1 {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "token is empty",
		})
	}
	token := tokens[0]
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		path := "./config.yaml"
		cfg := configs.MustConfig(&path)

		secretKey := []byte(cfg.JWT.SecretAuth)

		return secretKey, nil
	})
	if err != nil {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{
			"error": fmt.Errorf("jwt token is not valid: %s", err.Error()),
		})
	}

	payload, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "payload is not a valid",
		})
	}

	t, ok := payload["ExpiresAt"].(float64)
	if !ok {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "expiresAt is empty",
		})
	}

	now := time.Now().Unix()

	if now > int64(t) {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "JWT token is expired",
		})
	}

	c.Locals("UserID", jwtToken.Header["kid"])

	return c.Next()
}

// CompareShortToken - проверка корректности JWT токена для восстановления пароля
func CompareShortToken(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()
	tokens, ok := headers[contextReset]
	if !ok || len(tokens) < 1 {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "token is empty",
		})
	}
	token := tokens[0]
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		path := "./config.yaml"
		cfg := configs.MustConfig(&path)

		secretKey := []byte(cfg.JWT.SecretReset)

		return secretKey, nil
	})
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": fmt.Errorf("jwt token is not valid: %w", err),
		})
	}

	payload, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "payload is not a valid",
		})
	}

	t, ok := payload["ExpiresAt"].(float64)
	if !ok {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "expiresAt is empty",
		})
	}

	now := time.Now().Unix()

	if now > int64(t) {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "JWT token is expired",
		})
	}

	c.Locals("UserID", jwtToken.Header["kid"])

	return c.Next()
}

// ClearJWT - очистка Reset токена
func ClearJWT(c *fiber.Ctx) error {
	c.Set(contextKeyUser, "")

	return c.Next()
}

// ClearShortJWT - очистка Reset токена
func ClearShortJWT(c *fiber.Ctx) error {
	c.Set(contextReset, "")

	return c.Next()
}

// CompareAdminToken - проверка корректности JWT токена для админа
func CompareAdminToken(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()
	tokens, ok := headers[contextKeyUser]
	if !ok || len(tokens) < 1 {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "token is empty",
		})
	}
	token := tokens[0]
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		path := "./config.yaml"
		cfg := configs.MustConfig(&path)

		secretKey := []byte(cfg.JWT.SecretAuth)

		return secretKey, nil
	})
	if err != nil {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{
			"error": fmt.Errorf("jwt token is not valid: %s", err.Error()),
		})
	}

	payload, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "payload is not a valid",
		})
	}

	t, ok := payload["ExpiresAt"].(float64)
	if !ok {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "expiresAt is empty",
		})
	}

	now := time.Now().Unix()

	if now > int64(t) {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "JWT token is expired",
		})
	}

	c.Locals("UserID", jwtToken.Header["kid"])

	return c.Next()
}
