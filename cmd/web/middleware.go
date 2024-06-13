package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)

//
//func WriteToConsole(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		fmt.Println("Hit the page")
//		next.ServeHTTP(w, r)
//	})
//}

// NoSurf adds CSRF protection to POST request
func NoSurf(next http.Handler) http.Handler {
	crsfHandler := nosurf.New(next)

	crsfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		//because we are not using https, in production moments must...
		Secure:   app.InProduction, //binding
		SameSite: http.SameSiteLaxMode,
	})
	return crsfHandler
}

// SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
