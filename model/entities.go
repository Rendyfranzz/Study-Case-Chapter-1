package model

type User struct {
	ID               string `json:"id" bson:"_id"`
	Email            string `json:"email" bson:"email"`
	Nama             string `json:"nama" bson:"nama"`
	Password         string `json:"-" bson:"password"`
	Nik              string `json:"nik" bson:"nik"`
	JenisPelakuUsaha string `json:"jenis_pelaku_usaha" bson:"jenis_pelaku_usaha"`
	Umk              string `json:"umk" bson:"umk"`
	NoPonsel         string `json:"no_ponsel" bson:"ponsel"`
	JenisKelamin     string `json:"jenis_kelamin" bson:"kelamin"`
	TanggalLahir     string `json:"tanggal_lahir" bson:"tanggal_lahir"`
	Alamat           string `json:"alamat" bson:"alamat"`
}

type NIB struct {
	ID            string `json:"id" bson:"_id"`
	NIB           string `json:"nib" bson:"nib"`
	Nama          string `json:"nama_perusahaan" bson:"nama"`
	StatusAktif   string `json:"status_keaktifan" bson:"aktif"`
	StatusMigrasi string `json:"status_migrasi" bson:"migrasi"`
}
