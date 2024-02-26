package request

type Input_Transaksi_Inventory_Request struct {
	Co                       int     `json:"co"`
	Kode_transaksi_inventory string  `json:"kode_transaksi_inventory"`
	Nama_supplier            string  `json:"nama_supplier"`
	Nomor_telpon_supplier    string  `json:"nomor_telpon_supplier"`
	Tanggal                  string  `json:"tanggal"`
	Kode_nota                string  `json:"kode_nota"`
	Harga_ongkos_kirim       int64   `json:"harga_ongkos_kirim"`
	Ppn                      float64 `json:"ppn"`
	Kode_user                string  `json:"kode_user"`
	Jenis_transaksi          int     `json:"jenis_transaksi"`
	Uuid_session             string  `json:"uuid_session"`
}

type Input_Barang_Transaksi_Inventory_Request struct {
	Kode_inventory string `json:"kode_inventory"`
	Jumlah         string `json:"jumlah"`
	Harga          string `json:"harga"`
	Kode_refund    string `json:"kode_refund"`
}

type Input_Barang_Transaksi_Inventory_V2_Request struct {
	Co                              int     `json:"co"`
	Kode_barang_transaksi_inventory string  `json:"kode_barang_transaksi_inventory"`
	Kode_transaksi_inventory        string  `json:"kode_transaksi_inventory"`
	Kode_inventory                  string  `json:"kode_inventory"`
	Jumlah                          float64 `json:"jumlah"`
	Harga                           int64   `json:"harga"`
	Sub_total                       int64   `json:"sub_total"`
	Kode_refund                     string  `json:"kode_refund"`
}

type Read_Transaksi_Inventory_Request struct {
	Uuid_session string `json:"uuid_session"`
}

type Read_Transaksi_Inventory_Filter_Request struct {
	Tanggal_awal  string `json:"tanggal_awal"`
	Tanggal_akhir string `json:"tanggal_akhir"`
	Nama_supplier string `json:"nama_supplier"`
}

type Update_Header_Transaksi_Inventory_Kode_Request struct {
	Kode_transaksi_inventory string `json:"kode_transaksi_inventory"`
	Uuid_session             string `json:"uuid_session"`
}

type Update_Header_Transaksi_Inventory_Request struct {
	Nama_supplier         string  `json:"nama_supplier"`
	Nomor_telpon_supplier string  `json:"nomor_telpon_supplier"`
	Kode_nota             string  `json:"kode_nota"`
	Harga_ongkos_kirim    int64   `json:"harga_ongkos_kirim"`
	Ppn                   float64 `json:"ppn"`
	Total_harga           int64   `json:"total_harga"`
}

type Update_Barang_Transaksi_Inventory_Kode_Request struct {
	Kode_barang_transaksi_inventory string `json:"kode_barang_transaksi_inventory"`
	Uuid_session                    string `json:"uuid_session"`
}

type Update_Barang_Transaksi_Inventory_Request struct {
	Kode_inventory string  `json:"kode_inventory"`
	Jumlah         float64 `json:"jumlah"`
	Harga          int64   `json:"harga"`
	Sub_total      int64   `json:"sub_total"`
}

type Update_Status_Transaksi_Inventory_Request struct {
	Status int `json:"status"`
}
