package update

import (
	"encoding/json"
	bossignup "golang/bosSignup"
	"golang/login"
	"golang/postgres"
	"golang/register"
	"net/http"
)

type UserModel struct {
	Id       int
	Username string
	Password string
	Email    string
}
type LoginModel struct {
	Username string
	Password string
	Email    string
}

var userr UserModel

func Update(w http.ResponseWriter, r *http.Request) {
	register.Register(w, r)
	login.Login(w, r)

	openConnention := postgres.OpenConnention()
	db := openConnention
	r.ParseForm()

	userr.Username = r.FormValue("username")
	userr.Password = r.FormValue("password")
	userr.Email = r.FormValue("email")
	db.Exec("UPDATE userr SET username=$1,password=$2,eposta=$3 WHERE id=$4 ", userr.Username, userr.Password, userr.Email, userr.Id)

	peopleByte, _ := json.MarshalIndent(userr, "", "\t")

	w.Header().Set("Content-Type", "application/json")

	w.Write(peopleByte)

	defer db.Close()

	bossignup.BosSignup(w, r)
	db.Close()
}
