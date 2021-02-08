package main

import (
	"database/sql"
"fmt"
_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "Thanks001"
	database = "users_db"
)


func main()  {
	//Database initialization information
	initInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s database=%s sslmode=disable",
		host, port, user, password, database)

	db, err := sql.Open("postgres", initInfo)
	if err != nil{
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil{
		panic(err)
	}

	fmt.Printf("Established connection with %s database", database)
}
