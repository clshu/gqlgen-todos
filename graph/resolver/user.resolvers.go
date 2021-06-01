package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strings"

	"github.com/clshu/gqlgen-todos/graph/model"
	"github.com/clshu/gqlgen-todos/utils"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *mutationResolver) CreateUser(ctx context.Context, data model.UserInput) (*model.UserView, error) {
	user := model.User{}
	user = model.User{FirstName: data.FirstName, LastName: data.LastName, Email: data.Email, Password: data.Password}
	err := mgm.Coll(&user).Create(&user)
	if err != nil {
		return nil, err
	}
	viewer := model.UserView{ID: user.ID.Hex(), FirstName: user.FirstName, LastName: user.LastName, Email: user.Email}
	return &viewer, nil
}

func (r *mutationResolver) LogIn(ctx context.Context, data model.LogInInput) (*model.AuthPayload, error) {
	mctx := mgm.Ctx()
	email := strings.ToLower(data.Email)

	result := mgm.Coll(&model.User{}).FindOne(mctx, bson.M{"email": email})
	if result.Err() != nil {
		// most likely
		// "mongo: no documents in result"
		// Do not want to let the user know if it is
		// email wrong or password wrong. Send
		// the same error message Unable to Login
		return nil, fmt.Errorf("Unable to Login")
		// return nil, result.Err()
	}
	ret := &model.User{}

	err := result.Decode(ret)
	if err != nil {
		return nil, err
	}

	if ret.Email != email {
		return nil, fmt.Errorf("Unable to Login")
	}
	perr := utils.ComparePassword(ret.Password, data.Password)

	if perr != nil {
		return nil, fmt.Errorf("Unable to Login")
	}

	token, err := utils.CreateToken(ret.ID.Hex())
	if err != nil {
		return nil, err
	}
	payload := &model.AuthPayload{
		User: &model.UserView{
			ID:        ret.ID.Hex(),
			Email:     ret.Email,
			FirstName: ret.FirstName,
			LastName:  ret.LastName,
		},
		Token: token,
	}
	return payload, nil

}

func (r *queryResolver) Profile(ctx context.Context) (*model.UserView, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.UserView, error) {
	panic(fmt.Errorf("not implemented"))
}
