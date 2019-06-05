package assert

import (
	"errors"
	"regexp"
	"testing"
	"time"
)

func TestImplementsWrapper(t *testing.T) {
	assert := New(new(testing.T))

	if !Implements((*AssertionTesterInterface)(nil), new(AssertionTesterConformingObject)) {
		t.Error("Implements method should return true: AssertionTesterConformingObject implements AssertionTesterInterface")
	}
	if Implements((*AssertionTesterInterface)(nil), new(AssertionTesterNonConformingObject)) {
		t.Error("Implements method should return false: AssertionTesterNonConformingObject does not implements AssertionTesterInterface")
	}
}

func TestIsTypeWrapper(t *testing.T) {
	assert := New(new(testing.T))

	if !IsType(new(AssertionTesterConformingObject), new(AssertionTesterConformingObject)) {
		t.Error("IsType should return true: AssertionTesterConformingObject is the same type as AssertionTesterConformingObject")
	}
	if IsType(new(AssertionTesterConformingObject), new(AssertionTesterNonConformingObject)) {
		t.Error("IsType should return false: AssertionTesterConformingObject is not the same type as AssertionTesterNonConformingObject")
	}

}

func TestEqualWrapper(t *testing.T) {
	assert := New(new(testing.T))

	if !Equal("Hello World", "Hello World") {
		t.Error("Equal should return true")
	}
	if !Equal(123, 123) {
		t.Error("Equal should return true")
	}
	if !Equal(123.5, 123.5) {
		t.Error("Equal should return true")
	}
	if !Equal([]byte("Hello World"), []byte("Hello World")) {
		t.Error("Equal should return true")
	}
	if !Equal(nil, nil) {
		t.Error("Equal should return true")
	}
}

func TestEqualValuesWrapper(t *testing.T) {
	assert := New(new(testing.T))

	if !EqualValues(uint32(10), int32(10)) {
		t.Error("EqualValues should return true")
	}
}

func TestNotNilWrapper(t *testing.T) {
	assert := New(new(testing.T))

	if !NotNil(new(AssertionTesterConformingObject)) {
		t.Error("NotNil should return true: object is not nil")
	}
	if NotNil(nil) {
		t.Error("NotNil should return false: object is nil")
	}

}

func TestNilWrapper(t *testing.T) {
	assert := New(new(testing.T))

	if !Nil(nil) {
		t.Error("Nil should return true: object is nil")
	}
	if Nil(new(AssertionTesterConformingObject)) {
		t.Error("Nil should return false: object is not nil")
	}

}

func TestTrueWrapper(t *testing.T) {
	assert := New(new(testing.T))

	if !True(true) {
		t.Error("True should return true")
	}
	if True(false) {
		t.Error("True should return false")
	}

}

func TestFalseWrapper(t *testing.T) {
	assert := New(new(testing.T))

	if !False(false) {
		t.Error("False should return true")
	}
	if False(true) {
		t.Error("False should return false")
	}

}

func TestExactlyWrapper(t *testing.T) {
	assert := New(new(testing.T))

	a := float32(1)
	b := float64(1)
	c := float32(1)
	d := float32(2)

	if Exactly(a, b) {
		t.Error("Exactly should return false")
	}
	if Exactly(a, d) {
		t.Error("Exactly should return false")
	}
	if !Exactly(a, c) {
		t.Error("Exactly should return true")
	}

	if Exactly(nil, a) {
		t.Error("Exactly should return false")
	}
	if Exactly(a, nil) {
		t.Error("Exactly should return false")
	}

}

func TestNotEqualWrapper(t *testing.T) {

	assert := New(new(testing.T))

	if !NotEqual("Hello World", "Hello World!") {
		t.Error("NotEqual should return true")
	}
	if !NotEqual(123, 1234) {
		t.Error("NotEqual should return true")
	}
	if !NotEqual(123.5, 123.55) {
		t.Error("NotEqual should return true")
	}
	if !NotEqual([]byte("Hello World"), []byte("Hello World!")) {
		t.Error("NotEqual should return true")
	}
	if !NotEqual(nil, new(AssertionTesterConformingObject)) {
		t.Error("NotEqual should return true")
	}
}

