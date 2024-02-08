package request

type Login_Request struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Sign_Up_Request struct {
	Co              int    `json:"co"`
	Kode_user       string `json:"kode_user"`
	Nama_lengkap    string `json:"nama_lengkap"`
	Birth_date      string `json:"birth_date"`
	Email           string `json:"email"`
	Category_bisnis string `json:"category_bisnis"`
	Nama_bisnis     string `json:"nama_bisnis"`
	Alamat_bisnis   string `json:"alamat_bisnis"`
	Instagram       string `json:"instagram"`
	Facebook        string `json:"facebook"`
	Password        string `json:"password"`
}
type Profile_User_Request struct {
	Kode_user string `json:"kode_user"`
}
