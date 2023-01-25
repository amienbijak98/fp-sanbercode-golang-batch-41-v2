package models

import "time"

type Patient struct {
	ID             int        `json:"id"`
	Name           string     `json:"name"`
	Gender         string     `json:"gender"`
	BirthDate      time.Time  `json:"birth_date"`
	MobileNumber   string     `json:"mobile_number"`
	MedicalHistory string     `json:"medical_history"`
	CreatedBy      int        `json:"created_by"`
	UpdatedBy      *int       `json:"updated_by,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
}

type Doctor struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	MobileNumber   string    `json:"mobile_number"`
	Qualifications string    `json:"qualifications"`
	DepartementID  int       `json:"departement_id"`
	CreatedBy      int       `json:"created_by"`
	UpdatedBy      int       `json:"updated_by,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}

type Departement struct {
	ID              int       `json:"id"`
	DepartementName string    `json:"departement_name"`
	Location        string    `json:"location"`
	CreatedBy       int       `json:"created_by"`
	UpdatedBy       int       `json:"updated_by,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
}

type Appointment struct {
	ID            int       `json:"id"`
	AppoSchedule  time.Time `json:"appo_schedule"`
	Descriptions  string    `json:"descriptions"`
	PatientID     int       `json:"patient_id"`
	DoctorsID     int       `json:"doctors_id"`
	DepartementID int       `json:"departement_id"`
	CreatedBy     int       `json:"created_by"`
	UpdatedBy     int       `json:"updated_by,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
}

type UserAdmin struct {
	ID        int       `json:"id"`
	UserName  string    `json:"username"`
	Password  string    `json:"password"`
	Name      string    `json:"name,omitempty"`
	Role      string    `json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type UserCredentials struct {
	ID      int    `json:"id"`
	AdminID int    `json:"admin_id"`
	UUID    string `json:"uuid"`
}
