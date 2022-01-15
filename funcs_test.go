package goptional

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapSameType(t *testing.T) {
	cubed := func(v int) int {
		return v * v
	}

	present := New(5)
	assert.Equal(t, 5, present.value)
	mapped := Map(present, cubed)
	assert.Equal(t, cubed(5), mapped.value)
}

func TestMapDifferentType(t *testing.T) {
	toString := func(v int) string {
		return fmt.Sprintf(`"%v"`, v)
	}

	present := New(5)
	assert.Equal(t, 5, present.value)
	mapped := Map(present, toString)
	assert.Equal(t, `"5"`, mapped.value)
}

func TestMapEmpty(t *testing.T) {
	called := false
	cubed := func(v int) Optional[int] {
		called = true
		return New(v * v)
	}

	empty := None[int]()
	assert.False(t, empty.Present())
	mapped := MapInto(empty, cubed)
	assert.False(t, mapped.Present())
	assert.False(t, called)
}

func TestMapIntoPresentToPresent(t *testing.T) {
	cubed := func(v int) Optional[int] {
		return New(v * v)
	}

	present := New(5)
	assert.Equal(t, 5, present.value)
	mapped := MapInto(present, cubed)
	assert.True(t, mapped.Present())
	assert.Equal(t, 25, mapped.value)
}

func TestMapIntoPresentToEmpty(t *testing.T) {
	consume := func(v int) Optional[int] {
		return None[int]()
	}

	present := New(5)
	assert.Equal(t, 5, present.value)
	mapped := MapInto(present, consume)
	assert.False(t, mapped.Present())
}

func TestMapIntoEmpty(t *testing.T) {
	cubed := func(v int) Optional[int] {
		return New(v * v)
	}

	empty := None[int]()
	assert.False(t, empty.Present())
	mapped := MapInto(empty, cubed)
	assert.False(t, mapped.Present())
}
