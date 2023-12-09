package protected

import (
	"github.com/mousedownco/htmx-contact-app/auth"
	"github.com/mousedownco/htmx-contact-app/views"
	"net/http"
)

func HandleIndex(view *views.View) auth.UserHandlerFunc {
	return func(writer http.ResponseWriter, r *http.Request, u auth.User) {
		view.Render(writer, r, map[string]interface{}{"User": u})
	}
}
