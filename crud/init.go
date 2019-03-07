package crud

import (
	"database/sql"
	"github.com/volatiletech/sqlboiler/boil"
)

var (
	DB *sql.DB
)

func SetupDb() {
	DB, err := sql.Open("postgres", `postgresql://root@localhost:26257/votecube?sslmode=disable`)

	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	boil.SetDB(DB)
}
