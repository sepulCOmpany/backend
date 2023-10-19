package db

import "github.com/sepulCOmpany/backend/internal/models"

func (db *Db) GetAllGrimziks() ([]models.UserWithoutPassword, error) {
	const (
		grimzikRoleID = 2
		query         = `
			SELECT username, role_id
			FROM registred_users
				WHERE role_id=$1
`
	)

	var grimziks []models.UserWithoutPassword
	err := db.db.Select(&grimziks, query, grimzikRoleID)
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

func (db *Db) SetDeliveryState(sepulca models.Sepulca) error {
	const (
		query = `
			UPDATE sepulcas 
			    SET delivery_state_id=:delivery_state_id where id=:id
`
	)

	_, err := db.db.NamedExec(query, sepulca)
	if err != nil {
		return nil
	}

	return nil
}

func (db *Db) GetAllSepulcas() ([]models.Sepulca, error) {
	const (
		query = `
			SELECT sepulcas.id, size_id, shmurdik_id, grimzik_id, property_id, is_vaccinated, is_rubbered, delivery_state_id, r2.id as grimzik_id, r.id as shmurdik_id
			    FROM sepulcas
				INNER JOIN roles r on r.id = sepulcas.shmurdik_id
				INNER JOIN roles r2 on r2.id = sepulcas.grimzik_id
			
`
	)
	var dbData []models.Sepulca
	err := db.db.Select(&dbData, query)
	if err != nil {
		return nil, err
	}

	return dbData, nil
}
