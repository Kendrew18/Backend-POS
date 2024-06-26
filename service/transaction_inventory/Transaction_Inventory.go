package transaction_inventory

import (
	"Bakend-POS/db"
	"Bakend-POS/models/request"
	"Bakend-POS/models/response"
	"database/sql"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"
)

func Compare(data []request.Input_barang_Supplier_Request, kode_inventory string) bool {

	for j := 0; j < len(data); j++ {

		if data[j].Kode_inventory == kode_inventory {
			return true
		}

	}

	return false
}

func Input_Transaction_Inventory(Request request.Input_Transaksi_Inventory_Request, Request_barang []request.Input_Barang_Transaksi_Inventory_Request) (response.Response, error) {
	var res response.Response

	//Incoming Inventory = 0
	//Refund = 1

	con := db.CreateConGorm()

	co := 0

	err := con.Table("transaksi_inventory").Select("co").Limit(1).Order("co DESC").Scan(&co)

	Request.Co = co + 1
	Request.Kode_transaksi_inventory = "TI-" + strconv.Itoa(Request.Co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	date, _ := time.Parse("02-01-2006", Request.Tanggal)
	Request.Tanggal = date.Format("2006-01-02")

	if Request.Jenis_transaksi == 1 {
		err = con.Table("transaksi_inventory").Select("nama_supplier", "nomor_telpon_supplier").Where("kode_nota = ?", Request.Kode_nota).Scan(&Request)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

		Request.Harga_ongkos_kirim = 0
		Request.Ppn = 0
	}

	err = con.Table("transaksi_inventory").Select("co", "kode_transaksi_inventory", "nama_supplier", "nomor_telpon_supplier", "tanggal", "kode_nota", "harga_ongkos_kirim", "ppn", "kode_user", "jenis_transaksi").Create(&Request)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	co = 0

	err = con.Table("barang_transaksi_inventory").Select("co").Limit(1).Order("co DESC").Scan(&co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	total_harga := int64(0)
	total_jumlah := 0.0

	for i := 0; i < len(Request_barang); i++ {
		Request_barang[i].Co = co + 1 + i
		Request_barang[i].Kode_barang_transaksi_inventory = "BTI-" + strconv.Itoa(Request_barang[i].Co)
		Request_barang[i].Kode_transaksi_inventory = Request.Kode_transaksi_inventory

		if Request.Jenis_transaksi == 0 {
			Request_barang[i].Sub_total = int64(math.Round(float64(Request_barang[i].Harga) * Request_barang[i].Jumlah))
			total_harga = total_harga + Request_barang[i].Sub_total
		}

		total_jumlah = total_jumlah + Request_barang[i].Jumlah
	}

	if Request.Jenis_transaksi == 0 {
		total_harga = total_harga + Request.Harga_ongkos_kirim + int64(math.Round(float64(total_harga)*Request.Ppn/100))
	}

	err = con.Table("barang_transaksi_inventory").Select("co", "kode_barang_transaksi_inventory", "kode_transaksi_inventory", "kode_inventory", "jumlah", "harga", "sub_total", "kode_refund").Create(&Request_barang)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	err = con.Table("transaksi_inventory").Where("kode_transaksi_inventory = ?", Request.Kode_transaksi_inventory).Update("total_harga", total_harga).Update("total_barang", total_jumlah)

	// // add barang supplier yang belum masuk

	// var data_sup request.Input_barang_Supplier_Request
	// var arr_data_sup []request.Input_barang_Supplier_Request

	// co = 0

	// err = con.Table("barang_supplier").Select("co").Limit(1).Order("co DESC").Scan(&co)

	// if err.Error != nil {
	// 	res.Status = http.StatusNotFound
	// 	res.Message = "Status Not Found"
	// 	res.Data = Request
	// 	return res, err.Error
	// }

	// var data_stock []request.Input_barang_Supplier_Request

	// err = con.Table("barang_supplier").Select("kode_inventory").Joins("JOIN supplier s on s.kode_supplier = barang_supplier.kode_supplier").Where("barang_supplier.kode_supplier = ?", Request.Kode_supplier).Scan(&data_stock)

	// if err.Error != nil {
	// 	res.Status = http.StatusNotFound
	// 	res.Message = "Status Not Found"
	// 	res.Data = Request
	// 	return res, err.Error
	// }

	// x := 0

	// for i := 0; i < len(kode_inventory); i++ {
	// 	if Compare(data_stock, kode_inventory[i]) {

	// 	} else {
	// 		data_sup.Co = co + 1 + x
	// 		data_sup.Kode_barang_supplier = "BS-" + strconv.Itoa(data.Co)
	// 		data_sup.Kode_inventory = kode_inventory[i]
	// 		data_sup.Kode_supplier = Request.Kode_supplier

	// 		arr_data_sup = append(arr_data_sup, data_sup)
	// 		x++
	// 	}
	// }

	// if len(arr_data_sup) > 0 {
	// 	err = con.Table("barang_supplier").Select("co", "kode_barang_supplier", "kode_supplier", "kode_inventory").Create(&arr_data_sup)
	// 	if err.Error != nil {
	// 		res.Status = http.StatusNotFound
	// 		res.Message = "Status Not Found"
	// 		res.Data = Request
	// 		return res, err.Error
	// 	}
	// }

	// //Update data
	// err = con.Exec("UPDATE `inventory` JOIN barang_transaksi_inventory bsm ON bsm.kode_inventory = inventory.kode_inventory SET `jumlah_barang`=jumlah_barang + jumlah WHERE bsm.kode_transaksi_inventory = ?", Request.Kode_transaksi_inventory)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	} else {
		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = map[string]int64{
			"rows": err.RowsAffected,
		}
	}
	return res, nil
}

func Read_Transaction_Inventory(Request request.Read_Transaksi_Inventory_Request, Request_filter request.Read_Transaksi_Inventory_Filter_Request) (response.Response, error) {
	var res response.Response
	var arr_data []response.Read_Transaksi_Inventory_Response
	var data response.Read_Transaksi_Inventory_Response
	var rows *sql.Rows
	var err error

	con := db.CreateConGorm()

	statement := "transaksi_inventory.kode_user = '" + Request.Kode_user + "'"

	if Request_filter.Nama_supplier != "" {
		statement += " && transaksi_inventory.nama_supplier = '" + Request_filter.Nama_supplier + "'"
	}

	if Request_filter.Tanggal_awal != "" && Request_filter.Tanggal_akhir != "" {

		date, _ := time.Parse("02-01-2006", Request_filter.Tanggal_awal)
		Request_filter.Tanggal_awal = date.Format("2006-01-02")

		date2, _ := time.Parse("02-01-2006", Request_filter.Tanggal_akhir)
		Request_filter.Tanggal_akhir = date2.Format("2006-01-02")

		statement += " AND (tanggal >= '" + Request_filter.Tanggal_awal + "' && tanggal <= '" + Request_filter.Tanggal_akhir + "' )"

	} else if Request_filter.Tanggal_awal != "" {

		date, _ := time.Parse("02-01-2006", Request_filter.Tanggal_awal)
		Request_filter.Tanggal_awal = date.Format("2006-01-02")

		statement += " && tanggal = '" + Request_filter.Tanggal_awal + "'"

	}

	rows, err = con.Table("transaksi_inventory").Select("kode_transaksi_inventory", "DATE_FORMAT(tanggal, '%d-%m-%Y') AS tanggal", "kode_nota", "nama_supplier", "nomor_telpon_supplier", "harga_ongkos_kirim", "ppn", "total_harga", "total_barang", "status", "jenis_transaksi").Where(statement).Order("transaksi_inventory.co DESC").Rows()

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = arr_data
		return res, err
	}

	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(&data.Kode_transaksi_inventory, &data.Tanggal, &data.Kode_nota, &data.Nama_supplier, &data.Nomor_telpon_supplier, &data.Harga_ongkos_kirim, &data.Ppn, &data.Total_harga, &data.Total_barang, &data.Status, &data.Jenis_transaksi)

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = data
			return res, err
		}

		var arr_barang []response.Read_Barang_Transaksi_Inventory_Response

		err = con.Table("barang_transaksi_inventory").Select("kode_barang_transaksi_inventory", "barang_transaksi_inventory.kode_inventory", "nama_barang", "barang_transaksi_inventory.jumlah", "barang_transaksi_inventory.harga", "barang_transaksi_inventory.sub_total").Joins("join inventory s on s.kode_inventory = barang_transaksi_inventory.kode_inventory").Where("kode_transaksi_inventory = ?", data.Kode_transaksi_inventory).Scan(&arr_barang).Error

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = data
			return res, err
		}

		data.Barang_transaksi_inventory = arr_barang

		arr_data = append(arr_data, data)

	}

	if arr_data == nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = arr_data

	} else {
		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = arr_data
	}

	return res, nil
}

