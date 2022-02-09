package bossignup

import (
	"fmt"
	"net/http"

	helper "github.com/FurkanSamaraz/IsEmpty"
)

var uname, pwd, email string

func BosSignup(w http.ResponseWriter, r *http.Request) {
	unameCheck := helper.IsEmpty(uname)
	pwdCheck := helper.IsEmpty(pwd)
	mailCheck := helper.IsEmpty(email)

	if unameCheck || pwdCheck || mailCheck {

	} else {
		fmt.Fprintf(w, "Error Empty! \n")
	}
}
