package repository

import (
	"database/sql"

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

		err = rows.Scan(&patient.ID, &patient.Name, &patient.Gender, &patient.BirthDate, &patient.MobileNumber, &patient.MedicalHistory, &patient.CreatedBy, &patient.UpdatedBy, &patient.CreatedAt, &patient.UpdatedAt)
		if err != nil {
			panic(err)
		}

		patients = append(patients, patient)

	}
	return
}
