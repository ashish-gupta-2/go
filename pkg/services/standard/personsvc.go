package standard

import (
	"context"

	errpkg "ashish.com/m/internal/errors"
	"ashish.com/m/pkg/models"
	storage "ashish.com/m/pkg/storage"
	log "github.com/sirupsen/logrus"
)

// PersonService is the standard implementation of person service.
type PersonService struct {
	personStore storage.PersonStore
}

var (
	personIDEmptyErr  = "person id is empty"
	personNotFoundErr = "person not found"
)

// NewPersonService creates a new instance of person service.
func NewPersonService(s storage.PersonStore) *PersonService {
	return &PersonService{
		personStore: s,
	}
}

// GetPerson returns the person associated with the uid.
func (s *PersonService) GetPerson(ctx context.Context, uid string) (*models.Person, error) {
	fields := log.Fields{"person.id": uid}

	// check uid
	if len(uid) == 0 {
		log.WithContext(ctx).WithFields(fields).Error(personIDEmptyErr)
		return nil, errpkg.NewEmptyError("person id")
	}

	// get person from data store
	person, err := s.personStore.Get(ctx, uid)
	if err != nil {
		if _, ok := err.(*errpkg.RecordNotFoundError); ok {
			return nil, errpkg.NewResourceNotFoundError("person", uid)
		}
		return nil, err
	}
	return person, nil
}

// UpdatePersonOutput updates the person output with the uid.
func (s *PersonService) UpdatePerson(ctx context.Context, person *models.Person) (*models.Person, error) {
	fields := log.Fields{"person.id": person.UID}

	// get person from data store
	person, err := s.GetPerson(ctx, person.UID.String())
	if err != nil {
		return nil, err
	}

	err = s.personStore.Update(ctx, person.Base.UID.String(), *person)
	if err != nil {
		if _, ok := err.(*errpkg.RecordNotFoundError); ok {
			log.WithContext(ctx).WithFields(fields).Error(personNotFoundErr)
			return nil, errpkg.NewResourceNotFoundError("person", person.UID.String())
		}
		return nil, err
	}
	return person, nil
}

// UpdatePersonError updates the person error with the uid.
func (s *PersonService) CreatePerson(ctx context.Context, person *models.Person) (*models.Person, error) {
	emp, err := s.personStore.Create(ctx, *person)
	if err != nil {
		return nil, err
	}
	return emp, nil
}
