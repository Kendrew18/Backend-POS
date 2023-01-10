package models

import (
	"net/http"
	"project-1/db"
	str "project-1/struct"
)

func Filter_Transaksi(tanggal_pelunasan string, tipe_status int) (Response, error) {
	var res Response
	var arrobj []str.Read_Transaksi
	var obj str.Read_Transaksi

	tgl := ""
	if tanggal_pelunasan != "" {

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

		tgl += "WHERE tanggal_penjualan=\"" + bln_thn_sql + "\""
	}

	if tipe_status != 2 {
		if tgl == "" {
			if tipe_status == 0 {
				tgl += "WHERE status_transaksi=0"
			} else if tipe_status == 1 {
				tgl += "WHERE status_transaksi=1"
			}
		} else {
			if tipe_status == 0 {
				tgl += "&& status_transaksi=0"
			} else if tipe_status == 1 {
				tgl += "&& status_transaksi=1"
			}
		}
	}

	con := db.CreateCon()

	sqlStatement := "SELECT kode_transaksi, DATE_FORMAT(tanggal_penjualan, \"%d/%m/%Y\"), DATE_FORMAT(tanggal_pelunasan, \"%d/%m/%Y\"),status_transaksi,sub_total_harga,jumlah_barang FROM transaksi " + tgl

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Kode_transaksi, &obj.Tanggal_penjualan, &obj.Tanggal_pelunasan, &obj.Status_transaksi, &obj.Sub_total_harga, &obj.Jumlah_barang)
		total := String_Separator_To_Int(obj.Jumlah_barang)
		sub_total := 0
		for i := 0; i < len(total); i++ {
			sub_total += total[i]
		}
		obj.Total_jumlah_barang = sub_total
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

func Filter_Stock(tipe_urutan int) (Response, error) {
	var res Response
	var arr_invent []str.Read_Stock
	var invent str.Read_Stock

	con := db.CreateCon()
	tgl := ""
	if tipe_urutan != 2 {

		if tipe_urutan == 0 {
			tgl += " ORDER BY kode_stock ASC"
		} else if tipe_urutan == 1 {
			tgl += " ORDER BY kode_stock DESC"
		}
	}

	sqlStatement := "SELECT * FROM stock" + tgl

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&invent.Kode_stock, &invent.Nama_barang, &invent.Jumlah_barang, &invent.Harga_barang)
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

func Filter_Stock_Masuk(tanggal_pelunasan string, tipe_tanggal int, tipe_urutan int) (Response, error) {
	var res Response
	var arrobj []str.Read_Stock_Masuk
	var obj str.Read_Stock_Masuk

	tgl := ""
	if tanggal_pelunasan != "" {

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

		if tipe_tanggal == 0 {
			tgl += " WHERE tanggal_pelunasan=" + bln_thn_sql
		} else if tipe_tanggal == 1 {
			tmp := "%" + bln_thn_sql + "%"
			tgl += " WHERE tanggal_pelunasan like " + tmp
		} else if tipe_tanggal == 2 {
			tmp := "%" + bln_thn_sql + "%"
			tgl += " WHERE tanggal_pelunasan like " + tmp
		}

	}

	if tipe_urutan != 2 {
		if tipe_urutan == 0 {
			tgl += " ORDER BY kode_transaksi ASC"
		} else if tipe_urutan == 1 {
			tgl += " ORDER BY kode_transaksi DESC"
		}
	}

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM stock_masuk"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id_stock_masuk, &obj.Kode_supplier, &obj.Nama_penanggung_jawab, &obj.Kode_stock, &obj.Nama_stock,
			&obj.Tanggal_masuk, &obj.Jumlah_barang, &obj.Harga_barang)
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

