package repositories

import (
	"github.com/isette/appointments-concierge/models"
	"gorm.io/gorm"
)

type IAppointmentRepository interface {
	FindAll() ([]models.Appointment, error)
	Create(fields models.AppointmentWithDate) ([]models.Appointment, error)
	WithTransaction(tx *gorm.DB) AppointmentRepository
}

type AppointmentRepository struct {
	BaseRepository
}

func NewAppointmentRepository(db *gorm.DB) *AppointmentRepository {
	return &AppointmentRepository{
		BaseRepository: BaseRepository{db: db},
	}
}

func (r *AppointmentRepository) FindAll() ([]models.Appointment, error) {
	var appointments []models.Appointment
	err := r.db.Find(&appointments).Error
	return appointments, err
}

func (r *AppointmentRepository) Create(fields models.AppointmentWithDate) ([]models.Appointment, error) {
	fields.Appointment.InspectionDate = fields.Date.ToTime()

	if err := r.db.Create(&fields.Appointment).Error; err != nil {
		return nil, err
	}

	var appointments []models.Appointment
	r.db.Where("inspection_date = ?", fields.Date.ToTime()).Find(&appointments)

	return appointments, nil
}

func (r *AppointmentRepository) WithTransaction(tx *gorm.DB) *AppointmentRepository {
	return &AppointmentRepository{BaseRepository: *r.BaseRepository.WithTransaction(tx)}
}
