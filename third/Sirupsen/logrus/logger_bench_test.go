package logrus

import (
	"io/ioutil"
	"os"
	"testing"
)

// smallFields is a small size data set for benchmarking
var loggerFields = Fields{
	"foo":   "bar",
	"baz":   "qux",
	"one":   "two",
	"three": "four",
}

func BenchmarkDummyLogger(b *testing.B) {
	nullf, err := os.OpenFile("/dev/null", os.O_WRONLY, 0666)
	if err != nil {
		b.Fatalf("%v", err)
	}
	defer nullf.Close()
	doLoggerBenchmark(b, nullf, &TextFormatter{DisableColors: true}, smallFields)
}

func BenchmarkDummyLoggerNoLock(b *testing.B) {
	nullf, err := os.OpenFile("/dev/null", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		b.Fatalf("%v", err)
	}
	defer nullf.Close()
	doLoggerBenchmarkNoLock(b, nullf, &TextFormatter{DisableColors: true}, smallFields)
}

func doLoggerBenchmark(b *testing.B, out *os.File, formatter Formatter, fields Fields) {
	logger := Logger{
		Out:       out,
		Level:     InfoLevel,
		Formatter: formatter,
	}
	entry := WithFields(fields)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Info("aaa")
		}
	})
}

func doLoggerBenchmarkNoLock(b *testing.B, out *os.File, formatter Formatter, fields Fields) {
	logger := Logger{
		Out:       out,
		Level:     InfoLevel,
		Formatter: formatter,
	}
	SetNoLock()
	entry := WithFields(fields)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Info("aaa")
		}
	})
}

func BenchmarkLoggerJSONFormatter(b *testing.B) {
	doLoggerBenchmarkWithFormatter(b, &JSONFormatter{})
}

func BenchmarkLoggerTextFormatter(b *testing.B) {
	doLoggerBenchmarkWithFormatter(b, &TextFormatter{})
}

func doLoggerBenchmarkWithFormatter(b *testing.B, f Formatter) {
	b.SetParallelism(100)
	log := New()
	Formatter = f
	Out = ioutil.Discard
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Info("this is a dummy log")
		}
	})
}
