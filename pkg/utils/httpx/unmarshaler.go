package httpx

import (
	"errors"
	"reflect"
	"sync/atomic"
)

const (
	defaultKeyName = "key"
)

var (
	errValueNotStruct = errors.New("value type is not struct")
	keyUnmarshaler    = NewUnmarshaler(defaultKeyName)
	cacheKeys         atomic.Value
)

type (
	// A Unmarshaler is used to unmarshal with given tag key.
	Unmarshaler struct {
		key  string
		opts unmarshalOptions
	}

	// UnmarshalOption defines the method to customize a Unmarshaler.
	UnmarshalOption func(*unmarshalOptions)

	unmarshalOptions struct {
		fromString bool
	}

	keyCache map[string][]string
)

func init() {
	cacheKeys.Store(make(keyCache))
}

// NewUnmarshaler returns a Unmarshaler.
func NewUnmarshaler(key string, opts ...UnmarshalOption) *Unmarshaler {
	unmarshaler := Unmarshaler{
		key: key,
	}

	for _, opt := range opts {
		opt(&unmarshaler.opts)
	}

	return &unmarshaler
}

// WithStringValues customizes a Unmarshaler with number values from strings.
func WithStringValues() UnmarshalOption {
	return func(opt *unmarshalOptions) {
		opt.fromString = true
	}
}

// UnmarshalKey unmarshals m into v with tag key.
func UnmarshalKey(m map[string]interface{}, v interface{}) error {
	return keyUnmarshaler.Unmarshal(m, v)
}

// Unmarshal unmarshals m into v.
func (u *Unmarshaler) Unmarshal(m map[string]interface{}, v interface{}) error {
	return u.UnmarshalValuer(MapValuer(m), v)
}

// UnmarshalValuer unmarshals m into v.
func (u *Unmarshaler) UnmarshalValuer(m Valuer, v interface{}) error {
	return u.unmarshalWithFullName(m, v, "")
}

func (u *Unmarshaler) unmarshalWithFullName(m Valuer, v interface{}, fullName string) error {
	rv := reflect.ValueOf(v)
	if err := ValidatePtr(&rv); err != nil {
		return err
	}

	rte := reflect.TypeOf(v).Elem()
	if rte.Kind() != reflect.Struct {
		return errValueNotStruct
	}

	rve := rv.Elem()
	numFields := rte.NumField()
	for i := 0; i < numFields; i++ {
		field := rte.Field(i)
		if usingDifferentKeys(u.key, field) {
			continue
		}

		if err := u.processField(field, rve.Field(i), m, fullName); err != nil {
			return err
		}
	}

	return nil
}

func (u *Unmarshaler) processField(field reflect.StructField, value reflect.Value, m Valuer,
	fullName string) error {
	if usingDifferentKeys(u.key, field) {
		return nil
	}

	if field.Anonymous {
		return nil // return u.processAnonymousField(field, value, m, fullName)
	}

	return nil // return u.processNamedField(field, value, m, fullName)
}
