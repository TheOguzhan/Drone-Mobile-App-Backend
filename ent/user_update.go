// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/TheOguzhan/Drone-Mobile-App-Backend/ent/address"
	"github.com/TheOguzhan/Drone-Mobile-App-Backend/ent/order"
	"github.com/TheOguzhan/Drone-Mobile-App-Backend/ent/predicate"
	"github.com/TheOguzhan/Drone-Mobile-App-Backend/ent/user"
	"github.com/google/uuid"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUsername sets the "username" field.
func (uu *UserUpdate) SetUsername(s string) *UserUpdate {
	uu.mutation.SetUsername(s)
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetFirstName sets the "first_name" field.
func (uu *UserUpdate) SetFirstName(s string) *UserUpdate {
	uu.mutation.SetFirstName(s)
	return uu
}

// SetLastName sets the "last_name" field.
func (uu *UserUpdate) SetLastName(s string) *UserUpdate {
	uu.mutation.SetLastName(s)
	return uu
}

// SetIsUserConfirmed sets the "is_user_confirmed" field.
func (uu *UserUpdate) SetIsUserConfirmed(b bool) *UserUpdate {
	uu.mutation.SetIsUserConfirmed(b)
	return uu
}

// SetNillableIsUserConfirmed sets the "is_user_confirmed" field if the given value is not nil.
func (uu *UserUpdate) SetNillableIsUserConfirmed(b *bool) *UserUpdate {
	if b != nil {
		uu.SetIsUserConfirmed(*b)
	}
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetAuthTokens sets the "auth_tokens" field.
func (uu *UserUpdate) SetAuthTokens(s []string) *UserUpdate {
	uu.mutation.SetAuthTokens(s)
	return uu
}

// AppendAuthTokens appends s to the "auth_tokens" field.
func (uu *UserUpdate) AppendAuthTokens(s []string) *UserUpdate {
	uu.mutation.AppendAuthTokens(s)
	return uu
}

// ClearAuthTokens clears the value of the "auth_tokens" field.
func (uu *UserUpdate) ClearAuthTokens() *UserUpdate {
	uu.mutation.ClearAuthTokens()
	return uu
}

// AddUserAddressIDs adds the "user_addresses" edge to the Address entity by IDs.
func (uu *UserUpdate) AddUserAddressIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddUserAddressIDs(ids...)
	return uu
}

// AddUserAddresses adds the "user_addresses" edges to the Address entity.
func (uu *UserUpdate) AddUserAddresses(a ...*Address) *UserUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uu.AddUserAddressIDs(ids...)
}

// AddUserOrderIDs adds the "user_orders" edge to the Order entity by IDs.
func (uu *UserUpdate) AddUserOrderIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddUserOrderIDs(ids...)
	return uu
}

// AddUserOrders adds the "user_orders" edges to the Order entity.
func (uu *UserUpdate) AddUserOrders(o ...*Order) *UserUpdate {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return uu.AddUserOrderIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearUserAddresses clears all "user_addresses" edges to the Address entity.
func (uu *UserUpdate) ClearUserAddresses() *UserUpdate {
	uu.mutation.ClearUserAddresses()
	return uu
}

// RemoveUserAddressIDs removes the "user_addresses" edge to Address entities by IDs.
func (uu *UserUpdate) RemoveUserAddressIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveUserAddressIDs(ids...)
	return uu
}

// RemoveUserAddresses removes "user_addresses" edges to Address entities.
func (uu *UserUpdate) RemoveUserAddresses(a ...*Address) *UserUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uu.RemoveUserAddressIDs(ids...)
}

// ClearUserOrders clears all "user_orders" edges to the Order entity.
func (uu *UserUpdate) ClearUserOrders() *UserUpdate {
	uu.mutation.ClearUserOrders()
	return uu
}

// RemoveUserOrderIDs removes the "user_orders" edge to Order entities by IDs.
func (uu *UserUpdate) RemoveUserOrderIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveUserOrderIDs(ids...)
	return uu
}

