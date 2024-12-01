package handlers

import (
	"time"

	"github.com/isette/appointments-concierge/models"
	"github.com/isette/appointments-concierge/repositories"
	"gorm.io/gorm"
)

type AppointmentsRepo struct {
	repo repositories.AppointmentRepository
	db   *gorm.DB
}

func (appointmentRepository *AppointmentsRepo) GetAvailableAppointments(day models.Date, status string) ([]models.Appointment, error) {
	return appointmentRepository.repo.FindAll()
}

func (appointmentRepository *AppointmentsRepo) CreateAppointments(appointment models.AppointmentWithDate) ([]models.Appointment, error) {
	date := models.DateFull{Year: 2024, Month: 11, Day: 26}

	fields := models.AppointmentWithDate{
		Date: date,
		Appointment: models.Appointment{
			InspectionDate: time.Now(),
			InspectionTime: PointerToTime(time.Now()),
			InspectedID:    appointment.Appointment.InspectedID,
			InspectorID:    appointment.Appointment.InspectorID,
			EquipamentType: appointment.Appointment.EquipamentType,
		},
	}

	createdAppointments, err := appointmentRepository.repo.Create(fields)

	if err != nil {
		return nil, err
	}

	return createdAppointments, nil
}

func (appointmentRepository *AppointmentsRepo) CreateAppointmentsWithTransaction(day models.Date, appointment models.Appointment) error {
	return appointmentRepository.db.Transaction(func(tx *gorm.DB) error {
		repoWithTx := appointmentRepository.repo.WithTransaction(tx)

		date := models.DateFull{
			Year:  day.Year(),
			Month: int(day.Month()),
			Day:   day.Day(),
		}

		fields := models.AppointmentWithDate{
			Date: date,
			Appointment: models.Appointment{
				InspectedID:         1,
				InspectorID:         2,
				InspectionDate:      time.Now(),
				InspectionTime:      PointerToTime(time.Now()),
				EquipamentType:      new(models.EquipamentType),
				InspectionStartTime: PointerToTime(time.Now()), // Usando função para obter o ponteiro
				InspectionEndTime:   nil,                       // Se não tiver valor, pode ser nil
			},
		}

		if _, err := repoWithTx.Create(fields); err != nil {
			return err
		}

		return nil
	})
}

func PointerToTime(t time.Time) *time.Time {
	return &t
}
