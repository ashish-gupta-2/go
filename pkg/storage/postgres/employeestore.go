package postgres

import (
	"context"
	"fmt"

	errpkg "ashish.com/m/internal/errors"
	"ashish.com/m/pkg/models"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// EmployeeStore is the implementation of employee store using postgres db.
type EmployeeStore struct {
	db *gorm.DB
}

func NewEmployeeStore(db *gorm.DB) *EmployeeStore {
	return &EmployeeStore{
		db: db,
	}
}

var (
	parseUUIDErr       = "Unable to parse uuid from string"
	recordRetrievalErr = "record retrieval error"
	recordUpdateDBErr  = "record updation error"
	recordNotFoundErr  = "no record updated by update query"
	recordUpdateInfo   = "record updated"
)

func (s *EmployeeStore) Get(ctx context.Context, uid string) (*models.Employee, error) {
	employee := models.Employee{}
	uuid, err := uuid.Parse(uid)
	if err != nil {
		log.WithContext(ctx).WithField("employee.uid", uid).Error(parseUUIDErr)
		return nil, errpkg.NewFormatError("employee id", uid)
	}
	where := models.Employee{Base: models.Base{UID: uuid}}
	result := s.db.Where(where).First(&employee)

	if result.Error != nil {
		switch result.Error {
		case gorm.ErrRecordNotFound:
			log.WithContext(ctx).WithField("employee.uid", uid).Warn("record not found")
			criteria := fmt.Sprintf("uid=%s", uid)
			return nil, errpkg.NewRecordNotFoundError("job", criteria)
		default:
			msg := fmt.Sprintf(recordRetrievalErr+": %v", result.Error)
			log.WithContext(ctx).WithField("employee.uid", uid).Error(msg)
			return nil, errpkg.NewDatabaseError()
		}
	}
	return &employee, nil
}

func (s *EmployeeStore) Update(ctx context.Context, uid string, updates models.Employee, fields ...string) error {
	employee := models.Employee{}
	uuid, err := uuid.Parse(uid)
	if err != nil {
		log.WithContext(ctx).WithField("employee.uid", uid).Error(parseUUIDErr)
		return errpkg.NewFormatError("employee id", uid)
	}
	where := models.Employee{Base: models.Base{UID: uuid}}
	clauses := clause.Returning{}
	result := s.db.Model(&employee).Select(fields).Clauses(clauses).Where(where).Updates(updates)

	criteria := fmt.Sprintf("employee.id=%s", uid)
	if result.Error != nil {
		msg := fmt.Sprintf(recordUpdateDBErr+": %v", result.Error)
		log.WithContext(ctx).WithField("search.criteria", criteria).Error(msg)
		return errpkg.NewDatabaseError()
	} else if result.RowsAffected == 0 {
		log.WithContext(ctx).WithField("search.criteria", criteria).Warn(recordNotFoundErr)
		return errpkg.NewRecordNotFoundError("employee ", criteria)
	}

	log.WithContext(ctx).WithFields(log.Fields{"employee.uid": uid}).Info(recordUpdateInfo)
	return nil
}

func (s *EmployeeStore) Create(ctx context.Context, employee models.Employee) (*models.Employee, error) {
	result := s.db.Create(&employee)
	if result.Error != nil {
		log.WithContext(ctx).Errorf("record insertion error: %v", result.Error)
		return nil, errpkg.NewDatabaseError()
	}

	msg := "new employee record created"
	log.WithContext(ctx).WithField("employee.uid", employee.UID).Info(msg)
	return &employee, nil
}
