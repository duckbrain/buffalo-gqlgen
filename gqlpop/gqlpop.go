package gqlpop

import (
	"context"

	"github.com/gobuffalo/pop/v5"
	"github.com/pkg/errors"
)

func conn(ctx context.Context, db *pop.Connection) *pop.Connection {
	// TODO check ctx for tx and fall back to db; add ctx to conn
	return db
}
func query(ctx context.Context, db *pop.Connection) *pop.Query {
	// TODO check ctx for q or tx and fall back to db; add ctx to conn
	return db.Q()
}

func parentScope(parent, model interface{}) pop.ScopeFunc {
	// TODO filter the query to be .Where("parent_id = ?", parent.ID)
	return func(q *pop.Query) *pop.Query {
		return q
	}
}
func childScope(parent, model interface{}) pop.ScopeFunc {
	// TODO filter the query to be .Find(model, parent.ModelID)
	return func(q *pop.Query) *pop.Query {
		return q
	}
}

func filterScope(filters interface{}) pop.ScopeFunc {
	// TODO filter the query based on filters using reflection
	return func(q *pop.Query) *pop.Query {
		return q
	}
}

func assign(input, model interface{}) ([]string, error) {
	// TODO, assign fields from the input to the model using reflection
	// Allow input to satisfy an interface to update the model
	return nil, nil
}

func Find(ctx context.Context, db *pop.Connection, id, model interface{}) error {
	q := query(ctx, db)
	return q.Find(model, id)
}

func FindChild(ctx context.Context, db *pop.Connection, parent, model interface{}) error {
	q := query(ctx, db)
	return q.Scope(childScope(parent, model)).First(model)
}

func All(ctx context.Context, db *pop.Connection, parent, filter, model interface{}) error {
	q := query(ctx, db)
	return q.Scope(parentScope(parent, model)).Scope(filterScope(filter)).All(model)
}

func Create(ctx context.Context, db *pop.Connection, input, id, model interface{}) error {
	db = conn(ctx, db)
	_, err := assign(input, model)
	if err != nil {
		return errors.Wrap(err, "assign fields")
	}
	return db.Create(model)
}

func Update(ctx context.Context, db *pop.Connection, input, id, model interface{}) error {
	db = conn(ctx, db)
	err := db.Find(model, id)
	if err != nil {
		return errors.Wrap(err, "find model")
	}
	cols, err := assign(input, model)
	if err != nil {
		return errors.Wrap(err, "assign fields")
	}
	return db.UpdateColumns(model, cols...)
}

func Destroy(ctx context.Context, db *pop.Connection, model interface{}) error {
	db = conn(ctx, db)
	return db.Destroy(model)
}
