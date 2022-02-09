package postgres

import (
	"database/sql"
	"fmt"

	helper "github.com/FurkanSamaraz/IsEmpty"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "172754"
	dbname   = "postgres"
)

func OpenConnention() *sql.DB {

	psq := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psq)

	if err != nil {
		helper.IsEmpty(err.Error())
	}
	err = db.Ping()
	if err != nil {
		helper.IsEmpty(err.Error())
	}

	return db
}
