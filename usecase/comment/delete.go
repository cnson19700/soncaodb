package comment

import (
	"context"

	"github.com/soncaodb/util/myerror"
)

type DeleteCommentRequest struct {
	ID int64 `json:"id"`
}

func (u *Usecase) Delete(ctx context.Context, req *DeleteCommentRequest) error {
	comment, err := u.commentRepo.GetById(ctx, req.ID)
	if err != nil {
		return myerror.ErrGetBook(err)
	}

	err = u.commentRepo.Delete(ctx, comment.ID)

	return err
}
