package _struct

type Read_Pembukuan_Transaksi_Bulanan struct {
	Id_pembukuann_transaksi_bulanan string `json:"id_pembukuann_transaksi_bulanan"`
	Kode_stock                      string `json:"kode_stock"`
	Nama_barang                     string `json:"nama_barang"`
	Jumlah_barang                   string `json:"jumlah_barang"`
	Harga_barang                    string `json:"harga_barang"`
	Tanggal_pelunasan               string `json:"tanggal_pelunasan"`
	Total_harga_penjualan           int64  `json:"total_harga_penjualan"`
}

type Read_Pembukuan_Transaksi_Bulanan_List struct {
	Id_pembukuann_transaksi_bulanan string   `json:"id_pembukuann_transaksi_bulanan"`
	Kode_stock                      []string `json:"kode_stock"`
	Nama_barang                     []string `json:"nama_barang"`
	Jumlah_barang                   []int    `json:"jumlah_barang"`
	Harga_barang                    []int    `json:"harga_barang"`
	Tanggal_pelunasan               string   `json:"tanggal_pelunasan"`
	Total_harga_penjualan           int64    `json:"total_harga_penjualan"`
}
