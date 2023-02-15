package scalars

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMarshalTime(t *testing.T) {
	// two example times
	ex1, _ := time.Parse(time.RFC3339Nano, "2023-02-15T02:32:54.6498638Z")
	ex2, _ := time.Parse(time.RFC3339Nano, "2023-02-14T21:32:54.6498638-05:00")
	loc, _ := time.LoadLocation("Europe/Amsterdam") // example timezone

	times := []time.Time{ex1, ex1.In(loc), ex2, ex2.In(loc)} // 4 "differently created" times

	for _, time := range times {
		var b bytes.Buffer
		MarshalTime(time).MarshalGQL(&b)
		actual := string(b.Bytes())

		assert.Equal(t, "\"2023-02-15T02:32:54.6498638Z\"", actual) // serialize to utc time
	}
}

func TestUnmarshalTime(t *testing.T) {
	_, err := UnmarshalTime(1) // invalid data type
	assert.Error(t, err)

	_, err = UnmarshalTime("") // invalid string
	assert.Error(t, err)

	target, _ := time.Parse(time.RFC3339Nano, "2023-02-15T02:32:54.6498638Z")

	parsed1, err := UnmarshalTime("2023-02-15T02:32:54.6498638Z")
	assert.Equal(t, target, parsed1)
	assert.NoError(t, err)

	parsed2, err := UnmarshalTime("2023-02-14T21:32:54.6498638-05:00")
	assert.Equal(t, target, parsed2)
	assert.NoError(t, err)
}
