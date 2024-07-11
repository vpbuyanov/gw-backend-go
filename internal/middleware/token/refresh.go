package token

import (
	"math/rand"

	"github.com/vpbuyanov/gw-backend-go/internal/entity"
)

func GenerateRefreshToken() string {
	res := make([]rune, entity.LenRefreshToken)

	for i := range res {
		res[i] = entity.RefreshTokenSymbol[rand.Intn(len(entity.RefreshTokenSymbol))]
	}

	return string(res)
}
