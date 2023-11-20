package auth

import (
	"github.com/mousedownco/htmx-contact-app/views"
	"net/http"
)

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
