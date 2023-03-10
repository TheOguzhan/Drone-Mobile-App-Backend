// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/TheOguzhan/Drone-Mobile-App-Backend/ent/predicate"
	"github.com/TheOguzhan/Drone-Mobile-App-Backend/ent/users"
	"github.com/google/uuid"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeUsers = "Users"
)

// UsersMutation represents an operation that mutates the Users nodes in the graph.
type UsersMutation struct {
	config
	op                Op
	typ               string
	id                *int
	user_id           *uuid.UUID
	username          *string
	email             *string
	first_name        *string
	last_name         *string
	is_user_confirmed *bool
	password          *string
	clearedFields     map[string]struct{}
	done              bool
	oldValue          func(context.Context) (*Users, error)
	predicates        []predicate.Users
}

var _ ent.Mutation = (*UsersMutation)(nil)

// usersOption allows management of the mutation configuration using functional options.
type usersOption func(*UsersMutation)

// newUsersMutation creates new mutation for the Users entity.
func newUsersMutation(c config, op Op, opts ...usersOption) *UsersMutation {
	m := &UsersMutation{
		config:        c,
		op:            op,
		typ:           TypeUsers,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUsersID sets the ID field of the mutation.
func withUsersID(id int) usersOption {
	return func(m *UsersMutation) {
		var (
			err   error
			once  sync.Once
			value *Users
		)
		m.oldValue = func(ctx context.Context) (*Users, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Users.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUsers sets the old Users of the mutation.
func withUsers(node *Users) usersOption {
	return func(m *UsersMutation) {
		m.oldValue = func(context.Context) (*Users, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UsersMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UsersMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UsersMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UsersMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Users.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetUserID sets the "user_id" field.
func (m *UsersMutation) SetUserID(u uuid.UUID) {
	m.user_id = &u
}

// UserID returns the value of the "user_id" field in the mutation.
func (m *UsersMutation) UserID() (r uuid.UUID, exists bool) {
	v := m.user_id
	if v == nil {
		return
	}
	return *v, true
}

// OldUserID returns the old "user_id" field's value of the Users entity.
// If the Users object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UsersMutation) OldUserID(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUserID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUserID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUserID: %w", err)
	}
	return oldValue.UserID, nil
}

// ResetUserID resets all changes to the "user_id" field.
func (m *UsersMutation) ResetUserID() {
	m.user_id = nil
}

// SetUsername sets the "username" field.
func (m *UsersMutation) SetUsername(s string) {
	m.username = &s
}

// Username returns the value of the "username" field in the mutation.
func (m *UsersMutation) Username() (r string, exists bool) {
	v := m.username
	if v == nil {
		return
	}
	return *v, true
}

// OldUsername returns the old "username" field's value of the Users entity.
// If the Users object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UsersMutation) OldUsername(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUsername is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUsername requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUsername: %w", err)
	}
	return oldValue.Username, nil
}

// ResetUsername resets all changes to the "username" field.
func (m *UsersMutation) ResetUsername() {
	m.username = nil
}

// SetEmail sets the "email" field.
func (m *UsersMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the value of the "email" field in the mutation.
func (m *UsersMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old "email" field's value of the Users entity.
// If the Users object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UsersMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ResetEmail resets all changes to the "email" field.
func (m *UsersMutation) ResetEmail() {
	m.email = nil
}

// SetFirstName sets the "first_name" field.
func (m *UsersMutation) SetFirstName(s string) {
	m.first_name = &s
}

// FirstName returns the value of the "first_name" field in the mutation.
func (m *UsersMutation) FirstName() (r string, exists bool) {
	v := m.first_name
	if v == nil {
		return
	}
	return *v, true
}

// OldFirstName returns the old "first_name" field's value of the Users entity.
// If the Users object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UsersMutation) OldFirstName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldFirstName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldFirstName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFirstName: %w", err)
	}
	return oldValue.FirstName, nil
}

// ResetFirstName resets all changes to the "first_name" field.
func (m *UsersMutation) ResetFirstName() {
	m.first_name = nil
}

// SetLastName sets the "last_name" field.
func (m *UsersMutation) SetLastName(s string) {
	m.last_name = &s
}

// LastName returns the value of the "last_name" field in the mutation.
func (m *UsersMutation) LastName() (r string, exists bool) {
	v := m.last_name
	if v == nil {
		return
	}
	return *v, true
}

// OldLastName returns the old "last_name" field's value of the Users entity.
// If the Users object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UsersMutation) OldLastName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLastName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLastName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLastName: %w", err)
	}
	return oldValue.LastName, nil
}

// ResetLastName resets all changes to the "last_name" field.
func (m *UsersMutation) ResetLastName() {
	m.last_name = nil
}

// SetIsUserConfirmed sets the "is_user_confirmed" field.
func (m *UsersMutation) SetIsUserConfirmed(b bool) {
	m.is_user_confirmed = &b
}

// IsUserConfirmed returns the value of the "is_user_confirmed" field in the mutation.
func (m *UsersMutation) IsUserConfirmed() (r bool, exists bool) {
	v := m.is_user_confirmed
	if v == nil {
		return
	}
	return *v, true
}

// OldIsUserConfirmed returns the old "is_user_confirmed" field's value of the Users entity.
// If the Users object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UsersMutation) OldIsUserConfirmed(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldIsUserConfirmed is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldIsUserConfirmed requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldIsUserConfirmed: %w", err)
	}
	return oldValue.IsUserConfirmed, nil
}

