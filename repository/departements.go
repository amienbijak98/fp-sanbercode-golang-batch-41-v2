package repository

import (
	"database/sql"
	"time"

	"github.com/wisedevguy/fp-sanbercode-golang-batch-41-v2/models"
)

func InsertDepartement(db *sql.DB, departement models.Departement) error {
	departement.CreatedBy = AdminLoggedIn

	now := time.Now()

	query := `INSERT INTO departement (departement_name, location, created_by, created_at) 
				VALUES ($1, $2, $3, $4)`

	_, err := db.Exec(query, departement.DepartementName, departement.Location, departement.CreatedBy, now)
	if err != nil {
		return err
	}
	return nil
}

func UpdateDepartement(db *sql.DB, departement models.Departement) error {
	now := time.Now()

	query := `UPDATE departement SET departement_name = $1, location = $2, updated_by = $3, updated_at = $4 WHERE id = $5`

	_, err := db.Exec(query, departement.DepartementName, departement.Location, departement.UpdatedBy, now, departement.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteDepartement(db *sql.DB, id int) error {
	query := `DELETE FROM departement WHERE id = $1`

	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func GetDepartementByID(db *sql.DB, id int) (departement models.Departement, err error) {
	query := `SELECT * FROM departement WHERE id = $1`

	err = db.QueryRow(query, id).Scan(&departement.ID, &departement.DepartementName, &departement.Location, &departement.CreatedBy, &departement.UpdatedBy, &departement.CreatedAt, &departement.UpdatedAt)
	if err != nil {
		return models.Departement{}, err
	}
	return
}

func GetAllDepartements(db *sql.DB) (departements []models.Departement, err error) {
	sqlStatements := "SELECT * FROM departement"

	rows, err := db.Query(sqlStatements)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var departement models.Departement

		var updatedBy sql.NullInt64
		var updatedAt sql.NullTime

		err = rows.Scan(&departement.ID, &departement.DepartementName, &departement.Location, &departement.CreatedBy, &updatedBy, &departement.CreatedAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		if updatedBy.Valid {
			temp := int(updatedBy.Int64)
			departement.UpdatedBy = &temp
		}
		if updatedAt.Valid {
			departement.UpdatedAt = &updatedAt.Time
		}
		departements = append(departements, departement)
	}
	return departements, nil
}
