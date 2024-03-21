package request

type Input_Inventory_Request struct {
	Co             int    `json:"co"`
	Kode_inventory string `json:"kode_inventory"`
	Nama_barang    string `json:"nama_barang"`
	Harga_jual     int64  `json:"harga_jual"`
	Satuan_barang  string `json:"satuan_barang"`
	Kode_user      string `json:"kode_user"`
	Path_photo     string `json:"path_photo"`
}

type Read_Inventory_Request struct {
	Kode_user string `json:"kode_user"`
}

type Update_Inventory_Request struct {
	Kode_inventory string `json:"kode_inventory"`
	Nama_barang    string `json:"nama_barang"`
	Harga_jual     int64  `json:"harga_jual"`
	Satuan_barang  string `json:"satuan_barang"`
	Kode_user      string `json:"kode_user"`
}

type Check_Nama_Inventory_Request struct {
	Kode_inventory string `json:"kode_inventory"`
	Nama_barang    string `json:"nama_barang"`
	Kode_user      string `json:"kode_user"`
}

type Dropdown_Inventory_transaksi_inventory_request struct {
	Kode_user string `json:"kode_user"`
	Kode_nota string `json:"kode_nota"`
}
