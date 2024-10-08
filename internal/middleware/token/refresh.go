package token

import (
	"crypto/rand"

	"github.com/vpbuyanov/gw-backend-go/internal/entity"
)

func GenerateRefreshToken() string {
	b := make([]byte, entity.LenRefreshToken)

	_, _ = rand.Read(b)

	token := make([]rune, entity.LenRefreshToken)
	for i, bytes := range b {
		token[i] = entity.RefreshTokenSymbol[bytes%uint8(len(entity.RefreshTokenSymbol))]
	}

	return string(token)
}
