package db

import (
	"fmt"

	"github.com/sepulCOmpany/backend/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (db *Db) CheckRole() {

}

func (db *Db) Register(user models.User) error {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("cannot get hash of password, err: %w", err)
	}

	user.Password = string(hashPass)

	const query = `
		INSERT INTO 
		    registred_users (username, password, role_id)
			VALUES (:username, :password, :role_id)
	`

	_, err = db.db.NamedExec(query, user)
	return err
}

func (db *Db) Login(user models.User) (*models.UserWithoutPassword, error) {
	const query = `
		SELECT (username, password, role_id) 
		FROM registred_users
				WHERE username=$1
			
	`

	var dbData models.User
	err := db.db.Get(&dbData, query, user.UserName)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbData.Password), []byte(user.Password)); err != nil {
		return nil, fmt.Errorf("incorrect password from db, err: %w", err)
	}

	return &models.UserWithoutPassword{
		UserName: dbData.UserName,
		RoleID:   dbData.RoleID,
	}, nil
}
