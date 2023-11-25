package auth

import (
	"fmt"
	"github.com/mousedownco/htmx-contact-app/views"
	"net/http"
)

func HandleAuth(h http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Authorization: %s\n", r.Header.Get("Authorization"))
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
