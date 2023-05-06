package service

import (
	"context"

	"github.com/Rendyfranzz/Study-Case-Chapter-1/model"
	"github.com/Rendyfranzz/Study-Case-Chapter-1/repo"
)

type NIBService struct {
	nibRepo *repo.NIBRepo
}

func NewNIBService(nibRepo *repo.NIBRepo) *NIBService {
	return &NIBService{nibRepo: nibRepo}
}

type GetNIBInput struct {
	CommonRequest
	ID string `json:"-"`
}

type GetNIBOutput struct {
	CommonResponse
	NIB model.NIB `json:"nib"`
}

func (n *NIBService) GetNIB(ctx context.Context, in GetNIBInput) GetNIBOutput {
	var out GetNIBOutput

	nib, err := n.nibRepo.GetBy(ctx, in.ID)
	if err != nil {
		out.SetMsg(400, err.Error())
		return out
	}

	out.NIB = nib
	out.SetMsg(200, "Success")
	return out
}

type NIBRegisterInput struct {
	CommonRequest
	NIB           string `json:"nib"`
	Nama          string `json:"nama_perusahaan"`
	StatusAktif   string `json:"status_keaktifan"`
	StatusMigrasi string `json:"status_migrasi"`
}

type NIBRegisterOutput struct {
	CommonResponse
}

func (n *NIBService) RegisterNIB(ctx context.Context, in NIBRegisterInput) NIBRegisterOutput {
	var out NIBRegisterOutput

	nib := model.NIB{
		NIB:           in.NIB,
		Nama:          in.Nama,
		StatusAktif:   in.StatusAktif,
		StatusMigrasi: in.StatusMigrasi,
	}

	err := n.nibRepo.Insert(ctx, nib)
	if err != nil {
		out.SetMsg(500, err.Error())
		return out
	}

	out.SetMsg(201, "Created")
	return out
}
