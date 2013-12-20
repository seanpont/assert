package assert

import (
	_ "fmt"
	"reflect"
	"testing"
)

// Usage:
// assert := assert.Assert(t)
// ...
// assert.Equal(actual, expected)
// assert.Nil(err, "X threw error")
func Assert(t *testing.T) *Assertion {
	return &Assertion{t}
}

type Assertion struct {
	t *testing.T
}

func (a *Assertion) True(b bool, message string, messageParams ...interface{}) {
	if !b {
		a.t.Fatalf(message, messageParams...)
	}
}

func (a *Assertion) False(b bool, message string, messageParams ...interface{}) {
	a.True(!b, message, messageParams...)
}

func (a *Assertion) Nil(val interface{}, message string, messageParams ...interface{}) {
	eq := reflect.DeepEqual(val, nil)
	a.True(eq, message, messageParams...)
}

func (a *Assertion) NotNil(val interface{}, message string, messageParams ...interface{}) {
	eq := reflect.DeepEqual(val, nil)
	a.True(!eq, message, messageParams...)
}

func (a *Assertion) Equal(actual, expected interface{}) {
	eq := reflect.DeepEqual(actual, expected)
	a.True(eq, "\nExpected: %v\nReceived: %v", expected, actual)
}

func (a *Assertion) NotEqual(actual, expected interface{}) {
	eq := reflect.DeepEqual(actual, expected)
	a.True(!eq, "Expected %v to not equal %v, but it did", expected, actual)
}
