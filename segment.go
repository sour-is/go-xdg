package xdg

import (
	"fmt"
	"os"
)

type Segment interface {
	fmt.Stringer
}

type Env string

func (e Env) String() string {
	return os.Getenv(string(e))
}

func (e Env) IsSet() bool {
	_, ok := os.LookupEnv(string(e))
	return ok
}

type Str string

func (s Str) String() string {
	return string(s)
}

type Fn func() string

func (fn Fn) String() string {
	return fn()
}
