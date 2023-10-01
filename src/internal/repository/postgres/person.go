package postgres

import (
	"database/sql"
	"persserv/internal/entity"
	myerror "persserv/internal/error-my"

	"github.com/jmoiron/sqlx"
)

const (
	queryGetAll = `
		SELECT *
		FROM person
	`
	queryGetById = `
		SELECT *
		FROM person p
		WHERE p.id = ?
	`
	queryCreate = `
		INSERT INTO person (name, age, address, work)
		VALUES (?, ?, ?, ?)
		RETURNING id
	`
	queryUpdate = `
		UPDATE person
		SET name = COALESCE(NULLIF(TRIM(?), ''), name),
				age = COALESCE(NULLIF(?, 0), age),
				address = COALESCE(NULLIF(TRIM(?), ''), address),
				work = COALESCE(NULLIF(TRIM(?), ''), work)
		WHERE id = ?
		RETURNING id, name, age, address, work
	`
	queryDelete = `
		DELETE
		FROM person
		WHERE id=$1
		RETURNING id
	`
)

const (
	getAllPersonFuncName  = "repository - postgres - persons - get all"
	getByIdPersonFuncName = "repository - postgres - persons - get by id"
	createPersonFuncName  = "repository - postgres - persons - create"
	updatePersonFuncName  = "repository - postgres - persons - update"
	deletePersonFuncName  = "repository - postgres - persons - delete"
)

type PersonRepository struct {
	db *sqlx.DB
}

func NewPersonRepository(db *sqlx.DB) *PersonRepository {
	return &PersonRepository{db: db}
}

func (r *PersonRepository) GetAll() ([]entity.Person, *myerror.ErrorFull) {
	var items []entity.Person

	query := r.db.Rebind(queryGetAll)
	err := r.db.Select(&items, query)

	if err != nil && err != sql.ErrNoRows {
		return nil, myerror.NewError(getAllPersonFuncName, err)
	}

	return items, nil
}

func (r *PersonRepository) GetById(personId int) (entity.Person, *myerror.ErrorFull) {
	var item entity.Person

	query := r.db.Rebind(queryGetById)
	err := r.db.Get(&item, query, personId)

	if err == sql.ErrNoRows {
		return item, myerror.NewError(getByIdPersonFuncName, myerror.NotFound)
	} else if err == nil {
		return item, nil
	} else {
		return item, myerror.NewError(getByIdPersonFuncName, err)
	}
}

func (r *PersonRepository) Create(person entity.Person) (int, *myerror.ErrorFull) {
	var itemId int

	query := r.db.Rebind(queryCreate)
	row := r.db.QueryRow(query, person.Name, person.Age.Int64, person.Address.String, person.Work.String)
	err := row.Scan(&itemId)

	if err != nil {
		return 0, myerror.NewError(createPersonFuncName, err)
	}

	return itemId, nil
}

func (r *PersonRepository) Update(personId int, person entity.Person) (entity.Person, *myerror.ErrorFull) {
	updatedPerson := entity.Person{}

	query := r.db.Rebind(queryUpdate)
	err := r.db.
		QueryRowx(query, person.Name, person.Age.Int64,
			person.Address.String, person.Work.String, personId).
		Scan(&updatedPerson.Id, &updatedPerson.Name, &updatedPerson.Age,
			&updatedPerson.Address, &updatedPerson.Work)

	if err == sql.ErrNoRows {
		return updatedPerson, myerror.NewError(updatePersonFuncName, myerror.NotFound)
	} else if err == nil {
		return updatedPerson, nil
	} else {
		return updatedPerson, myerror.NewError(updatePersonFuncName, err)
	}
}

func (r *PersonRepository) Delete(personId int) *myerror.ErrorFull {
	var id int

	query := r.db.Rebind(queryDelete)
	err := r.db.QueryRowx(query, personId).Scan(&id)

	if err == sql.ErrNoRows {
		return myerror.NewError(deletePersonFuncName, myerror.NotFound)
	} else if err == nil {
		return nil
	} else {
		return myerror.NewError(updatePersonFuncName, err)
	}
}
