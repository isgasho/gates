package gates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayToNative(t *testing.T) {
	a := []Value{nil, Int(1), nil}
	b := Array(a)
	a[0] = b
	a[2] = b

	x := b.ToNative().([]interface{})
	assert.Equal(t, x, x[0])
	assert.Equal(t, int64(1), x[1])
	assert.Equal(t, x, x[2])

	assert.Equal(t, []interface{}(nil), Array(nil).ToNative())
}

func TestArrayIterator(t *testing.T) {
	a := Array([]Value{}).(_Array)
	it := a.Iterator()
	value, ok := it.Next()
	assert.Equal(t, Null, value)
	assert.False(t, ok)

	a = Array([]Value{Int(42), String("foo")}).(_Array)
	it = a.Iterator()
	value, ok = it.Next()
	assert.Equal(t, Int(42), value)
	assert.True(t, ok)
	value, ok = it.Next()
	assert.Equal(t, String("foo"), value)
	assert.True(t, ok)
	value, ok = it.Next()
	assert.Equal(t, Null, value)
	assert.False(t, ok)
}
