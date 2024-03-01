package main

import (
	"database/sql"
	"fmt"

	"github.com/Patrikesm/basic-solid-api/internal/server"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306)/routes?parseTime=true")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	s := server.NewServer(db)

	fmt.Println("Started server on port 3000")
	s.Run()
}
