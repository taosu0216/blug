// Code generated by ent, DO NOT EDIT.

package ent

import (
	"blug/internal/data/ent/friend"
	"blug/internal/data/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FriendUpdate is the builder for updating Friend entities.
type FriendUpdate struct {
	config
	hooks    []Hook
	mutation *FriendMutation
}

// Where appends a list predicates to the FriendUpdate builder.
func (fu *FriendUpdate) Where(ps ...predicate.Friend) *FriendUpdate {
	fu.mutation.Where(ps...)
	return fu
}

// SetTitle sets the "Title" field.
func (fu *FriendUpdate) SetTitle(s string) *FriendUpdate {
	fu.mutation.SetTitle(s)
	return fu
}

// SetNillableTitle sets the "Title" field if the given value is not nil.
func (fu *FriendUpdate) SetNillableTitle(s *string) *FriendUpdate {
	if s != nil {
		fu.SetTitle(*s)
	}
	return fu
}

// SetDesc sets the "Desc" field.
func (fu *FriendUpdate) SetDesc(s string) *FriendUpdate {
	fu.mutation.SetDesc(s)
	return fu
}

// SetNillableDesc sets the "Desc" field if the given value is not nil.
func (fu *FriendUpdate) SetNillableDesc(s *string) *FriendUpdate {
	if s != nil {
		fu.SetDesc(*s)
	}
	return fu
}

// SetLink sets the "Link" field.
func (fu *FriendUpdate) SetLink(s string) *FriendUpdate {
	fu.mutation.SetLink(s)
	return fu
}

// SetNillableLink sets the "Link" field if the given value is not nil.
func (fu *FriendUpdate) SetNillableLink(s *string) *FriendUpdate {
	if s != nil {
		fu.SetLink(*s)
	}
	return fu
}

// SetAvatar sets the "Avatar" field.
func (fu *FriendUpdate) SetAvatar(s string) *FriendUpdate {
	fu.mutation.SetAvatar(s)
	return fu
}

// SetNillableAvatar sets the "Avatar" field if the given value is not nil.
func (fu *FriendUpdate) SetNillableAvatar(s *string) *FriendUpdate {
	if s != nil {
		fu.SetAvatar(*s)
	}
	return fu
}

// Mutation returns the FriendMutation object of the builder.
func (fu *FriendUpdate) Mutation() *FriendMutation {
	return fu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FriendUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, fu.sqlSave, fu.mutation, fu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FriendUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FriendUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FriendUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fu *FriendUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(friend.Table, friend.Columns, sqlgraph.NewFieldSpec(friend.FieldID, field.TypeUUID))
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.Title(); ok {
		_spec.SetField(friend.FieldTitle, field.TypeString, value)
	}
	if value, ok := fu.mutation.Desc(); ok {
		_spec.SetField(friend.FieldDesc, field.TypeString, value)
	}
	if value, ok := fu.mutation.Link(); ok {
		_spec.SetField(friend.FieldLink, field.TypeString, value)
	}
	if value, ok := fu.mutation.Avatar(); ok {
		_spec.SetField(friend.FieldAvatar, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{friend.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fu.mutation.done = true
	return n, nil
}

// FriendUpdateOne is the builder for updating a single Friend entity.
type FriendUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FriendMutation
}

// SetTitle sets the "Title" field.
func (fuo *FriendUpdateOne) SetTitle(s string) *FriendUpdateOne {
	fuo.mutation.SetTitle(s)
	return fuo
}

// SetNillableTitle sets the "Title" field if the given value is not nil.
func (fuo *FriendUpdateOne) SetNillableTitle(s *string) *FriendUpdateOne {
	if s != nil {
		fuo.SetTitle(*s)
	}
	return fuo
}

// SetDesc sets the "Desc" field.
func (fuo *FriendUpdateOne) SetDesc(s string) *FriendUpdateOne {
	fuo.mutation.SetDesc(s)
	return fuo
}

// SetNillableDesc sets the "Desc" field if the given value is not nil.
func (fuo *FriendUpdateOne) SetNillableDesc(s *string) *FriendUpdateOne {
	if s != nil {
		fuo.SetDesc(*s)
	}
	return fuo
}

// SetLink sets the "Link" field.
func (fuo *FriendUpdateOne) SetLink(s string) *FriendUpdateOne {
	fuo.mutation.SetLink(s)
	return fuo
}

// SetNillableLink sets the "Link" field if the given value is not nil.
func (fuo *FriendUpdateOne) SetNillableLink(s *string) *FriendUpdateOne {
	if s != nil {
		fuo.SetLink(*s)
	}
	return fuo
}

// SetAvatar sets the "Avatar" field.
func (fuo *FriendUpdateOne) SetAvatar(s string) *FriendUpdateOne {
	fuo.mutation.SetAvatar(s)
	return fuo
}

// SetNillableAvatar sets the "Avatar" field if the given value is not nil.
func (fuo *FriendUpdateOne) SetNillableAvatar(s *string) *FriendUpdateOne {
	if s != nil {
		fuo.SetAvatar(*s)
	}
	return fuo
}

// Mutation returns the FriendMutation object of the builder.
func (fuo *FriendUpdateOne) Mutation() *FriendMutation {
	return fuo.mutation
}

// Where appends a list predicates to the FriendUpdate builder.
func (fuo *FriendUpdateOne) Where(ps ...predicate.Friend) *FriendUpdateOne {
	fuo.mutation.Where(ps...)
	return fuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *FriendUpdateOne) Select(field string, fields ...string) *FriendUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated Friend entity.
func (fuo *FriendUpdateOne) Save(ctx context.Context) (*Friend, error) {
	return withHooks(ctx, fuo.sqlSave, fuo.mutation, fuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FriendUpdateOne) SaveX(ctx context.Context) *Friend {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FriendUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FriendUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fuo *FriendUpdateOne) sqlSave(ctx context.Context) (_node *Friend, err error) {
	_spec := sqlgraph.NewUpdateSpec(friend.Table, friend.Columns, sqlgraph.NewFieldSpec(friend.FieldID, field.TypeUUID))
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Friend.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, friend.FieldID)
		for _, f := range fields {
			if !friend.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != friend.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fuo.mutation.Title(); ok {
		_spec.SetField(friend.FieldTitle, field.TypeString, value)
	}
	if value, ok := fuo.mutation.Desc(); ok {
		_spec.SetField(friend.FieldDesc, field.TypeString, value)
	}
	if value, ok := fuo.mutation.Link(); ok {
		_spec.SetField(friend.FieldLink, field.TypeString, value)
	}
	if value, ok := fuo.mutation.Avatar(); ok {
		_spec.SetField(friend.FieldAvatar, field.TypeString, value)
	}
	_node = &Friend{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{friend.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fuo.mutation.done = true
	return _node, nil
}