package usecase

import (
	"persserv/internal/dto"
	myerror "persserv/internal/error-my"
	"persserv/internal/repository"
)

const (
	getAllPersonFuncName  = "usecase - persons - get all"
	getByIdPersonFuncName = "usecase - persons - get by id"
	createPersonFuncName  = "usecase - persons - create"
	updatePersonFuncName  = "usecase - persons - update"
	deletePersonFuncName  = "usecase - persons - delete"
)

type PersonUseCase struct {
	repo repository.Person
}

func NewPersonUseCase(repo repository.Person) *PersonUseCase {
	return &PersonUseCase{repo: repo}
}

func (s *PersonUseCase) GetAll() ([]dto.Person, *myerror.ErrorFull) {
	persons, err := s.repo.GetAll()

	personsDto := []dto.Person{}
	for _, elem := range persons {
		personsDto = append(personsDto, dto.ToPersonDTO(elem))
	}

	return personsDto, err
}

func (s *PersonUseCase) GetById(personId int) (dto.Person, *myerror.ErrorFull) {
	person, err := s.repo.GetById(personId)
	return dto.ToPersonDTO(person), err
}

func (s *PersonUseCase) Create(person dto.PersonCreate) (int, *myerror.ErrorFull) {
	return s.repo.Create(person.ToEntity())
}

func (s *PersonUseCase) Update(personId int, personInput dto.PersonUpdate) (dto.Person, *myerror.ErrorFull) {
	var updPerson dto.Person

	if err := personInput.Validate(); err != nil {
		return updPerson, myerror.NewError(updatePersonFuncName, err)
	}

	updPersonEntity, err := s.repo.Update(personId, personInput.ToEntity())

	return dto.ToPersonDTO(updPersonEntity), err
}

func (s *PersonUseCase) Delete(personId int) *myerror.ErrorFull {
	return s.repo.Delete(personId)
}
