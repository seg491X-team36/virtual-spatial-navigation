package experiment

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatusDone(t *testing.T) {
	t.Run("tc1", func(t *testing.T) {
		s1 := ExperimentStatus{
			RoundInProgress: false, // not in progress
			RoundNumber:     1,     // round number < round total
			RoundsTotal:     2,
		}

		assert.False(t, s1.Done())
	})

	t.Run("tc2", func(t *testing.T) {
		s1 := ExperimentStatus{
			RoundInProgress: true, // not in progress
			RoundNumber:     1,    // round number < round total
			RoundsTotal:     2,
		}

		assert.False(t, s1.Done())
	})

	t.Run("tc3", func(t *testing.T) {
		s1 := ExperimentStatus{
			RoundInProgress: true, // not in progress
			RoundNumber:     2,    // round number < round total
			RoundsTotal:     2,
		}

		assert.False(t, s1.Done())
	})

	t.Run("tc4", func(t *testing.T) {
		s1 := ExperimentStatus{
			RoundInProgress: false, // not in progress
			RoundNumber:     2,     // round number < round total
			RoundsTotal:     2,
		}

		assert.True(t, s1.Done())
	})
}
