package models

import (
	"net/http"
	"project-1/db"
	str "project-1/struct"
	"strconv"
	"time"
)

func Generate_Id_Transaksi() int {
	var obj str.Generate_Id

	con := db.CreateCon()

	sqlStatement := "SELECT id_transaksi FROM generate_id"

	_ = con.QueryRow(sqlStatement).Scan(&obj.Id)

	no := obj.Id
	no = no + 1

	sqlstatement := "UPDATE generate_id SET id_transaksi=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return -1
	}

	stmt.Exec(no)

	return no
}

func Input_Transaksi(kode_stock string, nama_barang string, jumlah_barang string, harga_barang string, status_transaksi string, tanggal_pelunasan string, sub_total_harga int64) (Response, error) {
	var res Response
	var tr str.Input_Transaksi

	con := db.CreateCon()

	nm := Generate_Id_Transaksi()

	nm_str := strconv.Itoa(nm)

	currentTime := time.Now()

	id := "TR-" + currentTime.Format("20060102") + nm_str

	ls := []string{}
	str1 := ""

	for i := 0; i < len(tanggal_pelunasan); i++ {
		if byte(tanggal_pelunasan[i]) >= 48 && byte(tanggal_pelunasan[i]) <= 57 {
			str1 += string(tanggal_pelunasan[i])
			if i == len(tanggal_pelunasan)-1 {
				ls = append(ls, str1)
			}
		} else if tanggal_pelunasan[i] == '-' {
			ls = append(ls, str1)
			str1 = ""
		}
	}

	j := len(ls)
	bln_thn_sql := ""

	for x := j - 1; x >= 0; x-- {
		bln_thn_sql += ls[x]
		if x != 0 {
			bln_thn_sql += "-"
		}
	}

	if status_transaksi == "0" {

		sqlStatement := "INSERT INTO transaksi (kode_transaksi,kode_stock,nama_barang,jumlah_barang,harga_barang,tanggal_penjualan,tanggal_pelunasan,status_transaksi,sub_total_harga) values(?,?,?,?,?,CURRENT_DATE,?,?,?)"

		stmt, err := con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		_, err = stmt.Exec(id, kode_stock, nama_barang, jumlah_barang, harga_barang, bln_thn_sql, 0, sub_total_harga)

		sqlStatement = "SELECT kode_stock,jumlah_barang,harga_barang,tanggal_penjualan,tanggal_pelunasan FROM stock_masuk WHERE id_stock_masuk=? "

		_ = con.QueryRow(sqlStatement, id).Scan(&tr.Kode_stock, &tr.Jumlah_barang, &tr.Harga_barang,
			&tr.Tanggal_penjualan)

		k_stock := String_Separator_To_String(kode_stock)

		j_barang := String_Separator_To_Int(jumlah_barang)

		for i := 0; i < len(k_stock); i++ {
			var obj str.Jumlah_Barang

			sqlStatement = "SELECT jumlah_barang FROM stock WHERE kode_stock=?"
			_ = con.QueryRow(sqlStatement, k_stock[i]).Scan(&obj.Jumlah_Barang)

			total := obj.Jumlah_Barang - j_barang[i]

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

	} else if status_transaksi == "1" {
		sqlStatement := "INSERT INTO transaksi (kode_transaksi,kode_stock,nama_barang,jumlah_barang,harga_barang,tanggal_penjualan,tanggal_pelunasan,status_transaksi,sub_total_harga) values(?,?,?,?,?,CURRENT_DATE,?,?,?)"

		stmt, err := con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		_, err = stmt.Exec(id, kode_stock, nama_barang, jumlah_barang, harga_barang, bln_thn_sql, 1, sub_total_harga)

		sqlStatement = "SELECT kode_stock,jumlah_barang,harga_barang,tanggal_penjualan,tanggal_pelunasan FROM stock_masuk WHERE id_stock_masuk=? "

		_ = con.QueryRow(sqlStatement, id).Scan(&tr.Kode_stock, &tr.Jumlah_barang, &tr.Harga_barang,
			&tr.Tanggal_penjualan)

		k_stock := String_Separator_To_String(kode_stock)

		j_barang := String_Separator_To_Int(jumlah_barang)

		for i := 0; i < len(k_stock); i++ {
			var obj str.Jumlah_Barang

			sqlStatement = "SELECT jumlah_barang FROM stock WHERE kode_stock=?"
			_ = con.QueryRow(sqlStatement, k_stock[i]).Scan(&obj.Jumlah_Barang)

			total := obj.Jumlah_Barang - j_barang[i]

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

	}

	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = tr

	return res, nil

}

func Read_Transaksi() (Response, error) {
	var res Response
	var arrobj []str.Read_Transaksi
	var obj str.Read_Transaksi

	con := db.CreateCon()

	sqlStatement := "SELECT kode_transaksi, DATE_FORMAT(tanggal_penjualan, \"%d/%m/%Y\"), DATE_FORMAT(tanggal_pelunasan, \"%d/%m/%Y\"),status_transaksi FROM transaksi"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Kode_transaksi, &obj.Tanggal_penjualan, &obj.Tanggal_pelunasan, &obj.Status_transaksi)
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

func Read_Detail_transaksi(kode_transaksi string) (Response, error) {
	var res Response
	var obj_str str.Detail_Stock_Masuk_String
	var obj str.Detail_Stock_Masuk
	var arrobj []str.Detail_Stock_Masuk

	con := db.CreateCon()

	sqlStatement := "SELECT kode_stock,nama_barang,jumlah_barang,harga_barang FROM transaksi WHERE kode_transaksi=?"

	err := con.QueryRow(sqlStatement, kode_transaksi).Scan(&obj_str.Kode_stock, &obj_str.Jumlah_barang, &obj_str.Harga_barang)

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = obj
	} else {

		k_stock := String_Separator_To_String(obj_str.Kode_stock)
		j_barang := String_Separator_To_Int(obj_str.Jumlah_barang)
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
			res.Status = http.StatusNotFound
			res.Message = "Not Found"
			res.Data = arrobj
		} else {
			res.Status = http.StatusOK
			res.Message = "Sukses"
			res.Data = arrobj
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
	}

	return res, nil
}

func Update_Status(kode_transaksi string, tanggal_pelunasan string) (Response, error) {
	var res Response

	ls := []string{}
	str1 := ""

	for i := 0; i < len(tanggal_pelunasan); i++ {
		if byte(tanggal_pelunasan[i]) >= 48 && byte(tanggal_pelunasan[i]) <= 57 {
			str1 += string(tanggal_pelunasan[i])
			if i == len(tanggal_pelunasan)-1 {
				ls = append(ls, str1)
			}
		} else if tanggal_pelunasan[i] == '-' {
			ls = append(ls, str1)
			str1 = ""
		}
	}

	j := len(ls)
	bln_thn_sql := ""

	for x := j - 1; x >= 0; x-- {
		bln_thn_sql += ls[x]
		if x != 0 {
			bln_thn_sql += "-"
		}
	}

	con := db.CreateCon()

	sqlstatement := "UPDATE transaksi SET tanggal_pelunasan=?, status_transaksi=? WHERE kode_transaksi=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(tanggal_pelunasan, 1, kode_transaksi)

	if err != nil {
		return res, err
	}

	rowschanged, err := result.RowsAffected()

	if err != nil {
		return res, err
	}

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Suksess"
	res.Data = map[string]int64{
		"rows": rowschanged,
	}

	return res, nil
}
