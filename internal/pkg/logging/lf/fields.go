package lf

import (
	"fmt"
	"time"
)

type Field struct {
	key   string
	value interface{}
}

func (f Field) Key() string {
	return f.key
}

func (f Field) Value() string {
	v, ok := f.value.(fmt.Stringer)
	if ok {
		return v.String()
	}
	return fmt.Sprint(f.value)
}

func String(key, value string) Field {
	return Field{key: key, value: value}
}

func Time(key string, value time.Time) Field {
	return Field{key: key, value: value}
}

func Int(key string, value int) Field {
	return Field{key: key, value: value}
}

func Any(key string, value interface{}) Field {
	return Field{key: key, value: value}
}

func Bool(key string, value bool) Field {
	return Field{key: key, value: value}
}

func Float64(key string, value float64) Field {
	return Field{key: key, value: value}
}

func Err(value error) Field {
	str := ""
	if value != nil {
		str = value.Error()
	}
	return Field{key: "error", value: str}
}
