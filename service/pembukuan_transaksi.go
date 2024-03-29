package models

// import (
// 	"net/http"
// 	"project-1/db"
// 	_struct "project-1/struct"
// 	"strconv"
// )

// func Penutupan_Pembukuan(tanggal string) (Response, error) {
// 	var res Response
// 	var obj _struct.Read_Pembukuan_Transaksi
// 	var obj_bln _struct.Read_Pembukuan_Transaksi_Bulanan
// 	var obj_thn _struct.Read_Pembukuan_Transaksi_Tahunan
// 	var obj_str _struct.Detail_Stock_Masuk_String
// 	var arrobj_str []_struct.Detail_Stock_Masuk_String

// 	ls := []string{}
// 	str1 := ""

// 	for i := 0; i < len(tanggal); i++ {
// 		if byte(tanggal[i]) >= 48 && byte(tanggal[i]) <= 57 {
// 			str1 += string(tanggal[i])
// 			if i == len(tanggal)-1 {
// 				ls = append(ls, str1)
// 			}
// 		} else if tanggal[i] == '-' {
// 			ls = append(ls, str1)
// 			str1 = ""
// 		}
// 	}

// 	j := len(ls)
// 	bln_thn_sql := ""

// 	id := "PEM-"

// 	for x := j - 1; x >= 0; x-- {
// 		bln_thn_sql += ls[x]
// 		id += ls[x]
// 		if x != 0 {
// 			bln_thn_sql += "-"
// 		}
// 	}

// 	con := db.CreateCon()

// 	sqlStatement := "SELECT kode_stock,nama_barang,jumlah_barang,harga_barang,satuan_barang FROM transaksi WHERE tanggal_pelunasan=?"

// 	rows, err := con.Query(sqlStatement, bln_thn_sql)

// 	defer rows.Close()

// 	if err != nil {
// 		return res, err
// 	}

// 	for rows.Next() {
// 		err = rows.Scan(&obj_str.Kode_stock, &obj_str.Nama_barang, &obj_str.Jumlah_barang, &obj_str.Harga_barang)
// 		if err != nil {
// 			return res, err
// 		}
// 		arrobj_str = append(arrobj_str, obj_str)
// 	}

// 	if arrobj_str != nil {

// 		var k_stk_all []string
// 		var n_brg_all []string
// 		var j_brg_all []float64
// 		var h_brg_all []int64
// 		var s_brg_all []string

// 		for i := 0; i < len(arrobj_str); i++ {
// 			k_stock := String_Separator_To_String(arrobj_str[i].Kode_stock)
// 			j_barang := String_Separator_To_float64(arrobj_str[i].Jumlah_barang)
// 			h_barang := String_Separator_To_Int(arrobj_str[i].Harga_barang)
// 			n_barang := String_Separator_To_String(arrobj_str[i].Nama_barang)
// 			s_barang := String_Separator_To_String(arrobj_str[i].Satuan_barang)

// 			for j := 0; j < len(k_stock); j++ {
// 				if len(k_stk_all) == 0 {

// 					co := 0

// 					for k := 0; k < len(k_stk_all); k++ {
// 						if k_stk_all[k] == k_stock[j] && n_brg_all[k] == n_barang[j] {
// 							j_brg_all[k] += j_barang[j]
// 							h := int64(h_barang[j])
// 							h_brg_all[k] += h
// 							co++
// 						}
// 					}

// 					if co == 0 {
// 						k_stk_all = append(k_stk_all, k_stock[j])
// 						n_brg_all = append(n_brg_all, n_barang[j])
// 						j_brg_all = append(j_brg_all, j_barang[j])
// 						h := int64(h_barang[j])
// 						h_brg_all = append(h_brg_all, h)
// 						s_brg_all = append(s_brg_all, s_barang[j])
// 					}

// 				} else {
// 					k_stk_all = append(k_stk_all, k_stock[j])
// 					n_brg_all = append(n_brg_all, n_barang[j])
// 					j_brg_all = append(j_brg_all, j_barang[j])
// 					h := int64(h_barang[j])
// 					h_brg_all = append(h_brg_all, h)
// 					s_brg_all = append(s_brg_all, s_barang[j])
// 				}
// 			}
// 		}

// 		var k_stk_pmbk string
// 		var n_brg_pmbk string
// 		var j_brg_pmbk string
// 		var h_brg_pmbk string
// 		var s_brg_pmbk string

