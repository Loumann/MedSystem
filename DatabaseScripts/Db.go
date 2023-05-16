package database

import "database/sql"

var DbQuer *sql.DB

func ConnectDB() {

	var err error
	DbQuer, err = sql.Open("postgres", `host=localhost port=5432 user=postgres password=1234 dbname=MedBase sslmode=disable`)
	if err != nil {
		panic(err)
	}

	err = DbQuer.Ping()
	if err != nil {
		panic(err)
	}
}
