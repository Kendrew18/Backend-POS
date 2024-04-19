package response

type Read_News_Response struct {
	Co         int    `json:"co"`
	Kode_news  string `json:"kode_news"`
	Date       string `json:"date"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Image_path string `json:"image_path"`
}

type Read_Content_Response struct {
	Kode_content string `json:"kode_content"`
	Content      string `json:"content"`
}
