package token

import (
	"crypto/rand"

	"github.com/vpbuyanov/gw-backend-go/internal/entity"
)

func GenerateRefreshToken() string {
	b := make([]byte, entity.LenRefreshToken)

	_, _ = rand.Read(b)

	token := make([]rune, entity.LenRefreshToken)
	for i := 0; i < entity.LenRefreshToken; i++ {
		index := int(b[i]) % len(entity.RefreshTokenSymbol)
		token[i] = entity.RefreshTokenSymbol[index]
	}

	return string(token)
}