// ResetIsUserConfirmed resets all changes to the "is_user_confirmed" field.
func (m *UsersMutation) ResetIsUserConfirmed() {
	m.is_user_confirmed = nil
}

// SetPassword sets the "password" field.
func (m *UsersMutation) SetPassword(s string) {
	m.password = &s
}

// Password returns the value of the "password" field in the mutation.
func (m *UsersMutation) Password() (r string, exists bool) {
	v := m.password
	if v == nil {
		return
	}
	return *v, true
}

// OldPassword returns the old "password" field's value of the Users entity.
// If the Users object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UsersMutation) OldPassword(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPassword is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPassword requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPassword: %w", err)
	}
	return oldValue.Password, nil
}

// ResetPassword resets all changes to the "password" field.
func (m *UsersMutation) ResetPassword() {
	m.password = nil
}

// Where appends a list predicates to the UsersMutation builder.
func (m *UsersMutation) Where(ps ...predicate.Users) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the UsersMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *UsersMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Users, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *UsersMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *UsersMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Users).
func (m *UsersMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UsersMutation) Fields() []string {
	fields := make([]string, 0, 7)
	if m.user_id != nil {
		fields = append(fields, users.FieldUserID)
	}
	if m.username != nil {
		fields = append(fields, users.FieldUsername)
	}
	if m.email != nil {
		fields = append(fields, users.FieldEmail)
	}
	if m.first_name != nil {
		fields = append(fields, users.FieldFirstName)
	}
	if m.last_name != nil {
		fields = append(fields, users.FieldLastName)
	}
	if m.is_user_confirmed != nil {
		fields = append(fields, users.FieldIsUserConfirmed)
	}
	if m.password != nil {
		fields = append(fields, users.FieldPassword)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UsersMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case users.FieldUserID:
		return m.UserID()
	case users.FieldUsername:
		return m.Username()
	case users.FieldEmail:
		return m.Email()
	case users.FieldFirstName:
		return m.FirstName()
	case users.FieldLastName:
		return m.LastName()
	case users.FieldIsUserConfirmed:
		return m.IsUserConfirmed()
	case users.FieldPassword:
		return m.Password()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UsersMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case users.FieldUserID:
		return m.OldUserID(ctx)
	case users.FieldUsername:
		return m.OldUsername(ctx)
	case users.FieldEmail:
		return m.OldEmail(ctx)
	case users.FieldFirstName:
		return m.OldFirstName(ctx)
	case users.FieldLastName:
		return m.OldLastName(ctx)
	case users.FieldIsUserConfirmed:
		return m.OldIsUserConfirmed(ctx)
	case users.FieldPassword:
		return m.OldPassword(ctx)
	}
	return nil, fmt.Errorf("unknown Users field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UsersMutation) SetField(name string, value ent.Value) error {
	switch name {
	case users.FieldUserID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUserID(v)
		return nil
	case users.FieldUsername:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUsername(v)
		return nil
	case users.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case users.FieldFirstName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFirstName(v)
		return nil
	case users.FieldLastName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLastName(v)
		return nil
	case users.FieldIsUserConfirmed:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIsUserConfirmed(v)
		return nil
	case users.FieldPassword:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPassword(v)
		return nil
	}
	return fmt.Errorf("unknown Users field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UsersMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UsersMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UsersMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Users numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UsersMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UsersMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UsersMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Users nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UsersMutation) ResetField(name string) error {
	switch name {
	case users.FieldUserID:
		m.ResetUserID()
		return nil
	case users.FieldUsername:
		m.ResetUsername()
		return nil
	case users.FieldEmail:
		m.ResetEmail()
		return nil
	case users.FieldFirstName:
		m.ResetFirstName()
		return nil
	case users.FieldLastName:
		m.ResetLastName()
		return nil
	case users.FieldIsUserConfirmed:
		m.ResetIsUserConfirmed()
		return nil
	case users.FieldPassword:
		m.ResetPassword()
		return nil
	}
	return fmt.Errorf("unknown Users field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UsersMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UsersMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UsersMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UsersMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UsersMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UsersMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UsersMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Users unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UsersMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Users edge %s", name)
}
