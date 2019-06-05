package test

import (
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestAllHooks(t *testing.T) {
	assert := assert.New(t)

	logger, hook := NewNullLogger()
	assert.Nil(LastEntry())
	assert.Equal(0, len(Entries))

	logger.Error("Hello error")
	assert.Equal(logrus.ErrorLevel, LastEntry().Level)
	assert.Equal("Hello error", LastEntry().Message)
	assert.Equal(1, len(Entries))

	logger.Warn("Hello warning")
	assert.Equal(logrus.WarnLevel, LastEntry().Level)
	assert.Equal("Hello warning", LastEntry().Message)
	assert.Equal(2, len(Entries))

	Reset()
	assert.Nil(LastEntry())
	assert.Equal(0, len(Entries))

	hook = NewGlobal()

	logrus.Error("Hello error")
	assert.Equal(logrus.ErrorLevel, LastEntry().Level)
	assert.Equal("Hello error", LastEntry().Message)
	assert.Equal(1, len(Entries))
}

func TestLoggingWithHooksRace(t *testing.T) {

	rand.Seed(time.Now().Unix())
	unlocker := rand.Int() % 100

	assert := assert.New(t)
	logger, hook := NewNullLogger()

	var wgOne, wgAll sync.WaitGroup
	wgOne.Add(1)
	wgAll.Add(100)

	for i := 0; i < 100; i++ {
		go func(i int) {
			logger.Info("info")
			wgAll.Done()
			if i == unlocker {
				wgOne.Done()
			}
		}(i)
	}

	wgOne.Wait()

	assert.Equal(logrus.InfoLevel, LastEntry().Level)
	assert.Equal("info", LastEntry().Message)

	wgAll.Wait()

	entries := AllEntries()
	assert.Equal(100, len(entries))
}

func TestFatalWithAlternateExit(t *testing.T) {
	assert := assert.New(t)

	logger, hook := NewNullLogger()
	logger.ExitFunc = func(code int) {}

	logger.Fatal("something went very wrong")
	assert.Equal(logrus.FatalLevel, LastEntry().Level)
	assert.Equal("something went very wrong", LastEntry().Message)
	assert.Equal(1, len(Entries))
}