func TestContainsWrapper(t *testing.T) {

	assert := New(new(testing.T))
	list := []string{"Foo", "Bar"}

	if !Contains("Hello World", "Hello") {
		t.Error("Contains should return true: \"Hello World\" contains \"Hello\"")
	}
	if Contains("Hello World", "Salut") {
		t.Error("Contains should return false: \"Hello World\" does not contain \"Salut\"")
	}

	if !Contains(list, "Foo") {
		t.Error("Contains should return true: \"[\"Foo\", \"Bar\"]\" contains \"Foo\"")
	}
	if Contains(list, "Salut") {
		t.Error("Contains should return false: \"[\"Foo\", \"Bar\"]\" does not contain \"Salut\"")
	}

}

func TestNotContainsWrapper(t *testing.T) {

	assert := New(new(testing.T))
	list := []string{"Foo", "Bar"}

	if !NotContains("Hello World", "Hello!") {
		t.Error("NotContains should return true: \"Hello World\" does not contain \"Hello!\"")
	}
	if NotContains("Hello World", "Hello") {
		t.Error("NotContains should return false: \"Hello World\" contains \"Hello\"")
	}

	if !NotContains(list, "Foo!") {
		t.Error("NotContains should return true: \"[\"Foo\", \"Bar\"]\" does not contain \"Foo!\"")
	}
	if NotContains(list, "Foo") {
		t.Error("NotContains should return false: \"[\"Foo\", \"Bar\"]\" contains \"Foo\"")
	}

}

func TestConditionWrapper(t *testing.T) {

	assert := New(new(testing.T))

	if !Condition(func() bool { return true }, "Truth") {
		t.Error("Condition should return true")
	}

	if Condition(func() bool { return false }, "Lie") {
		t.Error("Condition should return false")
	}

}

func TestDidPanicWrapper(t *testing.T) {

	if funcDidPanic, _ := didPanic(func() {
		panic("Panic!")
	}); !funcDidPanic {
		t.Error("didPanic should return true")
	}

	if funcDidPanic, _ := didPanic(func() {
	}); funcDidPanic {
		t.Error("didPanic should return false")
	}

}

func TestPanicsWrapper(t *testing.T) {

	assert := New(new(testing.T))

	if !Panics(func() {
		panic("Panic!")
	}) {
		t.Error("Panics should return true")
	}

	if Panics(func() {
	}) {
		t.Error("Panics should return false")
	}

}

func TestNotPanicsWrapper(t *testing.T) {

	assert := New(new(testing.T))

	if !NotPanics(func() {
	}) {
		t.Error("NotPanics should return true")
	}

	if NotPanics(func() {
		panic("Panic!")
	}) {
		t.Error("NotPanics should return false")
	}

}

func TestNoErrorWrapper(t *testing.T) {
	assert := New(t)
	mockAssert := New(new(testing.T))

	// start with a nil error
	var err error

	True(NoError(err), "NoError should return True for nil arg")

	// now set an error
	err = errors.New("Some error")

	False(NoError(err), "NoError with error should return False")

}

func TestErrorWrapper(t *testing.T) {
	assert := New(t)
	mockAssert := New(new(testing.T))

	// start with a nil error
	var err error

	False(Error(err), "Error should return False for nil arg")

	// now set an error
	err = errors.New("Some error")

	True(Error(err), "Error with error should return True")

}

func TestEqualErrorWrapper(t *testing.T) {
	assert := New(t)
	mockAssert := New(new(testing.T))

	// start with a nil error
	var err error
	False(EqualError(err, ""),
		"EqualError should return false for nil arg")

	// now set an error
	err = errors.New("some error")
	False(EqualError(err, "Not some error"),
		"EqualError should return false for different error string")
	True(EqualError(err, "some error"),
		"EqualError should return true")
}

