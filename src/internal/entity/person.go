package entity

import (
	"database/sql"
)

// type Person struct {
// 	Id      int    `json:"-" db:"id"`
// 	Name    string `json:"name" db:"name" binding:"required"`
// 	Age     int    `json:"age" db:"age" binding:"required"`
// 	Address string `json:"address" db:"address" binding:"required"`
// 	Work    string `json:"work" db:"work" binding:"required"`
// }

type Person struct {
	Id      sql.NullInt64  `db:"id"`
	Name    string         `db:"name"`
	Age     sql.NullInt64  `db:"age"`
	Address sql.NullString `db:"address"`
	Work    sql.NullString `db:"work"`
}
