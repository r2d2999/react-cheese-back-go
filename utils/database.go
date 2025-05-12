package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB 

func InitDB(){
	// Conexión
	connStr := "host=localhost port=5432 user=postgres password=123 dbname=pruebaDB sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error al abrir la conexión:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}
	fmt.Println("Conexión exitosa a PostgreSQL!")

}