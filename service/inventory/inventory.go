package inventory

import (
	"Bakend-POS/db"
	"Bakend-POS/models/request"
	"Bakend-POS/models/response"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func Input_Inventory(Request request.Input_Inventory_Request, writer http.ResponseWriter, request *http.Request) (response.Response, error) {
	var res response.Response

	con := db.CreateConGorm()

	nama_barang := ""

	err := con.Table("inventory").Select("nama_barang").Where("nama_barang = ? AND kode_user = ?", Request.Nama_barang, Request.Kode_user).Scan(&nama_barang)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	if nama_barang == "" {

		co := 0

		err := con.Table("inventory").Select("co").Order("co DESC").Limit(1).Scan(&co)

		Request.Co = co + 1
		Request.Kode_inventory = "IN-" + strconv.Itoa(Request.Co)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

		request.ParseMultipartForm(10 * 1024 * 1024)
		file, handler, err2 := request.FormFile("photo")

		if file != nil {

			defer file.Close()

			fmt.Println("File Info")
			fmt.Println("File Name : ", handler.Filename)
			fmt.Println("File Size : ", handler.Size)
			fmt.Println("File Type : ", handler.Header.Get("Content-Type"))

			var tempFile *os.File
			path := ""

			if strings.Contains(handler.Filename, "jpg") {
				path = "uploads/foto_inventory/" + Request.Kode_inventory + "-" + Request.Kode_user + ".jpg"
				tempFile, err2 = ioutil.TempFile("uploads/foto_news/", "Read"+"*.jpg")
			}
			if strings.Contains(handler.Filename, "jpeg") {
				path = "uploads/foto_inventory/" + Request.Kode_inventory + "-" + Request.Kode_user + ".jpeg"
				tempFile, err2 = ioutil.TempFile("uploads/foto_inventory/", "Read"+"*.jpeg")
			}
			if strings.Contains(handler.Filename, "png") {
				path = "uploads/foto_inventory/" + Request.Kode_inventory + "-" + Request.Kode_user + ".png"
				tempFile, err2 = ioutil.TempFile("uploads/foto_inventory/", "Read"+"*.png")
			}

			if err2 != nil {
				return res, err2
			}

			fileBytes, err2 := ioutil.ReadAll(file)
			if err2 != nil {
				return res, err2
			}

			_, err2 = tempFile.Write(fileBytes)
			if err2 != nil {
				return res, err2
			}

			fmt.Println("Success!!")
			fmt.Println(tempFile.Name())
			tempFile.Close()

			err2 = os.Rename(tempFile.Name(), path)
			if err2 != nil {
				fmt.Println(err)
			}

			defer tempFile.Close()

			fmt.Println("new path:", tempFile.Name())

			Request.Path_photo = path
		} else {
			Request.Path_photo = "uploads/foto_inventory/box.jpg"
		}

		err = con.Table("inventory").Select("co", "kode_inventory", "nama_barang", "harga_jual", "satuan_barang", "kode_user", "path_photo").Create(&Request)

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
		res.Message = "Nama Barang Telah Digunakan"
		res.Data = Request
		return res, err.Error
	}

	return res, nil
}

func Read_Inventory(Request request.Read_Inventory_Request) (response.Response, error) {
	var res response.Response
	var arr_invent []response.Read_Inventory_Response

	con := db.CreateConGorm()

	err := con.Table("inventory").Select("kode_inventory", "nama_barang", "jumlah_barang", "satuan_barang", "harga_jual").Where("kode_user = ?", Request.Kode_user).Scan(&arr_invent)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	for i := 0; i < len(arr_invent); i++ {
		err := con.Table("detail_inventory").Select("kode_barang_transaksi_inventory", "detail_inventory.kode_transaksi_inventory", "nama_supplier", "jumlah", "harga").Joins("JOIN transaksi_inventory ti on ti.kode_transaksi_inventory=detail_inventory.kode_transaksi_inventory").Where("kode_inventory = ?", arr_invent[i].Kode_inventory).Scan(&arr_invent[i].Detail_inventory)

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

func Update_Inventory(Request request.Update_Inventory_Request, writer http.ResponseWriter, request *http.Request) (response.Response, error) {
	var res response.Response

	con := db.CreateConGorm()

	kode_inventory := ""

	fmt.Println(Request)

	err := con.Table("barang_transaksi_inventory").Select("kode_inventory").Where("kode_inventory = ?", Request.Kode_inventory).Limit(1).Scan(&kode_inventory)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	// kode_barang_pembukuan := ""

	// err = con.Table("barang_pembukuan").Select("kode_barang_pembukuan").Where("kode_stock = ?", Request.Kode_stock).Limit(1).Scan(&kode_barang_pembukuan)

	// if err.Error != nil {
	// 	res.Status = http.StatusNotFound
	// 	res.Message = "Status Not Found"
	// 	res.Data = Request
	// 	return res, err.Error
	// }

	if kode_inventory == "" {

		err := con.Table("inventory").Where("kode_inventory = ?", Request.Kode_inventory).Select("nama_barang", "satuan_barang").Updates(&Request)

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
	}

	path := ""

	err = con.Table("inventory").Select("path_photo").Where("kode_inventory=?", Request.Kode_inventory).Scan(&path)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	request.ParseMultipartForm(10 * 1024 * 1024)
	file, handler, err2 := request.FormFile("photo")

	if file != nil {

		if path != "uploads/foto_inventory/box.jpg" {
			path = "./" + path
			_ = os.Remove(path)

		}

		defer file.Close()

		fmt.Println("File Info")
		fmt.Println("File Name : ", handler.Filename)
		fmt.Println("File Size : ", handler.Size)
		fmt.Println("File Type : ", handler.Header.Get("Content-Type"))

		var tempFile *os.File
		path := ""

		if strings.Contains(handler.Filename, "jpg") {
			path = "uploads/foto_inventory/" + Request.Kode_inventory + "-" + Request.Kode_user + ".jpg"
			tempFile, err2 = ioutil.TempFile("uploads/foto_news/", "Read"+"*.jpg")
		}
		if strings.Contains(handler.Filename, "jpeg") {
			path = "uploads/foto_inventory/" + Request.Kode_inventory + "-" + Request.Kode_user + ".jpeg"
			tempFile, err2 = ioutil.TempFile("uploads/foto_inventory/", "Read"+"*.jpeg")
		}
		if strings.Contains(handler.Filename, "png") {
			path = "uploads/foto_inventory/" + Request.Kode_inventory + "-" + Request.Kode_user + ".png"
			tempFile, err2 = ioutil.TempFile("uploads/foto_inventory/", "Read"+"*.png")
		}

		if err2 != nil {
			return res, err2
		}

		fileBytes, err2 := ioutil.ReadAll(file)
		if err2 != nil {
			return res, err2
		}

		_, err2 = tempFile.Write(fileBytes)
		if err2 != nil {
			return res, err2
		}

		fmt.Println("Success!!")
		fmt.Println(tempFile.Name())
		tempFile.Close()

		err2 = os.Rename(tempFile.Name(), path)
		if err2 != nil {
			fmt.Println(err)
		}

		defer tempFile.Close()

		fmt.Println("new path:", tempFile.Name())

		Request.Path_photo = path
	}

	err = con.Table("inventory").Where("kode_inventory = ?", Request.Kode_inventory).Select("harga_jual", "path_photo").Updates(&Request)

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

func Check_Nama_Inventory(Request request.Check_Nama_Inventory_Request) (response.Response, error) {
	var res response.Response
	var check response.Check_Nama_Inventory_Response

	con := db.CreateConGorm()

	err := con.Table("inventory").Select("nama_barang").Where("kode_user = ? AND nama_barang = ? AND kode_inventory != ", Request.Kode_user, Request.Nama_barang, Request.Kode_inventory).Scan(&check)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	if check.Nama_barang == "" {
		check.Status = true
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = check
	} else {
		check.Status = false
		res.Status = http.StatusNotFound
		res.Message = "Nama Sudah Ada"
		res.Data = check
	}

	return res, nil
}

// func Dropdown_Inventory_transaksi_inventory(Request request.Dropdown_Inventory_transaksi_inventory_request) (response.Response, error) {
// 	var res response.Response
// 	var arr_invent []response.Read_Inventory_Response

// 	_, condition := session_checking.Session_Checking(Request.Uuid_session)

// 	if condition {
// 		con := db.CreateConGorm()

// 		err := con.Table("inventory").Select("kode_inventory", "nama_barang", "jumlah_barang", "satuan_barang", "harga_jual").Scan(&arr_invent)

// 		if err.Error != nil {
// 			res.Status = http.StatusNotFound
// 			res.Message = "Status Not Found"
// 			res.Data = Request
// 			return res, err.Error
// 		}

// 		if arr_invent == nil {
// 			res.Status = http.StatusNotFound
// 			res.Message = "Not Found"
// 			res.Data = arr_invent
// 		} else {
// 			res.Status = http.StatusOK
// 			res.Message = "Sukses"
// 			res.Data = arr_invent
// 		}
// 	} else {
// 		res.Status = http.StatusNotFound
// 		res.Message = "Nama Barang Telah Digunakan"
// 		res.Data = Request
// 	}

// 	return res, nil
// }
