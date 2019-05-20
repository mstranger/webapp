package controllers

import (
	"fmt"
	"net/http"

	"../views"
)

// NewUsers is used to create a new Users controller.
// This function will panic if the templates are not
// parsed correctly, and should only be used during
// initial setup.
func NewUsers() *Users {
	return &Users{
		// NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
		NewView: views.NewView("bootstrap", "users/new"),
	}
}

type Users struct {
	NewView *views.View
}

// New is used to render the form where a user can create
// a new user account.
//
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, nil)
}

type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema: "password"`
}

// Create is used to precess the signup form when a user
// submits it. This is used to create a new user account.
//
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}

	// if err := r.ParseForm(); err != nil {
	// 	panic(err)
	// }

	// dec := schema.NewDecoder()
	// if err := dec.Decode(&form, r.PostForm); err != nil {
	// 	panic(err)
	// }

	fmt.Fprintln(w, form)

	// r.PostForm = map[string][]string
	// fmt.Fprintln(w, r.PostForm["email"])
	// fmt.Fprintln(w, r.PostFormValue("email"))
	// fmt.Fprintln(w, r.PostForm["password"])
	// fmt.Fprintln(w, r.PostFormValue("password"))
}
