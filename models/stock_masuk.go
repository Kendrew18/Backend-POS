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

func Input_Stock_Masuk(kode_supplier string, kode_stock string, nama_supplier string, jumlah_barang string, harga_barang string) (Response, error) {
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

	k_stock := String_Separator_To_String(SM.Kode_stock)

	j_barang := String_Separator_To_Int(SM.Jumlah_barang)

	for i := 0; i < len(k_stock); i++ {
		var obj str.Jumlah_Barang

		sqlStatement = "SELECT jumlah_barang FROM stock WHERE kode_stock=?"
		_ = con.QueryRow(sqlStatement, k_stock[i]).Scan(&obj.Jumlah_Barang)

		total := obj.Jumlah_Barang + j_barang[i]

		sqlstatement := "UPDATE stock SET jumlah_barang=? WHERE kode_stock=?"

		stmt, err = con.Prepare(sqlstatement)

		if err != nil {
			return res, err
		}

		_, err := stmt.Exec(total, k_stock[i])

		if err != nil {
			return res, err
		}
	}

	stmt.Close()

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

	sqlStatement := "SELECT * FROM stock_masuk"

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

func Read_Detail_Stock_Masuk(id_stock_masuk string) (Response, error) {
	var res Response
	var obj_str str.Detail_Stock_Masuk_String
	var obj str.Detail_Stock_Masuk
	var arrobj []str.Detail_Stock_Masuk

	con := db.CreateCon()

	sqlStatement := "SELECT kode_stock,jumlah_barang,harga_barang FROM stock_masuk WHERE id_stock_masuk=?"

	err := con.QueryRow(sqlStatement, id_stock_masuk).Scan(&obj_str.Kode_stock, &obj_str.Jumlah_barang, &obj_str.Harga_barang)

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = obj
	}

	k_stock := String_Separator_To_String(obj_str.Kode_stock)
	j_barang := String_Separator_To_Int(obj_str.Jumlah_barang)
	h_barang := String_Separator_To_Int(obj_str.Harga_barang)

	for i := 0; i < len(k_stock); i++ {
		obj.Kode_stock = k_stock[i]
		obj.Jumlah_barang = j_barang[i]
		obj.Harga_barang = h_barang[i]
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
