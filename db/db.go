package db

import (
	"fmt"
	"os"
	"strconv"
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
)

type DBCredentials struct {
	Host string
	Port int
	Name string
	Username string
	Password string
}

type Company struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Zip string `json:"zip"`
	Website string `json:"website"`
}

type DBConnectionError struct{
	message string
}

func (e *DBConnectionError) Error() string {
	return e.message
}

func NewDBConnectionError(message string) error {
	return &DBConnectionError{message}
}

func CreateDBConnection(dbCredentials DBCredentials) (*sql.DB, error) { 
	var connectionString = fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", 
		dbCredentials.Host, 
		dbCredentials.Port, 
		dbCredentials.Username, 
		dbCredentials.Password, 
		dbCredentials.Name)

	db, err := sql.Open("postgres", connectionString)
	
	if err != nil {
		return nil, NewDBConnectionError("Failed to connect to database. Check if the credentials are correct and try again.")
	}

	err = db.Ping()

	if err != nil {
		return nil, NewDBConnectionError("Fail to ping to database. Check if the database is alive and try again.")
	}

	return db, nil
}

func GetDBConnection() (*sql.DB, error) {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))

	if err != nil {
		return nil, errors.New("DB_PORT environment variable not found.")
	}

	return CreateDBConnection(DBCredentials{
		Host: os.Getenv("DB_HOST"),
		Port: port,
		Name: os.Getenv("DB_NAME"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
	})
}

func executeSQL(sql string) error {
	db, err := GetDBConnection()

	if err != nil {
		return err
	}

	_, err = db.Exec(sql)

	if err != nil {
		fmt.Println(err)
		return err
	}

	defer db.Close()

	return nil
}

func CreateCompanyTable() error {
	return executeSQL(`CREATE TABLE IF NOT EXISTS companies (
		id SERIAL PRIMARY KEY, 
		name VARCHAR(255) NOT NULL, 
		zip VARCHAR(5) NOT NULL,
		website VARCHAR(255)
	);`)
}

func DropCompanyTable() error {
	return executeSQL(`DROP TABLE IF EXISTS companies`)
}

func GetQuantityOfCompanies() (int, error) {
	db, err := GetDBConnection()

	if err != nil {
		return 0, err
	}

	rows, err := db.Query("SELECT COUNT(*) FROM companies;")

	if err != nil {
		return 0, err
	}

	defer rows.Close()

	var quantity int = 0

	for rows.Next() {
		err := rows.Scan(&quantity)
		if err != nil {
			return 0, err
		} else {
			return quantity, nil
		}
	}
	
	return 0, errors.New("No results found")
}

func InsertCompany(name string, zip string) error {
	db, err := GetDBConnection()

	if err != nil {
		return err
	}

	insert, err := db.Prepare("INSERT INTO companies (name, zip) VALUES ($1, $2)")
	
	if err != nil {
		fmt.Println(err)
		return err
	}
	
	_, insertError := insert.Exec(name, zip)

	if insertError != nil {
		fmt.Println(insertError)
		return insertError
	}

	defer db.Close()

	return nil
}

func DeleteCompany(name string, zip string) error {
	db, err := GetDBConnection()

	if err != nil {
		return err
	}

	delete, err := db.Prepare("DELETE FROM companies WHERE name=$1 AND zip=$2")
	
	if err != nil {
		fmt.Println(err)
		return err
	}
	
	_, deleteError := delete.Exec(name, zip)

	if deleteError != nil {
		fmt.Println(deleteError)
		return deleteError
	}

	defer db.Close()

	return nil
}

func GetCompanyIdByNameAndZip(name string, zip string) (int, error) {
	db, err := GetDBConnection()

	if err != nil {
		return 0, err
	}

	rows, err := db.Query("SELECT id FROM companies WHERE LOWER(name) LIKE '%' || $1 || '%' AND zip = $2 LIMIT 1", name, zip)

	if err != nil {
		return 0, err
	}

	defer rows.Close()

	var id int = 0

	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return 0, err
		} else {
			return id, nil
		}
	}
	
	return 0, nil
}

func UpdateCompanyWebsite(id int, website string) error {
	db, err := GetDBConnection()

	if err != nil {
		return err
	}

    update, err := db.Prepare("UPDATE companies SET website=$1 WHERE id=$2")
	
	if err != nil {
		return err
	}

    result, err := update.Exec(website, id)
    if err != nil {
		return err
	}

	affect, err := result.RowsAffected()
	
    if err != nil {
		return err
	}

	if affect != 1 {
		return errors.New(fmt.Sprintf("%d rows affected", affect))
	}

	return nil
}

func GetCompanyByNameAndZip(name string, zip string) (*Company, error) {
	db, err := GetDBConnection()

	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT id, name, COALESCE(website,'') FROM companies WHERE LOWER(name) LIKE '%' || $1 || '%' AND zip = $2 LIMIT 1", name, zip)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var id int
	var fullname string
	var website string

	for rows.Next() {
		err := rows.Scan(&id, &fullname, &website)
		if err != nil {
			return nil, err
		} else {
			return &Company{
				Id: id,
				Name: fullname,
				Zip: zip,
				Website: website}, nil
		}
	}
	
	return nil, nil
}
