package main

import (
	"database/sql"

	"goJson/handler"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var dbCon *sql.DB

func main() {
	db, err := sql.Open("postgres", "user=postgres host=localhost password=12345678 dbname=project-db sslmode=disable")
	if err != nil {
		log.Fatal("Cannot start app, Error when connect to DB ", err.Error())
	}
	dbCon = db

	http.HandleFunc("/getservice", handler.GetServiceHandler)
	http.HandleFunc("/GetSevicebyId", handler.GetServiceByIdHandler)
	http.HandleFunc("/AddService", handler.AddServiceHandler)
	http.HandleFunc("/UpdateService", handler.UpdateServiceHandler)
	http.HandleFunc("/DeleteService", handler.DeleteServiceHandler)
	http.HandleFunc("Addtransaction", handler.AddTransactionHandler)
	http.HandleFunc("Alltransaction", handler.GetAllTransactionHandler)

	http.ListenAndServe(":8080", nil)
}
