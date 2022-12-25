package _struct

type Read_Transaksi struct {
	Kode_transaksi      string `json:"kode_transaksi"`
	Tanggal_penjualan   string `json:"tanggal_penjualan"`
	Tanggal_pelunasan   string `json:"tanggal_pelunasan"`
	Status_transaksi    string `json:"status_transaksi"`
	Sub_total_harga     int64  `json:"sub_total_harga"`
	Jumlah_barang       string `json:"jumlah_barang"`
	Total_jumlah_barang int    `json:"total_jumlah_barang"`
}
