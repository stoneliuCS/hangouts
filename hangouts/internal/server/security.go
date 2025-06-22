package server

import (
	"context"
	"errors"
	api "hangouts/internal/api"
	"log/slog"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type SecurityHandler struct {
	secretKey string
	logger    *slog.Logger
}

// HandleBearerAuth implements api.SecurityHandler.
func (s SecurityHandler) HandleBearerAuth(ctx context.Context, operationName api.OperationName, t api.BearerAuth) (context.Context, error) {
	bearerToken := t.GetToken()
	if bearerToken == "" {
		return ctx, errors.New("bearer token is empty")
	}
	tokenString := strings.TrimPrefix(bearerToken, "Bearer ")
	keyFunc := func(token *jwt.Token) (any, error) {
		err := token.Method.Verify(tokenString, token.Signature, s.secretKey)
		return token, err
	}
	_, err := jwt.Parse(tokenString, keyFunc)
	if err != nil {
		return ctx, err
	}
	return ctx, nil
}

func createSecurityHandler(key string, logger *slog.Logger) api.SecurityHandler {
	return SecurityHandler{secretKey: key, logger: logger}
}
