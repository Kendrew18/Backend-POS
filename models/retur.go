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

func Input_Retur(id_supplier string, nama_supplier string, kode_stock string, nama_barang string, jumlah_barang float64) (Response, error) {
	var res Response
	var obj str.Insert_Retur

	con := db.CreateCon()

	nm := Generate_Id_Retur()

	nm_str := strconv.Itoa(nm)

	currentTime := time.Now()

	id := "TR-" + currentTime.Format("2006-01-02") + "-" + nm_str

	sqlStatement := "INSERT INTO retur (co,id_retur,id_supplier,nama_supplier,kode_stock,nama_barang,tanggal_retur,jumlah_barang,status_retur) values(?,?,?,?,?,?,current_date,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(nm, id, id_supplier, nama_supplier, kode_stock, nama_barang, jumlah_barang, 0)

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

	sqlStatement := "SELECT id_retur,id_supplier,nama_supplier,kode_stock,nama_barang,tanggal_retur,jumlah_barang,status_retur FROM retur ORDER BY co ASC "

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

func Read_Kode_Nama_Barang(id_supplier string) (Response, error) {
	var res Response
	var arrobj []str.Kode_nama_sup
	var obj str.Kode_nama_sup
	var obj_text str.Kode_nama_sup_text

	con := db.CreateCon()

	sqlStatement := "SELECT kode_stock,nama_barang FROM supplier WHERE kode_supplier=?"

	_ = con.QueryRow(sqlStatement, id_supplier).Scan(&obj_text.Kode_stock, &obj_text.Nama_barang)

	k_stock := String_Separator_To_String(obj_text.Kode_stock)
	n_barang := String_Separator_To_String(obj_text.Nama_barang)

	if len(k_stock) > 0 {
		for i := 0; i < len(k_stock); i++ {
			obj.Kode_stock = k_stock[i]
			obj.Nama_barang = n_barang[i]
			arrobj = append(arrobj, obj)
		}
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

func Read_Max_Jumlah(id_supplier string, kode_stock string) (Response, error) {
	var res Response
	var arrobj []str.Read_Stock_Masuk
	var obj str.Read_Stock_Masuk

	var arrobj_max []str.Read_max_reture
	var obj_max str.Read_max_reture

	con := db.CreateCon()

	sqlStatement := "SELECT id_stock_masuk,kode_supplier,nama_penanggung_jawab,kode_stock,nama_stock,tanggal_masuk,jumlah_barang,satuan_barang,harga_barang FROM stock_masuk WHERE kode_supplier=? ORDER BY `stock_masuk`.`tanggal_masuk` DESC"

	rows, _ := con.Query(sqlStatement, id_supplier)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&obj.Id_stock_masuk, &obj.Kode_supplier, &obj.Nama_penanggung_jawab, &obj.Kode_stock, &obj.Nama_stock,
			&obj.Tanggal_masuk, &obj.Jumlah_barang, &obj.Harga_barang)
		if err != nil {
			return res, err
		}
		arrobj = append(arrobj, obj)
	}

	k_stock := String_Separator_To_String(obj.Kode_stock)
	j_barang := String_Separator_To_float64(obj.Jumlah_barang)

	for i := 0; i < len(k_stock); i++ {
		if k_stock[i] == kode_stock {
			obj_max.Kode_supplier = id_supplier
			obj_max.Max_barang = j_barang[i]
			obj_max.Kode_stock = kode_stock
			i = len(k_stock)
		}
	}
	arrobj_max = append(arrobj_max, obj_max)

	if arrobj_max == nil {
		arrobj_max = append(arrobj_max, obj_max)
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arrobj_max
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arrobj_max
	}

	return res, nil
}
