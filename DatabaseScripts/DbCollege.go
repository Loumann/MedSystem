package database

import "database/sql"

var DbQuerCollege *sql.DB

func ConnectDBCollege() {

	var err error
	DbQuer, err = sql.Open("postgres", `host=10.14.206.28 port=5432 user=student password=1234 dbname=medbase sslmode=disable`)
	if err != nil {
		panic(err)
	}

	err = DbQuer.Ping()
	if err != nil {
		panic(err)
	}
}
