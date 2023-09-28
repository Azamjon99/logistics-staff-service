package services

import (
	"context"

	"github.com/Azamjon99/logistics-staff-service/src/domain/models"
	"github.com/Azamjon99/logistics-staff-service/src/domain/repository"
)

type DriverService interface {
	RegisterDriver(ctx context.Context, phoneNumber, password string) (string, error)
	ChangeDriverPassword(ctx context.Context, driverID, currentPassword, newPassword string) error
	LoginDriver(ctx context.Context, phoneNumber, password string) (string, error)
	ConfirmSMSCode(ctx context.Context, driverID, smsCode string) (*models.DriverProfile, error)
	GetDriverProfile(ctx context.Context, driverID string) (*models.DriverProfile, error)
	UpdateDriverProfile(ctx context.Context, driverID, name, imageUrl string) (*models.DriverProfile, error)
	}

type driverSvcImpl struct {
	driverRepo repository.DriverRepository
}

func NewDriverSvc(driverRepo repository.DriverRepository) DriverService {
	return &driverSvcImpl{
		driverRepo: driverRepo,
	}
}

