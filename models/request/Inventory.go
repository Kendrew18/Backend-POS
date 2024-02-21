package request

type Input_Inventory_Request struct {
	Co             int    `json:"co"`
	Kode_inventory string `json:"kode_inventory"`
	Nama_barang    string `json:"nama_barang"`
	Harga_jual     int64  `json:"harga_jual"`
	Satuan_barang  string `json:"satuan_barang"`
	Kode_user      string `json:"kode_user"`
	Uuid_session   string `json:"uuid_session"`
}

type Read_Inventory_Request struct {
	Uuid_session string `json:"uuid_session"`
}

type Update_Inventory_Request struct {
	Kode_inventory string `json:"kode_inventory"`
	Nama_barang    string `json:"nama_barang"`
	Harga_jual     int64  `json:"harga_jual"`
	Satuan_barang  string `json:"satuan_barang"`
	Uuid_session   string `json:"uuid_session"`
}

type Check_Nama_Inventory_Request struct {
	Kode_inventory string `json:"kode_inventory"`
	Nama_barang    string `json:"nama_barang"`
	Kode_user      string `json:"kode_user"`
}
