package main

import (
	"database/sql"
	"fmt"

	"github.com/Natt-S10/Natthaphon_agnos_backend_internship_2023/srcs/routes"
	_ "github.com/lib/pq"
)

const (
	DBTYPE            string = "postgres"
	DBHOST            string = "localhost"
	DBPORT                   = 5438
	DBUSER            string = "dev"
	DBNAME            string = "api_log"
	DBPASSWORD        string = "12345678"
	SCHEMATEMPLATE    string = "CREATE SCHEMA IF NOT EXISTS logging AUTHORIZATION "
	USESCHEMATEMPLATE string = "set search_path to logging;"
	TABLETEMPLATE     string = `CREATE TABLE IF NOT EXISTS strong_passwd_steps_log (
		timestamp TIMESTAMP PRIMARY KEY,
		route VARCHAR(50) NOT NULL,
		status INT NOT NULL,
		init_password VARCHAR(50),
		num_of_steps INT,
		error_flag BOOLEAN NOT NULL
	 )`
)

func getPsqlInfo() string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		DBHOST, DBPORT, DBUSER, DBPASSWORD, DBNAME)
	return psqlInfo
}

func connectPostgres() *sql.DB {
	db, err := sql.Open(DBTYPE, getPsqlInfo())
	if err != nil {
		panic(err)
	}
	//ping test
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Postgres Connected Successfully")
	return db
}

func createTableIfNotExists(db *sql.DB) {
	_, err := db.Exec(TABLETEMPLATE)
	if err != nil {
		fmt.Println("Problem on createTableIfNotExists: strong_passwd_steps_log")
		panic(err)
	}
	fmt.Println("Assure existance of strong_passwd_steps_log table")
}

func createSchemaIfNotExists(db *sql.DB) {
	_, err := db.Exec(SCHEMATEMPLATE + DBUSER)
	if err != nil {
		fmt.Println("Problem on createSchemaIfNotExists: logging")
		panic(err)
	}
	fmt.Println("Assure existance of logging schema")
}

func useLoggingSchema(db *sql.DB) {
	_, err := db.Exec(USESCHEMATEMPLATE)
	if err != nil {
		fmt.Println("Problem on useLoggingSchema: logging")
		panic(err)
	}
	fmt.Println("Use logging schema")
}

func main() {
	// connecting postgres
	db := connectPostgres()
	defer db.Close()
	createSchemaIfNotExists(db)
	useLoggingSchema(db)
	createTableIfNotExists(db)

	// r := routes.setupRouter()
	r := routes.SetupRouter(db)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