// 		for i := 0; i < len(k_stk_all); i++ {
// 			k_stk_pmbk += "|" + k_stk_all[i] + "|"
// 			n_brg_pmbk += "|" + n_brg_all[i] + "|"
// 			str := strconv.FormatFloat(j_brg_all[i], 'E', -1, 32)
// 			j_brg_pmbk += "|" + str + "|"
// 			s := strconv.FormatInt(h_brg_all[i], 10)
// 			h_brg_pmbk += "|" + s + "|"
// 			s_brg_pmbk += "|" + s_brg_all[i] + "|"
// 		}

// 		var total int64

// 		for i := 0; i < len(h_brg_all); i++ {
// 			total += h_brg_all[i]
// 		}

// 		sqlStatement = "INSERT INTO pembukuan_transaksi (id_pembukuan_transaksi,kode_stock,nama_barang,jumlah_barang,harga_barang,tanggal_pelunasan,total_harga_penjualan,satuan_barang) values(?,?,?,?,?,?,?,?)"

// 		stmt, err := con.Prepare(sqlStatement)

// 		if err != nil {
// 			return res, err
// 		}

// 		_, err = stmt.Exec(id, k_stk_pmbk, n_brg_pmbk, j_brg_pmbk, h_brg_pmbk, bln_thn_sql, total, s_brg_pmbk)

// 		bln := ls[2] + "-" + ls[1]

// 		sqlStatement = "SELECT * FROM pembukuan_transaksi_bulanan WHERE DATE_FORMAT(tanggal_pelunasan, \"%Y-%m\")=?"

// 		_ = con.QueryRow(sqlStatement, bln).Scan(&obj_bln.Id_pembukuan_transaksi, &obj_bln.Kode_stock,
// 			&obj_bln.Nama_barang, &obj_bln.Jumlah_barang, &obj_bln.Satuan_barang, &obj_bln.Harga_barang, &obj_bln.Tanggal_pelunasan,
// 			&obj_bln.Total_harga_penjualan)

// 		if obj_bln.Tanggal_pelunasan != bln_thn_sql {

// 			if obj_bln.Id_pembukuan_transaksi == "" {

// 				sqlStatement := "SELECT kode_stock,nama_barang,jumlah_barang,harga_barang,satuan_barang FROM pembukuan_transaksi WHERE tanggal_pelunasan=?"

// 				_ = con.QueryRow(sqlStatement, bln_thn_sql).Scan(&obj_str.Kode_stock, &obj_str.Nama_barang, &obj_str.Jumlah_barang, &obj_str.Harga_barang, &obj_str.Satuan_barang)

// 				var k_stk_all []string
// 				var n_brg_all []string
// 				var j_brg_all []float64
// 				var h_brg_all []int64
// 				var s_brg_all []string

// 				k_stock := String_Separator_To_String(obj_str.Kode_stock)
// 				j_barang := String_Separator_To_float64(obj_str.Jumlah_barang)
// 				h_barang := String_Separator_To_Int(obj_str.Harga_barang)
// 				n_barang := String_Separator_To_String(obj_str.Nama_barang)
// 				s_barang := String_Separator_To_String(obj_str.Satuan_barang)

// 				for j := 0; j < len(k_stock); j++ {
// 					k_stk_all = append(k_stk_all, k_stock[j])
// 					n_brg_all = append(n_brg_all, n_barang[j])
// 					j_brg_all = append(j_brg_all, j_barang[j])
// 					h := int64(h_barang[j])
// 					h_brg_all = append(h_brg_all, h)
// 					s_brg_all = append(s_brg_all, s_barang[j])
// 				}

// 				var k_stk_pmbk string
// 				var n_brg_pmbk string
// 				var j_brg_pmbk string
// 				var h_brg_pmbk string
// 				var s_brg_pmbk string

// 				for i := 0; i < len(k_stk_all); i++ {
// 					k_stk_pmbk += "|" + k_stk_all[i] + "|"
// 					n_brg_pmbk += "|" + n_brg_all[i] + "|"
// 					str := strconv.FormatFloat(j_brg_all[i], 'E', -1, 32)
// 					j_brg_pmbk += "|" + str + "|"
// 					s := strconv.FormatInt(h_brg_all[i], 10)
// 					h_brg_pmbk += "|" + s + "|"
// 					s_brg_pmbk += "|" + s_brg_all[i] + "|"
// 				}

// 				var total int64

// 				for i := 0; i < len(h_brg_all); i++ {
// 					total += h_brg_all[i]
// 				}

