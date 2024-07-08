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
	Key             string `json:"key"`
	Status          int    `json:"status"`
}

type Profile_User_Request struct {
	Kode_user string `json:"kode_user"`
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

type Otp_Request struct {
	Kode_user    string `json:"kode_user"`
	Nama_lengkap string `json:"nama_lengkap"`
	Email        string `json:"email"`
	Kode_otp     string `json:"kode_otp"`
	Time_sent    string `json:"time_sent"`
}

type Resend_OTP_Request struct {
	Email        string `json:"email"`
	Nama_lengkap string `json:"nama_lengkap"`
	Time_sent    string `json:"time_sent"`
}

type Update_OTP_Request struct {
	Kode_otp  string `json:"kode_otp"`
	Time_sent string `json:"time_sent"`
}

type Activate_Account_Request struct {
	Email    string `json:"email"`
	Kode_otp string `json:"kode_otp"`
}

type Sign_Up_Google struct {
	Authtoken string `json:"authtoken"`
	Aud       string `json:"aud"`
}

type ProjectInfo struct {
	ProjectNumber string `json:"project_number"`
	ProjectID     string `json:"project_id"`
	StorageBucket string `json:"storage_bucket"`
}

type ClientInfo struct {
	MobileSDKAppID    string `json:"mobilesdk_app_id"`
	AndroidClientInfo struct {
		PackageName string `json:"package_name"`
	} `json:"android_client_info"`
}

type OAuthClient struct {
	ClientID    string `json:"client_id"`
	ClientType  int    `json:"client_type"`
	AndroidInfo struct {
		PackageName     string `json:"package_name"`
		CertificateHash string `json:"certificate_hash"`
	} `json:"android_info,omitempty"`
}

type APIKey struct {
	CurrentKey string `json:"current_key"`
}

type AppInviteService struct {
	OtherPlatformOAuthClient []OAuthClient `json:"other_platform_oauth_client"`
}

type Services struct {
	AppInviteService AppInviteService `json:"appinvite_service"`
}

type Client struct {
	ClientInfo  ClientInfo    `json:"client_info"`
	OAuthClient []OAuthClient `json:"oauth_client"`
	APIKey      []APIKey      `json:"api_key"`
	Services    Services      `json:"services"`
}

type Configuration struct {
	ProjectInfo          ProjectInfo `json:"project_info"`
	Client               []Client    `json:"client"`
	ConfigurationVersion string      `json:"configuration_version"`
}

type Condition struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Status   int    `json:"status"`
}
