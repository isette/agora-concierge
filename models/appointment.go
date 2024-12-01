package models

import "time"

type DateFull struct {
	Year  int
	Month int
	Day   int
}

type EquipamentType string

const (
	HORSE EquipamentType = "horse"
	TRUCK EquipamentType = "truck"
)

func (d DateFull) ToTime() time.Time {
	return time.Date(d.Year, time.Month(d.Month), d.Day, 0, 0, 0, 0, time.UTC)
}

type Appointment struct {
	ID                  *uint           `gorm:"primaryKey;autoIncrement"`
	InspectedID         uint            `gorm:"column:inspected_id;index"`
	InspectorID         uint            `gorm:"column:inspector_id;index"`
	InspectionDate      time.Time       `gorm:"type:date"`
	InspectionTime      *time.Time      `gorm:"type:time"`
	InspectionStartTime *time.Time      `gorm:"column:start_time;type:timestamp"`
	InspectionEndTime   *time.Time      `gorm:"column:end_time;type:timestamp"`
	EquipamentType      *EquipamentType `gorm:"column:equipament_type;type:varchar(255)"`
}

type AppointmentWithDate struct {
	Date        DateFull
	Appointment Appointment
}
