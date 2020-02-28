package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/micro/cli"
	"github.com/micro/micro/api"
	"github.com/micro/micro/plugin"
	"log"
	"net/http"
)

func init() {
	api.Register(plugin.NewPlugin(
		plugin.WithName("example"),
		plugin.WithInit(func(ctx *cli.Context) error {
			log.Println("Got value for example_flag", ctx.String("example_flag"))
			return nil
		}),
		plugin.WithHandler(jwtAuthWrapper()),
	))
}

var (
	publicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA4f5wg5l2hKsTeNem/V41
fGnJm6gOdrj8ym3rFkEU/wT8RDtnSgFEZOQpHEgQ7JL38xUfU0Y3g6aYw9QT0hJ7
mCpz9Er5qLaMXJwZxzHzAahlfA0icqabvJOMvQtzD6uQv6wPEyZtDTWiQi9AXwBp
HssPnpYGIn20ZZuNlX2BrClciHhCPUIIZOQn/MmqTD31jSyjoQoV7MhhMTATKJx2
XrHhR+1DcKJzQBSTAGnpYVaqpsARap+nwRipr3nUTuxyGohBTSmjJ2usSeQXHI3b
ODIRe1AuTyHceAbewn8b462yEWKARdpd9AjQW5SIVPfdsz5B6GlYQ5LdYKtznTuy
7wIDAQAB
-----END PUBLIC KEY-----`
)

func jwtAuthWrapper() plugin.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == `/user/login` {
				h.ServeHTTP(w, r)
				return
			}
			userID, err := getUserIdForJwt(r.Header.Get("Json-Web-Token"))
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(err.Error()))
				return
			}
			r.Header.Set("user_id", userID)
			h.ServeHTTP(w, r)
		})
	}
}
func getUserIdForJwt(tokenStr string) (string, error) {
	_, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		verifyKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
		if err != nil {
			return nil, err
		}
		return verifyKey, nil
	})
	if err != nil {
		return "", err
	}
	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(tokenStr, claims, nil)
	return claims[`user_id`].(string), nil
}
