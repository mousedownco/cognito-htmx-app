package protected

import (
	"github.com/mousedownco/htmx-contact-app/views"
	"net/http"
)

func HandleIndex(view *views.View) http.HandlerFunc {
	return func(writer http.ResponseWriter, r *http.Request) {
		view.Render(writer, r, map[string]interface{}{})
	}
}
