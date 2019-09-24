package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	llctx "../context"

	"../models"
	"github.com/gorilla/csrf"
	"golang.org/x/oauth2"
)

// NewOAuths is used to create a new OAuths controller.
// This function will panic if the templates are not
// parsed correctly, and should only be used during
// initial setup.
func NewOAuths(os models.OAuthService, dbxConfig *oauth2.Config) *OAuths {
	return &OAuths{
		os:       os,
		dbxOAuth: dbxConfig,
	}
}

type OAuths struct {
	os       models.OAuthService
	dbxOAuth *oauth2.Config
}

func (o *OAuths) DropboxConnect(w http.ResponseWriter, r *http.Request) {
	state := csrf.Token(r)
	cookie := http.Cookie{
		Name:     "oauth_state",
		Value:    state,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	url := o.dbxOAuth.AuthCodeURL(state)
	http.Redirect(w, r, url, http.StatusFound)
}

func (o *OAuths) DropboxCallback(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	state := r.FormValue("state")
	cookie, err := r.Cookie("oauth_state")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if cookie == nil || cookie.Value != state {
		http.Error(w, "Invalid state provided", http.StatusBadRequest)
		return
	}
	cookie.Value = ""
	cookie.Expires = time.Now()
	http.SetCookie(w, cookie)

	code := r.FormValue("code")
	token, err := o.dbxOAuth.Exchange(context.TODO(), code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user := llctx.User(r.Context())
	existing, err := o.os.Find(user.ID, models.OAuthDropbox)
	if err == models.ErrNotFound {
		// noop
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		o.os.Delete(existing.ID)
	}
	userOAuth := models.OAuth{
		UserID:  user.ID,
		Token:   *token,
		Service: models.OAuthDropbox,
	}
	err = o.os.Create(&userOAuth)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%+v", token)
	fmt.Fprintln(w, r.FormValue("code"), " state: ", r.FormValue("state"))
}

func (o *OAuths) DropboxTest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	path := r.FormValue("path")

	user := llctx.User(r.Context())
	userOAuth, err := o.os.Find(user.ID, models.OAuthDropbox)
	if err != nil {
		panic(err)
	}
	token := userOAuth.Token

	data := struct {
		Path string `json:"path"`
	}{
		Path: path,
	}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	client := o.dbxOAuth.Client(context.TODO(), &token)
	req, err := http.NewRequest(http.MethodPost, "https://api.dropbox.com/2/files/list_folder", bytes.NewReader(dataBytes))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.Copy(w, resp.Body)
}
