package interceptor

import (
	"github.com/NARUBROWN/spine/pkg/httperr"
	"github.com/janghanul090801/spine-clean-architecture/config"
	"github.com/janghanul090801/spine-clean-architecture/internal/tokenutil"
	"strings"

	"github.com/NARUBROWN/spine/core"
	"go.uber.org/zap"
)

type AuthInterceptor struct {
	logger *zap.Logger
}

func NewAuthInterceptor(logger *zap.Logger) *AuthInterceptor {
	return &AuthInterceptor{
		logger: logger,
	}
}

func (i *AuthInterceptor) PreHandle(ctx core.ExecutionContext, _ core.HandlerMeta) error {
	authHeader := ctx.Header("Authorization")
	if authHeader == "" {
		return httperr.Unauthorized("Authorization header is empty")
	}

	t := strings.Split(authHeader, " ")
	if len(t) == 2 {
		authToken := t[1]
		authorized, err := tokenutil.IsAuthorized(authToken, config.E.AccessTokenSecret)
		if authorized {
			userID, err := tokenutil.ExtractIDFromToken(authToken, config.E.AccessTokenSecret)
			if err != nil {
				return httperr.Unauthorized(err.Error())
			}
			ctx.Set("id", userID)
			return nil
		}
		return httperr.Unauthorized(err.Error())
	}
	return httperr.Unauthorized("Not authorized")
}

func (i *AuthInterceptor) PostHandle(core.ExecutionContext, core.HandlerMeta) {}

func (i *AuthInterceptor) AfterCompletion(core.ExecutionContext, core.HandlerMeta, error) {}
