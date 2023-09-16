package v1

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	errpkg "ashish.com/m/internal/errors"
	httppkg "ashish.com/m/internal/http"
	"ashish.com/m/internal/utils"
	"ashish.com/m/pkg/models"
	"ashish.com/m/pkg/services"
)

// PersonHTTPService is the http service implementation of person service.
type PersonHTTPService struct {
	personSvc services.PersonService
}

// NewPersonHTTPService creates a new instance of asset http service.
func NewPersonHTTPService(personSvc services.PersonService) *PersonHTTPService {
	return &PersonHTTPService{
		personSvc: personSvc,
	}
}

// GetPerson returns the person from person store.
func (s *PersonHTTPService) GetPerson(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) {
	personID := params["id"]

	// fetch person
	person, err := s.personSvc.GetPerson(ctx, personID)
	if err != nil {
		w.Header().Set(httppkg.HeaderContentType, httppkg.ContentTypeJSON)
		status := utils.MapToGRPCStatus(errpkg.NewResourceNotFoundError("person", personID))
		w.WriteHeader(httppkg.GetStatusCode(status.Code()))
		_ = json.NewEncoder(w).Encode(status.Proto())
		return
	}

	// return http response
	w.Header().Set(httppkg.HeaderContentType, httppkg.ContentTypeJSON)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(person)
}

// UpdatePersonOutput updates the person output in person store.
func (s *PersonHTTPService) UpdatePerson(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) {
	content, _ := ioutil.ReadAll(r.Body)
	defer func() { _ = r.Body.Close() }()

	var person *models.Person
	if len(content) > 0 {
		err := json.Unmarshal(content, person)
		if err != nil {
			w.Header().Set(httppkg.HeaderContentType, httppkg.ContentTypeJSON)
			status := utils.MapToGRPCStatus(errpkg.NewUnMarshalError(string(content)))
			w.WriteHeader(httppkg.GetStatusCode(status.Code()))
			_ = json.NewEncoder(w).Encode(status.Proto())
			return
		}
	}

	// fetch person
	_, err := s.personSvc.UpdatePerson(ctx, person)
	if err != nil {
		w.Header().Set(httppkg.HeaderContentType, httppkg.ContentTypeJSON)
		status := utils.MapToGRPCStatus(errpkg.NewResourceNotFoundError("person", person.UID.String()))
		w.WriteHeader(httppkg.GetStatusCode(status.Code()))
		_ = json.NewEncoder(w).Encode(status.Proto())
		return
	}

	// return http response
	w.Header().Set(httppkg.HeaderContentType, httppkg.ContentTypeJSON)
	w.WriteHeader(http.StatusOK)
}

// UpdatePersonOutput updates the person output in person store.
func (s *PersonHTTPService) CreatePerson(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) {
	content, _ := ioutil.ReadAll(r.Body)
	defer func() { _ = r.Body.Close() }()

	var person *models.Person
	if len(content) > 0 {
		err := json.Unmarshal(content, person)
		if err != nil {
			w.Header().Set(httppkg.HeaderContentType, httppkg.ContentTypeJSON)
			status := utils.MapToGRPCStatus(errpkg.NewUnMarshalError(string(content)))
			w.WriteHeader(httppkg.GetStatusCode(status.Code()))
			_ = json.NewEncoder(w).Encode(status.Proto())
			return
		}
	}

	// fetch person
	_, err := s.personSvc.CreatePerson(ctx, person)
	if err != nil {
		w.Header().Set(httppkg.HeaderContentType, httppkg.ContentTypeJSON)
		status := utils.MapToGRPCStatus(errpkg.NewDatabaseError())
		w.WriteHeader(httppkg.GetStatusCode(status.Code()))
		_ = json.NewEncoder(w).Encode(status.Proto())
		return
	}

	// return http response
	w.Header().Set(httppkg.HeaderContentType, httppkg.ContentTypeJSON)
	w.WriteHeader(http.StatusOK)
}