func TestEmptyWrapper(t *testing.T) {
	assert := New(t)
	mockAssert := New(new(testing.T))

	True(Empty(""), "Empty string is empty")
	True(Empty(nil), "Nil is empty")
	True(Empty([]string{}), "Empty string array is empty")
	True(Empty(0), "Zero int value is empty")
	True(Empty(false), "False value is empty")

	False(Empty("something"), "Non Empty string is not empty")
	False(Empty(errors.New("something")), "Non nil object is not empty")
	False(Empty([]string{"something"}), "Non empty string array is not empty")
	False(Empty(1), "Non-zero int value is not empty")
	False(Empty(true), "True value is not empty")

}

func TestNotEmptyWrapper(t *testing.T) {
	assert := New(t)
	mockAssert := New(new(testing.T))

	False(NotEmpty(""), "Empty string is empty")
	False(NotEmpty(nil), "Nil is empty")
	False(NotEmpty([]string{}), "Empty string array is empty")
	False(NotEmpty(0), "Zero int value is empty")
	False(NotEmpty(false), "False value is empty")

	True(NotEmpty("something"), "Non Empty string is not empty")
	True(NotEmpty(errors.New("something")), "Non nil object is not empty")
	True(NotEmpty([]string{"something"}), "Non empty string array is not empty")
	True(NotEmpty(1), "Non-zero int value is not empty")
	True(NotEmpty(true), "True value is not empty")

}

func TestLenWrapper(t *testing.T) {
	assert := New(t)
	mockAssert := New(new(testing.T))

	False(Len(nil, 0), "nil does not have length")
	False(Len(0, 0), "int does not have length")
	False(Len(true, 0), "true does not have length")
	False(Len(false, 0), "false does not have length")
	False(Len('A', 0), "Rune does not have length")
	False(Len(struct{}{}, 0), "Struct does not have length")

	ch := make(chan int, 5)
	ch <- 1
	ch <- 2
	ch <- 3

	cases := []struct {
		v interface{}
		l int
	}{
		{[]int{1, 2, 3}, 3},
		{[...]int{1, 2, 3}, 3},
		{"ABC", 3},
		{map[int]int{1: 2, 2: 4, 3: 6}, 3},
		{ch, 3},

		{[]int{}, 0},
		{map[int]int{}, 0},
		{make(chan int), 0},

		{[]int(nil), 0},
		{map[int]int(nil), 0},
		{(chan int)(nil), 0},
	}

	for _, c := range cases {
		True(Len(c.v, c.l), "%#v have %d items", c.v, c.l)
	}
}

func TestWithinDurationWrapper(t *testing.T) {
	assert := New(t)
	mockAssert := New(new(testing.T))
	a := time.Now()
	b := a.Add(10 * time.Second)

	True(WithinDuration(a, b, 10*time.Second), "A 10s difference is within a 10s time difference")
	True(WithinDuration(b, a, 10*time.Second), "A 10s difference is within a 10s time difference")

	False(WithinDuration(a, b, 9*time.Second), "A 10s difference is not within a 9s time difference")
	False(WithinDuration(b, a, 9*time.Second), "A 10s difference is not within a 9s time difference")

	False(WithinDuration(a, b, -9*time.Second), "A 10s difference is not within a 9s time difference")
	False(WithinDuration(b, a, -9*time.Second), "A 10s difference is not within a 9s time difference")

	False(WithinDuration(a, b, -11*time.Second), "A 10s difference is not within a 9s time difference")
	False(WithinDuration(b, a, -11*time.Second), "A 10s difference is not within a 9s time difference")
}

