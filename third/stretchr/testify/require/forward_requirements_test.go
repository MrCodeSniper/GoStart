package require

import (
	"errors"
	"testing"
	"time"
)

func TestImplementsWrapper(t *testing.T) {
	require := New(t)

	Implements((*AssertionTesterInterface)(nil), new(AssertionTesterConformingObject))

	mockT := new(MockT)
	mockRequire := New(mockT)
	Implements((*AssertionTesterInterface)(nil), new(AssertionTesterNonConformingObject))
	if !Failed {
		t.Error("Check should fail")
	}
}

func TestIsTypeWrapper(t *testing.T) {
	require := New(t)
	IsType(new(AssertionTesterConformingObject), new(AssertionTesterConformingObject))

	mockT := new(MockT)
	mockRequire := New(mockT)
	IsType(new(AssertionTesterConformingObject), new(AssertionTesterNonConformingObject))
	if !Failed {
		t.Error("Check should fail")
	}
}

func TestEqualWrapper(t *testing.T) {
	require := New(t)
	Equal(1, 1)

	mockT := new(MockT)
	mockRequire := New(mockT)
	Equal(1, 2)
	if !Failed {
		t.Error("Check should fail")
	}
}

func TestNotEqualWrapper(t *testing.T) {
	require := New(t)
	NotEqual(1, 2)

	mockT := new(MockT)
	mockRequire := New(mockT)
	NotEqual(2, 2)
	if !Failed {
		t.Error("Check should fail")
	}
}

func TestExactlyWrapper(t *testing.T) {
	require := New(t)

	a := float32(1)
	b := float32(1)
	c := float64(1)

	Exactly(a, b)

	mockT := new(MockT)
	mockRequire := New(mockT)
	Exactly(a, c)
	if !Failed {
		t.Error("Check should fail")
	}
}

func TestNotNilWrapper(t *testing.T) {
	require := New(t)
	NotNil(t, new(AssertionTesterConformingObject))

	mockT := new(MockT)
	mockRequire := New(mockT)
	NotNil(nil)
	if !Failed {
		t.Error("Check should fail")
	}
}

func TestNilWrapper(t *testing.T) {
	require := New(t)
	Nil(nil)

	mockT := new(MockT)
	mockRequire := New(mockT)
	Nil(new(AssertionTesterConformingObject))
	if !Failed {
		t.Error("Check should fail")
	}
}

func TestTrueWrapper(t *testing.T) {
	require := New(t)
	True(true)

	mockT := new(MockT)
	mockRequire := New(mockT)
	True(false)
	if !Failed {
		t.Error("Check should fail")
	}
}

func TestFalseWrapper(t *testing.T) {
	require := New(t)
	False(false)

	mockT := new(MockT)
	mockRequire := New(mockT)
	False(true)
	if !Failed {
		t.Error("Check should fail")
	}
}

func TestContainsWrapper(t *testing.T) {
	require := New(t)
	Contains("Hello World", "Hello")

	mockT := new(MockT)
	mockRequire := New(mockT)
	Contains("Hello World", "Salut")
	if !Failed {
		t.Error("Check should fail")
	}
}

func TestNotContainsWrapper(t *testing.T) {
	require := New(t)
	NotContains("Hello World", "Hello!")

	mockT := new(MockT)
	mockRequire := New(mockT)
	NotContains("Hello World", "Hello")
	if !Failed {
		t.Error("Check should fail")
	}
}

func TestPanicsWrapper(t *testing.T) {
	require := New(t)
	Panics(func() {
		panic("Panic!")
	})

	mockT := new(MockT)
	mockRequire := New(mockT)
	Panics(func() {})
	if !Failed {
		t.Error("Check should fail")
	}
}

func TestNotPanicsWrapper(t *testing.T) {
	require := New(t)
	NotPanics(func() {})

	mockT := new(MockT)
	mockRequire := New(mockT)
	NotPanics(func() {
		panic("Panic!")
	})
	if !Failed {
		t.Error("Check should fail")
	}
}

func TestNoErrorWrapper(t *testing.T) {
	require := New(t)
	NoError(nil)

	mockT := new(MockT)
	mockRequire := New(mockT)
	NoError(errors.New("some error"))
	if !Failed {
		t.Error("Check should fail")
	}
}

func TestErrorWrapper(t *testing.T) {
	require := New(t)
	Error(errors.New("some error"))

	mockT := new(MockT)
	mockRequire := New(mockT)
	Error(nil)
	if !Failed {
		t.Error("Check should fail")
	}
}

func TestEqualErrorWrapper(t *testing.T) {
	require := New(t)
	EqualError(errors.New("some error"), "some error")

	mockT := new(MockT)
	mockRequire := New(mockT)
	EqualError(errors.New("some error"), "Not some error")
	if !Failed {
		t.Error("Check should fail")
	}
}

func TestEmptyWrapper(t *testing.T) {
	require := New(t)
	Empty("")

	mockT := new(MockT)
	mockRequire := New(mockT)
	Empty("x")
	if !Failed {
		t.Error("Check should fail")
	}
}

