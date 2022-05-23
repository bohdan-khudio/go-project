package user

type Service interface {
	FindAll() ([]User, error)
	FindOne(id int64) (*User, error)
}

type service struct {
	repo *Repository
}

func NewService(r *Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) FindAll() ([]User, error) {
	return (*s.repo).FindAll()
}

func (s *service) FindOne(id int64) (*User, error) {
	return (*s.repo).FindOne(id)
}
