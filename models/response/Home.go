package response

type Home_Response struct {
	Total_pemasukan          string              `json:"total_pemasukan"`
	Total_pengeluaran        string              `json:"total_pengeluaran"`
	Total_pembayaran_pending string              `json:"total_pembayaran_pending"`
	Chart_Pemasukan          []Chart_Pemasukan   `json:"chart_pemasukan" gorm:"-"`
	Chart_Pengeluaran        []Chart_Pengeluaran `json:"chart_pengeluaran" gorm:"-"`
}
type Chart_Pengeluaran struct {
	Tanggal           string `json:"tanggal"`
	Total_pengeluaran string `json:"total_pengeluaran"`
}

type Chart_Pemasukan struct {
	Tanggal         string `json:"tanggal"`
	Total_pemasukan string `json:"total_pemasukan"`
}
