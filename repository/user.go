package repository

import (
	"database/sql"
	"fmt"
	"microservice/model"
	"strings"
)

type UserRepo struct {
	DB *sql.DB
}

func (o *UserRepo) Insert(userRequest model.UserRegister) (model.User, error) {
	user := model.User{}
	query := "INSERT INTO users (first_name, last_name, email,password) VALUES (?, ?, ?, ?)"
	result, err := o.DB.Exec(query, userRequest.FirstName, userRequest.LastName, userRequest.Email, userRequest.Password)
	if err != nil {
		return user, err
	}
	insertedID, err := result.LastInsertId()
	if err != nil {
		return user, err
	}
	user, err = o.GetByID(insertedID)
	if err != nil {
		return user, err
	}

	return user, nil
}
func (u *UserRepo) IsExist(fields map[string]string) (bool, error) {
	fmt.Println(fields)
	query := "SELECT"
	var keys []string
	var values []string

	for key, _ := range fields {

		keys = append(keys, key)
	}

	for _, val := range fields {

		values = append(values, val)
	}
	for _, key := range keys {

		query = fmt.Sprintf("%s %s", query, key)
	}
	query = fmt.Sprintf("%s %s", query, "FROM users WHERE")

	for _, key := range keys {
		query = fmt.Sprintf("%s %s", query, key+" = ? AND ")
	}
	query = removeLastAnd(query)
	fmt.Print(query)
	var emailPlaceholder string
	var row *sql.Row
	switch length := len(values); length {
	case 2:
		row = u.DB.QueryRow(query, values[0], values[1])
	case 3:
		row = u.DB.QueryRow(query, values[0], values[1], values[2])
	default:
		row = u.DB.QueryRow(query, values[0])
	}
	if err := row.Scan(&emailPlaceholder); err != nil {
		if err == sql.ErrNoRows {
			fmt.Println(err)
			return false, nil
		}
		fmt.Println(err)
		return false, fmt.Errorf("albumsById")
	}

	return true, nil
}
func removeLastAnd(query string) string {
	// Trim trailing whitespaces
	query = strings.TrimSpace(query)

	// Check if the query ends with "AND"
	if strings.HasSuffix(query, "AND") {
		// Find the last occurrence of "AND" and remove it
		lastAndIndex := strings.LastIndex(query, "AND")
		query = query[:lastAndIndex]
		// Trim trailing whitespaces again
		query = strings.TrimSpace(query)
	}

	return query
}
func (u *UserRepo) Get(email string) (model.User, error) {

	var user model.User
	row := u.DB.QueryRow("SELECT * FROM users WHERE email = ?", email)
	if err := row.Scan(&user.UserID, &user.Email, &user.FirstName, &user.LastName, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("user %s: no such user", email)
		}
	}
	return user, nil
}
func (u *UserRepo) GetByID(ID int64) (model.User, error) {
	user := model.User{}
	err := u.DB.QueryRow("SELECT first_name, last_name, email, password FROM users WHERE id = ?", ID).Scan(&user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		return user, err
	}
	return user, nil

}
func (u *UserRepo) UpdateByID() {

}
func (u *UserRepo) DeleteByID() {

}
