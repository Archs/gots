// Tideland Go Testing Support - Asserts - Unit Tests
//
// Copyright (C) 2012-2014 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package asserts_test

//--------------------
// IMPORTS
//--------------------

import (
	"errors"
	"io"
	"testing"

	"github.com/tideland/gots/v3/asserts"
)

//--------------------
// TESTS
//--------------------

// TestAssertTrue tests the True() assertion.
func TestAssertTrue(t *testing.T) {
	assert := createValueAssertion(t)

	assert.True(true, "should not fail")
	if assert.True(false, "should fail and be logged") {
		t.Errorf("True() returned true")
	}
}

// TestAssertFalse tests the False() assertion.
func TestAssertFalse(t *testing.T) {
	assert := createValueAssertion(t)

	assert.False(false, "should not fail")
	if assert.False(true, "should fail and be logged") {
		t.Errorf("False() returned true")
	}
}

// TestAssertNil tests the Nil() assertion.
func TestAssertNil(t *testing.T) {
	assert := createValueAssertion(t)
	assert.Nil(nil, "should not fail")
	if assert.Nil("not nil", "should fail and be logged") {
		t.Errorf("Nil() returned true")
	}
}

// TestAssertNotNil tests the NotNil() assertion.
func TestAssertNotNil(t *testing.T) {
	assert := createValueAssertion(t)

	assert.NotNil("not nil", "should not fail")
	if assert.NotNil(nil, "should fail and be logged") {
		t.Errorf("NotNil() returned true")
	}
}

// TestAssertEqual tests the Equal() assertion.
func TestAssertEqual(t *testing.T) {
	assert := createValueAssertion(t)
	m := map[string]int{"one": 1, "two": 2, "three": 3}

	assert.Equal(nil, nil, "should not fail")
	assert.Equal(true, true, "should not fail")
	assert.Equal(1, 1, "should not fail")
	assert.Equal("foo", "foo", "should not fail")
	assert.Equal(map[string]int{"one": 1, "three": 3, "two": 2}, m, "should not fail")
	if assert.Equal("one", 1, "should fail and be logged") {
		t.Errorf("Equal() returned true")
	}
	if assert.Equal("two", "2", "should fail and be logged") {
		t.Errorf("Equal() returned true")
	}
}

// TestAssertDifferent tests the Different() assertion.
func TestAssertDifferent(t *testing.T) {
	assert := createValueAssertion(t)
	m := map[string]int{"one": 1, "two": 2, "three": 3}

	assert.Different(nil, "nil", "should not fail")
	assert.Different("true", true, "should not fail")
	assert.Different(1, 2, "should not fail")
	assert.Different("foo", "bar", "should not fail")
	assert.Different(map[string]int{"three": 3, "two": 2}, m, "should not fail")
	if assert.Different("one", "one", "should fail and be logged") {
		t.Errorf("Different() returned true")
	}
	if assert.Different(2, 2, "should fail and be logged") {
		t.Errorf("Different() returned true")
	}
}

// TestAssertSubstring tests the Substring() assertion.
func TestAssertSubstring(t *testing.T) {
	assert := createValueAssertion(t)

	assert.Substring("this is assert test", "is assert", "should not fail")
	assert.Substring("this is 1 test", "test", "should not fail")
	if assert.Substring("this is assert test", "foo", "should fail and be logged") {
		t.Errorf("Substring() returned true")
	}
	if assert.Substring("this is assert test", "this   is   assert   test", "should fail and be logged") {
		t.Errorf("Substring() returned true")
	}
}

// TestAssertMatch tests the Match() assertion.
func TestAssertMatch(t *testing.T) {
	assert := createValueAssertion(t)

	assert.Match("this is assert test", "this.*test", "should not fail")
	assert.Match("this is 1 test", "this is [0-9] test", "should not fail")
	if assert.Match("this is assert test", "foo", "should fail and be logged") {
		t.Errorf("Match() returned true")
	}
	if assert.Match("this is assert test", "this*test", "should fail and be logged") {
		t.Errorf("Match() returned true")
	}
}

// TestAssertErrorMatch tests the ErrorMatch() assertion.
func TestAssertErrorMatch(t *testing.T) {
	assert := createValueAssertion(t)
	err := errors.New("oops, an error")

	assert.ErrorMatch(err, "oops, an error", "should not fail")
	assert.ErrorMatch(err, "oops,.*", "should not fail")
	if assert.ErrorMatch(err, "foo", "should fail and be logged") {
		t.Errorf("ErrorMatch() returned true")
	}
}

// TestAssertImplementor tests the Implementor() assertion.
func TestAssertImplementor(t *testing.T) {
	assert := createTypeAssertion(t)

	var err error
	var w io.Writer

	assert.Implementor(errors.New("error test"), &err, "should not fail")
	if assert.Implementor("string test", &err, "should fail and be logged") {
		t.Errorf("Implementor() returned true")
	}
	if assert.Implementor(errors.New("error test"), &w, "should fail and be logged") {
		t.Errorf("Implementor() returned true")
	}
}