func Update_Header_Transaction_Inventory(Request request.Update_Header_Transaksi_Inventory_Request, Request_kode request.Update_Header_Transaksi_Inventory_Kode_Request) (response.Response, error) {
	var res response.Response

	check := -1
	con_check := db.CreateConGorm()

	err := con_check.Table("transaksi_inventory").Select("status").Where("kode_transaksi_inventory = ?", Request_kode.Kode_transaksi_inventory).Scan(&check)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Update Error"
		res.Data = Request
		return res, err.Error
	}

	if check == 0 || check == 2 {

		con := db.CreateConGorm()

		total_harga := int64(0)

		err = con.Table("barang_transaksi_inventory").Select("SUM(sub_total)").Where("kode_transaksi_inventory = ?", Request_kode.Kode_transaksi_inventory).Scan(&total_harga)
		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Update Error"
			res.Data = Request
			return res, err.Error
		}

		Request.Total_harga = total_harga + Request.Harga_ongkos_kirim + int64(math.Round(float64(total_harga)*Request.Ppn/100))

		err = con.Table("transaksi_inventory").Where("kode_transaksi_inventory = ?", Request_kode.Kode_transaksi_inventory).Select("nama_supplier", "nomor_telpon_supplier", "kode_nota", "harga_ongkos_kirim", "ppn", "total_harga").Updates(&Request)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		} else {
			res.Status = http.StatusOK
			res.Message = "Suksess"
			res.Data = map[string]int64{
				"rows": err.RowsAffected,
			}
		}
	} else {
		res.Status = http.StatusNotFound
		res.Message = "Barang Tidak dapat di update"
		res.Data = Request
		return res, err.Error
	}
	return res, nil
}

