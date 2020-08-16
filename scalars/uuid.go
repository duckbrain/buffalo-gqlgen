package scalars

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gofrs/uuid"
)

func MarshalUUID(id uuid.UUID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		w.Write([]byte(strconv.Quote(id.String())))
	})
}

func UnmarshalUUID(v interface{}) (uuid.UUID, error) {
	switch v := v.(type) {
	case string:
		return uuid.FromString(v)
	default:
		return uuid.Nil, fmt.Errorf("%T is not a bool", v)
	}
}
