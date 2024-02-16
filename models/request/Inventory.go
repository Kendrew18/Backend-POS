package request

type Input_Inventory_Request struct {
	Co             int    `json:"co"`
	Kode_inventory string `json:"kode_inventory"`
	Nama_barang    string `json:"nama_barang"`
	Harga_jual     int64  `json:"harga_jual"`
	Satuan_barang  string `json:"satuan_barang"`
	Kode_user      string `json:"kode_user"`
}

type Read_Inventory_Request struct {
	Kode_user string `json:"kode_user"`
}

type Read_Inventory_Filter_Request struct {
	Status_ASC_DESC int    `json:"status_asc_desc"`
	Tanggal_awal    string `json:"tanggal_awal"`
	Tanggal_akhir   string `json:"tanggal_akhir"`
}

type Update_Inventory_Request struct {
	Kode_inventory string `json:"kode_inventory"`
	Nama_barang    string `json:"nama_barang"`
	Harga_jual     int64  `json:"harga_jual"`
	Satuan_barang  string `json:"satuan_barang"`
}

type Check_Nama_Inventory_Request struct {
	Kode_inventory string `json:"kode_inventory"`
	Nama_barang    string `json:"nama_barang"`
	Kode_user      string `json:"kode_user"`
}
