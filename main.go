package main

import (
	"fmt"
	"net/http"

	"./controllers"
	"./middleware"
	"./models"

	"github.com/gorilla/mux"
)

// var (
// 	homeView    *views.View
// 	contactView *views.View
//  signupView  *views.View
// )

// func home(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	must(homeView.Render(w, nil))
// err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil)
// if err != nil {
// 	panic(err)
// }
// }

// func contact(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	must(contactView.Render(w, nil))
// err := contactView.Template.ExecuteTemplate(w, contactView.Layout, nil)
// if err != nil {
// 	panic(err)
// }
// }

// func signup(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	must(signupView.Render(w, nil))
// }

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Sorry, but we couldn't find the page you are looking for</h1>")
}

const (
	host     = "localhost"
	port     = 5432
	user     = "mstranger"
	password = "password"
	dbname   = "webapp_dev"
)

func main() {
	// homeView = views.NewView("bootstrap", "views/home.gohtml")
	// contactView = views.NewView("bootstrap", "views/contact.gohtml")
	// signupView = views.NewView("bootstrap", "views/signup.gohtml")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	services, err := models.NewServices(psqlInfo)
	must(err)

	// TODO: Fix this
	defer services.Close()
	// services.DestructiveReset()
	services.AutoMigrate()

	// reset the DB
	// us.DestructiveReset()

	r := mux.NewRouter()
	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers(services.User)
	galleriesC := controllers.NewGalleries(services.Gallery, r)

	userMw := middleware.User{
		UserService: services.User,
	}
	requireUserMw := middleware.RequireUser{
		User: userMw,
	}

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

	// routes
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	// r.Handle("/signup", usersC.NewView).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	r.Handle("/login", usersC.LoginView).Methods("GET")
	r.HandleFunc("/login", usersC.Login).Methods("POST")
	// r.HandleFunc("/cookietest", usersC.CookieTest).Methods("GET")

	// Gallery routes
	r.Handle("/galleries", requireUserMw.ApplyFn(galleriesC.Index)).Methods("GET")
	r.Handle("/galleries/new", requireUserMw.Apply(galleriesC.New)).Methods("GET")
	r.HandleFunc("/galleries", requireUserMw.ApplyFn(galleriesC.Create)).Methods("POST")
	r.HandleFunc("/galleries/{id:[0-9]+}", galleriesC.Show).Methods("GET").Name(controllers.ShowGallery)
	r.HandleFunc("/galleries/{id:[0-9]+}/edit", requireUserMw.ApplyFn(galleriesC.Edit)).Methods("GET").Name(controllers.EditGallery)
	r.HandleFunc("/galleries/{id:[0-9]+}/update", requireUserMw.ApplyFn(galleriesC.Update)).Methods("POST")
	r.HandleFunc("/galleries/{id:[0-9]+}/delete", requireUserMw.ApplyFn(galleriesC.Delete)).Methods("POST")
	// r.Handle("/galleries/new", galleriesC.New).Methods("GET")
	// r.HandleFunc("/galleries", galleriesC.Create).Methods("POST")

	// 404 page
	r.NotFoundHandler = http.HandlerFunc(notFound)

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", userMw.Apply(r))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
