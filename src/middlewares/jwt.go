package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/lexgalante/go.customers/src/controllers"
	"github.com/lexgalante/go.customers/src/infrastructures"
	"github.com/lexgalante/go.customers/src/models"
)

//JwtSecurityMiddleware -> authenticathion middleware
func JwtSecurityMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//when route is /register or /login the middleware cannot execute
		if r.URL.Path == "/register" || r.URL.Path == "/login" {
			//call next handler
			next.ServeHTTP(w, r)
		} else {
			accessToken := r.Header.Get("Authorization")
			accessToken = strings.ReplaceAll(accessToken, "Bearer ", "")

			token, err := validateToken(accessToken)
			if err != nil {
				controllers.Unauthorized(w, r, models.MakeUnauthorizedError())
				return
			}
			//check token is valid
			if claims, ok := token.Claims.(*infrastructures.UserCustomClaims); ok && token.Valid {
				//check user is an administrator
				if r.Method == http.MethodDelete && !claims.IsAdmin {
					controllers.Forbidden(w, r, models.MakeForbiddenError())
					return
				}
				//call next handler
				next.ServeHTTP(w, r)
			} else {
				controllers.Unauthorized(w, r, models.MakeUnauthorizedError())
				return
			}
		}
	})
}

//validateToken -> validate access token see https://pkg.go.dev/github.com/golang-jwt/jwt#example-Parse-Hmac
func validateToken(accessToken string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(accessToken, &infrastructures.UserCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})
}
