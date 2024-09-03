package catch

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTryCatch(t *testing.T) {
	t.Run("No exception", func(t *testing.T) {
		var i int
		Try(func() { i = 100 }).Run()
		assert.Equal(t, 100, i)
	})

	t.Run("No exception with finally", func(t *testing.T) {
		var i int
		Try(func() { i = 100 }).Finally(func() { i = 1000 }).Run()
		assert.Equal(t, 1000, i)
	})

	t.Run("Exception without catch", func(t *testing.T) {
		assert.PanicsWithValue(t, "exception", Try(func() {
			panic("exception")
		}).Run)
	})

	t.Run("Exception with finally", func(t *testing.T) {
		var i int
		assert.PanicsWithError(t, "exception", Try(func() { panic(errors.New("exception")) }).Finally(func() { i = 100 }).Run)
		assert.Equal(t, 100, i)
	})

	t.Run("Exception with catch", func(t *testing.T) {
		assert.NotPanics(t, Try(func() { panic(errors.New("exception")) }).Catch(errors.New(""), func(a error) {
		}).Run)
	})

	t.Run("Exception with catch all", func(t *testing.T) {
		assert.NotPanics(t, Try(func() { panic(errors.New("exception")) }).CatchAll(func(a error) {
		}).Run)
	})
}
