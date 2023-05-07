package service

import (
	"context"

	"github.com/Rendyfranzz/Study-Case-Chapter-1/model"
	"github.com/Rendyfranzz/Study-Case-Chapter-1/repo"
)

type SubmissionService struct {
	submissionRepo *repo.SubmissionRepo
}

func NewSubmissionService(submissionRepo *repo.SubmissionRepo) *SubmissionService {
	return &SubmissionService{submissionRepo: submissionRepo}
}

type SubmissionRegisterInput struct {
	CommonRequest
	Nik          string `json:"nik"`
	Nama         string `json:"nama"`
	NoPonsel     string `json:"no_ponsel"`
	JenisKelamin string `json:"jenis_kelamin"`
	Alamat       string `json:"alamat"`
	Npwp         string `json:"npwp"`
	NoBpjs       string `json:"no_bpjs"`
}

type SubmissionRegisterOutput struct {
	CommonResponse
}

func (n *SubmissionService) RegisterSubmission(ctx context.Context, in SubmissionRegisterInput) SubmissionRegisterOutput {
	var out SubmissionRegisterOutput

	submission := model.Submission{
		Nik:          in.Nik,
		Nama:         in.Nama,
		NoPonsel:     in.NoPonsel,
		JenisKelamin: in.JenisKelamin,
		Alamat:       in.Alamat,
		Npwp:         in.Npwp,
		NoBpjs:       in.NoBpjs,
	}

	err := n.submissionRepo.Insert(ctx, submission)
	if err != nil {
		out.SetMsg(500, err.Error())
		return out
	}

	out.SetMsg(201, "Created")
	return out
}

type SubmissionEditInput struct {
	CommonRequest
	Nik          string `json:"nik"`
	Nama         string `json:"nama"`
	NoPonsel     string `json:"no_ponsel"`
	JenisKelamin string `json:"jenis_kelamin"`
	Alamat       string `json:"alamat"`
	Npwp         string `json:"npwp"`
	NoBpjs       string `json:"no_bpjs"`
}
type SubmissionEditKey struct {
	CommonRequest
	Search string `json:"-"`
}
type SubmissionEditOutput struct {
	CommonResponse
}

func (n *SubmissionService) EditSubmission(ctx context.Context, in SubmissionEditInput, key SubmissionEditKey) SubmissionEditOutput {
	var out SubmissionEditOutput

	submission := model.Submission{
		Nik:          in.Nik,
		Nama:         in.Nama,
		NoPonsel:     in.NoPonsel,
		JenisKelamin: in.JenisKelamin,
		Alamat:       in.Alamat,
		Npwp:         in.Npwp,
		NoBpjs:       in.NoBpjs,
	}
	err := n.submissionRepo.Update(ctx, submission, key.Search)
	if err != nil {
		out.SetMsg(500, err.Error())
		return out
	}

	out.SetMsg(201, "Updated")
	return out
}
