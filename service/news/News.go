package news

import (
	"Bakend-POS/db"
	"Bakend-POS/models/request"
	"Bakend-POS/models/response"
	"Bakend-POS/tools/session_checking"
	"net/http"
	"strconv"
	"time"
)

func Input_News(Request request.Input_News_Request) (response.Response, error) {
	var res response.Response
	_, condition := session_checking.Session_Checking(Request.Uuid_session)

	if condition {

		con := db.CreateConGorm()
		co := 0

		err := con.Table("news").Select("co").Order("co DESC").Scan(&co)

		Request.Co = co + 1
		Request.Kode_news = "NW-" + strconv.Itoa(Request.Co)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

		date, _ := time.Parse("02-01-2006", Request.Date)
		Request.Date = date.Format("2006-01-02")

		err = con.Select("co", "kode_news", "date", "title", "content").Create(&Request)

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
		res.Message = "Session Invalid"
		res.Data = Request
	}

	return res, nil
}

// func Read_Supplier(Request request.Read_Supplier_Request) (response.Response, error) {
// 	var res response.Response
// 	var data []response.Read_Supplier_Response

// 	User, condition := session_checking.Session_Checking(Request.Uuid_session)

// 	if condition {

// 		con := db.CreateConGorm().Table("supplier")

// 		err := con.Select("kode_supplier", "nama_supplier", "email_supplier", "nomor_telepon").Where("kode_user = ?", User.Kode_user).Order("co ASC").Scan(&data).Error

// 		if err != nil {
// 			res.Status = http.StatusNotFound
// 			res.Message = "Status Not Found"
// 			res.Data = Request
// 			return res, err
// 		}

// 		if data == nil {
// 			res.Status = http.StatusNotFound
// 			res.Message = "Not Found"
// 			res.Data = data
// 		} else {
// 			res.Status = http.StatusOK
// 			res.Message = "Sukses"
// 			res.Data = data
// 		}
// 	} else {
// 		res.Status = http.StatusNotFound
// 		res.Message = "Session Invalid"
// 		res.Data = data
// 	}

// 	return res, nil
// }
