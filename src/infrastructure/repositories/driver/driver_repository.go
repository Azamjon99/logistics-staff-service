package driver

import (
	"context"
	"github.com/Azamjon99/logistics-staff-service/src/domain/models"
	repositories "github.com/Azamjon99/logistics-staff-service/src/domain/repository"

	"gorm.io/gorm"
)

const(

  driverTable = "support.drivers"
)

type driverrepoImpl struct {
	db *gorm.DB
}
 
func NewDriverRepository(db *gorm.DB) repositories.DriverRepository{
	return &driverrepoImpl{
		db:db,
	}
}

func (r *driverrepoImpl)Savedriver(ctx context.Context, driver *models.Driver) error{
	result := r.db.WithContext(ctx).Table(driverTable).Create(&driver)

	if result.Error != nil{
		return result.Error
	}

	return nil;
}

func (r *driverrepoImpl)UpdateDriver(ctx context.Context, driver *models.Driver) error{
	result := r.db.WithContext(ctx).Table(driverTable).Save(&driver)

	if result.Error != nil{
		return result.Error
	}

	return nil;
}



func (r *driverrepoImpl) DeleteDriver(ctx context.Context, driverID string) error {
	result := r.db.WithContext(ctx).Table(driverTable).Where("driver_id = ?", driverID).Delete(&models.Driver{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *driverrepoImpl)ListDriverByDriver(ctx context.Context, driverID string)([]*models.Driver, error){
	var driver []*models.Driver
	result := r.db.WithContext(ctx).Table(driverTable).First(&driver, "order_id = ?", driverID)

	if result.Error != nil {
		return nil, result.Error
	}

	return driver, nil;
}

func (r *driverrepoImpl)GetDriver(ctx context.Context, driverID string)(*models.Driver, error){
	var driver *models.Driver
	result := r.db.WithContext(ctx).Table(driverTable).First(&driver, "driver_id = ?", driverID)

	if result.Error != nil {
		return nil, result.Error
	}

	return driver, nil;
}