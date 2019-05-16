package main

import (
	"fmt"
	"net/http"

	"./views"

	"github.com/gorilla/mux"
)

var (
	homeView    *views.View
	contactView *views.View
	signupView  *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
	// err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil)
	// if err != nil {
	// 	panic(err)
	// }
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))
	// err := contactView.Template.ExecuteTemplate(w, contactView.Layout, nil)
	// if err != nil {
	// 	panic(err)
	// }
}

func signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(signupView.Render(w, nil))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Sorry, but we couldn't find the page you are looking for</h1>")
}

func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	signupView = views.NewView("bootstrap", "views/signup.gohtml")

	/*
		var err error
		homeTemplate, err = template.ParseFiles(
			"views/home.gohtml",
			"views/layouts/footer.gohtml",
		)
		if err != nil {
			panic(err)
		}

		contactTemplate, err = template.ParseFiles(
			"views/contact.gohtml",
			"views/layouts/footer.gohtml",
		)
		if err != nil {
			panic(err)
		}
	*/

	r := mux.NewRouter()
	// routes
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/signup", signup)
	// 404 page
	r.NotFoundHandler = http.HandlerFunc(notFound)

	http.ListenAndServe(":3000", r)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
