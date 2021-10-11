package handler

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"net/http"
	"projectgo/app/logg"
	"projectgo/app/model"
	"strconv"
	"text/template"
	"time"
)

// HttpHandler - Обработчик http запросов.
func HttpHandler() {
	http.HandleFunc("/users", showAllUsers)
	http.HandleFunc("/users/add", createUser)
	http.HandleFunc("/users/delete", DropUserById)
	http.HandleFunc("/users/find", GetUserById)
	http.HandleFunc("/users/edit", EditUserById)

	http.ListenAndServe(":8080", nil)
}

// showAllUsers - Показать список всех пользователей.
func showAllUsers(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("app/web/html/user.html"))
	users, err := model.FindAllUsers()
	if err != nil {
		logg.LogE(errors.New("список пользователей не найден"), err.Error())
	}
	var dataUsers []model.User
	for _, user := range users {
		if err != nil {
			logg.LogE(errors.New("ошибка парсинга даты"), err.Error())
		}
		dataUsers = append(dataUsers, model.User{
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
	err = tmpl.Execute(w, model.UserList{Users: dataUsers})
	if err != nil {
		logg.LogE(errors.New("ошибка выполнения шаблона"), err.Error())
	}
}

// createUser - Создать пользователя.
func createUser(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("app/web/html/user.html"))
	if r.Method != http.MethodPost {
		if err := tmpl.Execute(w, nil); err != nil {
			logg.LogE(errors.New("ошибка выполнения шаблона"), err.Error())
		}
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dob := r.FormValue("dateOfBirth")
	dateOfBirth, err := time.Parse("2006-01-02", dob)
	if err != nil {
		logg.LogE(errors.New("ошибка, не удалось распарсить дату"), err.Error())
	}
	permission, err := strconv.Atoi(r.FormValue("permission"))
	if err != nil {
		logg.LogE(errors.New("ошибка конвертирования"), err.Error())
	}
	dataAddUser := model.User{
		Id:          uuid.UUID{},
		FirstName:   r.FormValue("firstName"),
		SecondName:  r.FormValue("secondName"),
		DateOfBirth: dateOfBirth,
		Gender:      r.FormValue("gender"),
		NumberPhone: r.FormValue("numberPhone"),
		Permission:  permission,
		Email:       r.FormValue("email"),
		Password:    r.FormValue("password"),
	}
	err = model.AddUser(&dataAddUser)
	if err != nil {
		logg.LogE(errors.New("ошибка добавления пользователя"), err.Error())
	}
	http.Redirect(w, r, "/users", http.StatusFound)
}

// DropUserById - Удалить пользователя из БД.
func DropUserById(w http.ResponseWriter, r *http.Request) {
	_ = template.Must(template.ParseFiles("app/web/html/user.html"))
	idUser, err := uuid.Parse(r.URL.Query().Get("id"))
	if err != nil {
		logg.LogE(errors.New("ошибка не удалось распарсить uuid"), err.Error())
	}
	err = model.DeleteUserById(idUser)
	if err != nil {
		logg.LogE(errors.New("ошибка удаления пользователя"), err.Error())
	}
	http.Redirect(w, r, "/users", http.StatusFound)
}

func EditUserById(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("app/web/html/edit.html"))
	if r.Method != http.MethodPost {
		if err := tmpl.Execute(w, nil); err != nil {
			logg.LogE(errors.New("ошибка выполнения шаблона"), err.Error())
		}
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	idUser, err := uuid.Parse(r.URL.Query().Get("id"))
	if err != nil {
		logg.LogE(errors.New("ошибка пользователя не существует"), err.Error())
	}
	dob := r.FormValue("dateOfBirth")
	dateOfBirth, err := time.Parse("2006-01-02", dob)
	if err != nil {
		logg.LogE(errors.New("ошибка, не удалось распарсить дату"), err.Error())
	}
	permission, err := strconv.Atoi(r.FormValue("permission"))
	if err != nil {
		logg.LogE(errors.New("ошибка конвертирования"), err.Error())
	}
	userId, err := uuid.Parse(r.FormValue("id"))
	if err != nil {
		return
	}
	dataEditUser := model.User{
		Id:          userId,
		FirstName:   r.FormValue("firstName"),
		SecondName:  r.FormValue("secondName"),
		DateOfBirth: dateOfBirth,
		Gender:      r.FormValue("gender"),
		NumberPhone: r.FormValue("numberPhone"),
		Permission:  permission,
		Email:       r.FormValue("email"),
		Password:    r.FormValue("password"),
	}
	err = model.UpdateUserById(idUser, dataEditUser)
	if err != nil {
		logg.LogE(errors.New("ошибка обновления пользователя"), err.Error())
	}
	http.Redirect(w, r, "/users", http.StatusFound)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idUser, err := uuid.Parse(r.URL.Query().Get("id"))
	if err != nil {
		logg.LogE(errors.New("ошибка парсинга"), err.Error())
	}
	user, err := model.FindUserById(idUser)
	if err != nil {
		logg.LogE(errors.New("ошибка пользователя не существует"), err.Error())
	}
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		return
	}
}