// 				id_bln := "PEM-BLN-" + bln

// 				sqlStatement = "INSERT INTO pembukuan_transaksi_bulanan (id_pembukuan_transaksi_bulanan,kode_stock,nama_barang,jumlah_barang,harga_barang,tanggal_pelunasan,total_harga_penjualan,satuan_barang) values(?,?,?,?,?,?,?,?)"

// 				stmt, err := con.Prepare(sqlStatement)

// 				if err != nil {
// 					return res, err
// 				}

// 				_, err = stmt.Exec(id_bln, k_stk_pmbk, n_brg_pmbk, j_brg_pmbk, h_brg_pmbk, bln_thn_sql, total, s_brg_pmbk)

// 			} else {

// 				var obj_str_bln _struct.Detail_Stock_Masuk_String

// 				sqlStatement := "SELECT kode_stock,nama_barang,jumlah_barang,harga_barang,satuan_barang FROM pembukuan_transaksi WHERE tanggal_pelunasan=?"

// 				_ = con.QueryRow(sqlStatement, bln_thn_sql).Scan(&obj_str.Kode_stock, &obj_str.Nama_barang, &obj_str.Jumlah_barang, &obj_str.Harga_barang, &obj_str.Satuan_barang)

// 				sqlStatement = "SELECT kode_stock,nama_barang,jumlah_barang,harga_barang,satuan_barang FROM pembukuan_transaksi_bulanan WHERE DATE_FORMAT(tanggal_pelunasan, \"%Y-%m\")=?"

// 				_ = con.QueryRow(sqlStatement, bln).Scan(&obj_str_bln.Kode_stock, &obj_str_bln.Nama_barang, &obj_str_bln.Jumlah_barang, &obj_str_bln.Harga_barang, &obj_str_bln.Satuan_barang)

// 				k_stk_all := String_Separator_To_String(obj_str_bln.Kode_stock)
// 				n_brg_all := String_Separator_To_String(obj_str_bln.Nama_barang)
// 				j_brg_all := String_Separator_To_float64(obj_str_bln.Jumlah_barang)
// 				h_brg_all := String_Separator_To_Int64(obj_str_bln.Harga_barang)
// 				s_brg_all := String_Separator_To_String(obj_str_bln.Satuan_barang)

// 				k_stock := String_Separator_To_String(obj_str.Kode_stock)
// 				j_barang := String_Separator_To_float64(obj_str.Jumlah_barang)
// 				h_barang := String_Separator_To_Int(obj_str.Harga_barang)
// 				n_barang := String_Separator_To_String(obj_str.Nama_barang)
// 				s_barang := String_Separator_To_String(obj_str.Satuan_barang)

// 				for j := 0; j < len(k_stock); j++ {

// 					co := 0

// 					for k := 0; k < len(k_stk_all); k++ {
// 						if k_stk_all[k] == k_stock[j] && n_brg_all[k] == n_barang[j] {
// 							j_brg_all[k] += j_barang[j]
// 							h := int64(h_barang[j])
// 							h_brg_all[k] += h
// 							co++
// 						}
// 					}

// 					if co == 0 {
// 						k_stk_all = append(k_stk_all, k_stock[j])
// 						n_brg_all = append(n_brg_all, n_barang[j])
// 						j_brg_all = append(j_brg_all, j_barang[j])
// 						h := int64(h_barang[j])
// 						h_brg_all = append(h_brg_all, h)
// 						s_brg_all = append(s_brg_all, s_barang[j])
// 					}
// 				}

// 				var k_stk_pmbk string
// 				var n_brg_pmbk string
// 				var j_brg_pmbk string
// 				var h_brg_pmbk string
// 				var s_brg_pmbk string

// 				for i := 0; i < len(k_stk_all); i++ {
// 					k_stk_pmbk += "|" + k_stk_all[i] + "|"
// 					n_brg_pmbk += "|" + n_brg_all[i] + "|"
// 					str := strconv.FormatFloat(j_brg_all[i], 'E', -1, 32)
// 					j_brg_pmbk += "|" + str + "|"
// 					s := strconv.FormatInt(h_brg_all[i], 10)
// 					h_brg_pmbk += "|" + s + "|"
// 					s_brg_pmbk += "|" + s_brg_all[i] + "|"
// 				}

// 				var total int64

// 				for i := 0; i < len(h_brg_all); i++ {
// 					total += h_brg_all[i]
// 				}

// 				id_bln := "PEM-BLN-" + bln