func Filter_Read_Pembukuan(tanggal string, tanggal2 string, tipe int) (Response, error) {
	var res Response
	if tipe == 1 {
		var arrobj []str.Read_Pembukuan_Transaksi
		var obj str.Read_Pembukuan_Transaksi

		var arrobj_fix []str.Read_Pembukuan_Transaksi_List
		var obj_fix str.Read_Pembukuan_Transaksi_List

		ls := []string{}
		ls2 := []string{}
		str1 := ""
		str2 := ""

		for i := 0; i < len(tanggal); i++ {
			if byte(tanggal[i]) >= 48 && byte(tanggal[i]) <= 57 {
				str1 += string(tanggal[i])
				str2 += string(tanggal2[i])
				if i == len(tanggal)-1 {
					ls = append(ls, str1)
					ls2 = append(ls2, str2)
				}
			} else if tanggal[i] == '-' {
				ls = append(ls, str1)
				ls2 = append(ls2, str2)
				str1 = ""
				str2 = ""
			}
		}

		j := len(ls)
		bln_thn_sql := ""
		bln_thn_sql2 := ""

		for x := j - 1; x >= 0; x-- {
			bln_thn_sql += ls[x]
			bln_thn_sql2 += ls2[x]
			if x != 0 {
				bln_thn_sql += "-"
				bln_thn_sql2 += "-"
			}
		}

		con := db.CreateCon()

		sqlStatement := "SELECT id_pembukuan_transaksi,kode_stock,nama_barang,jumlah_barang,harga_barang,Date_Format(tanggal_pelunasan,\"%d-%m-%Y\"),total_harga_penjualan FROM pembukuan_transaksi WHERE tanggal_pelunasan>=? && tanggal_pelunasan<=? "

		rows, err := con.Query(sqlStatement, bln_thn_sql, bln_thn_sql2)

		defer rows.Close()

		if err != nil {
			return res, err
		}

		for rows.Next() {
			err = rows.Scan(&obj.Id_pembukuan_transaksi, &obj.Kode_stock, &obj.Nama_barang, &obj.Jumlah_barang,
				&obj.Harga_barang, &obj.Tanggal_pelunasan, &obj.Total_harga_penjualan)
			if err != nil {
				return res, err
			}
			arrobj = append(arrobj, obj)
		}

		for i := 0; i < len(arrobj); i++ {
			obj_fix.Id_pembukuan_transaksi = arrobj[i].Id_pembukuan_transaksi
			obj_fix.Tanggal_pelunasan = arrobj[i].Tanggal_pelunasan
			obj_fix.Total_harga_penjualan = arrobj[i].Total_harga_penjualan
			obj_fix.Nama_barang = String_Separator_To_String(arrobj[i].Nama_barang)
			obj_fix.Kode_stock = String_Separator_To_String(arrobj[i].Kode_stock)
			obj_fix.Jumlah_barang = String_Separator_To_Int(arrobj[i].Jumlah_barang)
			obj_fix.Harga_barang = String_Separator_To_Int(arrobj[i].Harga_barang)
			arrobj_fix = append(arrobj_fix, obj_fix)
		}

		if arrobj_fix == nil {
			res.Status = http.StatusNotFound
			res.Message = "Not Found"
			res.Data = arrobj_fix
		} else {
			res.Status = http.StatusOK
			res.Message = "Sukses"
			res.Data = arrobj_fix
		}
	} else if tipe == 2 {

		var arrobj []str.Read_Pembukuan_Transaksi_Bulanan
		var obj str.Read_Pembukuan_Transaksi_Bulanan

		var arrobj_fix []str.Read_Pembukuan_Transaksi_Bulanan_List
		var obj_fix str.Read_Pembukuan_Transaksi_Bulanan_List

		ls := []string{}
		str1 := ""

		for i := 0; i < len(tanggal); i++ {
			if byte(tanggal[i]) >= 48 && byte(tanggal[i]) <= 57 {
				str1 += string(tanggal[i])
				if i == len(tanggal)-1 {
					ls = append(ls, str1)
				}
			} else if tanggal[i] == '-' {
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

		sqlStatement := "SELECT id_pembukuan_transaksi_bulanan,kode_stock,nama_barang,jumlah_barang,harga_barang,Date_Format(tanggal_pelunasan,\"%d-%m-%Y\"),total_harga_penjualan FROM pembukuan_transaksi_bulanan WHERE Date_Format(tanggal_pelunasan,\"%Y-%m\")=?"

		rows, err := con.Query(sqlStatement, bln_thn_sql)

		defer rows.Close()

		if err != nil {
			return res, err
		}

		for rows.Next() {
			err = rows.Scan(&obj.Id_pembukuann_transaksi, &obj.Kode_stock, &obj.Nama_barang, &obj.Jumlah_barang,
				&obj.Harga_barang, &obj.Tanggal_pelunasan, &obj.Total_harga_penjualan)
			if err != nil {
				return res, err
			}
			arrobj = append(arrobj, obj)
		}

		for i := 0; i < len(arrobj); i++ {
			obj_fix.Id_pembukuann_transaksi = arrobj[i].Id_pembukuann_transaksi
			obj_fix.Tanggal_pelunasan = arrobj[i].Tanggal_pelunasan
			obj_fix.Total_harga_penjualan = arrobj[i].Total_harga_penjualan
			obj_fix.Nama_barang = String_Separator_To_String(arrobj[i].Nama_barang)
			obj_fix.Kode_stock = String_Separator_To_String(arrobj[i].Kode_stock)
			obj_fix.Jumlah_barang = String_Separator_To_Int(arrobj[i].Jumlah_barang)
			obj_fix.Harga_barang = String_Separator_To_Int(arrobj[i].Harga_barang)
			arrobj_fix = append(arrobj_fix, obj_fix)
		}

		if arrobj_fix == nil {
			res.Status = http.StatusNotFound
			res.Message = "Not Found"
			res.Data = arrobj_fix
		} else {
			res.Status = http.StatusOK
			res.Message = "Sukses"
			res.Data = arrobj_fix
		}

	} else if tipe == 3 {

		var arrobj []str.Read_Pembukuan_Transaksi_Tahunan
		var obj str.Read_Pembukuan_Transaksi_Tahunan

		var arrobj_fix []str.Read_Pembukuan_Transaksi_Tahunan_List
		var obj_fix str.Read_Pembukuan_Transaksi_Tahunan_List

		ls := []string{}
		str1 := ""

		for i := 0; i < len(tanggal); i++ {
			if byte(tanggal[i]) >= 48 && byte(tanggal[i]) <= 57 {
				str1 += string(tanggal[i])
				if i == len(tanggal)-1 {
					ls = append(ls, str1)
				}
			} else if tanggal[i] == '-' {
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

		sqlStatement := "SELECT id_pembukuan_transaksi_tahunan,kode_stock,nama_barang,jumlah_barang,harga_barang,Date_Format(tanggal_pelunasan,\"%d-%m-%Y\"),total_harga_penjualan FROM pembukuan_transaksi_tahunan WHERE Date_Format(tanggal_pelunasan,\"%Y\")=?"

		rows, err := con.Query(sqlStatement, bln_thn_sql)

		defer rows.Close()

		if err != nil {
			return res, err
		}

		for rows.Next() {
			err = rows.Scan(&obj.Id_pembukuann_transaksi, &obj.Kode_stock, &obj.Nama_barang, &obj.Jumlah_barang,
				&obj.Harga_barang, &obj.Tanggal_pelunasan, &obj.Total_harga_penjualan)
			if err != nil {
				return res, err
			}
			arrobj = append(arrobj, obj)
		}

		for i := 0; i < len(arrobj); i++ {
			obj_fix.Id_pembukuann_transaksi = arrobj[i].Id_pembukuann_transaksi
			obj_fix.Tanggal_pelunasan = arrobj[i].Tanggal_pelunasan
			obj_fix.Total_harga_penjualan = arrobj[i].Total_harga_penjualan
			obj_fix.Nama_barang = String_Separator_To_String(arrobj[i].Nama_barang)
			obj_fix.Kode_stock = String_Separator_To_String(arrobj[i].Kode_stock)
			obj_fix.Jumlah_barang = String_Separator_To_Int(arrobj[i].Jumlah_barang)
			obj_fix.Harga_barang = String_Separator_To_Int(arrobj[i].Harga_barang)
			arrobj_fix = append(arrobj_fix, obj_fix)
		}

		if arrobj_fix == nil {
			res.Status = http.StatusNotFound
			res.Message = "Not Found"
			res.Data = arrobj_fix
		} else {
			res.Status = http.StatusOK
			res.Message = "Sukses"
			res.Data = arrobj_fix
		}

	}
	return res, nil
}
