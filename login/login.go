package login

import (
	"encoding/json"
	"fmt"
	"golang/postgres"
	"net/http"
)

type UserModel struct {
	Id       int
	Username string
	Password string
	Email    string
}

var userr UserModel

func Login(w http.ResponseWriter, r *http.Request) {
	openConnention := postgres.OpenConnention()
	var people []UserModel
	db := openConnention
	r.ParseForm()
	uname := r.FormValue("username")
	pwd := r.FormValue("password")
	email := r.FormValue("email")
	rows, _ := db.Query("SELECT * FROM userr")
	for rows.Next() {
		rows.Scan(&userr.Id, &userr.Username, &userr.Password, &userr.Email)
		people = append(people, userr)

	}
	if uname == userr.Username && pwd == userr.Password && email == userr.Email {
		fmt.Fprintf(w, "Login successful\n")
		fmt.Fprintln(w, "Hello", uname)
		peopleByte, _ := json.MarshalIndent(userr, "", "\t")
		w.Write(peopleByte)

	}

	defer db.Close()
}
