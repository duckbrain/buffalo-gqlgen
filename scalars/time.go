package scalars

import (
	"fmt"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gobuffalo/buffalo/binding/decoders"
	"github.com/gobuffalo/nulls"
)

func MarshalTime(t time.Time) graphql.Marshaler {
	return stringWriter(t.Format(time.RFC3339))
}

var decodeTime = decoders.TimeDecoderFn()

func UnmarshalTime(v interface{}) (t time.Time, err error) {
	switch v := v.(type) {
	case string:
		var r interface{}
		r, err = decodeTime([]string{v})
		if err != nil {
			break
		}
		t = r.(time.Time)
	case int:
		t = time.Unix(0, int64(v))
	default:
		err = fmt.Errorf("%T is not a bool", v)
	}
	return
}

func MarshalNullsTime(t nulls.Time) graphql.Marshaler {
	if t.Valid {
		return MarshalTime(t.Time)
	}
	return nullWriter
}

func UnmarshalNullsTime(v interface{}) (nulls.Time, error) {
	if v == nil {
		return nulls.Time{}, nil
	}
	t, err := UnmarshalTime(v)
	if err != nil {
		return nulls.Time{}, err
	}
	return nulls.Time{Valid: true, Time: t}, nil
}

func MarshalWeekday(t time.Weekday) graphql.Marshaler {
	return stringWriter(t.String())
}

var weekdays = map[string]time.Weekday{
	"Sun":       time.Sunday,
	"Mon":       time.Monday,
	"Tue":       time.Tuesday,
	"Wed":       time.Wednesday,
	"Thu":       time.Thursday,
	"Fri":       time.Friday,
	"Sat":       time.Saturday,
	"Sunday":    time.Sunday,
	"Monday":    time.Monday,
	"Tuesday":   time.Tuesday,
	"Wednesday": time.Wednesday,
	"Thursday":  time.Thursday,
	"Friday":    time.Friday,
	"Saturday":  time.Saturday,
}

func UnmarshalWeekday(v interface{}) (w time.Weekday, err error) {
	switch v := v.(type) {
	case string:
		var ok bool
		w, ok = weekdays[v]
		if !ok {
			err = fmt.Errorf("unknown weekday \"%v\"", v)
		}
	case int:
		w = time.Weekday(v)
	}
	return
}
