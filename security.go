package main

import (
	"context"
	"net/http"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware/security/jwt"
	"github.com/yudppp/viron-goa/app"
)

const (
	JWT_SECRET               = "your_secret_key"
	TOKEN_EXPIRED_DURICATION = time.Hour * 24
)

func NewJWTMiddleware() goa.Middleware {
	validationHandler, _ := goa.NewMiddleware(func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		token := jwt.ContextJWT(ctx)
		claims, ok := token.Claims.(jwtgo.MapClaims)
		if !ok {
			return jwt.ErrJWTError("unsupported claims shape")
		}
		err := claims.Valid()
		if err != nil {
			return jwt.ErrJWTError(err)
		}
		email, ok := claims["email"].(string)
		if !ok || email == "" {
			return jwt.ErrJWTError("unauthorized email")
		}
		goa.LogInfo(ctx, "auth", "email", email)
		return nil
	})
	return jwt.New(JWT_SECRET, validationHandler, app.NewJWTSecurity())
}

func checkAuth(email, password string) (bool, error) {
	// TODO: implement
	// all ok now!!
	return true, nil
}

func CreateJWT(c map[string]interface{}) string {
	claim := jwtgo.MapClaims(c)
	claim["exp"] = time.Now().Add(TOKEN_EXPIRED_DURICATION).Unix()
	claim["iat"] = time.Now().Unix()
	claim["nbf"] = time.Now().Unix()
	claim["iss"] = "virongoa"
	claim["sub"] = "access_token"
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claim)
	tokenString, _ := token.SignedString([]byte(JWT_SECRET))
	return tokenString
}
