package request

type Input_News_Request struct {
	Co         int    `json:"co"`
	Kode_news  string `json:"kode_news"`
	Date       string `json:"date"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Image_path string `json:"image_path"`
	Kode_user  string `json:"kode_user"`
}

type Read_News_Request struct {
	Kode_user string `json:"kode_user"`
}
