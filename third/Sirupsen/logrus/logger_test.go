package logrus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFieldValueError(t *testing.T) {
	buf := &bytes.Buffer{}
	l := &Logger{
		Out:       buf,
		Formatter: new(JSONFormatter),
		Hooks:     make(LevelHooks),
		Level:     DebugLevel,
	}
	Info("test")
	fmt.Println(buf.String())
	var data map[string]interface{}
	json.Unmarshal(buf.Bytes(), &data)
	_, ok := data[FieldKeyLogrusError]
	require.True(t, ok)
}

func TestNoFieldValueError(t *testing.T) {
	buf := &bytes.Buffer{}
	l := &Logger{
		Out:       buf,
		Formatter: new(JSONFormatter),
		Hooks:     make(LevelHooks),
		Level:     DebugLevel,
	}
	Info("test")
	fmt.Println(buf.String())
	var data map[string]interface{}
	json.Unmarshal(buf.Bytes(), &data)
	_, ok := data[FieldKeyLogrusError]
	require.False(t, ok)
}

func TestWarninglnNotEqualToWarning(t *testing.T) {
	buf := &bytes.Buffer{}
	bufln := &bytes.Buffer{}

	formatter := new(TextFormatter)
	DisableTimestamp = true
	DisableLevelTruncation = true

	l := &Logger{
		Out:       buf,
		Formatter: formatter,
		Hooks:     make(LevelHooks),
		Level:     DebugLevel,
	}
	Warning("hello,", "world")

	SetOutput(bufln)
	Warningln("hello,", "world")

	assert.NotEqual(t, buf.String(), bufln.String(), "Warning() and Wantingln() should not be equal")
}
