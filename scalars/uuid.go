package scalars

import (
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gobuffalo/nulls"
	"github.com/gofrs/uuid"
)

func MarshalUUID(id uuid.UUID) graphql.Marshaler {
	return stringWriter(id.String())
}

func UnmarshalUUID(v interface{}) (uuid.UUID, error) {
	switch v := v.(type) {
	case string:
		return uuid.FromString(v)
	default:
		return uuid.Nil, fmt.Errorf("%T is not a bool", v)
	}
}

func MarshalNullUUID(id uuid.NullUUID) graphql.Marshaler {
	if id.Valid {
		return MarshalUUID(id.UUID)
	}
	return nullWriter
}

func UnmarshalNullUUID(v interface{}) (uuid.NullUUID, error) {
	if v == nil {
		return uuid.NullUUID{}, nil
	}
	id, err := UnmarshalUUID(v)
	if err != nil {
		return uuid.NullUUID{}, err
	}
	return uuid.NullUUID{Valid: true, UUID: id}, nil
}
func MarshalNullsUUID(id nulls.UUID) graphql.Marshaler {
	if id.Valid {
		return MarshalUUID(id.UUID)
	}
	return nullWriter
}

func UnmarshalNullsUUID(v interface{}) (nulls.UUID, error) {
	if v == nil {
		return nulls.UUID{}, nil
	}
	id, err := UnmarshalUUID(v)
	if err != nil {
		return nulls.UUID{}, err
	}
	return nulls.UUID{Valid: true, UUID: id}, nil
}
