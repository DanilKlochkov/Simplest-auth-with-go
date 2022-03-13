package main

import (
	"fmt"
	"go-auth/users"
	"html/template"
	"net/http"
)

func getSignInPage(w http.ResponseWriter, r *http.Request) {
	templating(w, "sign_in.html", nil)
}

func getSignUpPage(w http.ResponseWriter, r *http.Request) {
	templating(w, "sign_up.html", nil)
}

func templating(w http.ResponseWriter, fileName string, data interface{}) {
	t, _ := template.ParseFiles(fileName)
	t.ExecuteTemplate(w, fileName, data)
}

func signInUser(w http.ResponseWriter, r *http.Request) {
	newUser := getUser(r)
	ok := users.DefaultUserService.VerifyUser(newUser)
	if !ok {
		fileName := "sign_in.html"
		t, _ := template.ParseFiles(fileName)
		t.ExecuteTemplate(w, fileName, fmt.Sprintf("%s Sign-in Failure.", newUser.Email))
		return
	}
	fileName := "sign_in.html"
	t, _ := template.ParseFiles(fileName)
	t.ExecuteTemplate(w, fileName, fmt.Sprintf("%s Sign-in Success.", newUser.Email))
	return
}

func signUpUser(w http.ResponseWriter, r *http.Request) {
	newUser := getUser(r)
	err := users.DefaultUserService.CreateUser(newUser)
	if err != nil {
		fileName := "sign_up.html"
		t, _ := template.ParseFiles(fileName)
		t.ExecuteTemplate(w, fileName, fmt.Sprintf("New User %s Sign-up Failure.", newUser.Email))
		return
	}
	fileName := "sign_up.html"
	t, _ := template.ParseFiles(fileName)
	t.ExecuteTemplate(w, fileName, fmt.Sprintf("New User %s Sign-up Success.", newUser.Email))
	return
}

func getUser(r *http.Request) users.User {
	email := r.FormValue("email")
	password := r.FormValue("password")
	return users.User{
		Email:    email,
		Password: password,
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/sign-in":
		signInUser(w, r)
	case "/sign-up":
		signUpUser(w, r)
	case "/sign-in-form":
		getSignInPage(w, r)
	case "/sign-up-form":
		getSignUpPage(w, r)
	}
}

func main() {
	http.HandleFunc("/", userHandler)
	http.ListenAndServe("localhost:5000", nil)
}
