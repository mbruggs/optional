package optional

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("with int", func(t *testing.T) {
		v := 5
		opt := New(v)

		got, valid := opt.Value()
		assert.Equal(t, v, got)
		assert.True(t, valid)

		v2 := opt.Valid()
		assert.True(t, v2)
	})

	t.Run("with string", func(t *testing.T) {
		s := "abcd"
		opt := New(s)

		got, valid := opt.Value()
		assert.Equal(t, s, got)
		assert.True(t, valid)

		v2 := opt.Valid()
		assert.True(t, v2)
	})

	t.Run("with own struct", func(t *testing.T) {
		type s struct {
			f1 int
			f2 string
		}

		st := s{5, "abcd"}

		opt := New(st)

		got, valid := opt.Value()
		assert.Equal(t, st, got)
		assert.True(t, valid)

		v2 := opt.Valid()
		assert.True(t, v2)
	})
}

func TestEmpty(t *testing.T) {
	t.Run("with int", func(t *testing.T) {
		opt := Empty[int]()

		got, valid := opt.Value()
		assert.Equal(t, 0, got)
		assert.False(t, valid)

		v2 := opt.Valid()
		assert.False(t, v2)
	})

	t.Run("with string", func(t *testing.T) {
		opt := Empty[string]()

		got, valid := opt.Value()
		assert.Equal(t, "", got)
		assert.False(t, valid)

		v2 := opt.Valid()
		assert.False(t, v2)
	})

	t.Run("with own struct", func(t *testing.T) {
		type s struct {
			f1 int
			f2 string
		}

		opt := Empty[s]()

		got, valid := opt.Value()
		assert.Empty(t, got)
		assert.False(t, valid)

		v2 := opt.Valid()
		assert.False(t, v2)
	})
}
