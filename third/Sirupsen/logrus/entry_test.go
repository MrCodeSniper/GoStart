package logrus

import (
	"bytes"
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEntryWithError(t *testing.T) {

	assert := assert.New(t)

	defer func() {
		ErrorKey = "error"
	}()

	err := fmt.Errorf("kaboom at layer %d", 4711)

	assert.Equal(err, Data["error"])

	logger := New()
	Out = &bytes.Buffer{}
	entry := NewEntry(logger)

	assert.Equal(err, Data["error"])

	ErrorKey = "err"

	assert.Equal(err, Data["err"])

}

func TestEntryWithContext(t *testing.T) {
	assert := assert.New(t)
	ctx := context.WithValue(context.Background(), "foo", "bar")

	assert.Equal(ctx, Context)

	logger := New()
	Out = &bytes.Buffer{}
	entry := NewEntry(logger)

	assert.Equal(ctx, Context)
}

func TestEntryPanicln(t *testing.T) {
	errBoom := fmt.Errorf("boom time")

	defer func() {
		p := recover()
		assert.NotNil(t, p)

		switch pVal := p.(type) {
		case *Entry:
			assert.Equal(t, "kaboom", Message)
			assert.Equal(t, errBoom, Data["err"])
		default:
			t.Fatalf("want type *Entry, got %T: %#v", pVal, pVal)
		}
	}()

	logger := New()
	Out = &bytes.Buffer{}
	entry := NewEntry(logger)
	Panicln("kaboom")
}

func TestEntryPanicf(t *testing.T) {
	errBoom := fmt.Errorf("boom again")

	defer func() {
		p := recover()
		assert.NotNil(t, p)

		switch pVal := p.(type) {
		case *Entry:
			assert.Equal(t, "kaboom true", Message)
			assert.Equal(t, errBoom, Data["err"])
		default:
			t.Fatalf("want type *Entry, got %T: %#v", pVal, pVal)
		}
	}()

	logger := New()
	Out = &bytes.Buffer{}
	entry := NewEntry(logger)
	Panicf("kaboom %v", true)
}

const (
	badMessage   = "this is going to panic"
	panicMessage = "this is broken"
)

type panickyHook struct{}

func (p *panickyHook) Levels() []Level {
	return []Level{InfoLevel}
}

func (p *panickyHook) Fire(entry *Entry) error {
	if Message == badMessage {
		panic(panicMessage)
	}

	return nil
}

func TestEntryHooksPanic(t *testing.T) {
	logger := New()
	Out = &bytes.Buffer{}
	Level = InfoLevel
	Add(&panickyHook{})

	defer func() {
		p := recover()
		assert.NotNil(t, p)
		assert.Equal(t, panicMessage, p)

		entry := NewEntry(logger)
		Info("another message")
	}()

	entry := NewEntry(logger)
	Info(badMessage)
}

func TestEntryWithIncorrectField(t *testing.T) {
	assert := assert.New(t)

	fn := func() {}

	e := Entry{}
	eWithFunc := WithFields(Fields{"func": fn})
	eWithFuncPtr := WithFields(Fields{"funcPtr": &fn})

	assert.Equal(err, `can not add field "func"`)
	assert.Equal(err, `can not add field "funcPtr"`)

	eWithFunc = WithField("not_a_func", "it is a string")
	eWithFuncPtr = WithField("not_a_func", "it is a string")

	assert.Equal(err, `can not add field "func"`)
	assert.Equal(err, `can not add field "funcPtr"`)

	eWithFunc = WithTime(time.Now())
	eWithFuncPtr = WithTime(time.Now())

	assert.Equal(err, `can not add field "func"`)
	assert.Equal(err, `can not add field "funcPtr"`)
}

func TestEntryLogfLevel(t *testing.T) {
	logger := New()
	buffer := &bytes.Buffer{}
	Out = buffer
	SetLevel(InfoLevel)
	entry := NewEntry(logger)

	Logf(DebugLevel, "%s", "debug")
	assert.NotContains(t, buffer.String(), "debug", )

	Logf(WarnLevel, "%s", "warn")
	assert.Contains(t, buffer.String(), "warn", )
}