// RemoveUserOrders removes "user_orders" edges to Order entities.
func (uu *UserUpdate) RemoveUserOrders(o ...*Order) *UserUpdate {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return uu.RemoveUserOrderIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, UserMutation](ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.FirstName(); ok {
		_spec.SetField(user.FieldFirstName, field.TypeString, value)
	}
	if value, ok := uu.mutation.LastName(); ok {
		_spec.SetField(user.FieldLastName, field.TypeString, value)
	}
	if value, ok := uu.mutation.IsUserConfirmed(); ok {
		_spec.SetField(user.FieldIsUserConfirmed, field.TypeBool, value)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uu.mutation.AuthTokens(); ok {
		_spec.SetField(user.FieldAuthTokens, field.TypeJSON, value)
	}
	if value, ok := uu.mutation.AppendedAuthTokens(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldAuthTokens, value)
		})
	}
	if uu.mutation.AuthTokensCleared() {
		_spec.ClearField(user.FieldAuthTokens, field.TypeJSON)
	}
	if uu.mutation.UserAddressesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserAddressesTable,
			Columns: []string{user.UserAddressesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: address.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedUserAddressesIDs(); len(nodes) > 0 && !uu.mutation.UserAddressesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserAddressesTable,
			Columns: []string{user.UserAddressesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: address.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.UserAddressesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserAddressesTable,
			Columns: []string{user.UserAddressesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: address.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.UserOrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserOrdersTable,
			Columns: []string{user.UserOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: order.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedUserOrdersIDs(); len(nodes) > 0 && !uu.mutation.UserOrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserOrdersTable,
			Columns: []string{user.UserOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: order.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.UserOrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserOrdersTable,
			Columns: []string{user.UserOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: order.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUsername sets the "username" field.
func (uuo *UserUpdateOne) SetUsername(s string) *UserUpdateOne {
	uuo.mutation.SetUsername(s)
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetFirstName sets the "first_name" field.
func (uuo *UserUpdateOne) SetFirstName(s string) *UserUpdateOne {
	uuo.mutation.SetFirstName(s)
	return uuo
}

// SetLastName sets the "last_name" field.
func (uuo *UserUpdateOne) SetLastName(s string) *UserUpdateOne {
	uuo.mutation.SetLastName(s)
	return uuo
}

// SetIsUserConfirmed sets the "is_user_confirmed" field.
func (uuo *UserUpdateOne) SetIsUserConfirmed(b bool) *UserUpdateOne {
	uuo.mutation.SetIsUserConfirmed(b)
	return uuo
}

// SetNillableIsUserConfirmed sets the "is_user_confirmed" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableIsUserConfirmed(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetIsUserConfirmed(*b)
	}
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetAuthTokens sets the "auth_tokens" field.
func (uuo *UserUpdateOne) SetAuthTokens(s []string) *UserUpdateOne {
	uuo.mutation.SetAuthTokens(s)
	return uuo
}

// AppendAuthTokens appends s to the "auth_tokens" field.
func (uuo *UserUpdateOne) AppendAuthTokens(s []string) *UserUpdateOne {
	uuo.mutation.AppendAuthTokens(s)
	return uuo
}

// ClearAuthTokens clears the value of the "auth_tokens" field.
func (uuo *UserUpdateOne) ClearAuthTokens() *UserUpdateOne {
	uuo.mutation.ClearAuthTokens()
	return uuo
}

// AddUserAddressIDs adds the "user_addresses" edge to the Address entity by IDs.
func (uuo *UserUpdateOne) AddUserAddressIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddUserAddressIDs(ids...)
	return uuo
}

// AddUserAddresses adds the "user_addresses" edges to the Address entity.
func (uuo *UserUpdateOne) AddUserAddresses(a ...*Address) *UserUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uuo.AddUserAddressIDs(ids...)
}

// AddUserOrderIDs adds the "user_orders" edge to the Order entity by IDs.
func (uuo *UserUpdateOne) AddUserOrderIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddUserOrderIDs(ids...)
	return uuo
}

// AddUserOrders adds the "user_orders" edges to the Order entity.
func (uuo *UserUpdateOne) AddUserOrders(o ...*Order) *UserUpdateOne {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return uuo.AddUserOrderIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearUserAddresses clears all "user_addresses" edges to the Address entity.
func (uuo *UserUpdateOne) ClearUserAddresses() *UserUpdateOne {
	uuo.mutation.ClearUserAddresses()
	return uuo
}

// RemoveUserAddressIDs removes the "user_addresses" edge to Address entities by IDs.
func (uuo *UserUpdateOne) RemoveUserAddressIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveUserAddressIDs(ids...)
	return uuo
}

// RemoveUserAddresses removes "user_addresses" edges to Address entities.
func (uuo *UserUpdateOne) RemoveUserAddresses(a ...*Address) *UserUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uuo.RemoveUserAddressIDs(ids...)
}

// ClearUserOrders clears all "user_orders" edges to the Order entity.
func (uuo *UserUpdateOne) ClearUserOrders() *UserUpdateOne {
	uuo.mutation.ClearUserOrders()
	return uuo
}

// RemoveUserOrderIDs removes the "user_orders" edge to Order entities by IDs.
func (uuo *UserUpdateOne) RemoveUserOrderIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveUserOrderIDs(ids...)
	return uuo
}

// RemoveUserOrders removes "user_orders" edges to Order entities.
func (uuo *UserUpdateOne) RemoveUserOrders(o ...*Order) *UserUpdateOne {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return uuo.RemoveUserOrderIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks[*User, UserMutation](ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.FirstName(); ok {
		_spec.SetField(user.FieldFirstName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.LastName(); ok {
		_spec.SetField(user.FieldLastName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.IsUserConfirmed(); ok {
		_spec.SetField(user.FieldIsUserConfirmed, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uuo.mutation.AuthTokens(); ok {
		_spec.SetField(user.FieldAuthTokens, field.TypeJSON, value)
	}
	if value, ok := uuo.mutation.AppendedAuthTokens(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldAuthTokens, value)
		})
	}
	if uuo.mutation.AuthTokensCleared() {
		_spec.ClearField(user.FieldAuthTokens, field.TypeJSON)
	}
	if uuo.mutation.UserAddressesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserAddressesTable,
			Columns: []string{user.UserAddressesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: address.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedUserAddressesIDs(); len(nodes) > 0 && !uuo.mutation.UserAddressesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserAddressesTable,
			Columns: []string{user.UserAddressesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: address.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.UserAddressesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserAddressesTable,
			Columns: []string{user.UserAddressesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: address.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.UserOrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserOrdersTable,
			Columns: []string{user.UserOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: order.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedUserOrdersIDs(); len(nodes) > 0 && !uuo.mutation.UserOrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserOrdersTable,
			Columns: []string{user.UserOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: order.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.UserOrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserOrdersTable,
			Columns: []string{user.UserOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: order.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
