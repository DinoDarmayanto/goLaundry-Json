package handler

import (
	"encoding/json"
	"goJson/myutils"
	"goJson/svcrepo"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetServiceHandler(w http.ResponseWriter, r *http.Request) {
	_, err := svcrepo.AllService(myutils.GetDBConnection())
	if err != nil {
		w.Write([]byte("Tidak mendapatkan data"))
	}
}

func GetServiceByIdHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Write([]byte("id tidak valid"))
		return
	}
	svc, err := svcrepo.ServiceById(id, myutils.GetDBConnection())
	if err != nil {
		w.Write([]byte("id tidak ada di database"))
		return
	}
	if svc == nil {
		w.Write([]byte("Service tidak ditemukan"))
		return
	}
	serviceJSON, err := json.Marshal(svc)
	if err != nil {
		w.Write([]byte("Error saat menghasilkan JSON"))
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(serviceJSON)
}

func UpdateServiceHandler(w http.ResponseWriter, r *http.Request) {
	httpMethod := r.Method
	if httpMethod != "PUT" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method untuk UpdateServiceHandler harus PUT"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Gagal membaca body request"))
		return
	}

	var service svcrepo.Service
	err = json.Unmarshal(body, &service)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Gagal memparsing data JSON"))
		return
	}

	err = svcrepo.UpdateService(myutils.GetDBConnection(), &service)
	if err != nil {
		w.Write([]byte("Gagal update data"))
		return
	}

	w.Write([]byte("Berhasil melakukan update"))
}

func AddServiceHandler(w http.ResponseWriter, r *http.Request) {
	httpMethod := r.Method
	if httpMethod != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method untuk AddServiceHandler harus POST"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Gagal membaca body request"))
		return
	}

	var service svcrepo.Service
	err = json.Unmarshal(body, &service)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Gagal memparsing data JSON"))
		return
	}

	err = svcrepo.AddService(myutils.GetDBConnection(), service)
	if err != nil {
		w.Write([]byte("Gagal menambahkan data"))
		return
	}

	w.Write([]byte("Berhasil melakukan penambahan"))
}

func DeleteServiceHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Write([]byte("id tidak valid"))
		return
	}

	err = svcrepo.DeleteService(id, myutils.GetDBConnection())
	if err != nil {
		w.Write([]byte("Gagal menghapus service"))
		return
	}

	w.Write([]byte("Service berhasil dihapus"))
}
