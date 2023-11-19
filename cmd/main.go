package main

import (
	"embed"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/gorilla/mux"
	"github.com/mousedownco/htmx-contact-app/auth"
	"github.com/mousedownco/htmx-contact-app/contacts"
	"github.com/mousedownco/htmx-contact-app/views"
	"log"
	"net/http"
	"os"
)

//go:embed static
var staticDir embed.FS

var port = ":8080"

func main() {
	cs := contacts.NewService()

	cog := auth.NewCognito(
		os.Getenv("COGNITO_ENDPOINT"),
		os.Getenv("COGNITO_CLIENT_ID"),
		os.Getenv("COGNITO_REDIRECT_URI"))

	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.FileServer(http.FS(staticDir)))
	r.Handle("/",
		http.RedirectHandler("/contacts", http.StatusTemporaryRedirect))
	r.Handle("/app-config.js",
		auth.HandleAppConfig(
			os.Getenv("COGNITO_POOL_ID"),
			os.Getenv("COGNITO_CLIENT_ID"),
			views.NewView("partial", "auth/app-config.gohtml"))).
		Methods("GET")
	r.Handle("/contacts",
		contacts.HandleIndex(cs,
			views.NewView("partial", "contacts/rows.gohtml"))).
		Headers("HX-Trigger", "search")
	// This handler differs from the book's implementation, see README for details
	r.Handle("/contacts/delete",
		contacts.HandleDeleteSelected(cs,
			views.NewView("layout", "contacts/index.gohtml", "contacts/rows.gohtml"))).Methods("POST")
	r.Handle("/contacts",
		contacts.HandleIndex(cs,
			views.NewView("layout", "contacts/index.gohtml", "contacts/rows.gohtml")))
	r.Handle("/contacts/count", contacts.HandleCountGet(cs)).Methods("GET")
	r.Handle("/contacts/new",
		contacts.HandleNew(views.NewView("layout", "contacts/new.gohtml"))).
		Methods("GET")
	r.Handle("/contacts/new",
		contacts.HandleNewPost(cs, views.NewView("layout", "contacts/new.gohtml"))).Methods("POST")
	r.Handle("/contacts/{id:[0-9]+}",
		contacts.HandleView(cs, views.NewView("layout", "contacts/show.gohtml"))).Methods("GET")
	r.Handle("/contacts/{id:[0-9]+}/edit",
		contacts.HandleEdit(cs, views.NewView("layout", "contacts/edit.gohtml"))).Methods("GET")
	r.Handle("/contacts/{id:[0-9]+}/edit",
		contacts.HandleEditPost(cs, views.NewView("layout", "contacts/edit.gohtml"))).Methods("POST")
	r.Handle("/contacts/{id:[0-9]+}/email", contacts.HandleEmailGet(cs)).Methods("GET")
	r.Handle("/contacts/{id:[0-9]+}",
		contacts.HandleDelete(cs, views.NewView("layout", "contacts/edit.gohtml"))).Methods("DELETE")

	r.Handle("/auth/code", auth.HandleCognitoCallback(cog, "/contacts")).Methods("GET")

	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != "" {
		log.Printf("Running Lambda function %s", os.Getenv("AWS_LAMBDA_FUNCTION_NAME"))
		muxLambda := gorillamux.New(r)
		lambda.Start(muxLambda.Proxy)
	} else {
		log.Printf("Starting server on port %s", port)
		http.Handle("/", r)
		_ = http.ListenAndServe(port, nil)
	}

}
