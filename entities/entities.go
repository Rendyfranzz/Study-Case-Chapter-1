package entities

type User struct {
	ID                 string `json:"id" bson:"_id"`
	Email              string `json:"email" bson:"email"`
	Nama               string `json:"nama" bson:"nama"`
	Password           string `json:"password" bson:"password"`
	Nik                string `json:"nik" bson:"nik"`
	Jenis_Pelaku_usaha string `json:"jenis_pelaku_usaha" bson:"jenis_pelaku_usaha"`
	Umk                string `json:"umk" bson:"umk"`
	No_ponsel          string `json:"no_ponsel" bson:"ponsel"`
	Jenis_Kelamin      string `json:"jenis_kelamin" bson:"kelamin"`
	Tanggal_Lahir      string `json:"tanggal_lahir" bson:"tanggal_lahir"`
	Alamat             string `json:"alamat" bson:"alamat"`
}

type Nib struct {
	ID             string `json:"id" bson:"_id"`
	NIB            string `json:"nib" bson:"nib"`
	Nama           string `json:"nama_perusahaan" bson:"nama"`
	Status_aktif   string `json:"status_keaktifan" bson:"aktif"`
	Status_migrasi string `json:"status_migrasi" bson:"migrasi"`
}

type Pengajuan struct {
	ID            string `json:"id" bson:"_id"`
	Nik           string `json:"nik" bson:"nik"`
	Nama          string `json:"nama" bson:"nama"`
	No_ponsel     string `json:"no_ponsel" bson:"ponsel"`
	Jenis_Kelamin string `json:"jenis_kelamin" bson:"kelamin"`
	Alamat        string `json:"alamat" bson:"alamat"`
	Npwp          string `json:"npwp" bson:"npwp"`
	No_bpjs       string `json:"no_bpjs" bson:"no_bpjs"`
}

type User_login struct {
	ID            string `json:"id" bson:"_id"`
	Nama          string `json:"nama" bson:"nama"`
	Nik           string `json:"nik" bson:"nik"`
	No_ponsel     string `json:"no_ponsel" bson:"ponsel"`
	Jenis_Kelamin string `json:"jenis_kelamin" bson:"kelamin"`
	Alamat        string `json:"alamat" bson:"alamat"`
}

type News struct {
	ID        string `json:"id" bson:"_id"`
	Judul     string `json:"judul" bson:"judul"`
	Deskripsi string `json:"deskripsi" bson:"deskripsi"`
	Tanggal   string `json:"tanggal" bson:"tanggal"`
}
