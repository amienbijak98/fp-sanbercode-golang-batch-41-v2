package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/wisedevguy/fp-sanbercode-golang-batch-41-v2/middleware"
	"github.com/wisedevguy/fp-sanbercode-golang-batch-41-v2/models"
	"golang.org/x/crypto/bcrypt"
)

var AdminLoggedIn int

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Register(db *sql.DB, admin models.UserAdmin) (err error) {
	sqlStatement := "INSERT INTO user_admin (name, username, password, role, created_at) VALUES ($1, $2, $3, $4, $5)"

	createdAt := time.Now()
	hash, _ := HashPassword(admin.Password)
	errs := db.QueryRow(sqlStatement, admin.Name, admin.UserName, hash, admin.Role, createdAt)
	return errs.Err()
}

func Login(db *sql.DB, admin models.UserAdmin) (interface{}, error) {
	sqlStatement := "SELECT id, username, password FROM user_admin WHERE username = $1"
	row := db.QueryRow(sqlStatement, admin.UserName)

	storedUser := struct {
		ID       int
		Username string
		Password string
	}{}

	err := row.Scan(&storedUser.ID, &storedUser.Username, &storedUser.Password)
	if err != nil {
		return nil, err
	}

	match := CheckPasswordHash(admin.Password, storedUser.Password)
	if !match {
		return nil, errors.New("password is incorrect")
	}

	sqlStatement = "DELETE FROM user_credentials WHERE admin_id = $1"
	db.QueryRow(sqlStatement, storedUser.ID)

	uuid := uuid.New()

	sqlStatement = "INSERT INTO user_credentials (admin_id, uuid) VALUES ($1, $2) RETURNING admin_id"
	insertRow := db.QueryRow(sqlStatement, storedUser.ID, uuid)
	if insertRow.Err() != nil {
		return nil, errors.New("failed to insert user credentials")
	}

	AdminLoggedIn = storedUser.ID

	token, err := middleware.GenerateToken(storedUser.ID, uuid.String())
	if err != nil {
		return nil, errors.New("failed to login")
	}

	return token, nil
}

func Logout(db *sql.DB, uuid string) error {
	sqlStatement := "DELETE FROM user_credentials WHERE uuid = $1"
	err := db.QueryRow(sqlStatement, uuid)

	AdminLoggedIn = 0

	return err.Err()
}

func GetAllAdmin(db *sql.DB) (admins []models.UserAdmin, err error) {
	sql := "SELECT * FROM user_admin"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var admin = models.UserAdmin{}

		err = rows.Scan(&admin.ID, &admin.UserName, &admin.Password, &admin.Name, &admin.Role, &admin.CreatedAt, &admin.UpdatedAt)
		if err != nil {
			panic(err)
		}

		admins = append(admins, admin)

	}
	return
}
