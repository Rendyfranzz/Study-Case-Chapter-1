package repo

import (
	"context"

	"github.com/Rendyfranzz/Study-Case-Chapter-1/config"
	"github.com/Rendyfranzz/Study-Case-Chapter-1/model"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
)

type SubmissionRepo struct {
	db *config.DB
}

func NewSubmissionRepo(db *config.DB) *SubmissionRepo {
	return &SubmissionRepo{db: db}
}

func (n *SubmissionRepo) GetBy(ctx context.Context, nik string) (model.Submission, error) {
	var submission model.Submission

	collection := n.db.Client.Database("oss").Collection("submission")

	err := collection.FindOne(ctx, bson.M{"nik": nik}).Decode(&submission)
	return submission, err
}

func (n *SubmissionRepo) Insert(ctx context.Context, submission model.Submission) error {
	submission.ID = xid.New().String()
	collection := n.db.Client.Database("oss").Collection("submission")
	_, err := collection.InsertOne(ctx, &submission)

	return err
}

func (n *SubmissionRepo) Update(ctx context.Context, submission model.Submission, key string) error {
	var temp model.Submission
	collection := n.db.Client.Database("oss").Collection("submission")
	err := collection.FindOne(ctx, bson.M{"nik": key}).Decode(&temp)
	if err != nil {
		return err
	}

	edited := bson.M{"npwp": submission.Npwp,
		"no_bpjs": submission.NoBpjs,
		"nik":     submission.Nik,
		"nama":    submission.Nama,
		"ponsel":  submission.NoPonsel,
		"alamat":  submission.Alamat,
	}

	_, errr := collection.UpdateOne(ctx, bson.M{"nik": key}, bson.M{"$set": edited})

	return errr
}
