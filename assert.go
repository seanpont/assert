package assert

import (
	"fmt"
	"reflect"
	"testing"
	"runtime"
	"strings"
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

func caller() string {
	fname, line, caller := "assert.go", 0, 1
	for ; strings.HasSuffix(fname, "assert.go") && caller<4; caller++ {
		_, fname, line, _ = runtime.Caller(caller)
	}
	fname = fname[strings.LastIndex(fname, "/")+1:]
	return fmt.Sprintf("\n%v:%v ", fname, line)
}

func (a *Assertion) True(b bool, message string, messageParams ...interface{}) {
	if !b {
		a.t.Fatalf(caller() + message, messageParams...)
	}
}

func (a *Assertion) False(b bool, message string, messageParams ...interface{}) {
	a.True(!b, message, messageParams...)
}

func (a *Assertion) Nil(val interface{}) {
	message := fmt.Sprintf("Expected nil but was %v", val)
	eq := reflect.DeepEqual(val, nil)
	a.True(eq, message) 
}

func (a *Assertion) NotNil(val interface{}) {
	eq := reflect.DeepEqual(val, nil)
	a.True(!eq, "Expected not nil but was nil")
}

func (a *Assertion) Equal(actual, expected interface{}) {
	eq := reflect.DeepEqual(actual, expected)
	a.True(eq, "\nExpected: %v\nReceived: %v", expected, actual)
}

func (a *Assertion) NotEqual(actual, expected interface{}) {
	eq := reflect.DeepEqual(actual, expected)
	a.True(!eq, "Expected %v to not equal %v, but it did", expected, actual)
}
