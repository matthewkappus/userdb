package main

import (
	"fmt"
	"net/http"
	"os"
)

var sessionUser *User

func greet(w http.ResponseWriter, r *http.Request) {
	if sessionUser == nil {
		http.Redirect(w, r, "/loginform", http.StatusTemporaryRedirect)
		return
	}
	fmt.Fprintf(w, "<h1>Welcome %s</h1><h2>(id: %d)</h2>", sessionUser.FirstName, sessionUser.ID)

}

func lf(w http.ResponseWriter, r *http.Request) {
	// form takes a "name" and posts it to "/login"
	fmt.Fprint(w, f)
}

func login(w http.ResponseWriter, r *http.Request) {

	var err error
	if sessionUser, err = Database.Get(r.FormValue("name")); err != nil {
		http.Error(w, http.StatusText(401), http.StatusUnauthorized)
		return
	}
	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

func main() {

	http.HandleFunc("/loginform", lf)
	http.HandleFunc("/login", login)
	http.HandleFunc("/", greet)

	fmt.Fprint(os.Stdout, "Command+Click URL to Launch Websit: http://127.0.0.1:8080\n")
	http.ListenAndServe(":8080", nil)
}

const f = `
<html><form method="POST" action="/login">
	<input type="text" placeholder="First Name" name="name">
	<input type="submit">
</form></html>
`
