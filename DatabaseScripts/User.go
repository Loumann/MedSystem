package database

import (
	"github.com/gin-gonic/gin"
	"log"
)

// CREATE TABLE DATABASE
const user = `
create table "User"
(
    "ID"         serial primary key,
    "FirstName"  varchar,
    "LastName"   varchar,
    "Patronymic" varchar,
    "SNILS"      varchar
);`
const userAnalise = `
create table "UserAnalise"
(
    "AnaliseId" integer   not null references "Analise",
    "UserId"    integer   not null references "User"
    constraint "UserAnalisePK"
        primary key ("AnaliseId", "UserId")
);`
const analise = `create table Analise
(
 "ID" integer primary key, 
"Dat" varchar,
"Bld" varchar,
"Ubg" varchar,
"Bil" varchar,
"Pro" varchar,
"Nit" varchar,
"Ket" varchar,
"Glu" varchar,
"PH"  varchar,
"SG"  varchar,
"Leu" varchar
)`
const UserLogin = `create table UserLogin
(
    Username varchar,
    Password varchar
    
)`

//CREATE TABLE DATABASE

type User struct {
	ID         int    `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Patronymic string `json:"patronymic"`
	SNILS      string `json:"snils"`
}

type Analise struct {
	ID   int    `json:"id"`
	Date string `json:"date"`
	Bld  string `json:"bld"`
	Ubg  string `json:"ubg"`
	Bil  string `json:"bil"`
	Pro  string `json:"pro"`
	Nit  string `json:"nit"`
	Ket  string `json:"ket"`
	Glu  string `json:"glu"`
	PH   string `json:"ph"`
	SG   string `json:"sh"`
	Leu  string `json:"leu"`
}

func Select(context *gin.Context) {
	var userAnalysis = User{}
	err := context.BindJSON(userAnalysis)
	if err != nil {
		log.Println(err)
		context.Status(404)
		return
	}

	var count int
	row := DbQuer.QueryRow(`select count(*) from "User"`)
	err = row.Scan(&count)
	if err != nil {
		log.Println(err)
		context.JSON(404, "Данных нет")
		return

	}
	if count == 0 {
		context.JSON(404, "Данных нет")

	}
}

func SelectUsers(search string) []User {
	query := `SELECT
    "ID", "FirstName", "LastName", "Patronymic", "SNILS"
	FROM "User"
    WHERE LOWER("FirstName") LIKE '%' || $1 || '%' OR
	LOWER("LastName") LIKE '%' || $1 || '%' OR
    LOWER("Patronymic") LIKE '%' || $1 || '%'
	ORDER BY "LastName", "FirstName", "Patronymic"`

	rows, e := DbQuer.Query(query, search)
	if e != nil {
		log.Println(e)
		return nil
	}

	defer rows.Close()

	users := make([]User, 0)

	for rows.Next() {
		user := User{}
		e = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Patronymic, &user.SNILS)
		if e != nil {
			log.Println(e)
			return nil
		}

		users = append(users, user)
	}

	return users
}

func SelectUserAnalise(id int) []Analise {
	rows, e := DbQuer.Query(`select "Date", "Bld", "Ubg", "Bil", "Pro", "Nit", "Ket", "Glu", "pH", "SG", "Leu"
		from "Analise" as a inner join "UserAnalise" UA on a."ID" = UA."AnaliseId" where "UserId"=$1`, id)
	if e != nil {
		log.Println(e)
		return nil
	}

	defer rows.Close()

	analyses := make([]Analise, 0)

	for rows.Next() {
		analise := Analise{}
		e = rows.Scan(
			&analise.Date,
			&analise.Bld,
			&analise.Ubg,
			&analise.Bil,
			&analise.Pro,
			&analise.Nit,
			&analise.Ket,
			&analise.Glu,
			&analise.PH,
			&analise.SG,
			&analise.Leu,
		)
		if e != nil {
			log.Println(e)
			return nil
		}

		analyses = append(analyses, analise)
	}

	return analyses
}
