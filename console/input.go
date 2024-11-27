package console

import (
	"context"
	"time"
)

type Input interface {
	Context() context.Context

	Arg(index int) string
	Args() []string

	GetString(name string) (string, error)
	GetInt(name string) (int, error)
	GetInt8(name string) (int8, error)
	GetInt16(name string) (int16, error)
	GetInt32(name string) (int32, error)
	GetInt64(name string) (int64, error)
	GetUint(name string) (uint, error)
	GetUint8(name string) (uint8, error)
	GetUint16(name string) (uint16, error)
	GetUint32(name string) (uint32, error)
	GetUint64(name string) (uint64, error)
	GetFloat32(name string) (float32, error)
	GetFloat64(name string) (float64, error)
	GetBool(name string) (bool, error)
	GetDuration(name string) (time.Duration, error)
	GetTime(name string, layout string) (time.Time, error)

	GetStringSlice(name string) ([]string, error)
	GetIntSlice(name string) ([]int, error)
	GetInt8Slice(name string) ([]int8, error)
	GetInt16Slice(name string) ([]int16, error)
	GetInt32Slice(name string) ([]int32, error)
	GetInt64Slice(name string) ([]int64, error)
	GetUintSlice(name string) ([]uint, error)
	GetUint8Slice(name string) ([]uint8, error)
	GetUint16Slice(name string) ([]uint16, error)
	GetUint32Slice(name string) ([]uint32, error)
	GetUint64Slice(name string) ([]uint64, error)
	GetFloat32Slice(name string) ([]float32, error)
	GetFloat64Slice(name string) ([]float64, error)
	GetBoolSlice(name string) ([]bool, error)
	GetDurationSlice(name string) ([]time.Duration, error)
	GetTimeSlice(name string, layout string) ([]time.Time, error)

	GetStringMap(name string) (map[string]string, error)
	GetIntMap(name string) (map[string]int, error)
	GetInt8Map(name string) (map[string]int8, error)
	GetInt16Map(name string) (map[string]int16, error)
	GetInt32Map(name string) (map[string]int32, error)
	GetInt64Map(name string) (map[string]int64, error)
	GetUintMap(name string) (map[string]uint, error)
	GetUint8Map(name string) (map[string]uint8, error)
	GetUint16Map(name string) (map[string]uint16, error)
	GetUint32Map(name string) (map[string]uint32, error)
	GetUint64Map(name string) (map[string]uint64, error)
	GetFloat32Map(name string) (map[string]float32, error)
	GetFloat64Map(name string) (map[string]float64, error)
	GetBoolMap(name string) (map[string]bool, error)
	GetDurationMap(name string) (map[string]time.Duration, error)
	GetTimeMap(name string, layout string) (map[string]time.Time, error)

	VisitAll(fn func(name string, typ string))
	Unmarshal(v any) error
}
