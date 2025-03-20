package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	error := godotenv.Load()
	if error != nil {
		panic(error.Error())
	}
	db, err := sql.Open("mysql", os.Getenv("DB_DSN_NAME"))
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS Users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,
		created_at DATETIME
	)
	`
	_, err = db.Exec(createUsersTable)
	if err != nil {
		log.Fatal("Failed to create Users table:", err)
	}

	createTasksTable := `
	CREATE TABLE IF NOT EXISTS Tasks (
		id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		description VARCHAR(255) NOT NULL,
		user_id INT NOT NULL,
		due_date DATETIME,
		FOREIGN KEY(user_id) REFERENCES Users(id)
	)
	`
	_, err = db.Exec(createTasksTable)
	if err != nil {
		log.Fatal("Failed to create Tasks table:", err)
	}

	fmt.Println("Tables creation done")
}
