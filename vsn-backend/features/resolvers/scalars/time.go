package scalars

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalTime(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, strconv.Quote(t.UTC().Format(time.RFC3339Nano)))
	})
}

func UnmarshalTime(v interface{}) (time.Time, error) {
	s, ok := v.(string)
	if !ok {
		var t time.Time
		return t, fmt.Errorf("invalid type %T, expect string", v)
	}

	t, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		return t, err
	}

	return t.UTC(), nil
}
