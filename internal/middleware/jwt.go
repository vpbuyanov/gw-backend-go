package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	"github.com/vpbuyanov/gw-backend-go/internal/configs"
	"github.com/vpbuyanov/gw-backend-go/internal/entity"
	"github.com/vpbuyanov/gw-backend-go/internal/logger"
	"github.com/vpbuyanov/gw-backend-go/internal/models"
)

const (
	contextKeyUser = "Authorization"
	day            = 24
)

// SignedToken - подписание JWT для авторизированного пользователя токена.
func SignedToken(ctx *fiber.Ctx) error {
	user, ok := ctx.Context().Value("User").(models.User)
	if !ok {
		entity.ErrorWithAbort(ctx, http.StatusBadRequest, "not header login", errors.New("not header login"))
		return nil
	}
	payload := jwt.MapClaims{
		"Issuer":    user.Name,
		"ExpiresAt": jwt.NewNumericDate(time.Now().UTC().Add(day * time.Hour)),
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
		entity.ErrorWithAbort(ctx, http.StatusBadRequest, "can not signed string", err)
		return nil
	}

	ctx.Set(contextKeyUser, t)
	err = ctx.SendStatus(http.StatusOK)
	if err != nil {
		logger.Log.Errorf(entity.ErrorSendRequest, err)
		return nil
	}

	return nil
}

// CompareToken - проверка корректности JWT токена.
func CompareToken(ctx *fiber.Ctx) error {
	headers := ctx.GetReqHeaders()
	tokens, ok := headers[contextKeyUser]
	if !ok || len(tokens) < 1 {
		entity.ErrorWithAbort(ctx, http.StatusUnauthorized, "token is empty", errors.New("token is empty"))
		return nil
	}

	token := tokens[0]
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok = token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		path := "./config.yaml"
		cfg := configs.MustConfig(&path)

		secretKey := []byte(cfg.JWT.SecretAuth)

		return secretKey, nil
	})
	if err != nil {
		entity.ErrorWithAbort(ctx, http.StatusForbidden, "jwt token is not valid", err)
		return nil
	}

	payload, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		entity.ErrorWithAbort(ctx, http.StatusUnauthorized, "payload is not a valid", errors.New("payload is not a valid"))
		return nil
	}

	t, ok := payload["ExpiresAt"].(float64)
	if !ok {
		entity.ErrorWithAbort(ctx, http.StatusUnauthorized, "expiresAt is empty", errors.New("expiresAt is empty"))
		return nil
	}

	now := time.Now().Unix()

	if now > int64(t) {
		entity.ErrorWithAbort(ctx, http.StatusUnauthorized, "JWT token is expired", errors.New("JWT token is expired"))
		return nil
	}

	ctx.Locals("UserID", jwtToken.Header["kid"])

	err = ctx.Next()
	if err != nil {
		logger.Log.Errorf(entity.ErrorNext, err)
		return nil
	}

	return nil
}

// ClearJWT - очистка Reset токена.
func ClearJWT(ctx *fiber.Ctx) error {
	ctx.Set(contextKeyUser, "")

	err := ctx.Next()
	if err != nil {
		logger.Log.Errorf(entity.ErrorNext, err)
		return nil
	}
	return nil
}

// CompareAdminToken - проверка корректности JWT токена для админа.
func CompareAdminToken(c *fiber.Ctx) error {
	err := CompareToken(c)
	if err != nil {
		logger.Log.Errorf("err: %v", err)
		return nil
	}

	return nil
}
