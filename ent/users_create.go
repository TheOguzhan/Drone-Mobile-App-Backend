// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/TheOguzhan/Drone-Mobile-App-Backend/ent/users"
	"github.com/google/uuid"
)

// UsersCreate is the builder for creating a Users entity.
type UsersCreate struct {
	config
	mutation *UsersMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (uc *UsersCreate) SetUserID(u uuid.UUID) *UsersCreate {
	uc.mutation.SetUserID(u)
	return uc
}

// SetUsername sets the "username" field.
func (uc *UsersCreate) SetUsername(s string) *UsersCreate {
	uc.mutation.SetUsername(s)
	return uc
}

// SetEmail sets the "email" field.
func (uc *UsersCreate) SetEmail(s string) *UsersCreate {
	uc.mutation.SetEmail(s)
	return uc
}

// SetFirstName sets the "first_name" field.
func (uc *UsersCreate) SetFirstName(s string) *UsersCreate {
	uc.mutation.SetFirstName(s)
	return uc
}

// SetLastName sets the "last_name" field.
func (uc *UsersCreate) SetLastName(s string) *UsersCreate {
	uc.mutation.SetLastName(s)
	return uc
}

// SetIsUserConfirmed sets the "is_user_confirmed" field.
func (uc *UsersCreate) SetIsUserConfirmed(b bool) *UsersCreate {
	uc.mutation.SetIsUserConfirmed(b)
	return uc
}

// SetNillableIsUserConfirmed sets the "is_user_confirmed" field if the given value is not nil.
func (uc *UsersCreate) SetNillableIsUserConfirmed(b *bool) *UsersCreate {
	if b != nil {
		uc.SetIsUserConfirmed(*b)
	}
	return uc
}

// SetPassword sets the "password" field.
func (uc *UsersCreate) SetPassword(s string) *UsersCreate {
	uc.mutation.SetPassword(s)
	return uc
}

// Mutation returns the UsersMutation object of the builder.
func (uc *UsersCreate) Mutation() *UsersMutation {
	return uc.mutation
}

// Save creates the Users in the database.
func (uc *UsersCreate) Save(ctx context.Context) (*Users, error) {
	uc.defaults()
	return withHooks[*Users, UsersMutation](ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UsersCreate) SaveX(ctx context.Context) *Users {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UsersCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UsersCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UsersCreate) defaults() {
	if _, ok := uc.mutation.IsUserConfirmed(); !ok {
		v := users.DefaultIsUserConfirmed
		uc.mutation.SetIsUserConfirmed(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UsersCreate) check() error {
	if _, ok := uc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Users.user_id"`)}
	}
	if _, ok := uc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "Users.username"`)}
	}
	if _, ok := uc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "Users.email"`)}
	}
	if v, ok := uc.mutation.Email(); ok {
		if err := users.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Users.email": %w`, err)}
		}
	}
	if _, ok := uc.mutation.FirstName(); !ok {
		return &ValidationError{Name: "first_name", err: errors.New(`ent: missing required field "Users.first_name"`)}
	}
	if _, ok := uc.mutation.LastName(); !ok {
		return &ValidationError{Name: "last_name", err: errors.New(`ent: missing required field "Users.last_name"`)}
	}
	if _, ok := uc.mutation.IsUserConfirmed(); !ok {
		return &ValidationError{Name: "is_user_confirmed", err: errors.New(`ent: missing required field "Users.is_user_confirmed"`)}
	}
	if _, ok := uc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "Users.password"`)}
	}
	return nil
}

func (uc *UsersCreate) sqlSave(ctx context.Context) (*Users, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UsersCreate) createSpec() (*Users, *sqlgraph.CreateSpec) {
	var (
		_node = &Users{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(users.Table, sqlgraph.NewFieldSpec(users.FieldID, field.TypeInt))
	)
	if value, ok := uc.mutation.UserID(); ok {
		_spec.SetField(users.FieldUserID, field.TypeUUID, value)
		_node.UserID = value
	}
	if value, ok := uc.mutation.Username(); ok {
		_spec.SetField(users.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := uc.mutation.Email(); ok {
		_spec.SetField(users.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := uc.mutation.FirstName(); ok {
		_spec.SetField(users.FieldFirstName, field.TypeString, value)
		_node.FirstName = value
	}
	if value, ok := uc.mutation.LastName(); ok {
		_spec.SetField(users.FieldLastName, field.TypeString, value)
		_node.LastName = value
	}
	if value, ok := uc.mutation.IsUserConfirmed(); ok {
		_spec.SetField(users.FieldIsUserConfirmed, field.TypeBool, value)
		_node.IsUserConfirmed = value
	}
	if value, ok := uc.mutation.Password(); ok {
		_spec.SetField(users.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	return _node, _spec
}

// UsersCreateBulk is the builder for creating many Users entities in bulk.
type UsersCreateBulk struct {
	config
	builders []*UsersCreate
}

// Save creates the Users entities in the database.
func (ucb *UsersCreateBulk) Save(ctx context.Context) ([]*Users, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*Users, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UsersMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UsersCreateBulk) SaveX(ctx context.Context) []*Users {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UsersCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UsersCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
