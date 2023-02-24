package experiment

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"

	"github.com/google/uuid"
	"github.com/spf13/afero"
)

type recorder interface {
	Record(round int, data experimentData)
}

type recorderParams struct {
	ExperimentId uuid.UUID
	TrackingId   uuid.UUID
}

type recorderFactory func(params recorderParams) recorder

type FSRecorder struct {
	sync.Once
	recorderParams
	fs afero.Fs
}

func NewFSRecorderFactory(fs afero.Fs) recorderFactory {
	return func(params recorderParams) recorder {
		return &FSRecorder{
			recorderParams: params,
			fs:             fs,
		}
	}
}

func (r *FSRecorder) Record(round int, data experimentData) {
	r.Once.Do(func() {
		// check the destination folder exists
		path := r.folderPath()
		if _, err := r.fs.Stat(path); err != nil {
			_ = r.fs.MkdirAll(path, 0700)
		}
	})

	// write frames and events in parallel
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		r.recordEvents(round, data.Events)
		wg.Done()
	}()

	go func() {
		r.recordFrames(round, data.Frames)
		wg.Done()
	}()

	wg.Wait()
}

func (r *FSRecorder) recordEvents(round int, events []event) {
	if len(events) == 0 {
		return
	}

	// open the file and setup the csv writer
	f, created := r.eventsFile(round)
	defer f.Close()
	w := csv.NewWriter(f)

	// write headers
	if created {
		eventWriter.WriteHeaders(w)
	}

	// write events
	eventWriter.Write(events, w)
}

func (r *FSRecorder) recordFrames(round int, frames []frame) {
	if len(frames) == 0 {
		return
	}

	// open the file and setup the csv writer
	f, created := r.framesFile(round) // open the
	defer f.Close()
	w := csv.NewWriter(f)

	// write headers
	if created {
		frameWriter.WriteHeaders(w)
	}

	// write frames
	frameWriter.Write(frames, w)
}

func (r *FSRecorder) folderPath() string {
	return filepath.Join(r.ExperimentId.String(), r.TrackingId.String())
}

func (r *FSRecorder) framesPath(round int) string {
	return filepath.Join(r.folderPath(), fmt.Sprintf("r%d-frames.csv", round))
}

func (r *FSRecorder) eventsPath(round int) string {
	return filepath.Join(r.folderPath(), fmt.Sprintf("r%d-events.csv", round))
}

func (r *FSRecorder) framesFile(round int) (w io.WriteCloser, created bool) {
	// check the path
	path := r.framesPath(round)
	_, err := r.fs.Stat(path)
	created = err != nil

	// open the file
	w, _ = r.fs.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	return w, created
}

func (r *FSRecorder) eventsFile(round int) (w io.WriteCloser, created bool) {
	// check the path
	path := r.eventsPath(round)
	_, err := r.fs.Stat(path)
	created = err != nil

	// open the file
	w, _ = r.fs.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	return w, created
}
