package xdg

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Path interface {
	Segments() []Segment
	String() string
	Paths() []Path
	First() Path
}

type path []Segment

var _ Path = (*path)(nil)
var _ Dirs = (*path)(nil)

func (p path) String() string {
	lis := make([]string, len(p))
	for i := range p {
		lis[i] = p[i].String()
	}

	return filepath.Join(lis...)
}
func (p path) Segments() []Segment {
	return p
}
func (p path) Paths() []Path {
	return []Path{p}
}
func (p path) First() Path {
	return p
}

type pathFn func() []Segment

var _ Path = (*pathFn)(nil)
var _ Dirs = (*pathFn)(nil)

func (p pathFn) String() string {
	return path(p()).String()
}
func (p pathFn) Segments() []Segment {
	return p()
}
func (p pathFn) Paths() []Path {
	return []Path{p}
}
func (p pathFn) First() Path {
	return p
}

func NewPath(lis ...Segment) Path {
	return path(lis)
}

func ParsePath(s string) Path {
	sp := strings.Split(s, string(os.PathSeparator))
	lis := make([]Segment, 0, len(sp))
	for i := range sp {
		switch {
		case sp[i] == "~":
			lis = append(lis, homeFn)
		case strings.HasPrefix(sp[i], "$"):
			lis = append(lis, Env(strings.TrimPrefix(sp[i], "$")))
		case strings.HasPrefix(sp[i], "%") && strings.HasSuffix(sp[i], "%"):
			lis = append(lis, Env(strings.Trim(sp[i], "%")))
		case sp[i] == "":
			lis = append(lis, Str("/"))
		default:
			lis = append(lis, Str(sp[i]))
		}
	}

	return path(lis)
}

func exists(name string) bool {
	_, err := fs.Stat(name)
	return err == nil || fs.IsExist(err)
}

func createPath(p Path) (string, error) {
	var searchedPaths []string
	dir := p.String()

	if exists(dir) {
		return dir, nil
	}
	if err := fs.MkdirAll(dir, os.ModeDir|0700); err == nil {
		return dir, nil
	}

	searchedPaths = append(searchedPaths, dir)

	return "", fmt.Errorf("could not create any of the following paths: %s",
		strings.Join(searchedPaths, ", "))
}
