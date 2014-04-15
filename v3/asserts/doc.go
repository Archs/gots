// Tideland Go Testing Support - Asserts
//
// Copyright (C) 2012-2014 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

// The asserts package provides a powerful and convenient set of assertions. Here
// expected and obtained values are compared. Individual descriptions as well as the
// output help to quickly find failed tests.
//
// In the beginning of a test function a new assertion instance is created with:
//
// assert := asserts.NewTestingAssertion(t, shallFail)
//
// Inside the test an assert looks like:
//
// assert.Equal(expected, obtained, "obtained value has to be like expected")
//
// If shallFail is set to true a failing assert also lets fail the Go test.
// Otherwise the failing is printed but the tests continue.
package asserts

// EOF
