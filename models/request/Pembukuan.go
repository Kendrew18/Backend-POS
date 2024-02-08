package request

type Input_Pembukuan_Request struct {
	Co                    int     `json:"co"`
	Kode_pembukuan        string  `json:"kode_pembukuan"`
	Kode_nota             string  `json:"kode_nota"`
	Tanggal               string  `json:"tanggal"`
	Kode_jenis_pembayaran string  `json:"kode_jenis_pembayaran"`
	Diskon                int64   `json:"diskon"`
	Total_harga           int64   `json:"total_harga"`
	Total_barang          float64 `json:"total_barang"`
	Kode_user             string  `json:"kode_user"`
}

type Input_Barang_Pembukuan_Request struct {
	Kode_stock    string `json:"kode_stock"`
	Jumlah        string `json:"jumlah"`
	Satuan_barang string `json:"satuan_barang"`
	Harga         string `json:"harga"`
	Sub_total     string `json:"sub_total"`
}

type Input_Barang_Pembukuan_V2_Request struct {
	Co                    int     `json:"co"`
	Kode_barang_pembukuan string  `json:"kode_barang_pembukuan"`
	Kode_pembukuan        string  `json:"kode_pembukuan"`
	Kode_stock            string  `json:"kode_stock"`
	Jumlah                float64 `json:"jumlah"`
	Satuan_barang         string  `json:"satuan_barang"`
	Harga                 int64   `json:"harga"`
	Sub_total             int64   `json:"sub_total"`
}

type Read_Pembukuan_Request struct {
	Status    int    `json:"status"`
	Kode_user string `json:"kode_user"`
}
type Read_Pembukuan_Filter_Request struct {
	Tanggal_awal  string `json:"tanggal_awal"`
	Tanggal_akhir string `json:"tanggal_akhir"`
}
