package model

type User struct {
	ID       string `json:"id" bson:"_id"`
	Email    string `json:"email" bson:"email"`
	Name     string `json:"name" bson:"name"`
	Password string `json:"password" bson:"password"`
}

type Nib struct {
	ID         string `json:"id" bson:"_id"`
	Perusahaan string `json:"perusahaan" bson:"perusahaan"`
	Name       string `json:"name" bson:"name"`
	Password   string `json:"password" bson:"password"`
}
