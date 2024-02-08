package _struct

type Insert_Stock_Masuk struct {
	Id_stock_masuk        string `json:"id_stock_masuk"`
	Kode_supplier         string `json:"kode_supplier"`
	Nama_penanggung_jawab string `json:"nama_penanggung_jawab"`
	Kode_stock            string `json:"kode_stock"`
	Nama_stock            string `json:"nama_stock"`
	Tanggal_masuk         string `json:"tanggal_masuk"`
	Jumlah_barang         string `json:"jumlah_barang"`
	Harga_barang          string `json:"harga_barang"`
}