func Update_Barang_Transaction_Inventory(Request request.Update_Barang_Transaksi_Inventory_Request, Request_kode request.Update_Barang_Transaksi_Inventory_Kode_Request) (response.Response, error) {
	var res response.Response

	check := -1
	con_check := db.CreateConGorm()

	err := con_check.Table("transaksi_inventory").Select("status").Joins("JOIN barang_transaksi_inventory bti on bti.kode_transaksi_inventory = transaksi_inventory.kode_transaksi_inventory").Where("kode_barang_transaksi_inventory = ?", Request_kode.Kode_barang_transaksi_inventory).Scan(&check)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Update Error"
		res.Data = Request
		return res, err.Error
	}

	if check == 0 || check == 2 {

		var res_update response.Update_Barang_Transaction_Inventory_Response

		con := db.CreateConGorm()

		Request.Sub_total = int64(math.Round(float64(Request.Harga) * Request.Jumlah))

		err = con.Table("barang_transaksi_inventory").Where("kode_barang_transaksi_inventory = ?", Request_kode.Kode_barang_transaksi_inventory).Select("jumlah", "harga", "sub_total").Updates(&Request)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Update Error"
			res.Data = Request
			return res, err.Error
		}

		err := con.Table("transaksi_inventory").Select("bti.kode_transaksi_inventory").Joins("JOIN barang_transaksi_inventory bti on bti.kode_transaksi_inventory = transaksi_inventory.kode_transaksi_inventory").Where("kode_barang_transaksi_inventory = ?", Request_kode.Kode_barang_transaksi_inventory).Scan(&res_update)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Update Error"
			res.Data = Request
			return res, err.Error
		}

		err = con.Table("transaksi_inventory").Select("SUM(sub_total) AS sub_total", "harga_ongkos_kirim", "ppn", "SUM(jumlah) AS total_barang").Joins("JOIN barang_transaksi_inventory bti ON bti.kode_transaksi_inventory = transaksi_inventory.kode_transaksi_inventory").Where("bti.kode_transaksi_inventory = ?", res_update.Kode_transaksi_inventory).Group("bti.kode_transaksi_inventory").Scan(&res_update)

		fmt.Println(res_update)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Update Error"
			res.Data = Request
			return res, err.Error
		}

		res_update.Total_harga = res_update.Sub_total + res_update.Harga_ongkos_kirim + int64(math.Round(float64(res_update.Sub_total)*res_update.Ppn/100))

		err = con.Table("transaksi_inventory").Where("kode_transaksi_inventory = ?", res_update.Kode_transaksi_inventory).Select("total_barang", "total_harga").Updates(&res_update)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		} else {
			res.Status = http.StatusOK
			res.Message = "Suksess"
			res.Data = map[string]int64{
				"rows": err.RowsAffected,
			}
		}
	} else {
		res.Status = http.StatusNotFound
		res.Message = "Barang Tidak dapat di update"
		res.Data = Request
		return res, err.Error
	}
	return res, nil
}

