package response

type Login_Response struct {
	Token  string `json:"token"`
	Status int    `json:"status"`
}

type User_Profile_Response struct {
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
	Status          int    `json:"status"`
}

type User_Session_Response struct {
	Kode_user string `json:"kode_user"`
	Status    int    `json:"status"`
}

type Sign_Up_Google_Response struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
