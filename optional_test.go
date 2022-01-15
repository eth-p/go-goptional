package goptional

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPresent(t *testing.T) {
	present := Optional[struct{}]{present: true}
	assert.True(t, present.Present())
}

func TestPresentEmpty(t *testing.T) {
	empty := Optional[struct{}]{present: false}
	assert.False(t, empty.Present())
}

func TestStringPresent(t *testing.T) {
	o := New[string]("foo")
	assert.Equal(t, "foo", o.String())
}

func TestStringEmpty(t *testing.T) {
	o := None[string]()
	assert.Equal(t, "(None)", o.String())
}

func TestUnwrap(t *testing.T) {
	o := New[string]("foo")
	value, err := o.Unwrap()
	assert.Equal(t, "foo", value)
	assert.Nil(t, err)
}

func TestUnwrapEmpty(t *testing.T) {
	o := None[string]()
	value, err := o.Unwrap()
	assert.Equal(t, "", value)
	assert.ErrorIs(t, EmptyOptional, err)
}

func TestUnwrapOr(t *testing.T) {
	o := New[string]("foo")
	value := o.UnwrapOr("bar")
	assert.Equal(t, "foo", value)
}

func TestUnwrapOrEmpty(t *testing.T) {
	o := None[string]()
	value := o.UnwrapOr("bar")
	assert.Equal(t, "bar", value)
}

func TestUnwrapOrElse(t *testing.T) {
	o := New[string]("foo")
	value := o.UnwrapOrElse(func() string {
		return "bar"
	})
	assert.Equal(t, "foo", value)
}

func TestUnwrapOrElseEmpty(t *testing.T) {
	o := None[string]()
	value := o.UnwrapOrElse(func() string {
		return "bar"
	})
	assert.Equal(t, "bar", value)
}

func TestExpect(t *testing.T) {
	o := New[string]("foo")
	assert.NotPanics(t, func() {
		value := o.Expect("no panic")
		assert.Equal(t, "foo", value)
	})
}

func TestExpectEmpty(t *testing.T) {
	o := None[string]()
	assert.Panics(t, func() {
		_ = o.Expect("panic")
	})
}
