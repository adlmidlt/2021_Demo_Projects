package sh

import (
	"github.com/google/uuid"
	"google.golang.org/genproto/googleapis/type/date"
	l "shophookah/pkg/logg"
)

type User struct {
	Id          uuid.UUID `json:"id"`
	FirstName   string    `json:"firstName"`
	SecondName  string    `json:"secondName"`
	DateOfBirth date.Date `json:"dateOfBirth"`
	Gender      string    `json:"gender"`
	NumberPhone string    `json:"numberPhone"`
	Permission  int       `json:"permission"`
	Login       string    `json:"login"`
	Password    string    `json:"password"`
}

func CreateUser(u *User) error {
	psqlDB := CreateConnToDB()
	query := `INSERT INTO users VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	if _, err := psqlDB.Exec(query,
		u.Id,
		u.FirstName,
		u.SecondName,
		u.DateOfBirth,
		u.Gender,
		u.NumberPhone,
		u.Login,
		u.Password,
		u.Permission,
	); err != nil {
		l.LogW("", err.Error())
	}
	defer CloseDB(psqlDB)
	return nil
}

/*func GetUserByLoginAndPasswd(login, passwd string) (User, error) {
	psqlDB := CreateConnToDB()

	query := "SELECT * FROM users WHERE login = $1 AND password = $2"
	rows, err := psqlDB.Query(query, login, passwd)
	CheckError(msg.ErrQueryToDB, err)
	CloseDB(psqlDB)

	var user User
	rows.Next()
	if err = rows.Scan(
		&user.Id,
		&user.FirstName,
		&user.SecondName,
		&user.DateOfBirth,
		&user.Gender,
		&user.NumberPhone,
		&user.Permission,
		&user.Login,
		&user.Password,
	); CheckError(msg.ErrRowsEmpty, err) {
		return User{}, err
	}
	CloseRows(rows, err)
	ErrRows(rows, err)
	return user, nil
}

func GetAllUsers() ([]User, error) {
	psqlDB := CreateConnToDB()

	query := "SELECT * FROM users"
	rows, err := psqlDB.Query(query)
	CheckError(msg.ErrQueryToDB, err)

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
			&user.Login,
			&user.Password,
		); CheckError(msg.ErrRowsEmpty, err) {
			return nil, err
		}
		users = append(users, user)
	}
	defer CloseDB(psqlDB)
	return users, nil
}
*/
