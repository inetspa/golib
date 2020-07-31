package web

import (
	"github.com/labstack/echo/v4"
	"strings"
)

// Extract token from http header
func GetTokenFromHeader(c echo.Context, tokenType string, header string) string {
	return extractToken(c.Request().Header.Get(header), tokenType)
}

func extractToken(token string, tokenType string) string {
	token = strings.TrimSpace(token)
	tokenType += " "
	if token == "" || len(token) < (len(tokenType)+1) || strings.ToLower(token[:len(tokenType)]) != strings.ToLower(tokenType) {
		return ""
	}
	token = strings.TrimSpace(token[len(tokenType):])
	return token
}
