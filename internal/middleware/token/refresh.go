package token

import (
	"math/rand/v2"

	"github.com/vpbuyanov/gw-backend-go/internal/entity"
)

func GenerateRefreshToken() string {
	res := make([]rune, entity.LenRefreshToken)

	for i := range res {
		res[i] = entity.RefreshTokenSymbol[rand.IntN(len(entity.RefreshTokenSymbol))]
	}

	return string(res)
}
