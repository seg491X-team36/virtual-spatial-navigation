package dataloader

import (
	"context"
	"errors"

	dl "github.com/graph-gophers/dataloader/v7"
)

type Loader[K, V any] interface {
	Get(ctx context.Context, key K) (V, error)
	GetMany(ctx context.Context, keys []K) []*V
}

type BatchFunc[K, V any] func(ctx context.Context, keys []K) []V

type KeyFunc[K, V any] func(V) K

type dataLoader[K comparable, V any] struct {
	loader *dl.Loader[K, V]
}

func (d *dataLoader[K, V]) Get(ctx context.Context, key K) (V, error) {
	return d.loader.Load(ctx, key)()
}

func (d *dataLoader[K, V]) GetMany(ctx context.Context, keys []K) []*V {
	data, errors := d.loader.LoadMany(ctx, keys)()
	results := make([]*V, len(keys))

	for i, result := range data {
		temp := result
		if errors[i] != nil {
			results[i] = nil
		} else {
			results[i] = &temp
		}
	}

	return results
}

func New[K comparable, V any](
	batchFn BatchFunc[K, V],
	keyFn KeyFunc[K, V],
) Loader[K, V] {
	loader := dl.NewBatchedLoader(
		batchFnWrapper(batchFn, keyFn),
		dl.WithCache[K, V](&dl.NoCache[K, V]{}))

	return &dataLoader[K, V]{
		loader: loader,
	}
}

/* batchFnWrapper
- handles sorting results so that the batchFn does not have to. SQL In query is not sorted. 
- if no result is found for a requested key it will add a result with a not found error
*/
func batchFnWrapper[K comparable, V any](batchFn BatchFunc[K, V], keyFn KeyFunc[K, V]) dl.BatchFunc[K, V] {
	return func(ctx context.Context, keys []K) []*dl.Result[V] {
		// index results by key
		results := map[K]V{}
		for _, result := range batchFn(ctx, keys) {
			results[keyFn(result)] = result
		}

		// result list
		resultList := make([]*dl.Result[V], len(keys))

		for i, key := range keys {
			// result and error for each key
			result, ok := results[key]
			var err error = nil
			if !ok {
				err = errors.New("not found")
			}

			resultList[i] = &dl.Result[V]{
				Data:  result,
				Error: err,
			}
		}

		return resultList
	}
}
