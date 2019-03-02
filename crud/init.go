package crud

import (
	"database/sql"
	"github.com/volatiletech/sqlboiler/boil"
)

var (
	db *sql.DB
)

func SetupDb() *sql.DB {
	db, err := sql.Open("postgres", `postgresql://root@localhost:26257/votecube?sslmode=disable`)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	boil.SetDB(db)

	return db
}
