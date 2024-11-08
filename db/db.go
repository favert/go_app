package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=abc123 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Erro ao conectar no DB:", err.Error())
		return nil
	}
	return db
}
