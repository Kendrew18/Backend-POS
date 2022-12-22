package models

import (
	"net/http"
	"project-1/db"
	str "project-1/struct"
	"strconv"
	"time"
)

func Generate_Id_Retur() int {
	var obj str.Generate_Id

	con := db.CreateCon()

	sqlStatement := "SELECT id_retur FROM generate_id"

	_ = con.QueryRow(sqlStatement).Scan(&obj.Id)

	no := obj.Id
	no = no + 1

	sqlstatement := "UPDATE generate_id SET id_retur=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return -1
	}

	stmt.Exec(no)

	return no
}

func Input_Retur(id_supplier string, nama_supplier string, kode_stock string, nama_barang string, jumlah_barang int) (Response, error) {
	var res Response

	var obj str.Insert_Retur

	con := db.CreateCon()

	nm := Generate_Id_Retur()

	nm_str := strconv.Itoa(nm)

	currentTime := time.Now()

	id := "TR-" + currentTime.Format("2006-01-02") + nm_str

	sqlStatement := "INSERT INTO retur (id_retur,id_supplier,nama_supplier,kode_stock,nama_barang,tanggal_retur,jumlah_barang,status_retur) values(?,?,?,?,?,current_date,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id, id_supplier, nama_supplier, kode_stock, nama_barang, jumlah_barang, 0)

	obj.Id_supplier = id_supplier
	obj.Nama_supplier = nama_supplier
	obj.Kode_stock = kode_stock
	obj.Nama_barang = nama_barang
	obj.Tanggal_retur = currentTime.Format("2006-01-02")
	obj.Jumlah_barang = jumlah_barang
	obj.Status_retur = 0

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = obj

	return res, nil
}

func Read_Retur() (Response, error) {
	var res Response
	var arrobj []str.Read_Retur
	var obj str.Read_Retur

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM retur"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id_retur, &obj.Id_supplier, &obj.Nama_supplier, &obj.Kode_stock,
			&obj.Nama_barang, &obj.Tanggal_retur, &obj.Jumlah_barang, &obj.Status_retur)
		if err != nil {
			return res, err
		}
		arrobj = append(arrobj, obj)
	}

	if arrobj == nil {
		arrobj = append(arrobj, obj)
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
