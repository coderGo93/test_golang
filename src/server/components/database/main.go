package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //Driver para SQL Server
)

//Variables constantes
const (
	ComponentName = "database"
	componentDir  = "server/components/" + ComponentName
	dbEngine      = "mysql"
	dbDataSource  = "user:password@/dbname;"
)

//Variables

var (
	db *sql.DB
)

//La funcion init() siempre corre primero en el archivo, es como el constructor
// Solo corre en programa genera, no cada vez en esta clase

func init() {
	fmt.Println("Success -> " + componentDir)
}

//OpenDB para conectar a la bd
func OpenDB() {
	var err error
	db, err = sql.Open(dbEngine, dbDataSource)

	if err = db.Ping(); err != nil {
		fmt.Println("Conexión fallida a la base de datos")
	} else {
		fmt.Println("La base de datos se ha conectado")
	}

}

//GetDatabase para obtener la bd
func GetDatabase() *sql.DB {
	return db
}
