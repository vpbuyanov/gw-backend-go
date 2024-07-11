package entity

const (

	// Tokens
	LenRefreshToken          = 32
	ExpiresMinuteAccessToken = 15
	ExpiresDayRefreshToken   = 30
	HeaderAccessToken        = "AccessToken"
	HeaderRefreshToken       = "RefreshToken"
)

var RefreshTokenSymbol = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
