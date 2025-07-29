package models

import "time"

type Appointment struct {
	DoctorName  string    `json:"doctor_name"`
	PatientName string    `json:"patient_name"`
	Diagnosis   string    `json:"diagnosis"`
	Date        time.Time `json:"date"`
	Type        string    `json:"type"`
	Platform    string    `json:"platform"`
}
