package repository

type Authorization interface {

}

type Agent interface {

}

type Repository struct {
	Authorization
	Agent
}

func NewRepository() *Repository {
	return &Repository{}
}
