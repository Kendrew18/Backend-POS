package request

type Login_Request struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Sign_Up_Request struct {
	Co              int    `json:"co"`
	Kode_user       string `json:"kode_user"`
	Nama_lengkap    string `json:"nama_lengkap"`
	Birth_date      string `json:"birth_date"`
	Gender          string `json:"gender"`
	Category_bisnis string `json:"category_bisnis"`
	Nama_bisnis     string `json:"nama_bisnis"`
	Alamat_bisnis   string `json:"alamat_bisnis"`
	Telepon_bisnis  string `json:"telepon_bisnis"`
	Email_bisnis    string `json:"email_bisnis"`
	Instagram       string `json:"instagram"`
	Facebook        string `json:"facebook"`
	Username        string `json:"username"`
	Password        string `json:"password"`
}

type Profile_User_Request struct {
	Uuid_session string `json:"uuid_session"`
}

type Update_Profile_User_Request struct {
	Kode_user       string `json:"kode_user"`
	Nama_lengkap    string `json:"nama_lengkap"`
	Birth_date      string `json:"birth_date"`
	Gender          string `json:"gender"`
	Category_bisnis string `json:"category_bisnis"`
	Nama_bisnis     string `json:"nama_bisnis"`
	Alamat_bisnis   string `json:"alamat_bisnis"`
	Telepon_bisnis  string `json:"telepon_bisnis"`
	Email_bisnis    string `json:"email_bisnis"`
	Instagram       string `json:"instagram"`
	Facebook        string `json:"facebook"`
}
