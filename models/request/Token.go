package request

type Token_Request struct {
	Token string `json:"token"`
}

type Token struct {
	Kode_user    string `json:"kode_user"`
	Uuid_session string `json:"uuid_session"`
}
