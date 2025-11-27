package oauth

import "context"

type OauthServiceInterface interface {
	GoogleOAuth(ctx context.Context, token string) (accessToken string, refreshToken string, exp int64, err error)
}
