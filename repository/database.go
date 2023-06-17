package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"somename/configs"
)

import (
	"errors"
	_ "github.com/lib/pq"
	"log"
	"somename/models"
	"time"
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

func (r *Repository) DeleteUser(id int) error {
	_, err := r.db.Exec(`DELETE FROM "User" WHERE "ID"=$1`, id)
	if err != nil {
		log.Println(err)
	}
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
		`select "Date", "Bld", "Ubg", "Bil", "Pro", "Nit", "Ket", "Glu", "ph", "SG", "Leu"
		from "Analise" as a inner join "UserAnalise" UA on a."ID" = UA."analiseid" where "userid"=$1 order by "Date" desc`, userID); err != nil {
		return nil, err
	}

	return analisys, nil
}

func (r *Repository) GetUserByID(ID int) (*models.User, error) {

	var user models.User

	log.Println("Норма")
	if err := r.db.Get(&user, `SELECT * FROM "User" WHERE "ID"=$1`, ID); err != nil {
		return nil, err
		log.Println("Норм")
	}

	return &user, nil

}

func (r *Repository) LinkUserWithAnalyse(userID int, analyse *models.Analysis) error {
	var analyseID int

	if err := r.db.Get(&analyseID, `INSERT INTO "Analise" ("Date", "Bld", "Ubg", "Bil", "Pro", "Nit", "Ket", "Glu", "ph", "SG", "Leu")
                   VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) returning "ID"`,
		time.Now(),
		analyse.Bld,
		analyse.Ubg,
		analyse.Bil,
		analyse.Pro,
		analyse.Nit,
		analyse.Ket,
		analyse.Glu,
		analyse.PH,
		analyse.SG,
		analyse.Leu); err != nil {
		return err
	}

	_, err := r.db.Exec(`INSERT INTO "UserAnalise" (AnaliseId, UserId) VALUES ($1, $2)`, analyseID, userID)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
