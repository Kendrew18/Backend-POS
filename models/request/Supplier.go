package request

type Input_Supplier_Request struct {
	Co             int    `json:"co"`
	Kode_supplier  string `json:"kode_supplier"`
	Nama_supplier  string `json:"nama_supplier"`
	Nomor_telepon  string `json:"nomor_telepon"`
	Email_supplier string `json:"email_supplier"`
	Kode_user      string `json:"kode_user"`
}

type Read_Supplier_Request struct {
	Kode_user string `json:"kode_user"`
}

type Input_barang_Supplier_Request struct {
	Co                   int    `json:"co"`
	Kode_barang_supplier string `json:"kode_barang_supplier"`
	Kode_supplier        string `json:"kode_supplier"`
	Kode_inventory       string `json:"kode_inventory"`
}

type Update_Supplier_Request struct {
	Kode_supplier  string `json:"kode_supplier"`
	Nomor_telepon  string `json:"nomor_telepon"`
	Email_supplier string `json:"email_supplier"`
}

type Delete_Supplier_Request struct {
	Kode_supplier string `json:"kode_supplier"`
}