func Delete_Barang_Transaksi_Inventory(Request request.Update_Barang_Transaksi_Inventory_Kode_Request) (response.Response, error) {
	var res response.Response

	check := -1
	con_check := db.CreateConGorm()

	err := con_check.Table("transaksi_inventory").Select("status").Joins("JOIN barang_transaksi_inventory bti on bti.kode_transaksi_inventory = transaksi_inventory.kode_transaksi_inventory").Where("kode_barang_transaksi_inventory = ?", Request.Kode_barang_transaksi_inventory).Scan(&check)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Update Error"
		res.Data = Request
		return res, err.Error
	}

	if check == 0 || check == 2 {

		con := db.CreateConGorm()

		data := ""

		err = con.Table("transaksi_inventory").Select("transaksi_inventory.kode_transaksi_inventory").Joins("JOIN barang_transaksi_inventory bpi ON bpi.kode_transaksi_inventory = transaksi_inventory.kode_transaksi_inventory ").Where("kode_barang_transaksi_inventory = ?", Request.Kode_barang_transaksi_inventory).Scan(&data)

		fmt.Println(data)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Update Error"
			res.Data = Request
			return res, err.Error
		}

		err = con.Table("barang_transaksi_inventory").Where("kode_barang_transaksi_inventory = ?", Request.Kode_barang_transaksi_inventory).Delete("")

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Update Error"
			res.Data = Request
			return res, err.Error
		}

		kode_barang := ""

		err = con.Table("barang_transaksi_inventory").Select("kode_barang_transaksi_inventory").Where("kode_transaksi_inventory=?", data).Limit(1).Scan(&kode_barang)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Update Error"
			res.Data = Request
			return res, err.Error
		}

		if kode_barang == "" {
			fmt.Println("masuk")

			err = con.Table("transaksi_inventory").Where("kode_transaksi_inventory = ?", data).Delete("")

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}
		} else {

			total_harga := int64(0)

			err = con.Table("barang_transaksi_inventory").Select("SUM(sub_total)").Where("kode_transaksi_inventory=?", data).Scan(&total_harga)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			total_jumlah := float64(0)

			err = con.Table("barang_transaksi_inventory").Select("SUM(jumlah)").Where("kode_transaksi_inventory=?", data).Scan(&total_jumlah)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			var temp request.Input_Transaksi_Inventory_Request

			err = con.Table("transaksi_inventory").Select("harga_ongkos_kirim", "ppn").Where("kode_transaksi_inventory=?", data).Scan(&temp)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			total_harga = total_harga + temp.Harga_ongkos_kirim + int64(math.Round(float64(total_harga)*temp.Ppn/100))

			err = con.Table("transaksi_inventory").Where("kode_transaksi_inventory = ?", data).Update("total_harga", total_harga).Update("total_barang", total_jumlah)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

		}

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		} else {
			res.Status = http.StatusOK
			res.Message = "Suksess"
			res.Data = map[string]int64{
				"rows": err.RowsAffected,
			}
		}
	} else {
		res.Status = http.StatusNotFound
		res.Message = "Barang Tidak dapat di update"
		res.Data = Request
		return res, err.Error
	}
	return res, nil
}

