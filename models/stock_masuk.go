package models

import (
	"fmt"
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

func Input_Stock_Masuk(kode_supplier string, nama_penanggung_jawab string, kode_stock string, nama_stock string,
	jumlah_barang string, satuan_barang string, harga_barang string) (Response, error) {
	var res Response
	var SM str.Insert_Stock_Masuk
	var sup str.Kode_nama_sup
	var k_sup = []string{}
	var n_sup = []string{}

	con := db.CreateCon()

	nm := Generate_Id_Stock_Masuk()

	nm_str := strconv.Itoa(nm)

	id := "SM-" + nm_str

	sqlStatement := "INSERT INTO stock_masuk (id_stock_masuk,kode_supplier,kode_stock,nama_stock,tanggal_masuk,nama_penanggung_jawab,jumlah_barang,satuan_barang,harga_barang) values(?,?,?,?,CURRENT_DATE,?,?,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id, kode_supplier, kode_stock, nama_stock, nama_penanggung_jawab, jumlah_barang, satuan_barang, harga_barang)

	sqlStatement = "SELECT * FROM stock_masuk WHERE id_stock_masuk=?"

	err = con.QueryRow(sqlStatement, id).Scan(&SM.Id_stock_masuk, &SM.Kode_supplier, &SM.Nama_penanggung_jawab,
		&SM.Kode_stock, &SM.Nama_stock, &SM.Tanggal_masuk, &SM.Jumlah_barang, &SM.Harga_barang)

	fmt.Println(SM.Nama_stock)

	k_stock := String_Separator_To_String(SM.Kode_stock)

	j_barang := String_Separator_To_float64(SM.Jumlah_barang)

	n_barang := String_Separator_To_String(SM.Nama_stock)

	sqlStatement = "SELECT kode_stock,nama_barang FROM supplier WHERE kode_supplier=?"

	_ = con.QueryRow(sqlStatement, SM.Kode_supplier).Scan(&sup.Kode_stock, &sup.Nama_barang)

	kd := 0
	if sup.Kode_stock == "" {
		kd = 0
	} else {
		kd = 1
		k_sup = String_Separator_To_String(sup.Kode_stock)
		n_sup = String_Separator_To_String(sup.Nama_barang)
	}

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

		tt := 0
		if kd == 1 {
			for j := 0; j < len(k_sup); j++ {
				if k_sup[j] != k_stock[i] && n_sup[j] != n_barang[i] {
					tt++
				} else if k_sup[j] == k_stock[i] && n_sup[j] != n_barang[i] {
					n_sup[j] = n_barang[i]
				}
			}
			if tt == len(k_sup) {
				k_sup = append(k_sup, k_stock[i])
				n_sup = append(n_sup, n_barang[i])
			}
		} else {
			k_sup = append(k_sup, k_stock[i])
			n_sup = append(n_sup, n_barang[i])
		}
	}
	k_sup_str := ""
	n_sup_str := ""
	for i := 0; i < len(k_sup); i++ {
		k_sup_str += "|" + k_sup[i] + "|"
		n_sup_str += "|" + n_sup[i] + "|"
	}

	sqlstatement := "UPDATE supplier SET kode_stock=?,nama_barang=? WHERE kode_supplier=?"

	stmt, err = con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(k_sup_str, n_sup_str, SM.Kode_supplier)

	if err != nil {
		return res, err
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
	var obj_fix str.Read_Stock_Masuk_fix
	var arrobj_fix []str.Read_Stock_Masuk_fix

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM stock_masuk"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id_stock_masuk, &obj.Kode_supplier, &obj.Nama_penanggung_jawab, &obj.Kode_stock, &obj.Nama_stock,
			&obj.Tanggal_masuk, &obj.Jumlah_barang, &obj.Satuan_barang, &obj.Harga_barang)
		if err != nil {
			return res, err
		}
		arrobj = append(arrobj, obj)
	}

	for i := 0; i < len(arrobj); i++ {
		tj := 0.0
		hb := 0
		k_stock := String_Separator_To_String(arrobj[i].Kode_stock)
		j_barang := String_Separator_To_float64(arrobj[i].Jumlah_barang)
		h_barang := String_Separator_To_Int(arrobj[i].Harga_barang)
		n_barang := String_Separator_To_String(arrobj[i].Nama_stock)
		for j := 0; j < len(j_barang); j++ {
			tj += j_barang[j]
			hb += h_barang[j]
		}
		obj_fix.Id_stock_masuk = arrobj[i].Id_stock_masuk
		obj_fix.Kode_supplier = arrobj[i].Kode_supplier
		obj_fix.Nama_penanggung_jawab = arrobj[i].Nama_penanggung_jawab
		obj_fix.Tanggal_masuk = arrobj[i].Tanggal_masuk
		obj_fix.Kode_stock = k_stock
		obj_fix.Jumlah_barang = j_barang
		obj_fix.Nama_barang = n_barang
		obj_fix.Harga_barang = h_barang
		obj_fix.Total_harga_barang = hb
		obj_fix.Total_Jumlah_barang = tj
		arrobj_fix = append(arrobj_fix, obj_fix)
	}

	if arrobj_fix == nil {
		arrobj_fix = append(arrobj_fix, obj_fix)
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arrobj_fix
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arrobj_fix
	}

	return res, nil
}

func Read_Detail_Stock_Masuk(id_stock_masuk string) (Response, error) {
	var res Response
	var obj_str str.Detail_Stock_Masuk_String
	var obj str.Detail_Stock_Masuk
	var arrobj []str.Detail_Stock_Masuk

	con := db.CreateCon()

	sqlStatement := "SELECT kode_stock,nama_barang,jumlah_barang,harga_barang FROM stock_masuk WHERE id_stock_masuk=?"

	err := con.QueryRow(sqlStatement, id_stock_masuk).Scan(&obj_str.Kode_stock, &obj_str.Nama_barang, &obj_str.Jumlah_barang, &obj_str.Harga_barang)

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = obj
	}

	k_stock := String_Separator_To_String(obj_str.Kode_stock)
	j_barang := String_Separator_To_float64(obj_str.Jumlah_barang)
	h_barang := String_Separator_To_Int(obj_str.Harga_barang)
	n_barang := String_Separator_To_String(obj_str.Nama_barang)

	for i := 0; i < len(k_stock); i++ {
		obj.Kode_stock = k_stock[i]
		obj.Nama_barang = n_barang[i]
		obj.Jumlah_barang = j_barang[i]
		obj.Harga_barang = h_barang[i]
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
