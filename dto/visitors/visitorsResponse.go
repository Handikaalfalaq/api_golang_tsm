package visitorsdto

type VisitorsResponse struct {
	ID          int    `json:"id"`
	Nama        string `json:"nama" form:"nama"`
	NoKTP       string `json:"no_ktp" form:"no_ktp"`
	TglLahir    string `json:"tgl_lahir" form:"tgl_lahir"`
	NamaGedung  string `json:"nama_gedung" form:"nama_gedung"`
	Suhu        string `json:"suhu" form:"suhu"`
	JamMasuk    string `json:"jam_masuk" form:"jam_masuk"`
	JamKeluar   string `json:"jam_keluar" form:"jam_keluar"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
}
