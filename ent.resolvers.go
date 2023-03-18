package main

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"

	"github.com/TheOguzhan/Drone-Mobile-App-Backend/ent"
	"github.com/TheOguzhan/Drone-Mobile-App-Backend/ent/user"
	"github.com/TheOguzhan/Drone-Mobile-App-Backend/graph"
	"github.com/google/uuid"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id uuid.UUID) (ent.Noder, error) {
	return r.client.Noder(ctx, id, ent.WithFixedNodeType(user.Table))
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []uuid.UUID) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids, ent.WithFixedNodeType(user.Table))
}

// Addresses is the resolver for the addresses field.
func (r *queryResolver) Addresses(ctx context.Context) ([]*ent.Address, error) {
	return r.client.Address.Query().All(ctx)
}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context) ([]*ent.Product, error) {
	return r.client.Product.Query().All(ctx)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*ent.User, error) {
	return r.client.User.Query().All(ctx)
}

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
