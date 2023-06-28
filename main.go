package main

import (
	"goJson/handler"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {

	http.HandleFunc("/getservice", handler.GetServiceHandler)
	http.HandleFunc("/GetSevicebyId", handler.GetServiceByIdHandler)
	http.HandleFunc("/AddService", handler.AddServiceHandler)
	http.HandleFunc("/UpdateService", handler.UpdateServiceHandler)
	http.HandleFunc("/DeleteService", handler.DeleteServiceHandler)
	http.HandleFunc("/Addtransaction", handler.AddTransactionHandler)
	http.HandleFunc("/Alltransaction", handler.GetAllTransactionHandler)

	http.ListenAndServe(":8080", nil)
}
