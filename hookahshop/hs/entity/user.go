package entity

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"hookahshop/hs/db"
	"hookahshop/hs/logg"
	"time"
)

// User - Пользователь.
type User struct {
	Id          uuid.UUID `json:"id"`
	FirstName   string    `json:"firstName"`
	SecondName  string    `json:"secondName"`
	DateOfBirth string    `json:"dateOfBirth"`
	Gender      string    `json:"gender"`
	NumberPhone string    `json:"numberPhone"`
	Permission  int       `json:"permission"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
}

// UserList - Список пользователей.
type UserList struct {
	Users []User `json:"users"`
}

// FindAllUsers - Найти всех пользователей.
func FindAllUsers() ([]User, error) {
	query := "SELECT * FROM hookah.users ORDER BY ID DESC;"
	rows, err := db.ConnToDB().Query(query)
	if err != nil {
		logg.LogE(errors.New("ошибка в запросе"), err.Error())
	}
	var users []User
	for rows.Next() {
		var user User
		if err = rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.SecondName,
			&user.DateOfBirth,
			&user.Gender,
			&user.NumberPhone,
			&user.Permission,
			&user.Email,
			&user.Password,
		); err != nil {
			logg.LogE(errors.New("количество значений не совпадает с количеством столбцов"),
				err.Error())
		}
		users = append(users, user)
	}
	defer db.CloseDB()
	return users, nil
}

// FindUserById - Найти пользователя по идентификатору.
func FindUserById(id uuid.UUID) (User, error) {
	user := User{}
	query := `SELECT * FROM hookah.users WHERE id = $1;`
	row := db.ConnToDB().QueryRow(query, id)

	if err := row.Scan(
		&user.Id,
		&user.FirstName,
		&user.SecondName,
		&user.DateOfBirth,
		&user.Gender,
		&user.NumberPhone,
		&user.Permission,
		&user.Email,
		&user.Password,
	); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("совпадений не найдено")
		}
		logg.LogE(errors.New("количество значений не совпадает с количеством столбцов"),
			err.Error())
	}
	defer db.CloseDB()
	return user, nil
}

// AddUser - Добавление пользователя.
func AddUser(u *User) error {
	var id uuid.UUID
	query := `INSERT INTO hookah.users VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id;`
	_ = db.ConnToDB().QueryRow(query,
		u.Id,
		u.FirstName,
		u.SecondName,
		u.DateOfBirth,
		u.Gender,
		u.NumberPhone,
		u.Permission,
		u.Email,
		u.Password,
	)
	u.Id = id

	defer db.CloseDB()
	return nil
}

// UpdateUser - Обновление данных пользователя.
func UpdateUser(id uuid.UUID, u User) (User, error) {
	user := User{}
	dateOfBirth, err := time.Parse("2006-01-02", user.DateOfBirth)
	if err != nil {
		return User{}, err
	}
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

	if err = db.ConnToDB().QueryRow(query, u.FirstName, u.SecondName, u.DateOfBirth, u.Gender,
		u.NumberPhone, u.Permission, u.Email, u.Password, id).
		Scan(
			&user.Id,
			&user.FirstName,
			&user.SecondName,
			&dateOfBirth,
			&user.Gender,
			&user.NumberPhone,
			&user.Permission,
			&user.Email,
			&user.Password,
		); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("совпадений не найдено")
		}
		logg.LogE(errors.New("количество значений не совпадает с количеством столбцов"),
			err.Error())
	}
	defer db.CloseDB()
	return user, nil
}

// DeleteUser - Удаление пользователя по идентификатору.
func DeleteUser(id uuid.UUID) error {
	query := `DELETE FROM hookah.users WHERE id = $1;`
	if _, err := db.ConnToDB().Exec(query, id); err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("совпадений не найдено")
		}
		return err
	}
	defer db.CloseDB()
	return nil
}
