package trxrepo

import (
	"database/sql"
	"fmt"
)

func AddTransaction(dbCon *sql.DB, trx *TransactionHeader) error {
	tx, err := dbCon.Begin()
	if err != nil {
		return fmt.Errorf("AddTransaction() Begin : %w", err)
	}

	qry := "INSERT INTO tr_header(start_date, end_date, cust_name, phone_no) VALUES($1, $2, $3, $4) RETURNING no"

	err = tx.QueryRow(qry, trx.StartDate, trx.EndDate, trx.CustName, trx.Phone).Scan(&trx.No)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("AddTransaction() Header : %w", err)
	}

	qry = "INSERT INTO tr_detail(trx_no,service_id, service_name, qty, uom, price) VALUES($1, $2, $3, $4, $5, $6)"
	for _, det := range trx.ArrDetail {
		_, err := tx.Exec(qry, trx.No, det.ServiceName, det.Qty, det.Uom, det.Price)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("AddTransaction() Detail : %w", err)
		}
	}

	tx.Commit()

	return nil
}
func ViewTransaction(dbCon *sql.DB, trxNo int) (*TransactionHeader, error) {
	trx := &TransactionHeader{}

	headerQuery := "SELECT no, start_date, end_date, cust_name, phone_no FROM tr_header WHERE no = $1"
	err := dbCon.QueryRow(headerQuery, trxNo).Scan(&trx.No, &trx.StartDate, &trx.EndDate, &trx.CustName, &trx.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("ViewTransaction() Header: No transaction found with the given transaction number")
		}
		return nil, fmt.Errorf("ViewTransaction() Header: %w", err)
	}

	detailQuery := "SELECT service_name, qty, uom, price FROM tr_detail WHERE trx_no = $1"
	rows, err := dbCon.Query(detailQuery, trxNo)
	if err != nil {
		return nil, fmt.Errorf("ViewTransaction() Detail: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		detail := &TransactionDetail{}
		err := rows.Scan(&detail.ServiceName, &detail.Qty, &detail.Uom, &detail.Price)
		if err != nil {
			return nil, fmt.Errorf("ViewTransaction() Detail: %w", err)
		}
		trx.ArrDetail = append(trx.ArrDetail, *detail)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("ViewTransaction() Detail: %w", err)
	}

	return trx, nil
}
