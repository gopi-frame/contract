package repository

import (
	"github.com/go-viper/mapstructure/v2"
	"time"
)

type Repository interface {
	Int64(path string, defaultValue ...int64) int64
	Int64s(path string, defaultValue ...[]int64) []int64
	Int64Map(path string, defaultValue ...map[string]int64) map[string]int64
	Int(path string, defaultValue ...int) int
	Ints(path string, defaultValue ...[]int) []int
	IntMap(path string, defaultValue ...map[string]int) map[string]int
	Float64(path string, defaultValue ...float64) float64
	Float64s(path string, defaultValue ...[]float64) []float64
	Float64Map(path string, defaultValue ...map[string]float64) map[string]float64
	Duration(path string, defaultValue ...time.Duration) time.Duration
	Time(path string, layout string, defaultValue ...time.Time) time.Time
	String(path string, defaultValue ...string) string
	Strings(path string, defaultValue ...[]string) []string
	StringMap(path string, defaultValue ...map[string]string) map[string]string
	StringsMap(path string, defaultValue ...map[string][]string) map[string][]string
	Bytes(path string, defaultValue ...[]byte) []byte
	Bool(path string, defaultValue ...bool) bool
	Bools(path string, defaultValue ...[]bool) []bool
	BoolMap(path string, defaultValue ...map[string]bool) map[string]bool
	Any(path string, defaultValue ...interface{}) interface{}
	Anys(path string, defaultValue ...[]any) []any
	AnyMap(path string, defaultValue ...map[string]any) map[string]any
	Keys() []string
	All() map[string]any
	Cut(path string) Repository
	Set(key string, value any)
	Has(key string) bool
	Get(key string) any
	Delete(key string)
	Unmarshal(path string, obj any, conf *mapstructure.DecoderConfig) error
	Load(provider Provider, parser Parser) error
	LoadAt(path string, provider Provider, parser Parser) error
	Merge(in Repository)
	MergeAt(path string, in Repository)
}
