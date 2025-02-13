package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/61lf0yl3/div-01/forum/internal/models"
)

// SignupHandler ...
func (c *Connect) SignupHandler(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("_session_")
	if err == nil {
		http.Error(w, "No, no, no! Get the f*ck out here!", http.StatusNotFound)
		return
	}

	if r.Method == "GET" {
		template, err := template.ParseFiles("templates/signup.html")
		if err != nil {
			log.Println(err)
		}
		template.Execute(w, nil)
	}

	if r.Method == "POST" {
		r.ParseForm()
		newUser := &models.User{
			Username: r.PostFormValue("username"),
			Email:    r.PostFormValue("email"),
			Password: r.PostFormValue("password"),
		}
		var confirmPwd string = r.PostFormValue("confirpwd")
		if ValidateInput(newUser, confirmPwd) != "filled correctly" {
			// if input not validated then show to the front for user
		}
		//if  err := c.Database.DB.
	}
}
