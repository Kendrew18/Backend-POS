package _struct

type Read_Retur struct {
	Id_retur      string  `json:"id_retur"`
	Id_supplier   string  `json:"id_supplier"`
	Nama_supplier string  `json:"nama_supplier"`
	Kode_stock    string  `json:"kode_stock"`
	Nama_barang   string  `json:"nama_barang"`
	Tanggal_retur string  `json:"tanggal_retur"`
	Jumlah_barang float64 `json:"jumlah_barang"`
	Status_retur  int     `json:"status_retur"`
	Keterangan    string  `json:"keterangan"`
}
