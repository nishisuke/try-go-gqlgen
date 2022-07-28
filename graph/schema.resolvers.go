package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"example/graph/generated"
	"example/graph/model"
	"example/graph/storage"
	"fmt"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	friend := model.User{
		ID:   "fri",
		Name: "friend",
	}
	user := model.User{
		ID:      "userid",
		Name:    "hey",
		Friends: []*model.User{&friend, &friend},
	}
	return []*model.Todo{
		{ID: "2", Text: "Wash", Done: false, User: &user},
		{ID: "3", Text: "Clean", Done: false, User: &user},
		{ID: "4", Text: "Eat", Done: true, User: &user},
	}, nil
}

// Friends is the resolver for the friends field.
func (r *userResolver) Friends(ctx context.Context, obj *model.User) ([]*model.User, error) {
	friendIDs := []string{"friA", "friB"}
	return storage.GetUser(ctx, friendIDs)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
