package response

type Read_Supplier_Response struct {
	Kode_supplier  string `json:"kode_supplier"`
	Nama_supplier  string `json:"nama_supplier"`
	Email_supplier string `json:"email_supplier"`
	Nomor_telepon  string `json:"nomor_telepon"`
}
