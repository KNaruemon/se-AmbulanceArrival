package entity

import (
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model

	Name        string
	Tel         string
	Email       string
	Assessments []Assessment `gorm:"foreignKey:PatientID"`
}

type AmbulanceOnDuty struct {
	gorm.Model

	Code       string
	OnDutyDate time.Time
	Passenger  uint

	AmbulanceID *uint
	Ambulance   Ambulance

	RecorderID *uint
	Recorder   Employee

	IncidentID *uint
	Incident   Incident
}

type AmbulanceArrival struct {
	gorm.Model
	Number_of_people int
	Distance         float32
	DateTime         time.Time

	RecorderID *uint
	Recorder   Employee

	PatientID *uint
	Patient   Patient

	AmbulanceOnDutyID *uint
	AmbulanceOnDuty   AmbulanceOnDuty
}

type Assessment struct {
	gorm.Model

	Symptom      string
	SymptomLevel uint
	Datetime     time.Time

	PatientID *uint
	Patient   Patient `gorm:"references:id"`

	RecorderID *uint
	Recorder   Employee `gorm:"references:id"`

	IncidentID *uint
	Incident   Incident `gorm:"references:id"`
}
type Ambulance struct {
	gorm.Model
	CarNumber    int
	Registration string
	DateTime     time.Time

	StatusID *uint
	Status   Status `gorm:"references:id"`

	AmbulanceTypeID *uint
	AmbulanceType   AmbulanceType `gorm:"references:id"`

	EmployeeID *uint
	Employee   Employee `gorm:"references:id"`
}
type Status struct {
	gorm.Model
	Detail  string
	Records []Ambulance `gorm:"foreignKey:StatusID"`
}
type AmbulanceType struct {
	gorm.Model
	Name    string
	Detail  string
	Records []Ambulance `gorm:"foreignKey:AmbulanceTypeID"`
}
type Employee struct {
	gorm.Model
	Name        string
	Tel         string
	Email       string `gorm:"uniqueIndex"`
	Password    string
	Records     []Ambulance  `gorm:"foreignKey:EmployeeID"`
	Assessments []Assessment `gorm:"foreignKey:RecorderID"`
}

type AmbulanceCheck struct {
	gorm.Model
	DateTime time.Time

	Cleanliness    float32
	Equipmentcheck bool
	Carcheck       bool
	Note           string

	//AmbulanceID ทำหน้าที่เป็น FK
	AmbulanceID *uint
	Ambulance   Ambulance `gorm:"references:id"`

	//RecorderID ทำหน้าที่เป็น FK
	RecorderID *uint
	Recorder   Employee `gorm:"references:id"`
}

type Illness struct {
	gorm.Model
	Value    string
	Incident []Incident `gorm:"foreignKey:IllnessID"`
}

type Incident struct {
	gorm.Model
	Title         string
	Informer      string
	Numberpatient int
	Location      string
	Datetime      time.Time
	EmployeeID    *uint
	Employee      Employee
	IllnessID     *uint
	Illness       Illness
	UrgencyID     *uint
	Urgency       Urgency
	Assessments   []Assessment `gorm:"foreignKey:IncidentID"`
}

type Urgency struct {
	gorm.Model
	Value    string
	Incident []Incident `gorm:"foreignKey:UrgencyID"`
}
