package user

// NewService return a new user service
func NewService(r Repository) *Service {
	srv := &Service{
		repo: r,
	}
	return srv
}
