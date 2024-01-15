package controller

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/flouressaint/todo-service/internal/service"
	"github.com/labstack/echo/v4"
)

const (
	userIdCtx = "userId"
)

type AuthMiddleware struct {
	authService service.Auth
}

func (h *AuthMiddleware) UserIdentity(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := bearerToken(c.Request())
		if !ok {
			log.Printf("AuthMiddleware.UserIdentity: bearerToken: %v", fmt.Errorf("invalid auth header"))
			newErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("invalid auth header").Error())
			return nil
		}

		userId, err := h.authService.ParseToken(token)
		if err != nil {
			log.Printf("AuthMiddleware.UserIdentity: h.authService.ParseToken: %v", err)
			newErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("cannot parse token: %v", err).Error())
			return err
		}
		c.Set(userIdCtx, userId)

		return next(c)
	}
}

func bearerToken(r *http.Request) (string, bool) {
	const prefix = "Bearer "

	header := r.Header.Get(echo.HeaderAuthorization)
	if header == "" {
		return "", false
	}

	if len(header) > len(prefix) && strings.EqualFold(header[:len(prefix)], prefix) {
		return header[len(prefix):], true
	}

	return "", false
}
