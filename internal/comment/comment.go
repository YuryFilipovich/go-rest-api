package comment

import (
	"context"
	"errors"
	"fmt"
)

type Store interface {
	GetComment(context.Context, string) (Comment, error)
	PostComment(context.Context, Comment) (Comment, error)
	DeleteComment(context.Context, string) error
	UpdateComment(context.Context, string, Comment) (Comment, error)
}

var (
	ErrFetchingComment = errors.New("failed to fetch comment")
	ErrNotImplemented  = errors.New("not implemented")
)

// Representation of the comment structure for the service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Struct on which all our logic will be built on top of
type Service struct {
	Store Store
}

// NewService returns a pointer to a new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("Retrieving a comment")

	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}
	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, ID string, updatedCmt Comment) (Comment, error) {
	cmt, err := s.Store.UpdateComment(ctx, ID, updatedCmt)

	if err != nil {
		fmt.Println("error updating comment")
		return Comment{}, err
	}

	return cmt, nil
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return s.Store.DeleteComment(ctx, id)
}

func (s *Service) PostComment(ctx context.Context, cmt Comment) (Comment, error) {
	insertedCmt, err := s.Store.PostComment(ctx, cmt)

	if err != nil {
		return Comment{}, err
	}

	return insertedCmt, nil
}
