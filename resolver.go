package graphql_gopher

import (
	"context"
	"errors"
	"github.com/graphql-gopher/database"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (*Todo, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateUser(ctx context.Context,input NewUser)(*User,error){
	collection := database.DB.Collection("user")
	if collection == nil {
		return nil,errors.New("asd")
	}
	_,err := collection.InsertOne(ctx,input)
	if err != nil {
		return nil,err
	}
	return &User{ID:input.ID,Name:input.Name},nil
}
type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]*Todo, error) {
	panic("not implemented")
}
