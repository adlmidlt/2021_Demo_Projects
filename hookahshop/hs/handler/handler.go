package handler

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"hookahshop/hs/db"
	"hookahshop/hs/entity"
	"hookahshop/hs/logg"
	"html/template"
	"net/http"
)

func Handler() {
	http.HandleFunc("/create", createUser)
	http.HandleFunc("/users", showUsers)
	http.HandleFunc("/users/edit", editUser)

	// Подключение static (*.html, *.png/jpg *.css файлов, *.js)
	http.Handle("/web/", http.StripPrefix("/web/",
		http.FileServer(http.Dir("hs/web/"))))

	http.ListenAndServe(":8080", nil)
}

// CreateUser - Создать пользователя
func createUser(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("hs/web/html/create.html"))

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	user := entity.User{
		Id:          uuid.New(),
		FirstName:   r.FormValue("firstName"),
		SecondName:  r.FormValue("secondName"),
		DateOfBirth: r.FormValue("dateOfBirth"),
		Gender:      r.FormValue("gender"),
		NumberPhone: r.FormValue("numberPhone"),
		Permission:  1,
		Email:       r.FormValue("email"),
		Password:    r.FormValue("password"),
	}

	if err := tmpl.Execute(w, entity.AddUser(&user)); err != nil {
		return
	}
}

// ShowUsers - Показать список пользователей.
func showUsers(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("hs/web/html/user.html"))
	findUsers, err := entity.FindAllUsers()
	if err != nil {
		logg.LogE(errors.New("список пользователей пуст"), err.Error())
	}
	var dataUsers []entity.User
	for _, user := range findUsers {
		dataUsers = append(dataUsers, entity.User{
			Id:          user.Id,
			FirstName:   user.FirstName,
			SecondName:  user.SecondName,
			DateOfBirth: user.DateOfBirth,
			Gender:      user.Gender,
			NumberPhone: user.NumberPhone,
			Permission:  user.Permission,
			Email:       user.Email,
			Password:    user.Password,
		})
	}
	if err = tmpl.Execute(w, entity.UserList{
		Users: dataUsers,
	}); err != nil {
		return
	}
}

func editUser(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("hs/web/html/edit.html"))
	query := `UPDATE hookah.users SET 
		first_name = $1, 
		second_name = $2,
		data_of_birth = $3,
        gender = $4,
		number_phone = $5,
        permission = $6,
		email = $7,
		password = $8
	WHERE id = $9 
	RETURNING id, first_name, second_name, data_of_birth, gender, 
	    number_phone, permission, email, password;`
	idUser, err := uuid.Parse(r.URL.Query().Get("idUser"))
	if err != nil {
		return
	}

	user, err := entity.FindUserById(idUser)
	if err != nil {
		return
	}
	id := user.Id
	firstName := r.FormValue("firstName")
	secondName := r.FormValue("firstName")
	dateOfBirth := r.FormValue("firstName")
	gender := r.FormValue("firstName")
	numberPhone := r.FormValue("firstName")
	permission := r.FormValue("firstName")
	email := r.FormValue("firstName")
	password := r.FormValue("firstName")
	res, err := db.ConnToDB().Exec(query, firstName, secondName, dateOfBirth, gender, numberPhone,
		permission, email, password, id)
	if err != nil {
		return

	}
	/*	userData := entity.User{
		Id:          user.Id,
		FirstName:   user.FirstName,
		SecondName:  user.SecondName,
		DateOfBirth: user.DateOfBirth,
		Gender:      user.Gender,
		NumberPhone: user.NumberPhone,
		Permission:  user.Permission,
		Email:       user.Email,
		Password:    user.Password,
	}*/
	// updateUser, err := entity.UpdateUser(idUser, res)

	if err = tmpl.Execute(w, res); err != nil {
		return
	}
}
