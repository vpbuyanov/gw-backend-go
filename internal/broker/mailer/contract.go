package mailer

import "github.com/vpbuyanov/gw-backend-go/internal/models"

type mailer interface {
	SendEmail(models.Gmail) error
}
