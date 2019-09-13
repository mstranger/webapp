package controllers

import (
	"log"
	"net/http"
	"time"

	"../context"
	"../email"
	"../models"
	"../rand"
	"../views"
)

// NewUsers is used to create a new Users controller.
// This function will panic if the templates are not
// parsed correctly, and should only be used during
// initial setup.
func NewUsers(us models.UserService, emailer *email.Client) *Users {
	return &Users{
		NewView:   views.NewView("bootstrap", "users/new"),
		LoginView: views.NewView("bootstrap", "users/login"),
		us:        us,
		emailer:   emailer,
	}
}

type Users struct {
	NewView   *views.View
	LoginView *views.View
	us        models.UserService
	emailer   *email.Client
}

// New is used to render the form where a user can create
// a new user account.
//
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	// type Alert struct {
	// 	Level   string
	// 	Message string
	// }

	// type Data struct {
	// 	Alert Alert
	// 	Yield interface{}
	// }

	// a := Alert{
	// 	Level:   "success",
	// 	Message: "Successfully rendered a dynamic alert!",
	// }

	// d := Data{
	// 	Alert: a,
	// 	Yield: "hello!",
	// }

	// d := views.Data{
	// 	Alert: &views.Alert{
	// 		Level:   views.AlertLvlError,
	// 		Message: "something went wrong",
	// 	},
	// 	Yield: "hello!!!",
	// }

	// if err := u.NewView.Render(w, r, nil); err != nil {
	// panic(err)
	// }

	u.NewView.Render(w, r, nil)
}

type SignupForm struct {
	Name     string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// Create is used to precess the signup form when a user
// submits it. This is used to create a new user account.
//
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var vd views.Data
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		log.Println(err)
		// vd.Alert = &views.Alert{
		// 	Level:   views.AlertLvlError,
		// 	Message: views.AlertMsgGeneric,
		// }
		vd.SetAlert(err)
		u.NewView.Render(w, r, vd)
		return
	}

	user := models.User{
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
	}

	if err := u.us.Create(&user); err != nil {
		// vd.Alert = &views.Alert{
		// 	Level:   views.AlertLvlError,
		// 	Message: err.Error(),
		// }
		vd.SetAlert(err)
		u.NewView.Render(w, r, vd)
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	u.emailer.Welcome(user.Name, user.Email)

	err := u.signIn(w, &user)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	alert := views.Alert{
		Level:   views.AlertLvlSuccess,
		Message: "Welcome to Website.com!",
	}

	views.RedirectAlert(w, r, "/galleries", http.StatusFound, alert)
}

type LoginForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// Login is use to verify the provided email address and
// password and then log the user in if they are correct.
//
// POST /login
func (u *Users) Login(w http.ResponseWriter, r *http.Request) {
	vd := views.Data{}
	form := LoginForm{}
	if err := parseForm(r, &form); err != nil {
		vd.SetAlert(err)
		u.LoginView.Render(w, r, vd)
		return
	}

	user, err := u.us.Authenticate(form.Email, form.Password)
	if err != nil {
		switch err {
		case models.ErrNotFound:
			vd.AlertError("Invalid email address")
		default:
			vd.SetAlert(err)
		}

		u.LoginView.Render(w, r, vd)
		return
	}

	// cookie := http.Cookie{
	// 	Name:  "email",
	// 	Value: user.Email,
	// }
	// http.SetCookie(w, &cookie)

	err = u.signIn(w, user)
	if err != nil {
		vd.SetAlert(err)
		u.LoginView.Render(w, r, vd)
		return
	}

	http.Redirect(w, r, "/galleries", http.StatusFound)
}

// Logout is used to delete a user session cookie (remember_token)
// and then will update the user reource with a new remember
// token.
//
// POST /logout
func (u *Users) Logout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "remember_token",
		Value:    "",
		Expires:  time.Now(),
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)

	user := context.User(r.Context())
	token, _ := rand.RememberToken()
	user.Remember = token
	u.us.Update(user)
	http.Redirect(w, r, "/", http.StatusFound)
}

// signIn is used to sign the given user in via cookies
func (u *Users) signIn(w http.ResponseWriter, user *models.User) error {
	if user.Remember == "" {
		token, err := rand.RememberToken()
		if err != nil {
			return err
		}

		user.Remember = token
		err = u.us.Update(user)
		if err != nil {
			return err
		}
	}

	cookie := http.Cookie{
		Name:     "remember_token",
		Value:    user.Remember,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	return nil
}

// CookieTest is used to display cookies set on the current user
/*
	func (u *Users) CookieTest(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("remember_token")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			// http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		user, err := u.us.ByRemember(cookie.Value)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			// http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, user)
	}
*/
