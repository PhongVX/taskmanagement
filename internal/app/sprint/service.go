package sprint

// NewService return a new task service
func NewService(r Repository) *Service {
	srv := &Service{
		repo: r,
	}
	return srv
}
