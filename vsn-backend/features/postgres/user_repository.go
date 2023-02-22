package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/seg491X-team36/vsn-backend/codegen/db"
	"github.com/seg491X-team36/vsn-backend/domain/model"
)

func convertUser(record db.User) model.User {
	return model.User{
		Id:     record.ID,
		Email:  record.Email,
		State:  model.UserAccountState(record.State),
		Source: record.Source,
	}
}

func convertUsers(records []db.User) []model.User {
	users := make([]model.User, len(records))
	for i, record := range records {
		users[i] = convertUser(record)
	}
	return users
}

type UserRepository struct {
	Query *db.Queries
}

func (repository *UserRepository) GetUser(
	ctx context.Context,
	id uuid.UUID,
) (model.User, error) {
	user, err := repository.Query.GetUser(ctx, id)
	return convertUser(user), err
}

func (repository *UserRepository) GetUsersByState(
	ctx context.Context,
	state model.UserAccountState,
) ([]model.User, error) {
	users, err := repository.Query.GetUsersByState(ctx, db.UserAccountState(state))
	return convertUsers(users), err
}

func (repository UserRepository) CreateUser(
	ctx context.Context,
	email string,
	source string,
) (model.User, error) {
	user, err := repository.Query.CreateUser(ctx, db.CreateUserParams{
		Email:  email,
		State:  db.UserAccountStateREGISTERED,
		Source: source,
	})
	return convertUser(user), err
}

func (repository UserRepository) UpdateUserState(
	ctx context.Context,
	input model.UserSelectInput,
) error {
	return errors.New("not implemented")
}
