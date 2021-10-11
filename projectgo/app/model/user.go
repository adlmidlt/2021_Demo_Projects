package model

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"projectgo/app/db"
	"projectgo/app/logg"
	"time"
)

// User - Пользователь.
type User struct {
	Id          uuid.UUID `json:"id"`
	FirstName   string    `json:"firstName"`
	SecondName  string    `json:"secondName"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Gender      string    `json:"gender"`
	NumberPhone string    `json:"numberPhone"`
	Permission  int       `json:"permission"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
}

type UserList struct {
	Users []User
}

// AddUser - Добавить пользователя.
func AddUser(u *User) error {
	sqlQuery := `INSERT INTO hookah.users VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`
	id := uuid.New()
	err := db.ConnToDB().QueryRow(sqlQuery, id, &u.FirstName, &u.SecondName, &u.DateOfBirth, &u.Gender, &u.NumberPhone,
		&u.Permission, &u.Email, &u.Password).Scan(&u.Id)
	if err != nil {
		logg.LogE(errors.New("ошибка в запросе"), err.Error())
	}
	defer db.CloseDB(db.ConnToDB())
	return err
}

// FindAllUsers - Найти всех пользователей.
func FindAllUsers() ([]User, error) {
	sqlQuery := `SELECT * FROM hookah.users`
	rows, err := db.ConnToDB().Query(sqlQuery)
	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.FirstName, &user.SecondName, &user.DateOfBirth, &user.Gender,
			&user.NumberPhone, &user.Permission, &user.Email, &user.Password)
		if err != nil {
			logg.LogE(errors.New("количество строк не совпадает с количеством колонок"), err.Error())
		}
		users = append(users, user)
	}
	defer db.CloseDB(db.ConnToDB())
	err = rows.Close()
	err = rows.Err()
	return users, err
}

func FindUserById(id uuid.UUID) (User, error) {
	var user User
	sqlQuery := `SELECT * FROM hookah.users WHERE id = $1;`
	err := db.ConnToDB().QueryRow(sqlQuery, id).Scan(&user.Id, &user.FirstName, &user.SecondName,
		&user.DateOfBirth, &user.NumberPhone, &user.Gender, &user.Permission, &user.Email, &user.Password)
	if err != nil {
		logg.LogE(errors.New("количество строк не совпадает с количеством колонок"), err.Error())
	}
	defer db.CloseDB(db.ConnToDB())
	return user, err
}

// DeleteUserById - Удалить пользователя по ИД.
func DeleteUserById(id uuid.UUID) error {
	sqlQuery := `DELETE FROM hookah.users WHERE id = $1;`
	_, err := db.ConnToDB().Exec(sqlQuery, id)
	if err != nil || id == uuid.Nil {
		logg.LogE(errors.New(fmt.Sprintf("такого идентификатора '%v' не существует", id)), err.Error())
	}
	defer db.CloseDB(db.ConnToDB())
	return err
}

func UpdateUserById(id uuid.UUID, u User) error {
	sqlQuery := `UPDATE hookah.users SET 
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
	insForm, err := db.ConnToDB().Prepare(sqlQuery)
	if err != nil {
		logg.LogE(errors.New("ошибка выполнения запроса"), err.Error())
	}
	_, err = insForm.Exec(&u.FirstName, &u.SecondName, &u.DateOfBirth, &u.NumberPhone, &u.Gender, &u.Permission,
		&u.Email, &u.Password, id)
	if err != nil {
		return err
	}
	defer db.CloseDB(db.ConnToDB())
	return err
}
