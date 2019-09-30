package go_graphql

import (
	"context"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

// Resolver struct
type Resolver struct{}

// Mutation resolver
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

// Query resolver
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

// User resolver
func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input NewUser) (*User, error) {
	u := &User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  input.Password,
	}

	err := u.Create()

	if err != nil {
		return nil, err
	}

	return u, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]*User, error) {
	return GetAllUsers()
}

type userResolver struct{ *Resolver }

func (r *userResolver) ID(ctx context.Context, obj *User) (int, error) {
	return int(obj.ID), nil
}
