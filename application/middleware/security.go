package middleware

import (
	"gintoki/application/handler"
	"gintoki/utils/throwable"
	"net/http"
)

type securityMiddleware struct {
	handler.Handler
}

func (s *securityMiddleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		tokens := make(map[string]bool)
		tokens["gintoki"] = true

		token := request.Header.Get("Authorization")

		if _, exists := tokens[token]; exists {
			next.ServeHTTP(writer, request)
		} else {
			err := throwable.NewUnauthorized("Authorization required!")
			s.Error(writer, err)
		}

	})
}

func NewSecurityMiddleware() *securityMiddleware {
	return &securityMiddleware{}
}