// 				sqlStatement = "UPDATE pembukuan_transaksi_bulanan SET kode_stock=?,nama_barang=?,jumlah_barang=?,harga_barang=?,total_harga_penjualan=?,tanggal_pelunasan=?,satuan_barang=? WHERE id_pembukuan_transaksi_bulanan=?"

// 				stmt, err := con.Prepare(sqlStatement)

// 				if err != nil {
// 					return res, err
// 				}

// 				_, err = stmt.Exec(k_stk_pmbk, n_brg_pmbk, j_brg_pmbk, h_brg_pmbk, total, bln_thn_sql, id_bln, s_brg_pmbk)
// 			}
// 		}

// 		thn := ls[2]

// 		sqlStatement = "SELECT * FROM pembukuan_transaksi_tahunan WHERE DATE_FORMAT(tanggal_pelunasan, \"%Y\")=?"

// 		_ = con.QueryRow(sqlStatement, thn).Scan(&obj_thn.Id_pembukuan_transaksi, &obj_thn.Kode_stock,
// 			&obj_thn.Nama_barang, &obj_thn.Jumlah_barang, obj.Satuan_barang, &obj_thn.Harga_barang, &obj_thn.Tanggal_pelunasan,
// 			&obj_thn.Total_harga_penjualan)

// 		if bln_thn_sql != obj_thn.Tanggal_pelunasan {
// 			if obj_thn.Id_pembukuan_transaksi == "" {

// 				sqlStatement := "SELECT kode_stock,nama_barang,jumlah_barang,harga_barang,satuan_barang FROM pembukuan_transaksi_tahunan WHERE tanggal_pelunasan=?"

// 				_ = con.QueryRow(sqlStatement, bln_thn_sql).Scan(&obj_str.Kode_stock, &obj_str.Nama_barang, &obj_str.Jumlah_barang, &obj_str.Harga_barang, &obj_str.Satuan_barang)

// 				var k_stk_all []string
// 				var n_brg_all []string
// 				var j_brg_all []float64
// 				var h_brg_all []int64
// 				var s_brg_all []string

// 				k_stock := String_Separator_To_String(obj_str.Kode_stock)
// 				j_barang := String_Separator_To_float64(obj_str.Jumlah_barang)
// 				h_barang := String_Separator_To_Int(obj_str.Harga_barang)
// 				n_barang := String_Separator_To_String(obj_str.Nama_barang)
// 				s_barang := String_Separator_To_String(obj_str.Satuan_barang)

// 				for j := 0; j < len(k_stock); j++ {
// 					k_stk_all = append(k_stk_all, k_stock[j])
// 					n_brg_all = append(n_brg_all, n_barang[j])
// 					j_brg_all = append(j_brg_all, j_barang[j])
// 					h := int64(h_barang[j])
// 					h_brg_all = append(h_brg_all, h)
// 					s_brg_all = append(s_brg_all, s_barang[j])
// 				}

// 				var k_stk_pmbk string
// 				var n_brg_pmbk string
// 				var j_brg_pmbk string
// 				var h_brg_pmbk string
// 				var s_brg_pmbk string

// 				for i := 0; i < len(k_stk_all); i++ {
// 					k_stk_pmbk += "|" + k_stk_all[i] + "|"
// 					n_brg_pmbk += "|" + n_brg_all[i] + "|"
// 					str := strconv.FormatFloat(j_brg_all[i], 'E', -1, 32)
// 					j_brg_pmbk += "|" + str + "|"
// 					s := strconv.FormatInt(h_brg_all[i], 10)
// 					h_brg_pmbk += "|" + s + "|"
// 					s_brg_pmbk += "|" + s_barang[i] + "|"
// 				}

// 				var total int64

// 				for i := 0; i < len(h_brg_all); i++ {
// 					total += h_brg_all[i]
// 				}

// 				id_bln := "PEM-THN-" + thn

// 				sqlStatement = "INSERT INTO pembukuan_transaksi_tahunan (id_pembukuan_transaksi_tahunan,kode_stock,nama_barang,jumlah_barang,harga_barang,tanggal_pelunasan,total_harga_penjualan,satuan_barang) values(?,?,?,?,?,?,?,?)"

// 				stmt, err := con.Prepare(sqlStatement)

// 				if err != nil {
// 					return res, err
// 				}

// 				_, err = stmt.Exec(id_bln, k_stk_pmbk, n_brg_pmbk, j_brg_pmbk, h_brg_pmbk, bln_thn_sql, total, s_brg_pmbk)

