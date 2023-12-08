package auth

import (
	"fmt"
	"github.com/mousedownco/htmx-contact-app/views"
	"net/http"
)

func HandleAuth(h http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Extract User Information from JWT
		jwt, e := DecodeJwtPayload(r.Header.Get("Authorization"))
		if e != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		fmt.Printf("JWT: %v\n", jwt)

		a := r.Header.Get("Authorization")
		fmt.Printf("Authorization: %s\n", a)
		h(w, r)
	})
}

func HandleAppConfig(userPoolId, clientId string, view *views.View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/javascript")
		view.Render(w, r, map[string]interface{}{
			"UserPoolId": userPoolId,
			"ClientId":   clientId,
		})
	}
}

func HandleSignUp(view *views.View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		view.Render(w, r, nil)
	}
}

func HandleSignUpConfirm(view *views.View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		sub := r.URL.Query().Get("sub")
		view.Render(w, r, map[string]interface{}{
			"Sub": sub,
		})
	}
}

func HandleSignIn(view *views.View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		view.Render(w, r, nil)
	}
}
