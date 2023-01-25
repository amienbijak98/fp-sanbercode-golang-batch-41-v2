package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/wisedevguy/fp-sanbercode-golang-batch-41-v2/models"
)

func InsertAppointment(db *sql.DB, appointment models.Appointment) error {
	appointment.CreatedBy = AdminLoggedIn

	now := time.Now()
	timeString := now.Format(time.RFC3339)

	query := `INSERT INTO appointment (appo_schedule, descriptions, patient_id, doctors_id, departement_id, created_by, created_at) 
				VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := db.Exec(query, appointment.AppoSchedule, appointment.Descriptions, appointment.PatientID, appointment.DoctorsID, appointment.DepartementID, appointment.CreatedBy, timeString)
	if err != nil {
		return err
	}
	return nil
}

func UpdateAppointment(db *sql.DB, appointment models.Appointment) error {
	if appointment.UpdatedBy == nil {
		return errors.New("updated_by field is required")
	}

	now := time.Now()
	timeString := now.Format(time.RFC3339)

	query := `UPDATE appointment SET appo_schedule = $1, descriptions = $2, patient_id = $3, doctors_id = $4, departement_id = $5, updated_by = $6, updated_at = $7 WHERE id = $8`

	_, err := db.Exec(query, appointment.AppoSchedule, appointment.Descriptions, appointment.PatientID, appointment.DoctorsID, appointment.DepartementID, *appointment.UpdatedBy, timeString, appointment.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAppointment(db *sql.DB, id int) error {
	query := "DELETE FROM appointment WHERE id = $1"

	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func GetAppointmentByID(db *sql.DB, id int) (models.Appointment, error) {
	var appointment models.Appointment
	var updatedBy sql.NullInt64

	query := "SELECT * FROM appointment WHERE id = $1"

	err := db.QueryRow(query, id).Scan(&appointment.ID, &appointment.AppoSchedule, &appointment.Descriptions, &appointment.PatientID, &appointment.DoctorsID, &appointment.DepartementID, &appointment.CreatedBy, &updatedBy, &appointment.CreatedAt, &appointment.UpdatedAt)
	if err != nil {
		return models.Appointment{}, err
	}

	if updatedBy.Valid {
		temp := int(updatedBy.Int64)
		appointment.UpdatedBy = &temp
	}

	return appointment, nil
}

func GetAllAppointments(db *sql.DB) (appointments []models.Appointment, err error) {
	sqlStatements := "SELECT * FROM appointment"

	rows, err := db.Query(sqlStatements)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var appointment models.Appointment
		var updatedBy sql.NullInt64

		err = rows.Scan(&appointment.ID, &appointment.AppoSchedule, &appointment.Descriptions, &appointment.PatientID, &appointment.DoctorsID, &appointment.DepartementID, &appointment.CreatedBy, &updatedBy, &appointment.CreatedAt, &appointment.UpdatedAt)
		if err != nil {
			return nil, err
		}
		if updatedBy.Valid {
			temp := int(updatedBy.Int64)
			appointment.UpdatedBy = &temp
		}
		appointments = append(appointments, appointment)
	}

	return appointments, nil
}
