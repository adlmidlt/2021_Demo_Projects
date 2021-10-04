package sh

import (
	"github.com/google/uuid"
	"net/http"
	l "shophookah/pkg/logg"
	msg "shophookah/pkg/logg/message"
	"text/template"
)

func UserSingUp(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("internal/sh/web/html/create.html"))

	if r.Method != http.MethodPost {
		if err := tmpl.Execute(w, nil); err != nil {
			l.LogE(msg.ErrParseTmpl, err.Error())
		}
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userAdd := User{
		Id:          uuid.New(),
		FirstName:   r.FormValue("firstName"),
		SecondName:  r.FormValue("secondName"),
		DateOfBirth: r.FormValue("dateOfBirth"),
		Gender:      r.FormValue("gender"),
		NumberPhone: r.FormValue("numberPhone"),
		Permission:  1,
		Login:       r.FormValue("login"),
		Password:    r.FormValue("password"),
	}

	if err := CreateUser(&userAdd); err != nil {

	}
}

/*
func HomePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("internal/sh/web/html/index.html"))

	tmpl.Execute(w, "dasdasddasdasdsadasdasd")
}

func UserSingIn(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("internal/sh/web/html/auth.html"))

	if r.Method != http.MethodGet {
		if err := tmpl.Execute(w, nil); err != nil {
			l.LogE(msg.ErrParseTmpl, err.Error())
		}
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	login := r.FormValue("login")
	passwd := r.FormValue("password")
	logAndPass, _ := GetUserByLoginAndPasswd(login, passwd)

	users, err := GetAllUsers()
	CheckError(errors.New(""), err)

	for _, user := range users {
		if user.Login == logAndPass.Login && user.Password == logAndPass.Password {
			http.Redirect(w, r, "/index", http.StatusOK)
		} else {
			fmt.Fprintln(w, "Введен некорректно логин или пароль")
			http.Redirect(w, r, "/auth", http.StatusBadGateway)
		}
	}

}
*/
