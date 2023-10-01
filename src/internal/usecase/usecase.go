package usecase

import (
	"persserv/internal/dto"
	myerror "persserv/internal/error-my"
	"persserv/internal/repository"
)

type Person interface {
	GetAll() ([]dto.Person, *myerror.ErrorFull)
	GetById(personId int) (dto.Person, *myerror.ErrorFull)
	Create(dto.PersonCreate) (int, *myerror.ErrorFull)
	Update(personId int, person dto.PersonUpdate) (dto.Person, *myerror.ErrorFull)
	Delete(personId int) *myerror.ErrorFull
}

type UseCase struct {
	Person
}

func NewUseCase(repos *repository.Repository) *UseCase {
	return &UseCase{
		Person: NewPersonUseCase(repos.Person),
	}
}