func TestNotEmptyWrapper(t *testing.T) {
	require := New(t)
	NotEmpty("x")

	mockT := new(MockT)
	mockRequire := New(mockT)
	NotEmpty("")
	if !Failed {
		t.Error("Check should fail")
	}
}

func TestWithinDurationWrapper(t *testing.T) {
	require := New(t)
	a := time.Now()
	b := a.Add(10 * time.Second)

	WithinDuration(a, b, 15*time.Second)

	mockT := new(MockT)
	mockRequire := New(mockT)
	WithinDuration(a, b, 5*time.Second)
	if !Failed {
		t.Error("Check should fail")
	}
}

func TestInDeltaWrapper(t *testing.T) {
	require := New(t)
	InDelta(1.001, 1, 0.01)

	mockT := new(MockT)
	mockRequire := New(mockT)
	InDelta(1, 2, 0.5)
	if !Failed {
		t.Error("Check should fail")
	}
}

func TestZeroWrapper(t *testing.T) {
	require := New(t)
	Zero(0)

	mockT := new(MockT)
	mockRequire := New(mockT)
	Zero(1)
	if !Failed {
		t.Error("Check should fail")
	}
}

func TestNotZeroWrapper(t *testing.T) {
	require := New(t)
	NotZero(1)

	mockT := new(MockT)
	mockRequire := New(mockT)
	NotZero(0)
	if !Failed {
		t.Error("Check should fail")
	}
}

func TestJSONEqWrapper_EqualSONString(t *testing.T) {
	mockT := new(MockT)
	mockRequire := New(mockT)

	JSONEq(`{"hello": "world", "foo": "bar"}`, `{"hello": "world", "foo": "bar"}`)
	if Failed {
		t.Error("Check should pass")
	}
}

func TestJSONEqWrapper_EquivalentButNotEqual(t *testing.T) {
	mockT := new(MockT)
	mockRequire := New(mockT)

	JSONEq(`{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
	if Failed {
		t.Error("Check should pass")
	}
}

func TestJSONEqWrapper_HashOfArraysAndHashes(t *testing.T) {
	mockT := new(MockT)
	mockRequire := New(mockT)

	JSONEq("{\r\n\t\"numeric\": 1.5,\r\n\t\"array\": [{\"foo\": \"bar\"}, 1, \"string\", [\"nested\", \"array\", 5.5]],\r\n\t\"hash\": {\"nested\": \"hash\", \"nested_slice\": [\"this\", \"is\", \"nested\"]},\r\n\t\"string\": \"foo\"\r\n}",
		"{\r\n\t\"numeric\": 1.5,\r\n\t\"hash\": {\"nested\": \"hash\", \"nested_slice\": [\"this\", \"is\", \"nested\"]},\r\n\t\"string\": \"foo\",\r\n\t\"array\": [{\"foo\": \"bar\"}, 1, \"string\", [\"nested\", \"array\", 5.5]]\r\n}")
	if Failed {
		t.Error("Check should pass")
	}
}

func TestJSONEqWrapper_Array(t *testing.T) {
	mockT := new(MockT)
	mockRequire := New(mockT)

	JSONEq(`["foo", {"hello": "world", "nested": "hash"}]`, `["foo", {"nested": "hash", "hello": "world"}]`)
	if Failed {
		t.Error("Check should pass")
	}
}

func TestJSONEqWrapper_HashAndArrayNotEquivalent(t *testing.T) {
	mockT := new(MockT)
	mockRequire := New(mockT)

	JSONEq(`["foo", {"hello": "world", "nested": "hash"}]`, `{"foo": "bar", {"nested": "hash", "hello": "world"}}`)
	if !Failed {
		t.Error("Check should fail")
	}
}

func TestJSONEqWrapper_HashesNotEquivalent(t *testing.T) {
	mockT := new(MockT)
	mockRequire := New(mockT)

	JSONEq(`{"foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
	if !Failed {
		t.Error("Check should fail")
	}
}

func TestJSONEqWrapper_ActualIsNotJSON(t *testing.T) {
	mockT := new(MockT)
	mockRequire := New(mockT)

	JSONEq(`{"foo": "bar"}`, "Not JSON")
	if !Failed {
		t.Error("Check should fail")
	}
}

func TestJSONEqWrapper_ExpectedIsNotJSON(t *testing.T) {
	mockT := new(MockT)
	mockRequire := New(mockT)

	JSONEq("Not JSON", `{"foo": "bar", "hello": "world"}`)
	if !Failed {
		t.Error("Check should fail")
	}
}

func TestJSONEqWrapper_ExpectedAndActualNotJSON(t *testing.T) {
	mockT := new(MockT)
	mockRequire := New(mockT)

	JSONEq("Not JSON", "Not JSON")
	if !Failed {
		t.Error("Check should fail")
	}
}

func TestJSONEqWrapper_ArraysOfDifferentOrder(t *testing.T) {
	mockT := new(MockT)
	mockRequire := New(mockT)

	JSONEq(`["foo", {"hello": "world", "nested": "hash"}]`, `[{ "hello": "world", "nested": "hash"}, "foo"]`)
	if !Failed {
		t.Error("Check should fail")
	}
}