func TestInDeltaWrapper(t *testing.T) {
	assert := New(new(testing.T))

	True(t, InDelta(1.001, 1, 0.01), "|1.001 - 1| <= 0.01")
	True(t, InDelta(1, 1.001, 0.01), "|1 - 1.001| <= 0.01")
	True(t, InDelta(1, 2, 1), "|1 - 2| <= 1")
	False(t, InDelta(1, 2, 0.5), "Expected |1 - 2| <= 0.5 to fail")
	False(t, InDelta(2, 1, 0.5), "Expected |2 - 1| <= 0.5 to fail")
	False(t, InDelta("", nil, 1), "Expected non numerals to fail")

	cases := []struct {
		a, b  interface{}
		delta float64
	}{
		{uint8(2), uint8(1), 1},
		{uint16(2), uint16(1), 1},
		{uint32(2), uint32(1), 1},
		{uint64(2), uint64(1), 1},

		{int(2), int(1), 1},
		{int8(2), int8(1), 1},
		{int16(2), int16(1), 1},
		{int32(2), int32(1), 1},
		{int64(2), int64(1), 1},

		{float32(2), float32(1), 1},
		{float64(2), float64(1), 1},
	}

	for _, tc := range cases {
		True(t, InDelta(tc.a, tc.b, tc.delta), "Expected |%V - %V| <= %v", tc.a, tc.b, tc.delta)
	}
}

func TestInEpsilonWrapper(t *testing.T) {
	assert := New(new(testing.T))

	cases := []struct {
		a, b    interface{}
		epsilon float64
	}{
		{uint8(2), uint16(2), .001},
		{2.1, 2.2, 0.1},
		{2.2, 2.1, 0.1},
		{-2.1, -2.2, 0.1},
		{-2.2, -2.1, 0.1},
		{uint64(100), uint8(101), 0.01},
		{0.1, -0.1, 2},
	}

	for _, tc := range cases {
		True(t, InEpsilon(tc.a, tc.b, tc.epsilon, "Expected %V and %V to have a relative difference of %v", tc.a, tc.b, tc.epsilon))
	}

	cases = []struct {
		a, b    interface{}
		epsilon float64
	}{
		{uint8(2), int16(-2), .001},
		{uint64(100), uint8(102), 0.01},
		{2.1, 2.2, 0.001},
		{2.2, 2.1, 0.001},
		{2.1, -2.2, 1},
		{2.1, "bla-bla", 0},
		{0.1, -0.1, 1.99},
	}

	for _, tc := range cases {
		False(t, InEpsilon(tc.a, tc.b, tc.epsilon, "Expected %V and %V to have a relative difference of %v", tc.a, tc.b, tc.epsilon))
	}
}

func TestRegexpWrapper(t *testing.T) {

	assert := New(new(testing.T))

	cases := []struct {
		rx, str string
	}{
		{"^start", "start of the line"},
		{"end$", "in the end"},
		{"[0-9]{3}[.-]?[0-9]{2}[.-]?[0-9]{2}", "My phone number is 650.12.34"},
	}

	for _, tc := range cases {
		True(t, Regexp(tc.rx, tc.str))
		True(t, Regexp(regexp.MustCompile(tc.rx), tc.str))
		False(t, NotRegexp(tc.rx, tc.str))
		False(t, NotRegexp(regexp.MustCompile(tc.rx), tc.str))
	}

	cases = []struct {
		rx, str string
	}{
		{"^asdfastart", "Not the start of the line"},
		{"end$", "in the end."},
		{"[0-9]{3}[.-]?[0-9]{2}[.-]?[0-9]{2}", "My phone number is 650.12a.34"},
	}

	for _, tc := range cases {
		False(t, Regexp(tc.rx, tc.str), "Expected \"%s\" to not match \"%s\"", tc.rx, tc.str)
		False(t, Regexp(regexp.MustCompile(tc.rx), tc.str))
		True(t, NotRegexp(tc.rx, tc.str))
		True(t, NotRegexp(regexp.MustCompile(tc.rx), tc.str))
	}
}

func TestZeroWrapper(t *testing.T) {
	assert := New(t)
	mockAssert := New(new(testing.T))

	for _, test := range zeros {
		True(Zero(test), "Zero should return true for %v", test)
	}

	for _, test := range nonZeros {
		False(Zero(test), "Zero should return false for %v", test)
	}
}

