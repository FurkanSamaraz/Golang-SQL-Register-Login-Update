package register

import (
	"encoding/json"
	"fmt"
	bossignup "golang/bosSignup"
	isvalid "golang/isValid"
	"golang/postgres"
	"net/http"
	"strings"

	emailControl "github.com/FurkanSamaraz/emailControl"
)

var uname, pwd, email string

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

func Register(w http.ResponseWriter, r *http.Request) {

	openConnention := postgres.OpenConnention()
	db := openConnention
	r.ParseForm()
	var lg LoginModel

	lg.Username = r.FormValue("username")
	lg.Password = r.FormValue("password")
	lg.Email = r.FormValue("email")

	rows, _ := db.Query("SELECT * FROM userr")
	for rows.Next() {
		rows.Scan(&userr.Id, &userr.Username, &userr.Password, &userr.Email)

	}

	uCheck := strings.Contains(lg.Password, lg.Username)
	eCheck := strings.Contains(lg.Password, lg.Email)

	if uCheck == true || eCheck == true {
		fmt.Fprintf(w, "Password must not contain username or email.")
	} else {

		if isvalid.IsVd(lg.Password) != true {
			fmt.Fprintf(w, "Use special characters, numbers, upper and lower case letters in the password.")
		} else {

			if lg.Username == "" || lg.Password == "" || lg.Email == "" {
				fmt.Fprintf(w, "cannot be empty")
			} else {
				if userr.Username == lg.Username {
					fmt.Fprintf(w, "username is used")
				} else {
					if emailControl.CheckEmail(lg.Email) == true {

						db.Exec("INSERT INTO userr(username,password,eposta) VALUES($1,$2,$3)", lg.Username, lg.Password, lg.Email)

						peopleByte, _ := json.MarshalIndent(lg, "", "\t")

						w.Header().Set("Content-Type", "application/json")

						w.Write(peopleByte)

						defer db.Close()

						bossignup.BosSignup(w, r)
						db.Close()
					} else {
						fmt.Fprintln(w, "record failed error email!! ", uname)
					}
				}
			}
		}
	}
}
