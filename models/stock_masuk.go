package models

import (
	"net/http"
	"project-1/db"
	str "project-1/struct"
	"strconv"
)

func Generate_Id_Stock_Masuk() int {
	var obj str.Generate_Id

	con := db.CreateCon()

	sqlStatement := "SELECT id_stock_masuk FROM generate_id"

	_ = con.QueryRow(sqlStatement).Scan(&obj.Id)

	no := obj.Id
	no = no + 1

	sqlstatement := "UPDATE generate_id SET id_stock_masuk=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return -1
	}

	stmt.Exec(no)

	return no
}

func Input_Stock_Masuk(kode_supplier string, kode_stock string, nama_supplier string, jumlah_barang int, harga_barang int) (Response, error) {
	var res Response
	var SM str.Insert_Stock_Masuk

	con := db.CreateCon()

	nm := Generate_Id_Stock_Masuk()

	nm_str := strconv.Itoa(nm)

	id := "SM-" + nm_str

	sqlStatement := "INSERT INTO stock_masuk (id_stock_masuk,kode_supplier,kode_stock,tanggal_masuk,nama_supplier,jumlah_barang,harga_barang) values(?,?,?,CURRENT_DATE,?,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id, kode_supplier, kode_stock, nama_supplier, jumlah_barang, harga_barang)

	sqlStatement = "SELECT * FROM stock_masuk WHERE id_stock_masuk=? "

	_ = con.QueryRow(sqlStatement, id).Scan(&SM.Id_stock_masuk, &SM.Kode_supplier, &SM.Kode_stock,
		&SM.Tanggal_masuk, &SM.Nama_supplier, &SM.Jumlah_barang, &SM.Harga_barang)

	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = SM

	return res, nil
}

func Read_Stock_Masuk() (Response, error) {
	var res Response
	var arrobj []str.Read_Stock_Masuk
	var obj str.Read_Stock_Masuk

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM inventory_stock"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id_stock_masuk, &obj.Kode_supplier, &obj.Kode_stock,
			&obj.Tanggal_masuk, &obj.Nama_supplier, &obj.Jumlah_barang, &obj.Harga_barang)
		if err != nil {
			return res, err
		}
		arrobj = append(arrobj, obj)
	}

	if arrobj == nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arrobj
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arrobj
	}

	return res, nil
}
