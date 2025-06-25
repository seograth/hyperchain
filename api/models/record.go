package models

type MedicalRecord struct {
	ID        string `json:"id"`
	Patient   string `json:"patient"`
	Doctor    string `json:"doctor"`
	Diagnosis string `json:"diagnosis"`
	Timestamp string `json:"timestamp"`
}
