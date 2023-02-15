package models

import (
	"net/http"
	"project-1/db"
	str "project-1/struct"
	"strconv"
)

func Generate_Id_Supplier() int {
	var obj str.Generate_Id

	con := db.CreateCon()

	sqlStatement := "SELECT id_supplier FROM generate_id"

	_ = con.QueryRow(sqlStatement).Scan(&obj.Id)

	no := obj.Id
	no = no + 1

	sqlstatement := "UPDATE generate_id SET id_supplier=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return -1
	}

	stmt.Exec(no)

	return no
}

func Input_Supplier(nama_supplier string, nomor_telpon string, email_supplier string) (Response, error) {
	var res Response
	var sup str.Insert_Supplier

	con := db.CreateCon()

	nm := Generate_Id_Supplier()

	nm_str := strconv.Itoa(nm)

	kode_suplier := "SUP-" + nm_str

	sqlStatement := "INSERT INTO supplier (co,kode_supplier,nama_supplier,nomor_telpon,email_supplier) values(?,?,?,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(nm, kode_suplier, nama_supplier, nomor_telpon, email_supplier)

	sup.Nama_supplier = nama_supplier
	sup.Email_supplier = email_supplier
	sup.Nomor_telpon = nomor_telpon

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = sup

	return res, nil
}

func Read_Supplier() (Response, error) {
	var res Response
	var arr_invent []str.Read_Supplier
	var invent str.Read_Supplier

	con := db.CreateCon()

	sqlStatement := "SELECT kode_supplier,nama_supplier,nomor_telpon,email_supplier FROM supplier ORDER BY co ASC "

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&invent.Kode_supplier, &invent.Nama_supplier, &invent.Nomor_telpon, &invent.Email_Supplier)
		if err != nil {
			return res, err
		}
		arr_invent = append(arr_invent, invent)
	}

	if arr_invent == nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr_invent
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr_invent
	}

	return res, nil
}

func Delete_Supplier(kode_supplier string) (Response, error) {
	var res Response
	var arrobj []str.Read_Kode_Stock_Masuk
	var obj str.Read_Kode_Stock_Masuk

	con := db.CreateCon()

	sqlStatement := "SELECT id_stock_masuk,kode_supplier FROM stock_masuk WHERE kode_supplier=? "

	rows, err := con.Query(sqlStatement, kode_supplier)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id_stock_masuk, &obj.Kode_supplier)
		if err != nil {
			return res, err
		}
		arrobj = append(arrobj, obj)
	}

	if arrobj == nil {

		sqlstatement := "DELETE FROM supplier WHERE kode_supplier=?"

		stmt, err := con.Prepare(sqlstatement)

		if err != nil {
			return res, err
		}

		result, err := stmt.Exec(kode_supplier)

		if err != nil {
			return res, err
		}

		rowsAffected, err := result.RowsAffected()

		if err != nil {
			return res, err
		}

		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = map[string]int64{
			"rows": rowsAffected,
		}

	} else {
		res.Status = http.StatusNotFound
		res.Message = "Tidak bisa di hapus"
		res.Data = arrobj
	}

	return res, nil
}
