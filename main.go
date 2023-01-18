package main

import (
	"database/sql"
	"fmt"

	"github.com/Natt-S10/Natthaphon_agnos_backend_internship_2023/srcs/routes"
	_ "github.com/lib/pq"
)

const (
	DBTYPE        = "postgres"
	DBHOST        = "localhost"
	DBPORT        = 5432
	DBUSER        = "dev"
	DBNAME        = "log"
	DBPASSWORD    = "12345678"
	TABLETEMPLATE = `CREATE TABLE IF NOT EXISTS log (
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
	defer db.Close()
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
		fmt.Println("Problem on CreateIfNotExistTable: log")
		panic(err)
	}
}

func main() {
	// connecting postgres
	db := connectPostgres()
	createTableIfNotExists(db)

	// r := routes.setupRouter()
	r := routes.SetupRouter(db)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
