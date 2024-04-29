package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/agusheryanto182/go-social-media/internal/helper"
	"github.com/agusheryanto182/go-social-media/internal/model/web"
	"github.com/agusheryanto182/go-social-media/internal/service"
	jwtSvc "github.com/agusheryanto182/go-social-media/utils/jwt"
	Njwt "github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

type AuthMiddleware struct {
	userService service.UserService
	jwtService  jwtSvc.IJwt
}

func NewAuthMiddleware(userService service.UserService, jwtService jwtSvc.IJwt) *AuthMiddleware {
	return &AuthMiddleware{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (s *AuthMiddleware) Protected(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		header := r.Header.Get("Authorization")

		if !strings.HasPrefix(header, "Bearer") {
			webResponse := &web.WebResponse{
				Code:    http.StatusUnauthorized,
				Message: "missing token",
			}
			helper.WriteResponse(w, webResponse)
			return
		}

		tokenString := strings.TrimPrefix(header, "Bearer ")
		token, err := s.jwtService.ValidateToken(tokenString)
		if err != nil {
			webResponse := &web.WebResponse{
				Code:    http.StatusUnauthorized,
				Message: "invalid token",
			}
			helper.WriteResponse(w, webResponse)
			return
		}

		claim, ok := token.Claims.(Njwt.MapClaims)
		if !ok || !token.Valid {
			webResponse := &web.WebResponse{
				Code:    http.StatusUnauthorized,
				Message: "invalid token",
			}
			helper.WriteResponse(w, webResponse)
			return
		}

		userID := uint64(claim["id"].(float64))

		user, err := s.userService.GetByID(ctx, userID)
		if err != nil {
			webResponse := &web.WebResponse{
				Code:    http.StatusUnauthorized,
				Message: "user not found",
			}
			helper.WriteResponse(w, webResponse)
			return
		}

		ctx = context.WithValue(ctx, "CurrentUser", *user)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func AuthMiddlewareFunc(next http.Handler) http.Handler {
	return (&AuthMiddleware{}).Protected(next)
}

func AuthMiddlewareMux(next http.Handler) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return AuthMiddlewareFunc(next)
	}
}
