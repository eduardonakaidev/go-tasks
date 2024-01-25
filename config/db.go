package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func ConnectPostgresDB() (*sql.DB, error) {

	url := fmt.Sprintf("host:%s user=%s password=%s dbname=%s port=%s sslmode=%s", os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSL"),
	)
	DB, err := sql.Open("postgres", url)
	if err != nil {
		panic(err)
	}

	// SQL Create DB
	_, err = DB.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		name TEXT,
		email TEXT,
		password TEXT
	);
	
	CREATE TABLE IF NOT EXISTS tasks (
		id TEXT PRIMARY KEY,
		title TEXT,
		description TEXT,
		is_completed BOOLEAN,
		id_user TEXT,
		created_at DATETIME,
		FOREIGN KEY (id_user) REFERENCES users (id)
	);`)

	if err != nil {
		log.Fatal("Erro ao criar o banco de dados: ", err)
	}

	fmt.Println("Tabelas criadas com sucesso!")
	err = DB.Ping()
	return DB, err
}
