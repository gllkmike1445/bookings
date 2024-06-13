package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/gllkmike1445/bookings/pkg/config"
	"github.com/gllkmike1445/bookings/pkg/handlers"
	"github.com/gllkmike1445/bookings/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

// This helps us to use app.(...) config.go 's properties
var app config.AppConfig

var session *scs.SessionManager

// main is the main application function
func main() {

	//Turn this true when in production
	app.InProduction = false

	//creating sessions
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true //the moment when someone closes the browser set to false
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction //binding

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = tc

	app.UsaCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.Newtamplate(&app)

	fmt.Printf(fmt.Sprintf("Starting server at port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
