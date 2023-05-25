package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"somename/configs"
	"somename/models"
)

type Repository struct {
	db *sqlx.DB
}

const dbDriverName = "postgres"

func GetDatabaseConnection(config *configs.Config) (*sqlx.DB, error) {
	return sqlx.Connect(
		dbDriverName, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			config.Host, config.Port, config.Username, config.Password, config.DBName, config.SSLMode))

}

func GetRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateUser(name, surname, patronymic, snils string) error {
	_, err := r.db.Exec(`INSERT INTO "User" ("FirstName", "LastName", "Patronymic", "SNILS")
            VALUES ($1, $2, $3, $4)`,
		name,
		surname,
		patronymic,
		snils)

	return err
}

func (r *Repository) UserIsExist(username string, password string) error {
	var count int8
	if err := r.db.Get(
		&count,
		`SELECT count(*) FROM "UserLogin" WHERE "Username"=$1 AND "Password"=$2`,
		username, password,
	); err != nil {
		return err
	}

	if count > 0 {
		return nil
	}
	return errors.New("не правильный пароль.логин")
}

func (r *Repository) GetUsers(search string) ([]models.User, error) {
	var users []models.User

	query := `SELECT
    "ID", "FirstName", "LastName", "Patronymic", "SNILS"
	FROM "User"
    WHERE LOWER("FirstName") LIKE '%' || $1 || '%' OR
	LOWER("LastName") LIKE '%' || $1 || '%' OR
    LOWER("Patronymic") LIKE '%' || $1 || '%'
	ORDER BY "LastName", "FirstName", "Patronymic"`
	if err := r.db.Select(&users, query, search); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Repository) GetAnalisys(userID int) ([]models.Analysis, error) {
	var analisys []models.Analysis

	if err := r.db.Select(&analisys,
		`select "Date", "Bld", "Ubg", "Bil", "Pro", "Nit", "Ket", "Glu", "pH", "SG", "Leu"
		from "Analise" as a inner join "UserAnalise" UA on a."ID" = UA."AnaliseId" where "UserId"=$1 order by "Date" desc`, userID); err != nil {
		return nil, err
	}

	return analisys, nil
}
