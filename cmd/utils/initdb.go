package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	userName = "root"
	password = "12345678"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "reservas"
)

func InitDB() *sql.DB {
	// Conexión de datos Golang: "Nombre de usuario: contraseña @ tcp (IP: número de puerto) / nombre de la base de datos? Charset = utf8"
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", userName, password, ip, dbName)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)

	}
	// Establecer el número máximo de conexiones a la base de datos
	db.SetConnMaxLifetime(10)
	// Establecer el número máximo de conexiones inactivas en la base de datos
	db.SetMaxIdleConns(5)
	// Verificar conexión
	if err := db.Ping(); err != nil {
		panic(err)
	}
	// Devuelve la referencia del puntero de la conexión a la base de datos
	return db
}
