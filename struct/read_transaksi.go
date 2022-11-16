package _struct

type Read_Transaksi struct {
	Kode_transaksi    string `json:"kode_transaksi"`
	Tanggal_penjualan string `json:"tanggal_penjualan"`
	Tanggal_pelunasan string `json:"tanggal_pelunasan"`
	Status_transaksi  string `json:"status_transaksi"`
}
