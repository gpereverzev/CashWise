package main

import (
	"database/sql"
	"github.com/gpereverzev/CashWise/src/api"
	"github.com/gpereverzev/CashWise/src/data/postgresql"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=8070 user=cashwise_db_user password=123456789 dbname=cashwise_db sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	masterDB := postgresql.NewMasterDB(db)

	api.Run(api.Config{
		MasterDB: masterDB,
	})
}
