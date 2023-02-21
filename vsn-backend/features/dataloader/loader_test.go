package dataloader

import (
	"context"
	"math/rand"
	"sync"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type exampleRecord struct {
	id    uuid.UUID
	value int
}

type exampleRepository struct {
	calls int
	data  map[uuid.UUID]exampleRecord
}

func (r *exampleRepository) GetMany(ctx context.Context, ids []uuid.UUID) []exampleRecord {
	results := []exampleRecord{}
	for _, id := range ids {
		result, ok := r.data[id]
		if ok {
			results = append(results, result)
		}
	}

	// result order may not be guaranteed
	rand.Shuffle(len(results), func(i, j int) {
		results[i], results[j] = results[j], results[i]
	})

	r.calls += 1

	return results
}

func TestLoader(t *testing.T) {
	id1 := uuid.MustParse("a5feb306-e4d0-4a12-9623-833f41b982ab")
	id2 := uuid.MustParse("f5f7a64b-934d-4acb-9259-b88838bda28b")
	id3 := uuid.MustParse("3dd67850-0fd4-4f65-a686-39c322eab241")

	invalid := uuid.MustParse("caf24e65-0f70-41d1-bd99-0d9a38736bc8")

	repository := &exampleRepository{
		calls: 0,
		data: map[uuid.UUID]exampleRecord{
			id1: {id1, 1},
			id2: {id2, 2},
			id3: {id3, 3},
		},
	}

	loader := New(repository.GetMany, func(r exampleRecord) uuid.UUID {
		return r.id
	})

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		// get a record that exists
		record, err := loader.Get(context.Background(), id3)
		assert.Equal(t, 3, record.value)
		assert.NoError(t, err)
		wg.Done()
	}()

	go func() {
		// get a record that does not exist
		record, err := loader.Get(context.Background(), invalid)
		assert.Equal(t, 0, record.value)
		assert.Error(t, err)
		wg.Done()
	}()

	go func() {
		// get multiple records
		records := loader.GetMany(context.Background(), []uuid.UUID{id1, id2, id3, invalid})
		assert.Equal(t, 1, records[0].value)
		assert.Equal(t, 2, records[1].value)
		assert.Equal(t, 3, records[2].value)
		assert.Nil(t, records[3])
		wg.Done()
	}()

	wg.Wait()

	// the GetMany function was only called once because of the batching
	assert.Equal(t, 1, repository.calls)

}