// 			} else {

// 				var obj_str_bln _struct.Detail_Stock_Masuk_String

// 				sqlStatement := "SELECT kode_stock,nama_barang,jumlah_barang,harga_barang,satuan_barang FROM pembukuan_transaksi WHERE tanggal_pelunasan=?"

// 				_ = con.QueryRow(sqlStatement, bln_thn_sql).Scan(&obj_str.Kode_stock, &obj_str.Nama_barang, &obj_str.Jumlah_barang, &obj_str.Harga_barang, &obj.Satuan_barang)

// 				sqlStatement = "SELECT kode_stock,nama_barang,jumlah_barang,harga_barang,satuan_barang FROM pembukuan_transaksi_tahunan WHERE DATE_FORMAT(tanggal_pelunasan, \"%Y\")=?"

// 				_ = con.QueryRow(sqlStatement, thn).Scan(&obj_str_bln.Kode_stock, &obj_str_bln.Nama_barang, &obj_str_bln.Jumlah_barang, &obj_str_bln.Harga_barang, &obj.Satuan_barang)

// 				k_stk_all := String_Separator_To_String(obj_str_bln.Kode_stock)
// 				n_brg_all := String_Separator_To_String(obj_str_bln.Nama_barang)
// 				j_brg_all := String_Separator_To_float64(obj_str_bln.Jumlah_barang)
// 				h_brg_all := String_Separator_To_Int64(obj_str_bln.Harga_barang)
// 				s_brg_all := String_Separator_To_String(obj_str_bln.Satuan_barang)

// 				k_stock := String_Separator_To_String(obj_str.Kode_stock)
// 				j_barang := String_Separator_To_float64(obj_str.Jumlah_barang)
// 				h_barang := String_Separator_To_Int(obj_str.Harga_barang)
// 				n_barang := String_Separator_To_String(obj_str.Nama_barang)
// 				s_barang := String_Separator_To_String(obj_str.Satuan_barang)

// 				for j := 0; j < len(k_stock); j++ {

// 					co := 0

// 					for k := 0; k < len(k_stk_all); k++ {
// 						if k_stk_all[k] == k_stock[j] && n_brg_all[k] == n_barang[j] {
// 							j_brg_all[k] += j_barang[j]
// 							h := int64(h_barang[j])
// 							h_brg_all[k] += h
// 							co++
// 						}
// 					}

// 					if co == 0 {
// 						k_stk_all = append(k_stk_all, k_stock[j])
// 						n_brg_all = append(n_brg_all, n_barang[j])
// 						j_brg_all = append(j_brg_all, j_barang[j])
// 						h := int64(h_barang[j])
// 						h_brg_all = append(h_brg_all, h)
// 						s_brg_all = append(s_brg_all, s_barang[j])
// 					}
// 				}

// 				var k_stk_pmbk string
// 				var n_brg_pmbk string
// 				var j_brg_pmbk string
// 				var h_brg_pmbk string
// 				var s_brg_pmbk string

// 				for i := 0; i < len(k_stk_all); i++ {
// 					k_stk_pmbk += "|" + k_stk_all[i] + "|"
// 					n_brg_pmbk += "|" + n_brg_all[i] + "|"
// 					str := strconv.FormatFloat(j_brg_all[i], 'E', -1, 32)
// 					j_brg_pmbk += "|" + str + "|"
// 					s := strconv.FormatInt(h_brg_all[i], 10)
// 					h_brg_pmbk += "|" + s + "|"
// 					s_brg_pmbk += "|" + s_brg_all[i] + "|"
// 				}

// 				var total int64

// 				for i := 0; i < len(h_brg_all); i++ {
// 					total += h_brg_all[i]
// 				}

// 				id_thn := "PEM-THN-" + thn

// 				sqlStatement = "UPDATE pembukuan_transaksi_tahunan SET kode_stock=?,nama_barang=?,jumlah_barang=?,harga_barang=?,total_harga_penjualan=?,tanggal_pelunasan=?, satuan_barang=? WHERE id_pembukuan_transaksi_tahunan=?"

// 				stmt, err := con.Prepare(sqlStatement)

// 				if err != nil {
// 					return res, err
// 				}

// 				_, err = stmt.Exec(k_stk_pmbk, n_brg_pmbk, j_brg_pmbk, h_brg_pmbk, total, bln_thn_sql, id_thn, s_brg_pmbk)
// 			}
// 		}
// 		sqlStatement = "SELECT * FROM pembukuan_transaksi WHERE tanggal_pelunasan=?"

