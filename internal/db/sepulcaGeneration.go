package db

import "github.com/sepulCOmpany/backend/internal/models"

func (db *Db) GetAllGrimziks() ([]models.UserWithoutPassword, error) {
	const (
		query = `
			SELECT username, role_id
			FROM registred_users
				WHERE role_id=2
`
	)

	var grimziks []models.UserWithoutPassword
	err := db.db.Get(&grimziks, query)
	if err != nil {
		return nil, err
	}

	return grimziks, nil
}

func (db *Db) CreateSepulca(sepulca models.Sepulca) error {
	const (
		query = `
			INSERT INTO sepulcas (size_id, shmurdik_id, grimzik_id, property_id)
			VALUES (:size_id, :shmurdik_id, :grimzik_id, :property_id)
`
	)

	_, err := db.db.NamedExec(query, sepulca)
	if err != nil {
		return nil
	}

	return nil
}

func (db *Db) VaccinateSepulca(sepulca models.Sepulca) error {
	const (
		query = `
			UPDATE sepulcas 
			    SET is_vaccinated=TRUE where id=:id
`
	)

	_, err := db.db.NamedExec(query, sepulca)
	if err != nil {
		return nil
	}

	return nil
}

func (db *Db) RubberSepulca(sepulca models.Sepulca) error {
	const (
		query = `
			UPDATE sepulcas 
			    SET is_rubbered=TRUE where id=:id
`
	)

	_, err := db.db.NamedExec(query, sepulca)
	if err != nil {
		return nil
	}

	return nil
}
