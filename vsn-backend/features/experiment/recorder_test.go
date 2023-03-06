package experiment

import (
	"io"
	"path/filepath"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestFSRecorder(t *testing.T) {
	fs := afero.NewMemMapFs() // in memory file system
	factory := NewFSRecorderFactory(fs)

	experimentId := uuid.New()
	trackingId := uuid.New()

	// create the recorder
	recorder := factory(recorderParams{
		ExperimentId: experimentId,
		TrackingId:   trackingId,
	})

	for i := 0; i < 2; i++ {
		// record twice to check headers don't get added twice
		recorder.Record(0, experimentData{
			Frames: []frame{{
				Timestamp: time.Now().UTC(),
				PositionX: 0,
				PositionY: 1,
				PositionZ: 2,
				RotationX: 0,
				RotationY: 1,
				RotationZ: 2,
			}},
			Events: []event{
				{Name: "TEST", Timestamp: time.Now().UTC()},
			},
		})
	}

	// try recording with empty arrays
	recorder.Record(0, experimentData{
		Frames: nil,
		Events: nil,
	})

	// open the frames file
	file, err := fs.Open(filepath.Join(experimentId.String(), trackingId.String(), "r0-frames.csv"))
	assert.NoError(t, err)

	data, _ := io.ReadAll(file)
	t.Log(string(data)) // just visually inspect

	// open the events file
	file, err = fs.Open(filepath.Join(experimentId.String(), trackingId.String(), "r0-events.csv"))
	assert.NoError(t, err)

	data, _ = io.ReadAll(file)
	t.Log(string(data)) // just visually inspect

}