// 		_ = con.QueryRow(sqlStatement, bln_thn_sql).Scan(&obj.Id_pembukuan_transaksi, &obj.Kode_stock,
// 			&obj.Nama_barang, &obj.Jumlah_barang, &obj.Satuan_barang, &obj.Harga_barang, &obj.Tanggal_pelunasan, &obj.Total_harga_penjualan)

// 		stmt.Close()

// 		res.Status = http.StatusOK
// 		res.Message = "Sukses"
// 		res.Data = obj
// 	} else {
// 		res.Status = http.StatusOK
// 		res.Message = "Not found transaction"
// 		res.Data = obj
// 	}

// 	return res, nil
// }

// func Read_Pembukuan(tanggal string) (Response, error) {
// 	var res Response
// 	var arrobj []_struct.Read_Pembukuan_Transaksi
// 	var obj _struct.Read_Pembukuan_Transaksi

// 	var arrobj_fix []_struct.Read_Pembukuan_Transaksi_List
// 	var obj_fix _struct.Read_Pembukuan_Transaksi_List

// 	ls := []string{}
// 	str1 := ""

// 	for i := 0; i < len(tanggal); i++ {
// 		if byte(tanggal[i]) >= 48 && byte(tanggal[i]) <= 57 {
// 			str1 += string(tanggal[i])
// 			if i == len(tanggal)-1 {
// 				ls = append(ls, str1)
// 			}
// 		} else if tanggal[i] == '-' {
// 			ls = append(ls, str1)
// 			str1 = ""
// 		}
// 	}

// 	j := len(ls)
// 	bln_thn_sql := ""

// 	for x := j - 1; x >= 0; x-- {
// 		bln_thn_sql += ls[x]
// 		if x != 0 {
// 			bln_thn_sql += "-"
// 		}
// 	}

// 	awal := ls[2] + "-" + ls[1] + "-" + "01"

// 	con := db.CreateCon()

// 	sqlStatement := "SELECT id_pembukuan_transaksi,kode_stock,nama_barang,jumlah_barang,satuan_barang,harga_barang,Date_Format(tanggal_pelunasan,\"%d-%m-%Y\"),total_harga_penjualan FROM pembukuan_transaksi WHERE tanggal_pelunasan<=?&&tanggal_pelunasan>=? ORDER BY id_pembukuan_transaksi DESC "

// 	rows, err := con.Query(sqlStatement, bln_thn_sql, awal)

// 	defer rows.Close()

// 	if err != nil {
// 		return res, err
// 	}

// 	for rows.Next() {
// 		err = rows.Scan(&obj.Id_pembukuan_transaksi, &obj.Kode_stock, &obj.Nama_barang, &obj.Jumlah_barang, &obj.Satuan_barang,
// 			&obj.Harga_barang, &obj.Tanggal_pelunasan, &obj.Total_harga_penjualan)
// 		if err != nil {
// 			return res, err
// 		}
// 		arrobj = append(arrobj, obj)
// 	}

// 	for i := 0; i < len(arrobj); i++ {
// 		obj_fix.Id_pembukuan_transaksi = arrobj[i].Id_pembukuan_transaksi
// 		obj_fix.Tanggal_pelunasan = arrobj[i].Tanggal_pelunasan
// 		obj_fix.Total_harga_penjualan = arrobj[i].Total_harga_penjualan
// 		obj_fix.Nama_barang = String_Separator_To_String(arrobj[i].Nama_barang)
// 		obj_fix.Kode_stock = String_Separator_To_String(arrobj[i].Kode_stock)
// 		obj_fix.Jumlah_barang = String_Separator_To_float64(arrobj[i].Jumlah_barang)
// 		obj_fix.Harga_barang = String_Separator_To_Int(arrobj[i].Harga_barang)
// 		obj_fix.Satuan_barang = String_Separator_To_String(arrobj[i].Satuan_barang)
// 		arrobj_fix = append(arrobj_fix, obj_fix)
// 	}

// 	if arrobj_fix == nil {
// 		res.Status = http.StatusNotFound
// 		res.Message = "Not Found"
// 		res.Data = arrobj_fix
// 	} else {
// 		res.Status = http.StatusOK
// 		res.Message = "Sukses"
// 		res.Data = arrobj_fix
// 	}

// 	return res, nil
// }
