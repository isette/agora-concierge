package models

import (
	"testing"
	"time"

	"github.com/isette/appointments-concierge/handlers"
	"github.com/isette/appointments-concierge/models"
	"github.com/isette/appointments-concierge/repositories"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&models.Appointment{}); err != nil {
		return nil, err
	}

	return db, nil
}

func TestAppointmentRepository_FindAll(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Erro ao configurar o banco de dados de teste: %v", err)
	}

	repo := repositories.NewAppointmentRepository(db)

	db.Create(&models.Appointment{
		InspectedID:    1,
		InspectorID:    2,
		InspectionDate: time.Now(),
		InspectionTime: handlers.PointerToTime(time.Now()),
		EquipamentType: new(models.EquipamentType),
	})
	db.Create(&models.Appointment{
		InspectedID:    3,
		InspectorID:    4,
		InspectionDate: time.Now(),
		InspectionTime: handlers.PointerToTime(time.Now()),
		EquipamentType: new(models.EquipamentType),
	})

	appointments, err := repo.FindAll()

	assert.NoError(t, err)
	assert.Len(t, appointments, 2)
	assert.Equal(t, uint(1), appointments[0].InspectedID)
	assert.Equal(t, uint(3), appointments[1].InspectedID)
}

func TestAppointmentRepository_Create(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Erro ao configurar o banco de dados de teste: %v", err)
	}

	repo := repositories.NewAppointmentRepository(db)

	date := models.DateFull{Year: 2024, Month: 11, Day: 26}
	appointment := models.AppointmentWithDate{
		Date: date,
		Appointment: models.Appointment{
			InspectedID:    1,
			InspectorID:    2,
			InspectionDate: time.Now(),
			InspectionTime: handlers.PointerToTime(time.Now()),
			EquipamentType: new(models.EquipamentType),
		},
	}

	appointments, err := repo.Create(appointment)

	assert.NoError(t, err)
	assert.Len(t, appointments, 1)
	assert.Equal(t, uint(1), appointments[0].InspectedID)
}

func TestAppointmentRepository_WithTransaction(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Erro ao configurar o banco de dados de teste: %v", err)
	}

	repo := repositories.NewAppointmentRepository(db)
	tx := db.Begin()

	repoWithTx := repo.WithTransaction(tx)

	assert.IsType(t, &handlers.AppointmentsRepo{}, repoWithTx)

	tx.Commit()
}
