package db

import (
	"fmt"

	"github.com/sepulCOmpany/backend/internal/models"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func (db *Db) Register(user models.User) (*models.UserWithoutPassword, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("cannot get hash of password, err: %w", err)
	}

	user.Password = string(hashPass)

	const query = `
		INSERT INTO 
		    registred_users (username, password, role_id)
			VALUES (:username, :password, :role_id)
	`
	const queryGetLast = `
		SELECT id, role_id, username
		FROM registred_users
				ORDER BY id DESC LIMIT 1
			
	`
	var dbData models.UserWithoutPassword
	tx := db.db.MustBegin()
	_, err1 := db.db.NamedExec(query, user)
	if err1 != nil {
		logrus.Debug(err)
	}
	err2 := db.db.Get(&dbData, queryGetLast)
	if err2 != nil {
		logrus.Debug(err)
	}
	tx.Commit()
	return &dbData, nil
}

func (db *Db) Login(user models.User) (*models.UserWithoutPassword, error) {
	const query = `
		SELECT role_id, username, password
		FROM registred_users
				WHERE username=$1
			
	`
	username := user.Username
	var dbData models.User
	err := db.db.Get(&dbData, query, username)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbData.Password), []byte(user.Password)); err != nil {
		return nil, fmt.Errorf("incorrect password from db, err: %w", err)
	}

	return &models.UserWithoutPassword{
		UserName: dbData.Username,
		RoleID:   dbData.RoleID,
	}, nil
}