func TestNotZeroWrapper(t *testing.T) {
	assert := New(t)
	mockAssert := New(new(testing.T))

	for _, test := range zeros {
		False(NotZero(test), "Zero should return true for %v", test)
	}

	for _, test := range nonZeros {
		True(NotZero(test), "Zero should return false for %v", test)
	}
}

func TestJSONEqWrapper_EqualSONString(t *testing.T) {
	assert := New(new(testing.T))
	if !JSONEq(`{"hello": "world", "foo": "bar"}`, `{"hello": "world", "foo": "bar"}`) {
		t.Error("JSONEq should return true")
	}

}

func TestJSONEqWrapper_EquivalentButNotEqual(t *testing.T) {
	assert := New(new(testing.T))
	if !JSONEq(`{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`) {
		t.Error("JSONEq should return true")
	}

}

func TestJSONEqWrapper_HashOfArraysAndHashes(t *testing.T) {
	assert := New(new(testing.T))
	if !JSONEq("{\r\n\t\"numeric\": 1.5,\r\n\t\"array\": [{\"foo\": \"bar\"}, 1, \"string\", [\"nested\", \"array\", 5.5]],\r\n\t\"hash\": {\"nested\": \"hash\", \"nested_slice\": [\"this\", \"is\", \"nested\"]},\r\n\t\"string\": \"foo\"\r\n}",
		"{\r\n\t\"numeric\": 1.5,\r\n\t\"hash\": {\"nested\": \"hash\", \"nested_slice\": [\"this\", \"is\", \"nested\"]},\r\n\t\"string\": \"foo\",\r\n\t\"array\": [{\"foo\": \"bar\"}, 1, \"string\", [\"nested\", \"array\", 5.5]]\r\n}") {
		t.Error("JSONEq should return true")
	}
}

func TestJSONEqWrapper_Array(t *testing.T) {
	assert := New(new(testing.T))
	if !JSONEq(`["foo", {"hello": "world", "nested": "hash"}]`, `["foo", {"nested": "hash", "hello": "world"}]`) {
		t.Error("JSONEq should return true")
	}

}

func TestJSONEqWrapper_HashAndArrayNotEquivalent(t *testing.T) {
	assert := New(new(testing.T))
	if JSONEq(`["foo", {"hello": "world", "nested": "hash"}]`, `{"foo": "bar", {"nested": "hash", "hello": "world"}}`) {
		t.Error("JSONEq should return false")
	}
}

func TestJSONEqWrapper_HashesNotEquivalent(t *testing.T) {
	assert := New(new(testing.T))
	if JSONEq(`{"foo": "bar"}`, `{"foo": "bar", "hello": "world"}`) {
		t.Error("JSONEq should return false")
	}
}

func TestJSONEqWrapper_ActualIsNotJSON(t *testing.T) {
	assert := New(new(testing.T))
	if JSONEq(`{"foo": "bar"}`, "Not JSON") {
		t.Error("JSONEq should return false")
	}
}

func TestJSONEqWrapper_ExpectedIsNotJSON(t *testing.T) {
	assert := New(new(testing.T))
	if JSONEq("Not JSON", `{"foo": "bar", "hello": "world"}`) {
		t.Error("JSONEq should return false")
	}
}

func TestJSONEqWrapper_ExpectedAndActualNotJSON(t *testing.T) {
	assert := New(new(testing.T))
	if JSONEq("Not JSON", "Not JSON") {
		t.Error("JSONEq should return false")
	}
}

func TestJSONEqWrapper_ArraysOfDifferentOrder(t *testing.T) {
	assert := New(new(testing.T))
	if JSONEq(`["foo", {"hello": "world", "nested": "hash"}]`, `[{ "hello": "world", "nested": "hash"}, "foo"]`) {
		t.Error("JSONEq should return false")
	}
}
