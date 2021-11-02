package entity

import (
	"time"

	"gorm.io/gorm"
)

type Nurse struct {
	gorm.Model
	Pid        string `gorm:"unqueIndex"`
	Nurse_name string
	Password   string
	Tel        string    `gorm:"unqueIndex"`
	Patient    []Patient `gorm:"foreignKey:NurseID"`
}

type Medicine struct {
	gorm.Model
	Med_name  string `gorm:"unqueIndex"`
	Type      string
	Med_price int
	Patients  []Patient `gorm:"foreignKey:MedicineID"`
}

type Disease struct {
	gorm.Model
	Disease_name string    `gorm:"unqueIndex"`
	Patients     []Patient `gorm:"foreignKey:DiseaseID"`
}

type Patient struct {
	gorm.Model

	Identification string `gorm:"unqueIndex"`
	Patient_name   string
	MedicineID     *uint
	Medicine       Medicine
	DiseaseID      *uint
	Disease        Disease
	NurseID        *uint
	Nurse          Nurse
	Date           time.Time
}
