package experiment

import (
	"encoding/csv"
	"fmt"
	"time"
)

type csvWriter[V any] struct {
	Headers   []string
	Serialize func(V) []string
}

func (r *csvWriter[V]) Write(records []V, w *csv.Writer) {
	for _, record := range records {
		w.Write(r.Serialize(record))
	}
	w.Flush()
}

func (r *csvWriter[V]) WriteHeaders(w *csv.Writer) {
	w.Write(r.Headers)
}

var (
	eventWriter = csvWriter[event]{
		Headers: []string{
			"timestamp",
			"event_name",
		},
		Serialize: func(e event) []string {
			// serialize the event
			return []string{
				e.Timestamp.Format(time.RFC3339Nano),
				e.Name,
			}
		},
	}

	frameWriter = csvWriter[frame]{
		Headers: []string{
			"timestamp",
			"position_x",
			"position_y",
			"position_z",
			"rotation_x",
			"rotation_y",
			"rotation_z",
		},
		Serialize: func(f frame) []string {
			// serialize the event
			return []string{
				f.Timestamp.Format(time.RFC3339Nano),
				fmt.Sprintf("%.7f", f.PositionX),
				fmt.Sprintf("%.7f", f.PositionY),
				fmt.Sprintf("%.7f", f.PositionZ),
				fmt.Sprintf("%.7f", f.RotationX),
				fmt.Sprintf("%.7f", f.RotationY),
				fmt.Sprintf("%.7f", f.RotationZ),
			}
		},
	}
)
