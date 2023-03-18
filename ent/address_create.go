// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/TheOguzhan/Drone-Mobile-App-Backend/ent/address"
	"github.com/TheOguzhan/Drone-Mobile-App-Backend/ent/user"
	"github.com/google/uuid"
)

// AddressCreate is the builder for creating a Address entity.
type AddressCreate struct {
	config
	mutation *AddressMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ac *AddressCreate) SetName(s string) *AddressCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetAddressLine sets the "address_line" field.
func (ac *AddressCreate) SetAddressLine(s string) *AddressCreate {
	ac.mutation.SetAddressLine(s)
	return ac
}

// SetLatitude sets the "latitude" field.
func (ac *AddressCreate) SetLatitude(f float64) *AddressCreate {
	ac.mutation.SetLatitude(f)
	return ac
}

// SetLongtitude sets the "longtitude" field.
func (ac *AddressCreate) SetLongtitude(f float64) *AddressCreate {
	ac.mutation.SetLongtitude(f)
	return ac
}

// SetID sets the "id" field.
func (ac *AddressCreate) SetID(u uuid.UUID) *AddressCreate {
	ac.mutation.SetID(u)
	return ac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ac *AddressCreate) SetNillableID(u *uuid.UUID) *AddressCreate {
	if u != nil {
		ac.SetID(*u)
	}
	return ac
}

// SetAddressMasterID sets the "address_master" edge to the User entity by ID.
func (ac *AddressCreate) SetAddressMasterID(id uuid.UUID) *AddressCreate {
	ac.mutation.SetAddressMasterID(id)
	return ac
}

// SetAddressMaster sets the "address_master" edge to the User entity.
func (ac *AddressCreate) SetAddressMaster(u *User) *AddressCreate {
	return ac.SetAddressMasterID(u.ID)
}

// Mutation returns the AddressMutation object of the builder.
func (ac *AddressCreate) Mutation() *AddressMutation {
	return ac.mutation
}

// Save creates the Address in the database.
func (ac *AddressCreate) Save(ctx context.Context) (*Address, error) {
	ac.defaults()
	return withHooks[*Address, AddressMutation](ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AddressCreate) SaveX(ctx context.Context) *Address {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AddressCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AddressCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AddressCreate) defaults() {
	if _, ok := ac.mutation.ID(); !ok {
		v := address.DefaultID()
		ac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AddressCreate) check() error {
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Address.name"`)}
	}
	if _, ok := ac.mutation.AddressLine(); !ok {
		return &ValidationError{Name: "address_line", err: errors.New(`ent: missing required field "Address.address_line"`)}
	}
	if _, ok := ac.mutation.Latitude(); !ok {
		return &ValidationError{Name: "latitude", err: errors.New(`ent: missing required field "Address.latitude"`)}
	}
	if _, ok := ac.mutation.Longtitude(); !ok {
		return &ValidationError{Name: "longtitude", err: errors.New(`ent: missing required field "Address.longtitude"`)}
	}
	if _, ok := ac.mutation.AddressMasterID(); !ok {
		return &ValidationError{Name: "address_master", err: errors.New(`ent: missing required edge "Address.address_master"`)}
	}
	return nil
}

func (ac *AddressCreate) sqlSave(ctx context.Context) (*Address, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AddressCreate) createSpec() (*Address, *sqlgraph.CreateSpec) {
	var (
		_node = &Address{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(address.Table, sqlgraph.NewFieldSpec(address.FieldID, field.TypeUUID))
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ac.mutation.Name(); ok {
		_spec.SetField(address.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ac.mutation.AddressLine(); ok {
		_spec.SetField(address.FieldAddressLine, field.TypeString, value)
		_node.AddressLine = value
	}
	if value, ok := ac.mutation.Latitude(); ok {
		_spec.SetField(address.FieldLatitude, field.TypeFloat64, value)
		_node.Latitude = value
	}
	if value, ok := ac.mutation.Longtitude(); ok {
		_spec.SetField(address.FieldLongtitude, field.TypeFloat64, value)
		_node.Longtitude = value
	}
	if nodes := ac.mutation.AddressMasterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   address.AddressMasterTable,
			Columns: []string{address.AddressMasterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_address_slaves = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AddressCreateBulk is the builder for creating many Address entities in bulk.
type AddressCreateBulk struct {
	config
	builders []*AddressCreate
}

// Save creates the Address entities in the database.
func (acb *AddressCreateBulk) Save(ctx context.Context) ([]*Address, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Address, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AddressMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AddressCreateBulk) SaveX(ctx context.Context) []*Address {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AddressCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AddressCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}