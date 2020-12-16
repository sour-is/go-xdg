package xdg

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Path interface {
	IsValid() bool
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
		if !p[i].IsValid() {
			return ""
		}
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
func (p path) IsValid() bool {
	for _, s := range p.Segments() {
		if !s.IsValid() {
			return false
		}
	}
	return true
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
func (p pathFn) IsValid() bool {
	return p() != nil
}

func NewPath(lis ...Segment) Path {
	return path(lis)
}

func ParsePath(s string) Path {
	sep := string(os.PathSeparator)
	sp := strings.Split(s, sep)
	lis := make([]Segment, 0, len(sp))
	for i := range sp {
		switch {
		case sp[i] == "~":
			lis = append(lis, homeFn)
		case strings.HasPrefix(sp[i], "$"):
			e := Env(strings.TrimPrefix(sp[i], "$"))
			lis = append(lis, e)
		case strings.HasPrefix(sp[i], "%") && strings.HasSuffix(sp[i], "%"):
			e := Env(strings.Trim(sp[i], "%"))
			lis = append(lis, e)
		case sp[i] == "":
			lis = append(lis, Str(sep))
		case len(sp[i]) == 2 && sp[i][1] == ':' && ((sp[i][0] >= 'a' && sp[i][0] <= 'z') || (sp[i][0] >= 'A' && sp[i][0] <= 'Z')):
			lis = append(lis, Str(sp[i]+sep))
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
