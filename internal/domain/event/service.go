package event

type Service interface {
	FindAll() ([]Event, error)
	FindOne(id int64) (*Event, error)
	Create(event Event) (string, error)
	Update(event Event) (string, error)
	Delete(id int64) (string, error)
}

type service struct {
	repo *Repository
}

func NewService(r *Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) FindAll() ([]Event, error) {
	return (*s.repo).FindAll()
}

func (s *service) FindOne(id int64) (*Event, error) {
	return (*s.repo).FindOne(id)
}

func (s *service) Create(event Event) (string, error) {
	return (*s.repo).Create(event)
}

func (s *service) Update(event Event) (string, error) {
	return (*s.repo).Update(event)
}

func (s *service) Delete(id int64) (string, error) {
	return (*s.repo).Delete(id)
}
