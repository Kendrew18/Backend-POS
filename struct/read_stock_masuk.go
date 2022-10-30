package _struct

type Read_Stock_Masuk struct {
	Id_stock_masuk string `json:"id_stock_masuk"`
	Kode_supplier  string `json:"kode_supplier"`
	Kode_stock     string `json:"kode_stock"`
	Tanggal_masuk  string `json:"tanggal_masuk"`
	Nama_supplier  string `json:"nama_supplier"`
	Jumlah_barang  string `json:"jumlah_barang"`
	Harga_barang   string `json:"harga_barang"`
}
