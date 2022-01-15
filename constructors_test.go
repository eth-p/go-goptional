package goptional

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNone(t *testing.T) {
	empty := None[string]()
	assert.False(t, empty.Present())
}

func TestNew(t *testing.T) {
	present := New[string]("foo")
	assert.True(t, present.Present())
	assert.Equal(t, "foo", present.value)
}

func TestFromTrue(t *testing.T) {
	present := From("foo", true)
	assert.True(t, present.Present())
	assert.Equal(t, "foo", present.value)
}

func TestFromFalse(t *testing.T) {
	present := From("foo", false)
	assert.False(t, present.Present())
	assert.Equal(t, "", present.value) // Default
}

func TestFromErrorValue(t *testing.T) {
	present := FromError("foo", nil)
	assert.True(t, present.Present())
	assert.Equal(t, "foo", present.value)
}

func TestFromErrorError(t *testing.T) {
	present := FromError("foo", errors.New("an error"))
	assert.False(t, present.Present())
	assert.Equal(t, "", present.value) // Default
}

func TestFromPtrValue(t *testing.T) {
	str := "foo"
	present := FromPtr(&str)
	assert.True(t, present.Present())
	assert.Equal(t, &str, present.value)
}

func TestFromPtrNil(t *testing.T) {
	present := FromPtr[string](nil)
	assert.False(t, present.Present())
	assert.Nil(t, present.value)
}