func Update_Status_Transaksi_Inventory(Request request.Body_Update_Status_Transaksi_inventory) (response.Response, error) {
	var res response.Response
	con := db.CreateConGorm()
	status := -1

	err := con.Table("transaksi_inventory").Select("status").Where("kode_transaksi_inventory = ?", Request.Update_header_transaksi_inventory_kode.Kode_transaksi_inventory).Scan(&status)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	if status != 1 {
		if Request.Update_status_transaksi_inventory.Status == 2 || Request.Update_status_transaksi_inventory.Status == 0 {
			Request_kode := Request.Update_header_transaksi_inventory_kode
			Request := Request.Update_status_transaksi_inventory

			con := db.CreateConGorm()

			err := con.Table("transaksi_inventory").Where("kode_transaksi_inventory = ?", Request_kode.Kode_transaksi_inventory).Select("status").Updates(&Request)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			} else {
				res.Status = http.StatusOK
				res.Message = "Suksess"
				res.Data = map[string]int64{
					"rows": err.RowsAffected,
				}
			}
		} else if Request.Update_status_transaksi_inventory.Status == 1 {
			Request_kode := Request.Update_header_transaksi_inventory_kode
			Request := Request.Update_status_transaksi_inventory

			con := db.CreateConGorm()

			err := con.Table("transaksi_inventory").Where("kode_transaksi_inventory = ?", Request_kode.Kode_transaksi_inventory).Select("status").Updates(&Request)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			jenis_transaksi := -1

			err = con.Table("transaksi_inventory").Select("jenis_transaksi").Where("kode_transaksi_inventory = ?", Request_kode.Kode_transaksi_inventory).Scan(&jenis_transaksi)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			if jenis_transaksi == 0 {

				var arr_data []request.Input_Barang_Transaksi_Inventory_Request

				err = con.Table("barang_transaksi_inventory").Select("co", "kode_barang_transaksi_inventory", "kode_transaksi_inventory", "kode_inventory", "jumlah", "harga", "sub_total").Where("kode_transaksi_inventory = ?", Request_kode.Kode_transaksi_inventory).Scan(&arr_data)

				if err.Error != nil {
					res.Status = http.StatusNotFound
					res.Message = "Status Not Found"
					res.Data = Request
					return res, err.Error
				}

				err = con.Table("detail_inventory").Select("co", "kode_barang_transaksi_inventory", "kode_transaksi_inventory", "kode_inventory", "jumlah", "harga", "sub_total").Create(&arr_data)

				if err.Error != nil {
					res.Status = http.StatusNotFound
					res.Message = "Status Not Found"
					res.Data = Request
					return res, err.Error
				}

				err = con.Exec("UPDATE `inventory` JOIN barang_transaksi_inventory bsm ON bsm.kode_inventory = inventory.kode_inventory SET `jumlah_barang`=jumlah_barang + jumlah WHERE bsm.kode_transaksi_inventory = ?", Request_kode.Kode_transaksi_inventory)

				if err.Error != nil {
					res.Status = http.StatusNotFound
					res.Message = "Update Gagal"
					res.Data = Request
					return res, err.Error
				} else {
					res.Status = http.StatusOK
					res.Message = "Suksess"
					res.Data = map[string]int64{
						"rows": 1,
					}
				}

			} else if jenis_transaksi == 1 {

				err = con.Exec("UPDATE `inventory` JOIN barang_transaksi_inventory bsm ON bsm.kode_inventory = inventory.kode_inventory SET `jumlah_barang`=jumlah_barang - jumlah WHERE bsm.kode_transaksi_inventory = ?", Request_kode.Kode_transaksi_inventory)

				if err.Error != nil {
					res.Status = http.StatusNotFound
					res.Message = "Update Gagal"
					res.Data = Request
					return res, err.Error
				}

				err = con.Exec("UPDATE `detail_inventory` JOIN barang_transaksi_inventory bti ON bti.kode_refund = detail_inventory.kode_barang_transaksi_inventory SET detail_inventory.jumlah = detail_inventory.jumlah - bti.jumlah WHERE bti.kode_transaksi_inventory = ?", Request_kode.Kode_transaksi_inventory)

				if err.Error != nil {
					res.Status = http.StatusNotFound
					res.Message = "Status Not Found"
					res.Data = Request
					return res, err.Error
				} else {
					res.Status = http.StatusOK
					res.Message = "Suksess"
					res.Data = map[string]int64{
						"rows": 1,
					}
				}

			}
		}
	} else {
		res.Status = http.StatusNotFound
		res.Message = "Tidah dapat di edit diakrenakan sudah sukses"
		res.Data = Request
	}
	return res, nil
}

func Dropdown_Transaksi_Inventory(Request request.Dropdown_Inventory_transaksi_inventory_request) (response.Response, error) {
	var res response.Response
	var arr_invent []response.Dropdown_Transaction_Inventory_Response

	con := db.CreateConGorm()

	if Request.Kode_nota != "" {

		err := con.Table("detail_inventory").Select("kode_barang_transaksi_inventory", "detail_inventory.kode_transaksi_inventory", "detail_inventory.kode_inventory", "nama_barang", "jumlah", "harga").Joins("JOIN inventory i on i.kode_inventory = detail_inventory.kode_inventory").Joins("JOIN transaksi_inventory ti on ti.kode_transaksi_inventory=detail_inventory.kode_transaksi_inventory").Where("kode_nota = ?", Request.Kode_nota).Scan(&arr_invent)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

	} else {
		err := con.Table("inventory").Select("kode_inventory", "nama_barang").Where("kode_user = ?", Request.Kode_user).Scan(&arr_invent)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

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
