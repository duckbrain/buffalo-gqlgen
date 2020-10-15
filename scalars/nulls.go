package scalars

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/gobuffalo/nulls"
)

type nullable interface {
	MarshalJSON() ([]byte, error)
	Interface() interface{}
}

func MarshalNullable(n nullable) graphql.Marshaler {
	if n.Interface() == nil {
		return nullWriter
	}
	b, err := n.MarshalJSON()
	if err != nil {
		panic(err)
	}
	return outputWriter(string(b))
}

func MarshalNullsInt(i nulls.Int) graphql.Marshaler {
	return MarshalNullable(i)
}
func UnmarshalNullsInt(v interface{}) (i nulls.Int, err error) {
	if v == nil {
		return
	}
	i.Int, err = graphql.UnmarshalInt(v)
	if err != nil {
		return
	}
	i.Valid = true
	return

}
func MarshalNullsInt32(i nulls.Int32) graphql.Marshaler {
	return MarshalNullable(i)
}
func UnmarshalNullsInt32(v interface{}) (i nulls.Int32, err error) {
	if v == nil {
		return
	}
	i.Int32, err = graphql.UnmarshalInt32(v)
	if err != nil {
		return
	}
	i.Valid = true
	return
}
func MarshalNullsInt64(i nulls.Int64) graphql.Marshaler {
	return MarshalNullable(i)
}
func UnmarshalNullsInt64(v interface{}) (i nulls.Int64, err error) {
	if v == nil {
		return
	}
	i.Int64, err = graphql.UnmarshalInt64(v)
	if err != nil {
		return
	}
	i.Valid = true
	return
}
func MarshalNullsUint32(i nulls.UInt32) graphql.Marshaler {
	return MarshalNullable(i)
}

// TODO
// func UnmarshalNullsUint32(v interface{}) (i nulls.UInt32, err error) {
// 	if v == nil {
// 		return
// 	}
// 	i.Uint32, err = graphql.UnmarshalUint32(v)
// 	if err != nil {
// 		return
// 	}
// 	i.Valid = true
// 	return
// }
func MarshalNullsBool(i nulls.Bool) graphql.Marshaler {
	return MarshalNullable(i)
}
func UnmarshalNullsBool(v interface{}) (i nulls.Bool, err error) {
	if v == nil {
		return
	}
	i.Bool, err = graphql.UnmarshalBoolean(v)
	if err != nil {
		return
	}
	i.Valid = true
	return
}
func MarshalNullsFloat32(i nulls.Float32) graphql.Marshaler {
	return MarshalNullable(i)
}
func UnmarshalNullsFloat32(v interface{}) (i nulls.Float32, err error) {
	if v == nil {
		return
	}
	f, err := graphql.UnmarshalFloat(v)
	if err != nil {
		return
	}
	i.Float32 = float32(f)
	i.Valid = true
	return
}
func MarshalNullsFloat64(i nulls.Float64) graphql.Marshaler {
	return MarshalNullable(i)
}
func UnmarshalNullsFloat64(v interface{}) (i nulls.Float64, err error) {
	if v == nil {
		return
	}
	i.Float64, err = graphql.UnmarshalFloat(v)
	if err != nil {
		return
	}
	i.Valid = true
	return
}
func MarshalNullsString(i nulls.String) graphql.Marshaler {
	return MarshalNullable(i)
}
func UnmarshalNullsString(v interface{}) (i nulls.String, err error) {
	if v == nil {
		return
	}
	i.String, err = graphql.UnmarshalString(v)
	if err != nil {
		return
	}
	i.Valid = true
	return
}
