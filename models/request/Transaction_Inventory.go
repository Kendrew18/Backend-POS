package request

type Input_Transaksi_Body_Request struct {
	Transaksi_inventory        Input_Transaksi_Inventory_Request          `json:"transaksi_inventory"`
	Barang_transaksi_inventory []Input_Barang_Transaksi_Inventory_Request `json:"barang_transaksi_inventory"`
}

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
}

type Input_Barang_Transaksi_Inventory_Request struct {
	Co                              int     `json:"co"`
	Kode_barang_transaksi_inventory string  `json:"kode_barang_transaksi_inventory"`
	Kode_transaksi_inventory        string  `json:"kode_transaksi_inventory"`
	Kode_inventory                  string  `json:"kode_inventory"`
	Jumlah                          float64 `json:"jumlah"`
	Harga                           int64   `json:"harga"`
	Sub_total                       int64   `json:"sub_total"`
	Kode_refund                     string  `json:"kode_refund"`
}

type Body_Read_Transaksi_Inventory_Request struct {
	Read_transaksi_inventory        Read_Transaksi_Inventory_Request        `json:"read_transaksi_inventory"`
	Read_transaksi_inventory_filter Read_Transaksi_Inventory_Filter_Request `json:"read_transaksi_inventory_filter"`
}

type Read_Transaksi_Inventory_Request struct {
	Kode_user string `json:"kode_user"`
}

type Read_Transaksi_Inventory_Filter_Request struct {
	Tanggal_awal  string `json:"tanggal_awal"`
	Tanggal_akhir string `json:"tanggal_akhir"`
	Nama_supplier string `json:"nama_supplier"`
}

type Body_Update_Header_Transaksi_Inventory_Request struct {
	Update_header_transaksi_inventory_kode Update_Header_Transaksi_Inventory_Kode_Request `json:"update_header_transaksi_inventory_kode"`
	Update_header_transaksi_inventory      Update_Header_Transaksi_Inventory_Request      `json:"update_header_transaksi_inventory"`
}

type Update_Header_Transaksi_Inventory_Kode_Request struct {
	Kode_transaksi_inventory string `json:"kode_transaksi_inventory"`
	Kode_user                string `json:"kode_user"`
}

type Update_Header_Transaksi_Inventory_Request struct {
	Nama_supplier         string  `json:"nama_supplier"`
	Nomor_telpon_supplier string  `json:"nomor_telpon_supplier"`
	Kode_nota             string  `json:"kode_nota"`
	Harga_ongkos_kirim    int64   `json:"harga_ongkos_kirim"`
	Ppn                   float64 `json:"ppn"`
	Total_harga           int64   `json:"total_harga"`
}

type Body_Update_Barang_Transaksi_Inventory struct {
	Update_barang_transaksi_inventory_kode Update_Barang_Transaksi_Inventory_Kode_Request `json:"update_barang_transaksi_inventory_kode"`
	Update_barang_transaksi_inventory      Update_Barang_Transaksi_Inventory_Request      `json:"update_barang_transaksi_inventory"`
}

type Update_Barang_Transaksi_Inventory_Kode_Request struct {
	Kode_barang_transaksi_inventory string `json:"kode_barang_transaksi_inventory"`
	Kode_user                       string `json:"kode_user"`
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

type Body_Update_Status_Transaksi_inventory struct {
	Update_header_transaksi_inventory_kode Update_Header_Transaksi_Inventory_Kode_Request `json:"update_header_transaksi_inventory_kode"`
	Update_status_transaksi_inventory      Update_Status_Transaksi_Inventory_Request      `json:"update_status_transaksi_inventory"`
}
