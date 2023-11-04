package store

import (
	"graphql/graph/model"
)

type Storer interface {
	CreateCompanyServices(company model.NewCompany) (*model.Company, error)
	CreateJobServices(job model.NewJob) (*model.Job, error)
	ViewAllCompaniesServices() ([]*model.Company, error)
	ViewAllJobsServices() ([]*model.Job, error)
	FindCompanyByIDServices(companyID string) (*model.Company, error)
	FindJobByJobIDServices(jobID string) (*model.Job, error)
	FindJobByCompanyIDServices(companyID string) ([]*model.Job, error)
	RegisterUserServices(input model.NewUser) (*model.User, error)
}

type Store struct {
	Storer
}

func NewStore(storer Storer) Store {
	return Store{Storer: storer}
}
