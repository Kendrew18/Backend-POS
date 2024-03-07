package news

import (
	"Bakend-POS/db"
	"Bakend-POS/models/request"
	"Bakend-POS/models/response"
	"Bakend-POS/tools/session_checking"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func Input_News(Request request.Input_News_Request, writer http.ResponseWriter, request *http.Request) (response.Response, error) {
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

		request.ParseMultipartForm(10 * 1024 * 1024)
		file, handler, err2 := request.FormFile("photo")
		if err2 != nil {
			fmt.Println(err2)
			return res, err2
		}

		defer file.Close()

		fmt.Println("File Info")
		fmt.Println("File Name : ", handler.Filename)
		fmt.Println("File Size : ", handler.Size)
		fmt.Println("File Type : ", handler.Header.Get("Content-Type"))

		var tempFile *os.File
		path := ""

		if strings.Contains(handler.Filename, "jpg") {
			path = "uploads/foto_laporan_vendor/" + Request.Kode_news + ".jpg"
			tempFile, err2 = ioutil.TempFile("uploads/foto_laporan_vendor/", "Read"+"*.jpg")
		}
		if strings.Contains(handler.Filename, "jpeg") {
			path = "uploads/foto_laporan_vendor/" + Request.Kode_news + ".jpeg"
			tempFile, err2 = ioutil.TempFile("uploads/foto_laporan_vendor/", "Read"+"*.jpeg")
		}
		if strings.Contains(handler.Filename, "png") {
			path = "uploads/foto_laporan_vendor/" + Request.Kode_news + ".png"
			tempFile, err2 = ioutil.TempFile("uploads/foto_laporan_vendor/", "Read"+"*.png")
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

		Request.Image_path = path

		err = con.Select("co", "kode_news", "date", "title", "content", "image_path").Create(&Request)

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

func Read_News(Request request.Read_News_Request) (response.Response, error) {
	var res response.Response
	var data []response.Read_News_Response

	_, condition := session_checking.Session_Checking(Request.Uuid_session)

	if condition {

		con := db.CreateConGorm().Table("news")

		err := con.Select("kode_news", "date", "title", "content", "image_path").Order("co DESC").Scan(&data).Error

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err
		}

		if data == nil {
			res.Status = http.StatusNotFound
			res.Message = "Not Found"
			res.Data = data
		} else {
			res.Status = http.StatusOK
			res.Message = "Sukses"
			res.Data = data
		}
	} else {
		res.Status = http.StatusNotFound
		res.Message = "Session Invalid"
		res.Data = data
	}

	return res, nil
}
