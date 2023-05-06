package service

import (
	"context"
	"time"

	"github.com/Rendyfranzz/Study-Case-Chapter-1/model"
	"github.com/Rendyfranzz/Study-Case-Chapter-1/repo"
	"golang.org/x/crypto/bcrypt"
)

type JWTGenerator interface {
	Generate(expr time.Duration, cred string) (string, error)
}

type AuthService struct {
	userRepo *repo.UserRepo
	jwt      JWTGenerator
}

func NewAuthService(repo *repo.UserRepo, jwt JWTGenerator) *AuthService {
	return &AuthService{userRepo: repo, jwt: jwt}
}

type LoginInput struct {
	CommonRequest
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginOutput struct {
	CommonResponse
}

func (a *AuthService) Login(ctx context.Context, in LoginInput) LoginOutput {
	var out LoginOutput

	if in.Password == "" {
		out.SetMsg(400, "password is required")
		return out
	}

	if in.Email == "" {
		out.SetMsg(400, "email is required")
		return out
	}

	user, err := a.userRepo.GetBy(ctx, in.Email)
	if (err != nil) || (user == (model.User{})) {
		out.SetMsg(400, "not found")
		return out
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password)); err != nil {
		out.SetMsg(400, "bad credentials")
		return out
	}

	token, err := a.jwt.Generate(time.Minute*2, user.Nama)
	if err != nil {
		out.SetMsg(500, "can't generate token")
		return out
	}

	out.SetAuthToken = token
	out.SetMsg(200, "Authenticated")

	return out
}

type RegisterInput struct {
	CommonRequest
	Email            string `json:"email"`
	Nama             string `json:"nama"`
	Password         string `json:"password"`
	Nik              string `json:"nik"`
	JenisPelakuUsaha string `json:"jenis_pelaku_usaha"`
	Umk              string `json:"umk"`
	NoPonsel         string `json:"no_ponsel"`
	JenisKelamin     string `json:"jenis_kelamin"`
	TanggalLahir     string `json:"tanggal_lahir"`
	Alamat           string `json:"alamat"`
}

type RegisterOutput struct {
	CommonResponse
}

func (a *AuthService) Register(ctx context.Context, in RegisterInput) RegisterOutput {
	var out RegisterOutput

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		out.SetMsg(500, err.Error())
		return out
	}

	user := model.User{
		Nama:             in.Nama,
		Email:            in.Email,
		Nik:              in.Nik,
		JenisPelakuUsaha: in.JenisPelakuUsaha,
		Umk:              in.Umk,
		NoPonsel:         in.NoPonsel,
		JenisKelamin:     in.JenisKelamin,
		TanggalLahir:     in.TanggalLahir,
		Alamat:           in.Alamat,
		Password:         string(hashPassword),
	}

	if err := a.userRepo.Insert(ctx, user); err != nil {
		out.SetMsg(500, err.Error())
		return out
	}

	out.SetMsg(200, "Registered")
	return out
}
