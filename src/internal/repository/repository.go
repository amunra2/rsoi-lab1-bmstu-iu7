package repository

import (
	"persserv/internal/entity"
	myerror "persserv/internal/error-my"
	"persserv/internal/repository/postgres"

	"github.com/jmoiron/sqlx"
)

type Person interface {
	GetAll() ([]entity.Person, *myerror.ErrorFull)
	GetById(personId int) (entity.Person, *myerror.ErrorFull)
	Create(person entity.Person) (int, *myerror.ErrorFull)
	Update(personId int, person entity.Person) (entity.Person, *myerror.ErrorFull)
	Delete(personId int) *myerror.ErrorFull
}

type Repository struct {
	Person
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Person: postgres.NewPersonRepository(db),
	}
}
