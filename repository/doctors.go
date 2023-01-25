package repository

import (
	"database/sql"
	"time"

	"github.com/wisedevguy/fp-sanbercode-golang-batch-41-v2/models"
)

func InsertDoctor(db *sql.DB, doctor models.Doctor) error {
	doctor.CreatedBy = AdminLoggedIn

	now := time.Now()
	timeString := now.Format(time.RFC3339)

	query := `INSERT INTO doctors (name, mobile_number, qualifications, departement_id, created_by, created_at) 
				VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := db.Exec(query, doctor.Name, doctor.MobileNumber, doctor.Qualifications, doctor.DepartementID, doctor.CreatedBy, timeString)
	if err != nil {
		return err
	}

	return nil
}

func UpdateDoctor(db *sql.DB, doctor models.Doctor) error {
	doctor.UpdatedBy = new(int)
	*doctor.UpdatedBy = AdminLoggedIn

	now := time.Now()
	timeString := now.Format(time.RFC3339)

	query := `UPDATE doctors SET name = $1, mobile_number = $2, qualifications = $3, departement_id = $4, updated_by = $5, updated_at = $6 WHERE id = $7`

	_, err := db.Exec(query, doctor.Name, doctor.MobileNumber, doctor.Qualifications, doctor.DepartementID, doctor.UpdatedBy, timeString, doctor.ID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteDoctor deletes an existing doctor from the database
func DeleteDoctor(db *sql.DB, id int) error {
	query := `DELETE FROM doctors WHERE id = $1`

	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func GetDoctorByID(db *sql.DB, id int) (models.Doctor, error) {
	var doctor models.Doctor

	query := `SELECT * FROM doctors WHERE id = $1`

	err := db.QueryRow(query, id).Scan(&doctor.ID, &doctor.Name, &doctor.MobileNumber, &doctor.Qualifications, &doctor.DepartementID, &doctor.CreatedBy, &doctor.UpdatedBy, &doctor.CreatedAt, &doctor.UpdatedAt)
	if err != nil {
		return doctor, err
	}

	return doctor, nil
}

func GetAllDoctors(db *sql.DB) ([]models.Doctor, error) {
	var doctors []models.Doctor

	query := `SELECT * FROM doctors`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var doctor models.Doctor

		var updatedBy sql.NullInt64
		var updatedAt sql.NullTime

		err = rows.Scan(&doctor.ID, &doctor.Name, &doctor.MobileNumber, &doctor.Qualifications, &doctor.DepartementID, &doctor.CreatedBy, &updatedBy, &doctor.CreatedAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		if updatedBy.Valid {
			temp := int(updatedBy.Int64)
			doctor.UpdatedBy = &temp
		}
		if updatedAt.Valid {
			doctor.UpdatedAt = &updatedAt.Time
		}

		doctors = append(doctors, doctor)
	}

	return doctors, nil
}