// TestAssertAssignable tests the Assignable() assertion.
func TestAssertAssignable(t *testing.T) {
	assert := createTypeAssertion(t)

	assert.Assignable(1, 5, "should not fail")
	if assert.Assignable("one", 5, "should fail and be logged") {
		t.Errorf("Assignable() returned true")
	}
}

// TestAssertUnassignable tests the Unassignable() assertion.
func TestAssertUnassignable(t *testing.T) {
	assert := createTypeAssertion(t)

	assert.Unassignable("one", 5, "should not fail")
	if assert.Unassignable(1, 5, "should fail and be logged") {
		t.Errorf("Unassignable() returned true")
	}
}

// TestAssertEmpty tests the Empty() assertion.
func TestAssertEmpty(t *testing.T) {
	assert := createValueAssertion(t)

	assert.Empty("", "should not fail")
	assert.Empty([]bool{}, "should also not fail")
	if assert.Empty("not empty", "should fail and be logged") {
		t.Errorf("Empty() returned true")
	}
	if assert.Empty([3]int{1, 2, 3}, "should also fail and be logged") {
		t.Errorf("Empty() returned true")
	}
	if assert.Empty(true, "illegal type has to fail") {
		t.Errorf("Empty() returned true")
	}
}

// TestAssertNotEmpty tests the NotEmpty() assertion.
func TestAsserNotEmpty(t *testing.T) {
	assert := createValueAssertion(t)

	assert.NotEmpty("not empty", "should not fail")
	assert.NotEmpty([3]int{1, 2, 3}, "should also not fail")
	if assert.NotEmpty("", "should fail and be logged") {
		t.Errorf("NotEmpty() returned true")
	}
	if assert.NotEmpty([]int{}, "should also fail and be logged") {
		t.Errorf("NotEmpty() returned true")
	}
	if assert.NotEmpty(true, "illegal type has to fail") {
		t.Errorf("NotEmpty() returned true")
	}
}

// TestAssertLength tests the Length() assertion.
func TestAssertLength(t *testing.T) {
	assert := createValueAssertion(t)

	assert.Length("", 0, "should not fail")
	assert.Length([]bool{true, false}, 2, "should also not fail")
	if assert.Length("not empty", 0, "should fail and be logged") {
		t.Errorf("Length() returned true")
	}
	if assert.Length([3]int{1, 2, 3}, 10, "should also fail and be logged") {
		t.Errorf("Length() returned true")
	}
	if assert.Length(true, 1, "illegal type has to fail") {
		t.Errorf("Length() returned true")
	}
}

// TestAssertPanics tests the Panics() assertion.
func TestAssertPanics(t *testing.T) {
	assert := createValueAssertion(t)

	if !assert.Panics(func() { panic("ouch") }, "should panic") {
		t.Errorf("Panics() returned false")
	}
	if assert.Panics(func() { _ = 1 + 1 }, "should not panic") {
		t.Errorf("Panics() returned true")
	}
}

// TestAssertFail tests the fail asserts.
func TestAssertFail(t *testing.T) {
	assert := createValueAssertion(t)

	assert.Fail("this should fail")
}

// TestTestingAssertion tests the testing assertion.
func TestTestingAssertion(t *testing.T) {
	assert := asserts.NewTestingAssertion(t, false)
	foo := func() {}
	bar := 4711

	assert.Assignable(47, 11, "should not fail")
	assert.Assignable(foo, bar, "should fail (but not the test)")
	assert.Assignable(foo, bar)
	assert.Assignable(foo, bar, "this", "should", "fail", "too")
}

// TestPanicAssertion tests if the panic assertions panic when they fail.
func TestPanicAssert(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Logf("panic worked: '%v'", err)
		}
	}()

	assert := asserts.NewPanicAssertion()
	foo := func() {}

	assert.Assignable(47, 11, "should not fail")
	assert.Assignable(47, foo, "should fail")

	t.Errorf("should not be reached")
}

//--------------------
// FAIL FUNCS
//--------------------

type logFailer struct {
	t        *testing.T
	describe bool
}

func (f *logFailer) Logf(format string, args ...interface{}) {
	f.t.Logf(format, args...)
}

func (f *logFailer) Fail(test asserts.Test, obtained, expected interface{}, msgs ...string) bool {
	if f.describe {
		f.Logf("testing assert '%s' failed: '%v' <> '%v' (%s)",
			test, asserts.ValueDescription(obtained), asserts.ValueDescription(expected), msgs)
	} else {
		f.Logf("testing assert '%s' failed: '%v' <> '%v' (%s)", test, obtained, expected, msgs)
	}
	return false
}

// createValueAssertion returns an assertion with assert value logging fail func.
func createValueAssertion(t *testing.T) asserts.Assertion {
	return asserts.NewAssertion(&logFailer{t, false})
}

// createTypeAssertion returns an assertion with assert value description (type) logging fail func.
func createTypeAssertion(t *testing.T) asserts.Assertion {
	return asserts.NewAssertion(&logFailer{t, true})
}

// EOF
