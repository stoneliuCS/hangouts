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

type Claims struct {
	jwt.RegisteredClaims
}

// HandleBearerAuth implements api.SecurityHandler.
func (s SecurityHandler) HandleBearerAuth(ctx context.Context, operationName api.OperationName, t api.BearerAuth) (context.Context, error) {
	bearerToken := t.GetToken()
	if bearerToken == "" {
		return ctx, errors.New("bearer token is empty")
	}
	tokenString := strings.TrimPrefix(bearerToken, "Bearer ")
	keyFunc := func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid JWT token")
		} else {
			return s.secretKey, nil
		}
	}
	_, err := jwt.ParseWithClaims(tokenString, &Claims{}, keyFunc)
	if err != nil {
		return nil, err
	}
	return ctx, err
}

func createSecurityHandler(key string, logger *slog.Logger) api.SecurityHandler {
	return SecurityHandler{secretKey: key, logger: logger}
}
