package models

type MigrateServices interface {
	MigrateTask() error
	MigrateUser() error
}

type Service struct {
	migrate MigrateServices
}

func NewMigrateService(Ms MigrateServices) *Service {
	return &Service{
		migrate: Ms,
	}
}

func (S *Service) MigrateTask() error {
	return S.migrate.MigrateTask()
}

func (S *Service) MigrateUser() error {
	return S.migrate.MigrateUser()
}
