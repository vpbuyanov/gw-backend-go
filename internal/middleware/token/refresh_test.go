package token

import (
	"testing"

	"github.com/gofiber/fiber/v2/log"
)

func TestGenerateRefreshToken(t *testing.T) {
	got := GenerateRefreshToken()
	log.Info(got)
}
