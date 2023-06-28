package svcrepo

import (
	"database/sql"
	"fmt"
)

func AllService(dbCon *sql.DB) ([]Service, error) {
	qry := "SELECT id, name, price, uom FROM ms_service"

	rows, err := dbCon.Query(qry)
	if err != nil {
		return nil, fmt.Errorf("GetAllService(): %w", err)
	}
	defer rows.Close()

	var arrService []Service
	for rows.Next() {
		var svc Service
		err := rows.Scan(&svc.Id, &svc.Name, &svc.Price, &svc.Uom)
		if err != nil {
			return nil, fmt.Errorf("GetAllService(): %w", err)
		}
		arrService = append(arrService, svc)
	}

	return arrService, nil
}
func AddService(db *sql.DB, svc Service) error {
	qry := "INSERT INTO ms_service (name, price, uom) VALUES ($1, $2, $3)"
	_, err := db.Exec(qry, svc.Name, svc.Price, svc.Uom)
	if err != nil {
		return fmt.Errorf("AddService(): %w", err)
	}
	return nil
}

func DeleteService(id int, conDB *sql.DB) error {
	qry := "DELETE FROM ms_service WHERE id = $1"
	_, err := conDB.Exec(qry, id)
	if err != nil {
		return fmt.Errorf("DeleteService(): %w", err)
	}
	return nil
}

func UpdateService(conDB *sql.DB, svc *Service) error {
	qry := "UPDATE ms_service SET name = $1, price = $2, uom = $3 WHERE id = $4"
	_, err := conDB.Exec(qry, svc.Name, svc.Price, svc.Uom, svc.Id)
	if err != nil {
		return fmt.Errorf("UpdateService(): %w", err)
	}
	return nil
}

func ServiceById(id int, conDB *sql.DB) (*Service, error) {
	qry := "SELECT id, name, price, uom FROM ms_service WHERE id =$1"

	svc := &Service{}
	err := conDB.QueryRow(qry, id).Scan(&svc.Id, &svc.Name, &svc.Price, &svc.Uom)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("GetServiceById() : %w", err)
	}
	return svc, nil
}
