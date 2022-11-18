package models

import (
	"net/http"
	"project-1/db"
	_struct "project-1/struct"
	"strconv"
)

func Penutupan_Pembukuan(tanggal string) (Response, error) {
	var res Response
	var obj _struct.Read_Pembukuan_Transaksi
	var obj_str _struct.Detail_Stock_Masuk_String
	var arrobj_str []_struct.Detail_Stock_Masuk_String

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

	id := "PEM-"

	for x := j - 1; x >= 0; x-- {
		bln_thn_sql += ls[x]
		id += ls[x]
		if x != 0 {
			bln_thn_sql += "-"
		}
	}

	con := db.CreateCon()

	sqlStatement := "SELECT kode_stock,nama_barang,jumlah_barang,harga_barang FROM transaksi WHERE tanggal_pelunasan=?"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj_str.Kode_stock, &obj_str.Nama_barang, &obj_str.Jumlah_barang, &obj_str.Harga_barang)
		if err != nil {
			return res, err
		}
		arrobj_str = append(arrobj_str, obj_str)
	}

	var k_stk_all []string
	var n_brg_all []string
	var j_brg_all []int
	var h_brg_all []int64

	for i := 0; i < len(arrobj_str); i++ {
		k_stock := String_Separator_To_String(arrobj_str[i].Kode_stock)
		j_barang := String_Separator_To_Int(arrobj_str[i].Jumlah_barang)
		h_barang := String_Separator_To_Int(arrobj_str[i].Harga_barang)
		n_barang := String_Separator_To_String(arrobj_str[i].Nama_barang)

		for j := 0; j < len(k_stock); j++ {
			if len(k_stk_all) == 0 {

				co := 0

				for k := 0; k < len(k_stk_all); k++ {
					if k_stk_all[k] == k_stock[j] && n_brg_all[k] == n_barang[j] {
						j_brg_all[k] += j_barang[j]
						h := int64(h_barang[j])
						h_brg_all[k] += h
						co++
					}
				}

				if co == 0 {
					k_stk_all = append(k_stk_all, k_stock[j])
					n_brg_all = append(n_brg_all, n_barang[j])
					j_brg_all = append(j_brg_all, j_barang[j])
					h := int64(h_barang[j])
					h_brg_all = append(h_brg_all, h)
				}

			} else {
				k_stk_all = append(k_stk_all, k_stock[j])
				n_brg_all = append(n_brg_all, n_barang[j])
				j_brg_all = append(j_brg_all, j_barang[j])
				h := int64(h_barang[j])
				h_brg_all = append(h_brg_all, h)
			}
		}
	}

	var k_stk_pmbk string
	var n_brg_pmbk string
	var j_brg_pmbk string
	var h_brg_pmbk string

	for i := 0; i < len(k_stk_all); i++ {
		k_stk_pmbk += "|" + k_stk_all[i] + "|"
		n_brg_pmbk += "|" + n_brg_all[i] + "|"
		str := strconv.Itoa(j_brg_all[i])
		j_brg_pmbk += "|" + str + "|"
		s := strconv.FormatInt(h_brg_all[i], 10)
		h_brg_pmbk += "|" + s + "|"
	}

	var total int64

	for i := 0; i < len(h_brg_all); i++ {
		total += h_brg_all[i]
	}

	sqlStatement = "INSERT INTO pembukuan_transaksi (id_pembukuan_transaksi,kode_stock,nama_barang,jumlah_barang,harga_barang,tanggal_penjualan,total_harga_penjualan) values(?,?,?,?,?,current_date,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id, k_stk_pmbk, n_brg_pmbk, j_brg_pmbk, h_brg_pmbk, total)

	sqlStatement = "SELECT * FROM pembukuan_transaksi WHERE tanggal_penjualan=?"

	_ = con.QueryRow(sqlStatement, bln_thn_sql).Scan(&obj.Id_pembukuan_transaksi, &obj.Kode_stock,
		&obj.Nama_barang, &obj.Jumlah_barang, &obj.Harga_barang, &obj.Tanggal_penjualan, &obj.Total_harga_penjualan)

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = obj

	return res, nil
}

func Read_Pembukuan(tanggal string) (Response, error) {
	var res Response
	var arrobj []_struct.Read_Pembukuan_Transaksi
	var obj _struct.Read_Pembukuan_Transaksi

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM pembukuan_transaksi"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id_pembukuan_transaksi, &obj.Kode_stock, &obj.Nama_barang, &obj.Jumlah_barang,
			&obj.Harga_barang, &obj.Tanggal_penjualan, &obj.Total_harga_penjualan)
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
