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

// PersonStore is the implementation of person store using postgres db.
type PersonStore struct {
	db *gorm.DB
}

func NewPersonStore(db *gorm.DB) *PersonStore {
	return &PersonStore{
		db: db,
	}
}

func (s *PersonStore) Get(ctx context.Context, uid string) (*models.Person, error) {
	person := models.Person{}
	uuid, err := uuid.Parse(uid)
	if err != nil {
		log.WithContext(ctx).WithField("person.uid", uid).Error(parseUUIDErr)
		return nil, errpkg.NewFormatError("person id", uid)
	}
	where := models.Person{Base: models.Base{UID: uuid}}
	result := s.db.Where(where).First(&person)

	if result.Error != nil {
		switch result.Error {
		case gorm.ErrRecordNotFound:
			log.WithContext(ctx).WithField("person.uid", uid).Warn("record not found")
			criteria := fmt.Sprintf("uid=%s", uid)
			return nil, errpkg.NewRecordNotFoundError("job", criteria)
		default:
			msg := fmt.Sprintf(recordRetrievalErr+": %v", result.Error)
			log.WithContext(ctx).WithField("person.uid", uid).Error(msg)
			return nil, errpkg.NewDatabaseError()
		}
	}
	return &person, nil
}

func (s *PersonStore) Update(ctx context.Context, uid string, updates models.Person, fields ...string) error {
	person := models.Person{}
	uuid, err := uuid.Parse(uid)
	if err != nil {
		log.WithContext(ctx).WithField("person.uid", uid).Error(parseUUIDErr)
		return errpkg.NewFormatError("person id", uid)
	}
	where := models.Person{Base: models.Base{UID: uuid}}
	clauses := clause.Returning{}
	result := s.db.Model(&person).Select(fields).Clauses(clauses).Where(where).Updates(updates)

	criteria := fmt.Sprintf("person.id=%s", uid)
	if result.Error != nil {
		msg := fmt.Sprintf(recordUpdateDBErr+": %v", result.Error)
		log.WithContext(ctx).WithField("search.criteria", criteria).Error(msg)
		return errpkg.NewDatabaseError()
	} else if result.RowsAffected == 0 {
		log.WithContext(ctx).WithField("search.criteria", criteria).Warn(recordNotFoundErr)
		return errpkg.NewRecordNotFoundError("person ", criteria)
	}

	log.WithContext(ctx).WithFields(log.Fields{"person.uid": uid}).Info(recordUpdateInfo)
	return nil
}

func (s *PersonStore) Create(ctx context.Context, person models.Person) (*models.Person, error) {
	result := s.db.Create(&person)
	if result.Error != nil {
		log.WithContext(ctx).Errorf("record insertion error: %v", result.Error)
		return nil, errpkg.NewDatabaseError()
	}

	msg := "new person record created"
	log.WithContext(ctx).WithField("person.uid", person.UID).Info(msg)
	return &person, nil
}
