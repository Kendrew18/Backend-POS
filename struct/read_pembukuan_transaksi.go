package _struct

type Read_Pembukuan_Transaksi struct {
	Id_pembukuan_transaksi string `json:"id_pembukuan_transaksi"`
	Kode_stock             string `json:"kode_stock"`
	Nama_barang            string `json:"nama_barang"`
	Jumlah_barang          string `json:"jumlah_barang"`
	Harga_barang           string `json:"harga_barang"`
	Tanggal_penjualan      string `json:"tanggal_penjualan"`
	Total_harga_penjualan  int64  `json:"total_harga_penjualan"`
}
