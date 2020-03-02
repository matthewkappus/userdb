package main

import (
	"fmt"
	"net/http"
)

var sessionUser *User

func greet(w http.ResponseWriter, r *http.Request) {
	if sessionUser == nil {
		http.Redirect(w, r, "/loginform", http.StatusTemporaryRedirect)
		return
	}
	fmt.Fprintf(w, "<h1>Welcome %s</h1>", r.FormValue("name"))

}

func lf(w http.ResponseWriter, r *http.Request) {
	// form takes a "name" and posts it to "/login"
	fmt.Fprint(w, f)
}

func login(w http.ResponseWriter, r *http.Request) {

	var err error
	if sessionUser, err = Database.Get(r.FormValue("name")); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

func main() {

	http.HandleFunc("/loginform", lf)
	http.HandleFunc("/login", login)
	http.HandleFunc("/", greet)

	http.ListenAndServe(":8080", nil)
}

const f = `
<html><form method="POST" action="/login">
	<input type="text" placeholder="First Name" name="name">
	<input type="submit">
</form></html>
`
