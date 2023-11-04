package mstore

import (
	"fmt"
	"graphql/graph/model"
	"graphql/models"
)

type Service struct {
	C models.Conn
}

func NewService(ms *models.Conn) Service {
	return Service{
		C: *ms,
	}
}

func (s *Service) CreateCompanyServices(cd model.NewCompany) (*model.Company, error) {
	fmt.Println("creating company in database") // Displaying a message to signify performing database manipulation
	compData := models.NewCompany{
		CompanyName: cd.CompanyName,
		FoundedYear: cd.FoundedYear,
		Location:    cd.Location,
	}
	comp, err := s.C.InsertCompany(compData)
	if err != nil {
		return nil, err
	}
	return &model.Company{
		CompanyName: comp.CompanyName,
		FoundedYear: comp.FoundedYear,
		Location:    comp.Location,
		Jobs:        nil,
	}, nil
}

func (s *Service) CreateJobServices(ni model.NewJob) (*model.Job, error) {
	fmt.Println("creating job in database")
	jobData := models.NewJob{
		Title:              ni.Title,
		ExperienceRequired: ni.ExperienceRequired,
		CompanyID:          ni.CompanyID,
	}

	job, err := s.C.InsertJob(jobData)
	if err != nil {
		return nil, err
	}
	return &model.Job{
		Title:              job.Title,
		ExperienceRequired: job.ExperienceRequired,
		CompanyID:          job.CompanyID,
	}, nil
}

func (s *Service) ViewAllCompaniesServices() ([]*model.Company, error) {
	cmp, err := s.C.FetchAllCompanies()
	if err != nil {
		return nil, err
	}
	var result []*model.Company

	for _, c := range cmp {
		result = append(result, &model.Company{
			ID:          c.ID,
			CompanyName: c.CompanyName,
			FoundedYear: c.FoundedYear,
			Location:    c.Location,
		})
	}
	return result, nil
}

func (s *Service) ViewAllJobsServices() ([]*model.Job, error) {
	job, err := s.C.FetchAllJobPostings()
	if err != nil {
		return nil, err
	}
	var result []*model.Job
	for _, c := range job {
		result = append(result, &model.Job{
			ID:                 c.ID,
			Title:              c.Title,
			ExperienceRequired: c.ExperienceRequired,
			CompanyID:          c.CompanyID,
		})
	}
	return result, nil
}

func (s *Service) FindCompanyByIDServices(companyID string) (*model.Company, error) {
	cmp, err := s.C.FetchJobsForCompany(companyID)
	if err != nil {
		return nil, err
	}
	return &model.Company{
		ID:          cmp.ID,
		CompanyName: cmp.CompanyName,
		FoundedYear: cmp.FoundedYear,
		Location:    cmp.Location,
	}, nil
}

func (s *Service) FindJobByJobIDServices(jobID string) (*model.Job, error) {
	job, err := s.C.FetchJobById(jobID)
	if err != nil {
		return nil, err
	}
	return &model.Job{
		ID:                 job.ID,
		Title:              job.Title,
		ExperienceRequired: job.ExperienceRequired,
		CompanyID:          job.CompanyID,
	}, nil
}

func (s *Service) FindJobByCompanyIDServices(companyID string) ([]*model.Job, error) {
	job, err := s.C.FetchJobByCompId(companyID)
	if err != nil {
		return nil, err
	}
	var result []*model.Job

	for _, c := range job {
		result = append(result, &model.Job{
			ID:                 c.ID,
			Title:              c.Title,
			ExperienceRequired: c.ExperienceRequired,
			CompanyID:          c.CompanyID,
		})
	}
	return result, nil
}

func (s *Service) RegisterUserServices(input model.NewUser) (*model.User, error) {
	fmt.Println("signup is in progress")
	userData := models.NewUser{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
	user, err := s.C.InsertUser(userData)
	if err != nil {
		return nil, err
	}
	return &model.User{
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: user.Password,
	}, nil
}
