package contacts

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mousedownco/htmx-contact-app/views"
	"net/http"
	"strconv"
	"time"
)

func HandleIndex(svc *Service, view *views.View) http.HandlerFunc {
	return func(writer http.ResponseWriter, r *http.Request) {
		var contacts []Contact
		q := r.URL.Query().Get("q")
		if q != "" {
			contacts = svc.Search(q)
			// Sleep just so the indicator appears
			time.Sleep(1 * time.Second)
		} else {
			contacts = svc.All()
		}
		data := map[string]interface{}{
			"Contacts": contacts,
			"Query":    q,
		}
		view.Render(writer, r, data)
	}
}

func HandleNew(view *views.View) http.HandlerFunc {
	return func(writer http.ResponseWriter, r *http.Request) {
		view.Render(writer, r, map[string]interface{}{
			"Contact": Contact{},
			"Errors":  map[string]string{},
		})
	}
}

func HandleNewPost(svc *Service, view *views.View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := Contact{
			First: r.FormValue("first_name"),
			Last:  r.FormValue("last_name"),
			Phone: r.FormValue("phone"),
			Email: r.FormValue("email"),
		}
		vErrors := svc.Validate(c)
		if len(vErrors) > 0 {
			view.Render(w, r, map[string]interface{}{
				"Contact": c,
				"Errors":  vErrors,
			})
		} else {
			e := svc.Save(c)
			if e != nil {
				view.Render(w, r, map[string]interface{}{
					"Contact": c,
					"Errors":  map[string]string{"General": e.Error()},
				})
			}
			views.Flash(w, r, "Created New Contact!")
			http.Redirect(w, r, "/contacts", http.StatusFound)
		}
	}
}

func HandleView(svc *Service, view *views.View) http.HandlerFunc {
	return func(writer http.ResponseWriter, r *http.Request) {
		mux.Vars(r)
		id, e := strconv.Atoi(mux.Vars(r)["id"])
		if e != nil {
			http.Error(writer, "Contact Not Found", http.StatusNotFound)
			return
		}
		c := svc.Find(id)
		if (c == Contact{}) {
			http.Error(writer, "Contact Not Found", http.StatusNotFound)
			return
		}
		view.Render(writer, r, map[string]interface{}{
			"Contact": c,
			"Errors":  map[string]string{},
		})
	}
}

func HandleEdit(svc *Service, view *views.View) http.HandlerFunc {
	return func(writer http.ResponseWriter, r *http.Request) {
		mux.Vars(r)
		id, e := strconv.Atoi(mux.Vars(r)["id"])
		if e != nil {
			http.Error(writer, "Contact Not Found", http.StatusNotFound)
			return
		}
		c := svc.Find(id)
		if (c == Contact{}) {
			http.Error(writer, "Contact Not Found", http.StatusNotFound)
			return
		}
		view.Render(writer, r, map[string]interface{}{
			"Contact": c,
			"Errors":  map[string]string{},
		})
	}
}

func HandleEditPost(svc *Service, view *views.View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mux.Vars(r)
		id, e := strconv.Atoi(mux.Vars(r)["id"])
		if e != nil {
			http.Error(w, "Contact Not Found", http.StatusNotFound)
			return
		}
		c := svc.Find(id)
		c.First = r.FormValue("first_name")
		c.Last = r.FormValue("last_name")
		c.Phone = r.FormValue("phone")
		c.Email = r.FormValue("email")

		vErrors := svc.Validate(c)
		if len(vErrors) > 0 {
			view.Render(w, r, map[string]interface{}{
				"Contact": c,
				"Errors":  vErrors,
			})
		} else {
			e = svc.Save(c)
			if e != nil {
				view.Render(w, r, map[string]interface{}{
					"Contact": c,
					"Errors":  map[string]string{"General": e.Error()},
				})
			}
			views.Flash(w, r, "Updated Contact!")
			http.Redirect(w, r, fmt.Sprintf("/contacts/%d", c.Id), http.StatusFound)
		}
	}
}

func HandleDelete(svc *Service, view *views.View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mux.Vars(r)
		id, e := strconv.Atoi(mux.Vars(r)["id"])
		if e != nil {
			http.Error(w, "Contact Not Found", http.StatusNotFound)
			return
		}
		c := svc.Find(id)
		if (c == Contact{}) {
			http.Error(w, "Contact Not Found", http.StatusNotFound)
			return
		}
		e = svc.Delete(id)
		if e != nil {
			view.Render(w, r, map[string]interface{}{
				"Contact": c,
				"Errors":  map[string]string{"General": e.Error()},
			})
		}
		if r.Header.Get("HX-Trigger") == "delete-btn" {
			views.Flash(w, r, "Deleted Contact!")
			http.Redirect(w, r, "/contacts", http.StatusSeeOther)
		} else {
			_, _ = w.Write([]byte(""))
		}
	}
}

func HandleDeleteSelected(svc *Service, view *views.View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		e := r.ParseForm()
		if e != nil {
			fmt.Sprintf("Error parsing form: %v", e)
			return
		}

		values := r.Form["selected_contact_ids"]
		_ = fmt.Sprintf("Selected Contact Ids: %v", values)
		for _, id := range r.Form["selected_contact_ids"] {
			idi, err := strconv.Atoi(id)
			if err != nil {
				_ = fmt.Sprintf("Error converting id: %v", err)
			} else {
				_ = svc.Delete(idi)
			}
		}
		view.Render(w, r, map[string]interface{}{"Contacts": svc.All()})
	}
}

func HandleEmailGet(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mux.Vars(r)
		id, e := strconv.Atoi(mux.Vars(r)["id"])
		if e != nil {
			http.Error(w, "Contact Not Found", http.StatusNotFound)
			return
		}
		c := svc.Find(id)
		if (c == Contact{}) {
			http.Error(w, "Contact Not Found", http.StatusNotFound)
			return
		}
		c.Email = r.FormValue("email")
		_, _ = w.Write([]byte(svc.Validate(c)["Email"]))
	}
}

func HandleCountGet(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Sleep just so the indicator appears
		time.Sleep(1 * time.Second)
		_, _ = w.Write([]byte(fmt.Sprintf(
			"(%d total Contacts)", len(svc.Contacts))))
	}
}
