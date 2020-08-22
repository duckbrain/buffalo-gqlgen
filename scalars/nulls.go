package scalars

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/gobuffalo/nulls"
)

type nullable interface {
	MarshalJSON() ([]byte, error)
	Interface() interface{}
}

func marshallNullable(n nullable) graphql.Marshaler {
	if n.Interface() == nil {
		return nullWriter
	}
	b, err := n.MarshalJSON()
	if err != nil {
		panic(err)
	}
	return outputWriter(string(b))
}

func MarshallNullsInt(i nulls.Int) graphql.Marshaler {
	return marshallNullable(i)
}
func MarshallNullsInt32(i nulls.Int32) graphql.Marshaler {
	return marshallNullable(i)
}
func MarshallNullsInt64(i nulls.Int64) graphql.Marshaler {
	return marshallNullable(i)
}
func MarshallNullsUint32(i nulls.UInt32) graphql.Marshaler {
	return marshallNullable(i)
}
func MarshallNullsBool(i nulls.Bool) graphql.Marshaler {
	return marshallNullable(i)
}
func MarshallNullsFloat32(i nulls.Float32) graphql.Marshaler {
	return marshallNullable(i)
}
func MarshallNullsFloat64(i nulls.Float64) graphql.Marshaler {
	return marshallNullable(i)
}
func MarshallNullsString(i nulls.String) graphql.Marshaler {
	return marshallNullable(i)
}
