Tideland Go Testing Support
===========================

Description
-----------

The *Tideland Go Testing Support* (GOTS) contains the two packages

- `asserts` to extend the standard Go testing package and
- `generators` for the generating of test data.

Installation
------------

    go get github.com/tideland/gots/v3/asserts
    go get github.com/tideland/gots/v3/generators

Usage
-----

**Asserts**

A typical assertion instance for testing is created with

    assert := asserts.NewTestingAssertion(t, true)

The true assures that a failing assertion lets also fail the Go test. Now
there are several methods like

- `assert.True()`
- `assert.Equal()`
- `assert.Length()`
- `assert.Implementor()`
- `assert.Empty()`
- ...

Beside the test they can take an optional last string argument that's
been printed in the report if the assert fails.

**Generators**

The generator helps to generate test data based on a random number generator. 
This can be ints, words, sentences, paragraphs, male and female names, domains, 
URLs, e-mail addresses, durations and times. The SimpleRand() returns a Rand 
based on the current time while FixedRand() returns always the same. Especially 
the latter helps to produce the same data in different runs:

    assert := asserts.NewTestingAssertion(t, false)
    genA := generators.New(generators.FixedRand())
    genB := generators.New(generators.FixedRand())

    urlA := genA.URL()
    urlB := genB.URL()
    assert.Equal(urlA, urlB)

And now have fun. ;)

Documentation
-------------

- http://godoc.org/github.com/tideland/gots/v3/asserts
- http://godoc.org/github.com/tideland/gots/v3/generators

Authors
------

- Frank Mueller - <mue@tideland.biz>

License
-------

*Tideland Go Testing Support* is distributed under the terms of the BSD 3-Clause license.
