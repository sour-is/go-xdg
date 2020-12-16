package xdg

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sour-is/go-xdg/vfs"
)

type Dirs interface {
	Paths() []Path
	String() string
	First() Path
}

type DirList []Path

var _ Dirs = (DirList)(nil)

func (d DirList) String() string {
	paths := d.Paths()
	lis := make([]string, len(paths))
	for i := range paths {
		if !paths[i].IsValid() {
			continue
		}
		lis[i] = paths[i].String()
	}

	return strings.Join(lis, string(os.PathListSeparator))
}

func (d DirList) Paths() []Path {
	return []Path(d)
}

func (d DirList) First() Path {
	lis := d.Paths()
	if len(lis) == 0 {
		return path{Str("")}
	}
	return lis[0]
}

type DirFn func() []Path

var _ Dirs = (*DirFn)(nil)

func (d DirFn) String() string {
	var lis DirList = d.Paths()
	return lis.String()
}

func (d DirFn) Paths() []Path {
	var lis []Path = d()
	return lis
}

func (d DirFn) First() Path {
	lis := d.Paths()
	if len(lis) == 0 {
		return path{Str("")}
	}
	return lis[0]
}

func Find(d Dirs, name string) (string, error) {
	for _, p := range AddSuffix(d, name).Paths() {
		path := p.String()
		if exists(path) {
			return path, nil
		}
	}

	return "", fmt.Errorf("could not locate `%s` in any of the following paths: %s",
		filepath.Base(name), d)
}

func FindAll(d Dirs, name string) ([]string, error) {
	paths := AddSuffix(d, name).Paths()
	foundPaths := make([]string, 0, len(paths))

	for _, p := range paths {
		path := p.String()
		if exists(path) {
			foundPaths = append(foundPaths, path)
		}
	}

	if len(foundPaths) == 0 {
		return nil, fmt.Errorf("could not locate `%s` in any of the following paths: %s",
			filepath.Base(name), d)
	}

	return foundPaths, nil
}

func Create(d Dirs, name string) (vfs.File, error) {
	p := d.First()

	_, err := createPath(p)
	if err != nil {
		return nil, err
	}

	f, err := fs.Create(AddSuffix(p, name).String())
	if err != nil {
		return nil, err
	}

	return f, f.Chmod(0600)
}

func ParseDirs(s string) Dirs {
	sep := string(os.PathListSeparator)
	sp := strings.Split(s, sep)
	lis := make([]Path, 0, len(sp))

	for i := range sp {
		path := ParsePath(sp[i])
		if path == nil {
			continue
		}
		lis = append(lis, path)
	}

	return DirList(lis)
}

func AddSuffix(d Dirs, s string) Dirs {
	return DirFn(func() []Path {
		paths := d.Paths()
		lis := make([]Path, len(paths))
		for i := range paths {
			lis[i] = path(append(paths[i].Segments(), Str(s)))
		}

		return lis
	})
}

func EnvDefault(env Env, defaultDirs Dirs) Dirs {
	return DirFn(func() []Path {
		if env.IsSet() {
			return ParseDirs(env.String()).Paths()
		}

		if defaultDirs == nil {
			return []Path(nil)
		}

		return defaultDirs.Paths()
	})
}

func Merge(lis ...Dirs) Dirs {
	return DirFn(func() []Path {
		n := 0

		for i := range lis {
			n += len(lis[i].Paths())
		}

		dirs := make([]Path, 0, n)

		for i := range lis {
			dirs = append(dirs, lis[i].Paths()...)
		}

		return dirs
	})
}
