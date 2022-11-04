package _struct

type Input_Transaksi struct {
	Kode_stock        string `json:"kode_stock"`
	Jumlah_barang     string `json:"jumlah_barang"`
	Harga_barang      string `json:"harga_barang"`
	Tanggal_penjualan string `json:"tanggal_penjualan"`
	Tanggal_pelunasan string `json:"tanggal_pelunasan"`
}
