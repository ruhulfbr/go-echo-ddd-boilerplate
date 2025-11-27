package bootstrap

import (
	"context"
	"fmt"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/config"
)

func InitOIDC(cfg *config.Config) (*oidc.Provider, *oidc.IDTokenVerifier, error) {
	provider, err := oidc.NewProvider(context.Background(), "https://accounts.google.com")
	if err != nil {
		return nil, nil, fmt.Errorf("oidc.NewProvider: %w", err)
	}

	verifier := provider.Verifier(&oidc.Config{
		ClientID: cfg.OAuth.ClientID,
	})

	return provider, verifier, nil
}
