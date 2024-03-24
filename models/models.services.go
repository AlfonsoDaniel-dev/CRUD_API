package models

type MigrateServices interface {
	Migrate() error
}

type Service struct {
	migrate MigrateServices
}

func NewMigrateService(Ms MigrateServices) *Service {
	return &Service{
		migrate: Ms,
	}
}

func (S *Service) Migrate() error {
	return S.migrate.Migrate()
}
