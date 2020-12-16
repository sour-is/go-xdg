package xdg

import (
	"fmt"
	"os"
)

type Segment interface {
	IsValid() bool
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
func (e Env) IsValid() bool {
	return e.IsSet()
}


type Str string

func (s Str) String() string {
	return string(s)
}
func (s Str) IsValid() bool {
	return true
}

type Fn func() (string, bool)

func (fn Fn) String() string {
	s, _ := fn()
	return s
}
func (fn Fn) IsValid() bool {
	_, ok := fn()
	return ok

}
