package handler

import (
	"encoding/json"
	"goJson/myutils"
	"goJson/trxrepo"
	"io/ioutil"
	"net/http"
	"strconv"
)

func AddTransactionHandler(w http.ResponseWriter, r *http.Request) {
	httpMethod := r.Method
	if httpMethod != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method untuk AddTransactionHandler harus POST"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Gagal membaca body request"))
		return
	}

	var transaction trxrepo.TransactionHeader
	err = json.Unmarshal(body, &transaction)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Gagal memparsing data JSON"))
		return
	}

	err = trxrepo.AddTransaction(myutils.GetDBConnection(), &transaction)
	if err != nil {
		w.Write([]byte("Gagal Add Data"))
		return
	}

	w.Write([]byte("Berhasil Add Data"))
}

func GetAllTransactionHandler(w http.ResponseWriter, r *http.Request) {
	httpMethod := r.Method
	if httpMethod != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method untuk GetAllTransactionHandler harus GET"))
		return
	}

	trxNoStr := r.FormValue("trxNo")
	trxNo, err := strconv.Atoi(trxNoStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Parameter trxNo tidak valid"))
		return
	}

	transactions, err := trxrepo.ViewTransaction(myutils.GetDBConnection(), trxNo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Gagal mendapatkan data transaksi"))
		return
	}

	transactionsJSON, err := json.Marshal(transactions)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Gagal menghasilkan JSON"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(transactionsJSON)
}
