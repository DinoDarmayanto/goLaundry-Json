package main

import (
	"goJson/handler"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {

	http.HandleFunc("/getservice", handler.GetServiceHandler)
	http.HandleFunc("/getSevicebyId", handler.GetServiceByIdHandler)
	http.HandleFunc("/addService", handler.AddServiceHandler)
	http.HandleFunc("/updateService", handler.UpdateServiceHandler)
	http.HandleFunc("/deleteService", handler.DeleteServiceHandler)
	http.HandleFunc("/addtransaction", handler.AddTransactionHandler)
	http.HandleFunc("/alltransaction", handler.GetAllTransactionHandler)

	http.ListenAndServe(":8080", nil)
}
