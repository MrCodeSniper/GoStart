package objx

import (
	"fmt"
	"strconv"
)

// Value provides methods for extracting interface{} data in various
// types.
type Value struct {
	// data contains the raw data being managed by this Value
	data interface{}
}

// Data returns the raw data contained by this Value
func (v *Value) Data() interface{} {
	return v.data
}

// String returns the value always as a string
func (v *Value) String() string {
	switch {
	case IsStr():
		return Str()
	case IsBool():
		return strconv.FormatBool(Bool())
	case IsFloat32():
		return strconv.FormatFloat(float64(Float32()), 'f', -1, 32)
	case IsFloat64():
		return strconv.FormatFloat(Float64(), 'f', -1, 64)
	case IsInt():
		return strconv.FormatInt(int64(Int()), 10)
	case IsInt():
		return strconv.FormatInt(int64(Int()), 10)
	case IsInt8():
		return strconv.FormatInt(int64(Int8()), 10)
	case IsInt16():
		return strconv.FormatInt(int64(Int16()), 10)
	case IsInt32():
		return strconv.FormatInt(int64(Int32()), 10)
	case IsInt64():
		return strconv.FormatInt(Int64(), 10)
	case IsUint():
		return strconv.FormatUint(uint64(Uint()), 10)
	case IsUint8():
		return strconv.FormatUint(uint64(Uint8()), 10)
	case IsUint16():
		return strconv.FormatUint(uint64(Uint16()), 10)
	case IsUint32():
		return strconv.FormatUint(uint64(Uint32()), 10)
	case IsUint64():
		return strconv.FormatUint(Uint64(), 10)
	}

	return fmt.Sprintf("%#v", v.Data())
}
