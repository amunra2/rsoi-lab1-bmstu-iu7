package dto

import (
	"database/sql"
	"persserv/internal/entity"
	myerror "persserv/internal/error-my"
)

// Create
type PersonCreate struct {
	Name    string `json:"name" binding:"required"`
	Age     int64  `json:"age,omitempty" binding:"gte=0,lte=130"`
	Address string `json:"address,omitempty"`
	Work    string `json:"work,omitempty"`
}

func (p *PersonCreate) ToEntity() entity.Person {
	return entity.Person{
		Name: p.Name,
		Age: sql.NullInt64{
			Int64: p.Age,
		},
		Address: sql.NullString{
			String: p.Address,
		},
		Work: sql.NullString{
			String: p.Work,
		},
	}
}

// Update
type PersonUpdate struct {
	Name    string `json:"name,omitempty"`
	Age     int64  `json:"age,omitempty" binding:"gte=0,lte=130"`
	Address string `json:"address,omitempty"`
	Work    string `json:"work,omitempty"`
}

func (p *PersonUpdate) Validate() error {
	if p.Name == "" && p.Age == 0 && p.Address == "" && p.Work == "" {
		return myerror.UpdateStructureIsEmpty
	}

	return nil
}

func (p *PersonUpdate) ToEntity() entity.Person {
	return entity.Person{
		Name: p.Name,
		Age: sql.NullInt64{
			Int64: p.Age,
		},
		Address: sql.NullString{
			String: p.Address,
		},
		Work: sql.NullString{
			String: p.Work,
		},
	}
}

// Get
type Person struct {
	Id      int64  `json:"id"`
	Name    string `json:"name,"`
	Age     int64  `json:"age"`
	Address string `json:"address"`
	Work    string `json:"work"`
}

func ToPersonDTO(p entity.Person) Person {
	return Person{
		Id:      p.Id.Int64,
		Name:    p.Name,
		Age:     p.Age.Int64,
		Address: p.Address.String,
		Work:    p.Work.String,
	}
}
