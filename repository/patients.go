package repository

import (
	"database/sql"
	"time"

	"github.com/wisedevguy/fp-sanbercode-golang-batch-41-v2/models"
)

func GetAllPatients(db *sql.DB) (patients []models.Patient, err error) {
	sql := "SELECT * FROM patients"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var patient = models.Patient{}

		//var timeString string

		err = rows.Scan(&patient.ID, &patient.Name, &patient.Gender, &patient.BirthDate, &patient.MobileNumber, &patient.MedicalHistory, &patient.CreatedBy, &patient.UpdatedBy, &patient.CreatedAt, &patient.UpdatedAt)

		if err != nil {
			panic(err)
		}

		patients = append(patients, patient)

	}
	return
}

func GetPatientByID(db *sql.DB, id int) (models.Patient, error) {
	var patient models.Patient

	query := `SELECT id, name, gender, birth_date, mobile_number, medical_history, created_by, updated_by, created_at, updated_at FROM patients WHERE id = $1`

	err := db.QueryRow(query, id).Scan(&patient.ID, &patient.Name, &patient.Gender, &patient.BirthDate, &patient.MobileNumber, &patient.MedicalHistory, &patient.CreatedBy, &patient.UpdatedBy, &patient.CreatedAt, &patient.UpdatedAt)

	if err != nil {
		return models.Patient{}, err
	}

	return patient, nil
}

func InsertPatient(db *sql.DB, patient models.Patient) error {
	patient.CreatedBy = AdminLoggedIn

	now := time.Now()
	timeString := now.Format(time.RFC3339)

	query := `INSERT INTO patients (name, gender, birth_date, mobile_number, medical_history, created_by, created_at) 
				VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := db.Exec(query, patient.Name, patient.Gender, patient.BirthDate, patient.MobileNumber, patient.MedicalHistory, patient.CreatedBy, timeString)
	if err != nil {
		return err
	}
	return nil
}

func UpdatePatient(db *sql.DB, patient models.Patient) error {
	patient.UpdatedBy = new(int)
	*patient.UpdatedBy = AdminLoggedIn

	now := time.Now()
	// Format time into ISO8601 format
	timeString := now.Format(time.RFC3339)

	query := `UPDATE patients SET name = $2, gender = $3, birth_date = $4, mobile_number = $5, medical_history = $6, updated_by = $7, updated_at = $8 WHERE id = $1`

	_, err := db.Exec(query, patient.ID, patient.Name, patient.Gender, patient.BirthDate, patient.MobileNumber, patient.MedicalHistory, patient.UpdatedBy, timeString)

	if err != nil {
		return err
	}

	return nil
}

func DeletePatient(db *sql.DB, id int) error {
	query := `DELETE FROM patients WHERE id = $1`
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
